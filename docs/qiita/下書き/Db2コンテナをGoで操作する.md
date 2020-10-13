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
MacやLinuxの場合はSQLのCLIドライバのパスを通す必要があるので、それに関しては後述します。

今回は、Db2との疎通確認に重きを置いておりますので、API化などは行っていません。
単純にDb2からデータを取ってきて、コンソールに出力するだけのプログラムを書いていきます。(いつかGoでREST APIの実装も紹介します。)

## 1. フォルダ構成の説明
まずはフォルダ構成を説明します。

```bash:project
project
├─go
|  ├─config
|  |      ├─env.yaml
|  |      └─config.go
|  ├─db
|  |  └─db.go
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
  - **/config**  
    - **env.yaml**
      データベース接続用の情報を格納
    - **config.go**  
      データベース接続情報を内部で取り出す構造体を定義
  - **/db/db.go**  
    データベースの構造体を定義
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
クライアントに公開している50000ポートはenv.yamlで定義する必要があるので、覚えておきます。

## 3. Goの実装方針
main.goでの実行流れは大体以下のようになります。

1. config.goでConfig構造体にenv.yamlの情報を注入したインスタンスを作成。
2. db.goにConfigインスタンスを渡し、DBとのコネクションを張る。
3. 張った情報を持つDB構造体のインスタンスを作成
4. main.go内でDBインスタンスのコネクションを用いてSQL文を発行しデータのやり取りをする。

本当は、ドメイン駆動設計とかで、ユーザードメインとか、インフラストラクチャとか作ってカッコいい設計をしたいんですが、それはまたの機会ということで。

## 4. インポートするパッケージ
利用するパッケージ
* github.com/ibmdb/go_ibm_db
* gopkg.in/yaml.v2

基本的にGoでDb2を利用する際は、`github.com/ibmdb/go_ibm_db`というパッケージを利用します。

Goをインストールしている方は`GOPATH`が通っているか確認しましょう。
私のWindowsでは`GOROOT`を`c:/go`に、`GOPATH`を`c:/users/usrname/go`に指定しており、Goのパッケージ関連は全て`GOPATH`である`c:/users/usrname/go`直下に落とされます。

確認の仕方は

```bash
$ go env
```

でできます。
設定していなかったら環境変数の追加などで対応しましょう。

もし余計なパッケージを環境に落としたくないという方は、`go mod`でもできます。

それでは`github.com/ibmdb/go_ibm_db`を落とします。
以下のコマンドを叩きます。

```bash
$ go get github.com/ibmdb/go_ibm_db
```

また、今回の実装では、yamlファイルから構造体のインスタンスを作る場面が現れるので、`gopkg.in/yaml.v2`も落としてきます。

```bash
$ go get gopkg.in/yaml.v2
```



## 5. Goの実装


## 6. 実行結果

## 7. Mac, Linuxでの実装

## 8. まとめ
