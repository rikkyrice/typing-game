# DB2のDBコンテナを立てる

## 概要
こんにちは。
今回はちょっとした開発やテストに使いたいときに、簡単にセットアップができて、簡単にデータを処理できるIBMのDB2コンテナの立て方と使い方を紹介します。

このコンテナは実際に私もプロジェクトの開発で利用していて、大変便利なので超おすすめです。

## 開発環境
* Win10 Enterprise
* docker (v19.03.13)
* ibmcom/db2 (v11.5.4.0)

## 前提
* dockerコマンドが打てる
* SQLの知識がある

## Db2とは？
IBMが1983年から商用で販売しているデータ管理ソフトウェア(DBMS)です。
基本的にエンタープライズ利用が多いソフトウェアだと思いますが、Docker Hubにコンテナを公開しており、こちらは制限付きで無料で使えます。

## Getting Started
それでは紹介していきます。
まず、環境はWin10なんですが、Macでも何不自由なくできます。
まぁ、Windows依存の設定やら環境やらは全く必要ないので(コンテナなのでOS依存がないのは当たり前ですが)気にせず大丈夫です。
Macの方もLinuxの方もこの記事通りに進めば上手く行くはずです。

## 1. Docker Hubからコンテナイメージをダウンロード

それではお使いの環境にDB2のコンテナイメージを落としてきましょう。
もちろんですが、Docker Hubの登録が必要になりますので、そちらは済ませておいてください。
Docker Hubのページは[こちら](https://hub.docker.com/r/ibmcom/db2)です。

PowerShell, CMD, Git Bash, Terminalいずれかで以下コマンドを打ちましょう。
最新バージョンをpullします。

```bash
$ docker pull ibmcom/db2:11.5.4.0
```

ダウンロードが始まって、コンテナイメージが落ちてきます。
2.69GBくらいの容量があるので、結構ダウンロードには時間がかかると思います。

pullが完了すると以下のように表示されます。

```bash
$ docker images
REPOSITORY               TAG                            IMAGE ID            CREATED             SIZE
ibmcom/db2               11.5.4.0                       d6b3abc02d57        3 months ago        2.69GB
```

## 2. 構成情報を用意
Db2を利用するに当たって、構成情報が必要になります。
データベースの名前とか、パスワードとかですね。
以下にサンプルを用意しました。

```env.list
LICENSE=accept
DB2INSTANCE=db2inst1
DB2INST1_PASSWORD=password
DBNAME=USERDB
BLU=false
ENABLE_ORACLE_COMPATIBILITY=false
TO_CREATE_SAMPLEDB=false
PERSISTENT_HOME=true
HADR_ENABLED=false
ETCD_ENDPOINT=
ETCD_USERNAME=
ETCD_PASSWORD=
```
* **LICENSE**:  
  Db2ソフトウェアを使う前に必要な項目
* **DB2INSTANCE**:  
  Db2インスタンスの名前。デフォルトはDB2INST1
* **DB2INST1_PASSWORD**:  
  Db2インスタンスを利用するためのパスワード。デフォルトは12文字の自動生成文字列
* **DBNAME**:  
  データベース名。
* **BLU**:  
  trueにするとBLUの利用ができる。
* **ENABLE_ORACLE_COMPATIBILITY**:  
  trueにすると、オラクルデータベースとの互換性が得られる
* **TO_CREATE_SAMPLEDB**:  
  trueにすると、あらかじめ出来上がったサンプルデータベースが作られる
* **PERSISTENT_HOME**:  
  デフォルトではtrue。Docker for Windowsを利用している場合false。
* **HADR_ENABLED**:  
  trueにすると、Db2 HADRが構成される。これをtrueにすると下の情報の入力が求められる。
* **ETCD_ENDPOINT**:  
  ETCD情報の入力。
* **ETCD_USERNAME**:  
  ETCDのusernameを入力
* **ETCD_PASSWORD**:  
  ETCDのpasswordを入力

## 3. コンテナを実行
それでは上記の情報を元にコンテナを実行していきます。

```bash
$ docker run --name test-db --detach --privileged=true -p 50000 --env-file env.list ibmcom/db2:11.5.4.0
```

* **--name**  
  コンテナ名を指定。今回はtest-db。指定しておくと、`docker stop test-db`とかで操作が楽
* **--detach**  
  コンテナをバックグラウンドで実行。
* **--privileged=true**  
  特権モードを指定
* **-p**  
  ポートを指定。今回は50000番
* **--env-file**  
  構成情報のファイルを指定
* **ibmcom/db2:11.5.4.0**  
  実行するコンテナイメージ

実行するとdockerコンテナがバックグラウンドで実行されます。
dbコンテナのセットアップには結構時間がかかります。
体感では2, 3分で完了します。
待っているだけじゃいつ終わったのかわからないので、コンテナのログを出力します。

```bash
$ docker logs -f test-db
```

長いログが流れて、`(*) Setup has completed.`が来ると完了です。

実際には、コンテナが実行されると裏で
`/var/db2_setup/lib/setup_db2_instance.sh`
というシェルスクリプトが実行されます。

このシェルは裏で
`/var/db2_setup/include/db2_common_functions`
というシェルが実行され、このシェルがDBのセットアップを行います。

さらに
`/var/custom`
というディレクトリを作成することで、その中にsqlファイルやcsvファイルを格納し、データを初期値としてデータベースに格納するシェルスクリプトを実行させることができます。
また別の記事で紹介しようと思います。

## DBコンテナに直接データを挿入
DBコンテナが作成できたので、実際にそこに入ってテーブルやデータを挿入してみます。
DBコンテナへの入り方は以下のコマンドで実行できます。

```bash
$ docker exec -it test-db bash -c "su - db2inst1"

$ db2 connect to userdb
```

これでDb2へ入れました。

試しにテーブルを確認してみます。

```bash
$ db2 "list tables"

Table/View                      Schema          Type  Creation time
------------------------------- --------------- ----- --------------------------

  0 record(s) selected.

```

当然ながらテーブルは一つもありませんが、ちゃんと表示されました。

次にテーブルを作ってデータを挿入してみます。

```bash
$ db2 "create table users(id varchar(20) not null primary key, name varchar(20) not null)"
DB20000I  The SQL command completed successfully.

$ db2 "insert into users values('rikkyrice', 'rikuhashiki')"
DB20000I  The SQL command completed successfully.

$ db2 "select * from users"

ID                   NAME
-------------------- --------------------
rikkyrice            rikuhashiki

  1 record(s) selected.

```

このような感じでデータを挿入できます。

## まとめ
データベースを欲しいときに起動し、簡単にセットアップができるのは大変便利ですね。
しかし起動がちょっと遅いのがうーむという感じですが、
しかしカスタムも簡単で、テスト用だったり開発用にはとても重宝します。

## 参考文献
* [DB2DBA: Docker上でDb2を動かしてみる](https://qiita.com/SVC34/items/71dec32ca68943432f76)
