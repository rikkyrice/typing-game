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
* gopkg.in/yaml.v2
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

無事落とせてこれた方は`PathToInstaller/installer/clidriver/bin`のパスを通す必要があるので、通しましょう。
これでgo_ibm_dbのセットアップは完了です。

もし余計なパッケージを環境に落としたくないという方は、`go mod`でもできます。
しかしその場合も、`sqlcli.h`は必要になりますので、インストールしてきたinstallerをプロジェクトにコピーしてきて、、シェルスクリプトなどで、`clidriver/bin`のパスを通し、moduleを指定してビルドすることで実行ファイルを生成できます。

### 3.2. yaml.v2
また、今回の実装では、yamlファイルから構造体のインスタンスを作る場面が現れるので、`gopkg.in/yaml.v2`も落としてきます。

```bash
$ go get gopkg.in/yaml.v2
```

### 3.3. errors
さらにエラーの実装もするので、`errors`パッケージも落としましょう。

```bash
$ go get github.com/pkg/errors
```

## 4. Goの実装
基本的に実装は本当に3で紹介した通りです。
main.goのmain関数を見ながら紹介します。

まずこのコード

```go:main.go
  c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}
```

configパッケージの`Init`関数を呼んでいます。
この関数は何をするかというと
`config/env.yaml`の内容を読んできて、それをConfig構造体に落としています。

## 5. 実行結果

## 6. Mac, Linuxでの実装

## 7. まとめ
