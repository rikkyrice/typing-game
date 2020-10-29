こんにちは。
今回はDockerで立てたDb2コンテナをGoで操作する方法について紹介します。
Db2コンテナの立て方や、セットアップ時にデータを挿入する方法については、以下の記事で紹介しておりますので、ぜひ参考にしてください。

[Db2のDBコンテナを立ててちょっとしたデータを挿入してみる](https://qiita.com/rikkyrice/items/fdf45eeba6f1ddb5353d)
[Db2/DBコンテナに初期テストデータを挿入した状態でセットアップ](https://qiita.com/rikkyrice/items/20946fc5b4e2153b0f87)

今回は、Db2コンテナをデータを挿入した状態で立ち上げ、Goでの実装を中心に紹介します。

## 概要
Db2コンテナにデータは挿入できたけど、実際にそのデータを取得してきて操作したり、データを更新したりするにはどうするんだという方向けの内容となっています。

今回はGo言語でDb2からデータを取ってくる方法をご紹介しようと思います。

## 開発環境
* Win10 Enterprise
* docker (v19.03.13)
* ibmcom/db2 (v11.5.4.0)
* Go (v1.12.9)
* Git Bash (git version 2.27.0.windows.1)

## 前提
* ibmcomのDb2コンテナの立ち上げ方がある程度分かる
* Go言語がある程度わかる

## Getting Started
開発環境はWindowsですが、MacでもLinuxでもできます。

今回は、Db2との疎通確認に重きを置いておりますので、API化などは行っていません。
単純にDb2からデータを取ってきて、コンソールに出力するだけのプログラムを書いていきます。(いつかGoでREST APIの実装も紹介します。)

## 1. フォルダ構成の説明
まずはフォルダ構成を説明します。

```bash:project
project
├─go
|  ├─model
|  |     ├─user.go
|  |     ├─tweet.go
|  |     └─reply.go
|  └─main.go
└─db  
   ├─data
   |    ├─users_insert.csv
   |    ├─tweets_insert.csv
   |    └─replys_insert.csv
   ├─sql
   |   ├─users_create.sql
   |   ├─tweets_create.sql
   |   └─replys_create.sql
   ├─createschema.sh
   ├─Dockerfile
   └─env.list
```
- **/go**
  - **/model**  
    - **user.go**
      ユーザーのDTOとDAO
    - **tweet.go**
      ツイートのDTOとDAO
    - **reply.go**  
      リプライのDTOとDAO
  - **main.go**  
    メインの関数  
- **/db:** データベースをセットアップするフォルダ
  - **/data**  
    初期にデータベースに登録するテストデータのフォルダ
  - **/sql**  
    テーブル作成するSQL文のフォルダ
  - **createschema.sh**  
    データベースセットアップ時に呼ばれるテーブル作成用スクリプト
  - **Dockerfile**  
    コンテナ定義
  - **env.list**  
    Db2コンテナ用の構成情報

本当は、ドメイン駆動設計とかで、ユーザードメインとか、インフラストラクチャとか作ってカッコいい設計をしたいんですが、それはまたの機会ということで。

## 2. コンテナの立ち上げ
まずは`Dockerfile`を用いて、コンテナイメージをビルドします。
実行するコマンドは以下です。

```bash
$ cd db
$ docker build -t test-db:v1.0 .
```

これで、コンテナイメージが出来上がるので、早速runしていきます。

```bash
$ docker run --name go-db --restart=always --detach --privileged=true -p 50000:50000 --env-file env.list test-db:v1.0
```

詳しい説明は[こちら](https://qiita.com/rikkyrice/items/20946fc5b4e2153b0f87)で紹介しています。

ここで大事なのはポートを50000:50000でポートフォワーディングしていることです。
クライアントに公開している50000ポートはDBと接続する時に指定する必要があるので、覚えておきます。

## 3. インポートするパッケージ
利用するパッケージ
* github.com/ibmdb/go_ibm_db
* github.com/pkg/errors

### 3.1. go_ibm_db
基本的にGoでDb2を利用する際は、`github.com/ibmdb/go_ibm_db`というパッケージを利用します。

以下のコマンドを叩きます。

```bash
$ go get github.com/ibmdb/go_ibm_db
```

またデータベースを操作するにあたって、SQLを操作するためのドライバが必要になります。
色々操作があるので順にやります。

まず、落としてきた`github.com/ibmdb/go_ibm_db`を見に行きます。
おそらく`GOPATH`配下に落とされていると思うので、こちらの階層を下ると、`installer`というフォルダにぶち当たります。
このフォルダ内`setup.go`がclidriverのダウンロードスクリプトになっています。

```bash
$ cd PathToInstaller/installer
$ go run setup.go
```

これでclidriverが`installer`配下にダウンロードできます。(パーミッションエラーが起きた方は、installerフォルダの権限を変えてみてください。)
結構時間がかかる気がします。

無事落とせてこれた方は`PathToInstaller/installer/clidriver/bin`のパスを通す必要があるので、通しましょう。
これでgo_ibm_dbのセットアップは完了です。

もし余計なパッケージを環境に落としたくないという方は、`go mod`でもできます。
しかしその場合も、`sqlcli.h`は必要になりますので、インストールしてきたinstallerをプロジェクトにコピーしてきて、、シェルスクリプトなどで、`clidriver/bin`のパスを通し、moduleを指定してビルドすることで実行ファイルを生成できます。

### 3.2. errors
また、エラーの実装もするので、`errors`パッケージも落としましょう。

```bash
$ go get github.com/pkg/errors
```

## 4. Goの実装
基本的に実装は本当に3で紹介した通りです。
main.goのmain関数を見ながら紹介します。

まずこのコード

```go:main.go
  config := "HOSTNAME=localhost;DATABASE=USERDB;PORT=50000;UID=db2inst1;PWD=password"
	conn, err := sql.Open("go_ibm_db", config)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()
```

configにDB接続情報を格納します。HOSTNAMEとPORT以外はenv.listに乗せてある情報を使います。
その下の`sql.Open`でDBとのコネクションを張ります。
一つ目の引数はドライバ名を指定します。今回は`go_ibm_db`です。
二つ目の引数はDB接続情報を指定します。エラーを取りうるので、エラー処理もかかせず行います。
コネクションは必ず終了する必要があるので、Goのプラクティスである`defer`を使ってコネクションを閉じましょう。

これでDb2コンテナとのコネクションが取得できました。
これを利用してデータを操作していきます。

まずはユーザーを全件取得して、情報をユーザー構造体に格納し、インスタンスの配列を作っています。

```go:main.go
users, err := model.GetAllUser(conn)
if err != nil {
  fmt.Printf("取得に失敗 %+v", err)
}
```

ではユーザーDAOとDTOを定義しているuser.goを見ていきます。

```go:user.go
// User is users entity
type User struct {
	id        string
	name      string
	mail      string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func (u *User) String() string {
	return fmt.Sprintf(
		"ユーザー名:%s",
		u.name,
	)
}

// GetID returns user's id
func (u *User) GetID() string {
	return u.id
}
```

ユーザー構造体はテーブル定義のカラムをフィールドに定義しています。
GetIDメソッドはユーザーのIDを取得するメソッドです。これは他のテーブルのクエリにIDを渡すためにユーザー構造体のフィールドがプライベートに指定されているため、書いています。
まぁここら辺は他の言語でも似たようなことやると思います。

その下、ユーザー全件取得メソッドですが、

```go:user.go
// GetAllUser returns all user instances
func GetAllUser(conn *sql.DB) ([]User, error) {
	selectAllUserQuery := `SELECT * FROM users`

	selectAllUserPstmt, err := conn.Prepare(selectAllUserQuery)
	if err != nil {
		return []User{}, errors.Wrapf(err, "ステートメントの作成に失敗しました")
	}

	var users []User

	rows, err := selectAllUserPstmt.Query()
	if err != nil {
		return []User{}, errors.Wrap(err, "クエリ実行に失敗")
	}
	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.id,
			&user.name,
			&user.mail,
			&user.password,
			&user.createdAt,
			&user.updatedAt,
		); err != nil {
			return []User{}, errors.Wrap(err, "結果読み込み失敗")
		}
		users = append(users, user)
	}
	return users, nil
}
```

ここは色んな書き方があるんですが、Prepare()メソッドでステートメントを用意してから、queryを実行する方法で書きます。

これを実行すると、取れてきたレコードが`rows`に格納されます。
`rows`はNextメソッドを持っていて、for文でそれぞれのレコードを回すことができます。
さらに`rows.Scan()`にユーザーインスタンスの情報を渡してあげると、そこにレコードの情報を格納してくれます。

これで、ユーザー情報をユーザーインスタンスに格納することができました。
ユーザーの配列を返します。

それではmainに戻ります。

次からはユーザーインスタンスからIDを取ってきて、Tweetの`WHERE句`に渡して挙げて、ユーザーに紐づくレコードを取ってきています。
取ってきたtweetレコードからさらにIDを取ってきて、それに紐づくReplyを取得し出力、それをユーザーレコード分行うといった処理をしています。

```go:main.go
// 件数少ないので3重for文で。
	for _, user := range users {
		fmt.Println(user.String())
		tweets, err := model.GetAllTweets(conn, user.GetID())
		if err != nil {
			fmt.Printf("取得に失敗 %+v", err)
		}
		for _, tweet := range tweets {
			fmt.Println(tweet.String())
			replys, err := model.GetAllReplys(conn, tweet.GetID())
			if err != nil {
				fmt.Printf("取得に失敗", err)
			}
			for _, reply := range replys {
				fmt.Println(reply.String())
			}
		}
	}
```

`WHERE句`にIDを渡すためにはSQL文を`SELECT * FROM Tweets WHERE user_id = ?`のように与えたいパラメータの箇所を`?`とします。
パラメータ分第2引数を与えることで、`WHERE句`をカスタムできます。

書き方は、
`rows, err := selectAllTweetPstmt.Query(userID)`
このような形です。


## 5. 実行結果
Windowsで実行すると、コンテナから値を受け取ってくる段階で、日本語箇所は文字化けして表示されてしまいます。
Db2で用いているコンテナがLinuxコンテナなので、文字コードがUTF-8のまま文字列が送られてくることに起因していると思われます。

実行結果は以下のようになります。

```bash
ユーザー名:hoge
ツイート本文:�����̓e�X�g�ł��B, 作成日:2020-10-09 12:00:00 +0900 JST
リプライユーザー名:fugaaaa, リプライ本文:�e�X�g�m�F���܂����B, 作成日:2020-10-11 12:00:00 +0900 JST
-----------------------
ユーザー名:fuga
ツイート本文:�����̓e�X�g�ł��B, 作成日:2020-10-10 12:00:00 +0900 JST
リプライユーザー名:hogeeee, リプライ本文:�e�X�g�m�F���܂����B, 作成日:2020-10-11 12:00:00 +0900 JST
-----------------------
```

まぁめっちゃ文字化けしてますね。
悲しいです。
このままだとあれなんで、Macで実行した結果も載せときます。

```bash
ユーザー名:hoge
ツイート本文:これはテストです。, 作成日:2020-10-09 12:00:00 +0900 JST
リプライユーザー名:fugaaaa, リプライ本文:テスト確認しました。, 作成日:2020-10-11 12:00:00 +0900 JST
-----------------------
ユーザー名:fuga
ツイート本文:これはテストです。, 作成日:2020-10-10 12:00:00 +0900 JST
リプライユーザー名:hogeeee, リプライ本文:テスト確認しました。, 作成日:2020-10-11 12:00:00 +0900 JST
-----------------------
```

こんな感じで、Db2から取得できています。

## 6. まとめ
文字コードの弊害がありながらも、GoでDb2コンテナに接続する手法を紹介しました。

これでAPI開発とか楽に行えますね。
