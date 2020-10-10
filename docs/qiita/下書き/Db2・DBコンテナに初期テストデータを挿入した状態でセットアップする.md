
こんにちは。
今回はIBMのDb2コンテナを使ってスムーズにデータベースのセットアップとテストデータの挿入を行う方法を紹介します。

## 概要
Db2のコンテナは落としてこれたけど、
立ち上げた段階でテーブルやテストデータをバックグラウンドで作っておいてほしいといった方にオススメのやり方です。

Db2コンテナの落とし方は以下の記事で紹介していますので、ぜひご覧ください。
[Db2のDBコンテナを立ててちょっとしたデータを挿入してみる](https://qiita.com/rikkyrice/items/fdf45eeba6f1ddb5353d)

この記事では、上記の記事の派生として、カスタムコンテナの作る手順を紹介しています。

## 開発環境
* Win10 Enterprise
* docker (v19.03.13)
* ibmcom/db2 (v11.5.4.0)

## 前提
* SQLコマンドが打てる
* シェルスクリプトが分かる

## Getting Started
それでは紹介していきます。
こちらもWindowsでもMacでもLinuxでも操作は変わらないと思います。
Macの環境も自宅にあるので、試してみましたが、難なく動作しました。

## 1. もろもろ準備
まずはもろもろ準備します。

* コンテナイメージ
* 構成情報ファイル
* SQLファイル
* テストデータ
* コンテナイメージ内でSQLとテストデータの挿入を実行するシェルスクリプト
* Dockerfile

ディレクトリ構成

```bash:/project
project
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

### 1.1. コンテナイメージのダウンロード
dockerのローカルリポジトリにDb2のコンテナイメージを保存します。
以下のコマンドを実行します。

```bash
$ docker pull ibmcom/db2:11.5.4.0
```

既にダウンロードしている人は必要ありません。
初回は結構時間かかります。

存在確認します。

```bash
$ docker images
REPOSITORY               TAG                            IMAGE ID            CREATED             SIZE
ibmcom/db2               11.5.4.0                       d6b3abc02d57        3 months ago        2.69GB
```

今回はこのコンテナイメージをベースイメージとして利用します。

### 1.2. 構成情報ファイル準備
コンテナを実行する際にもろもろ定義した構成ファイルを読み込ませる必要があります。
これは`-e`オプションでも指定できるんですが、ファイルに置いた方が、Git管理しやすいので、推奨しています。
以下にサンプルを用意しました。

このファイルの説明は[こちら](https://qiita.com/rikkyrice/items/fdf45eeba6f1ddb5353d)でしてます。

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

### 1.3. SQLファイル準備
今回はつぶやきアプリの作成を想定したデータの準備をしてみます。

簡単なテーブル定義は以下です。

**ユーザー**

* ID (PK)
* 名前
* メールアドレス
* パスワード
* 作成日
* 更新日

**ツイート**

* ID (PK)
* ユーザーID (FK)
* 本文
* 作成日
* 更新日

**リプライ**

* ID (PK)
* ツイートID (FK)
* ユーザーID (FK)
* 本文
* 作成日
* 更新日

上記のテーブル定義を元にSQLファイルを書いていきます。
今回は以下の感じで用意しました。

* **ユーザー**

```SQL:users_create.sql
CREATE TABLE users (
    id VARCHAR(36) NOT NULL,
    name VARCHAR(40) NOT NULL,
    mail VARCHAR(100) NOT NULL,
    password VARCHAR(30) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (
        id
    )
);
```

* **ツイート**

```SQL:tweets_create.sql
CREATE TABLE tweets (
    id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    body VARCHAR(300),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (
        id
    ),

    FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);
```

* **リプライ**

```sql:replys_create.sql
CREATE TABLE replys (
    id VARCHAR(36) NOT NULL,
    tweet_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    body VARCHAR(300),
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY (
        id
    ),

    FOREIGN KEY (tweet_id) REFERENCES tweets(id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION,
    FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);
```

### 1.4. テストデータ準備
テストデータの準備をします。
テストデータもSQLファイルでinsert文を書いていけば良いんですが、
メンテナンスが大変なので、今回はCSVファイルに書いていって、それをinsertするという方法を取ります。

挿入するCSVファイルは以下に用意しました。

* **ユーザー**

```csv:users_insert.csv
hogeeee,hoge,hoge@hoge.com,hogehoge,2020-10-09-12.00.00.000000,2020-10-09-12.00.00.000000
fugaaaa,fuga,fuga@fuga.com,fugafuga,2020-10-10-12.00.00.000000,2020-10-10-12.00.00.000000
```

* **ツイート**

```csv:tweets_insert.csv
dba11ffb-b8e0-642d-bb1b-4c8053bdb4bd,hogeeee,これはテストです。,2020-10-09-12.00.00.000000,2020-10-09-12.00.00.000000
b193cb79-0e0c-85d9-2f0a-32d9774bb0aa,fugaaaa,これはテストです。,2020-10-10-12.00.00.000000,2020-10-10-12.00.00.000000
```

* **リプライ**

```csv:replys_insert.csv
7e3991a6-d3da-252f-f14f-cfed35a512a7,b193cb79-0e0c-85d9-2f0a-32d9774bb0aa,fugaaaa,テスト確認しました。,2020-10-11-12.00.00.000000
b2da92cf-6bd2-ac1c-618a-b36ef8eb94b1,dba11ffb-b8e0-642d-bb1b-4c8053bdb4bd,hogeeee,テスト確認しました。,2020-10-11-12.00.00.000000

```

ユーザーを二人用意し、それぞれがテスト用にツイートをポストして、それぞれがそのツイートにリプライしたという想定です。

### 1.5. シェルスクリプト
[こちら](https://qiita.com/rikkyrice/items/fdf45eeba6f1ddb5353d)の記事でも補足しましたが、
ibmcom/db2:11.5.4.0は`docker run`をすると裏で、
まずは、`/var/db2_setup/lib/setup_db2_instance.sh`というシェルスクリプトが実行されます。
このシェルスクリプトは環境情報を元にDb2インスタンスの設定を行い、同時に
`/var/db2_setup/include/db2_common_functions`というシェルスクリプトの実行を行っています。
こちらのシェルは、Db2のセットアップを行っています。
さらに最後に

```shell:setup_db2_instance.sh
#!/bin/bash

....

# If the /var/custom directory exists, run all scripts there. It is for products that build on top of our base image
if [[ -d /var/custom ]]; then
    echo "(*) Running user-provided scripts ... "
    for script in `ls /var/custom`; do
       echo "(*) Running $script ..."
       /var/custom/$script
    done
fi

.....

```

という記述があるのですが、
つまり最終的にコンテナ内の`/var/custom/`ディレクトリにあるスクリプトが呼ばれるわけですね。
ここにデータ挿入のシェルスクリプトを呼ばせて、初期設定するわけです。

しかし、
デフォルトでは`/var/custom`は存在しないので、自分で作る必要があります。
これは後述のDockerfile作成で紹介します。

それではテーブルを作って、データを挿入する命令を出すシェルスクリプトを書いていきます。

```shell:createschema.sh
#!/bin/bash

export PATH=/database/config/db2inst1/sqllib/bin/:$PATH

db2 connect to USERDB user db2inst1 using password

# テーブル作成
db2 -tvf /var/custom/sql/users_create.sql
db2 -tvf /var/custom/sql/tweets_create.sql
db2 -tvf /var/custom/sql/replys_create.sql

# データを挿入
db2 import from /var/custom/data/users_insert.csv of del insert into users
db2 import from /var/custom/data/tweets_insert.csv of del insert into tweets
db2 import from /var/custom/data/replys_insert.csv of del insert into replys

# Terminate
db2 terminate
touch /tmp/end.txt

```

まずdb2コマンドを使えるようにPATHを通しています。
そして、SQLファイルを順番に実行していき、上で宣言したテーブルを作っていきます。

`db2 -tvf`の意味は

* -t
  セミコロン(;)をステートメントの終了文字として認識し、コマンドを実行するオプション
* -v
  コマンド・テキストを標準出力するオプション。ログでの確認に便利
* -f
  コマンドをファイルから読み込むオプション。今回はSQLファイルからコマンドを読み込むので指定

`db2 import from ${filename} of del insert into ${tablename}`の意味は

* import from filename
  ファイルからデータを読み込む構文。
* of del insert into
  csvファイルではカンマ区切りでデータを指定したため、カンマ区切りでデータの挿入が行われる

### 1.6. Dockerfileの準備
それでは最後にカスタムDb2コンテナのイメージをビルドする`Dockerfile`を作成します。

もろもろ設定あると思いますが、かなりそぎ落として簡単なDockerfileにします。

再掲ですが、プロジェクトのディレクトリ構成は以下です。

```bash:/project
project
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

```dockerfile:Dockerfile
FROM ibmcom/db2:11.5.4.0
RUN mkdir /var/custom
RUN mkdir /var/custom/sql
RUN mkdir /var/custom/data
COPY ./sql/*.sql /var/custom/sql/
COPY ./data/*.csv /var/custom/data/
COPY createschema.sh /var/custom/
RUN chmod 444 /var/custom/sql/*.sql
RUN chmod 444 /var/custom/data/*.csv
RUN chmod a+x /var/custom/createschema.sh
```

まずベースイメージとして、落としてきた`ibmcom/db2:11.5.4.0`を指定します。
このコンテナをカスタムしていきます。

すでに伝えたように`/var/custom/`ディレクトリを作ることで、
セットアップ時にそこに指定したスクリプトが呼ばれるため、`mkdir`をします。
sqlファイルなどは専用のディレクトリを作りました。

そこに用意したファイルをコピーしていきます。
さらに権限を読み取り専用で設定します。
createschema.shはシェルスクリプトを実行する必要があるので、実行権限を与えます。

これでカスタムイメージをビルドする準備ができました。

## 2. カスタムコンテナイメージ作成

続いて、準備したもろもろを使って、カスタマイズしたコンテナイメージを作成します。
みなさんご存知`docker build`ですね。

今回は識別しやすいように`test-db:v1.0`というタグを付けていきます。

```bash
$ docker build -t test-db:v1.0 .
```

ディレクトリは`Dockerfile`がある場所で実行してください。
初回はかなり時間がかかると思います。
キャッシュがある方は1秒くらいで終わりますかね。

ビルドが完了したら存在を確認します。

```bash
$ docker images
REPOSITORY               TAG                            IMAGE ID            CREATED             SIZE
test-db                  v1.0                           186064b82d09        28 minutes ago      2.69GB
```

できてますね。

## 3. カスタムコンテナ実行

それではカスタムしたコンテナを起動しましょう。

```bash
$ docker run --name test-db --restart=always --detach --privileged=true -p 50000 --env-file env.list test-db:v1.0
```

コンテナ名は`test-db`として、構成情報は用意した`env.list`を指定しています。

実行すると、コンテナIDを吐くだけで、いつセットアップが終わったのか分からないので、ログを見てみます。

```bash
$ docker logs -f test-db


SQL3109N  The utility is beginning to load data from file
"/var/custom/data/replys_insert.csv".
(*) Previous setup has not been detected. Creating the users...
(*) Creating users ...
(*) Creating instance ...
DB2 installation is being initialized.

 Total number of tasks to be performed: 4
Total estimated time for all tasks to be performed: 309 second(s)

Task #1 start
Description: Setting default global profile registry variables
Estimated time 1 second(s)
Task #1 end

Task #2 start
Description: Initializing instance list
Estimated time 5 second(s)
Task #2 end

Task #3 start
Description: Configuring DB2 instances
Estimated time 300 second(s)
Task #3 end

Task #4 start
Description: Updating global profile registry
Estimated time 3 second(s)
Task #4 end

The execution completed successfully.

.........省略.........

SQL3110N  The utility has completed processing.  "2" rows were read from the
input file.

SQL3221W  ...Begin COMMIT WORK. Input Record Count = "2".

SQL3222W  ...COMMIT of any database changes was successful.

SQL3149N  "2" rows were processed from the input file.  "2" rows were
successfully inserted into the table.  "0" rows were rejected.


Number of rows read         = 2
Number of rows skipped      = 0
Number of rows inserted     = 2
Number of rows updated      = 0
Number of rows rejected     = 0
Number of rows committed    = 2

DB20000I  The TERMINATE command completed successfully.
(*) Running data ...
/var/db2_setup/lib/setup_db2_instance.sh: line 201: /var/custom/data: Is a directory
(*) Running sql ...
/var/db2_setup/lib/setup_db2_instance.sh: line 201: /var/custom/sql: Is a directory
          from "/database/data/db2inst1/NODE0000/SQL00001/LOGSTREAM0000/".

2020-10-10-01.35.55.270890+000 E239025E525           LEVEL: Event
PID     : 18622                TID : 140605766231808 PROC : db2sysc 0
INSTANCE: db2inst1             NODE : 000            DB   : USERDB
APPHDL  : 0-7                  APPID: *LOCAL.db2inst1.201010013552
AUTHID  : DB2INST1             HOSTNAME: 99a855c216d7
EDUID   : 22                   EDUNAME: db2agent (idle) 0
FUNCTION: DB2 UDB, base sys utilities, sqeLocalDatabase::FreeResourcesOnDBShutdown, probe:16544
STOP    : DATABASE: USERDB   : DEACTIVATED: NO
```

こういう感じのログが吐かれると思います。

終了フラグとしては
`DB20000I  The TERMINATE command completed successfully.`
このログが出ていたら、セットアップは終了しています。

テーブルの作成は、少し上にさかのぼっていただくと、

```bash
CREATE TABLE users ( id VARCHAR(36) NOT NULL, name VARCHAR(40) NOT NULL, mail VARCHAR(100) NOT NULL, password VARCHAR(30) NOT NULL, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL, PRIMARY KEY ( id ) )
DB20000I  The SQL command completed successfully.
```
というログが出ていまして、何もエラーが無ければ`The SQL command completed successfully`で正常に終了しています。

データ挿入に関しては、`Number of rows committted = 2`とかのログが見えると思います。
これが挿入したデータがどういったステータスで終了したかを表すログになります。
今回は`inserted = 2`で、`rejected = 0`なので、すべて正常に挿入されています。
これが`rejected = 2`となっていたら、2件のデータ挿入に失敗しているので、調査が必要です。

もし途中で失敗していたり、実行しなおしたいと思ったときは、

```bash
# 実行コンテナの強制削除
$ docker rm -f test-db
test-db

# コンテナイメージの削除
$ docker rmi test-db:v1.0
```

で、きれいさっぱり削除できます。
コンテナを実行しなおしたいだけの時は上のコマンドを、SQLファイルを変更したり、Dockerfileを変更したりした時は、上と下のコマンド両方を実行します。

## 4. 挿入したデータの存在確認
それでは、挿入したデータが期待通りの値でちゃんと挿入されているか確認してみましょう。

実際にコンテナ内に入ってDBと接続し、SELECT文を打ってみます。

```bash
$ docker exec -it test-db bash -c "su - db2inst1"
Last login: Sat Oct 10 01:46:02 UTC 2020
[db2inst1@99a855c216d7 ~]$ db2 connect to userdb

   Database Connection Information

 Database server        = DB2/LINUXX8664 11.5.4.0
 SQL authorization ID   = DB2INST1
 Local database alias   = USERDB

[db2inst1@99a855c216d7 ~]$ db2 "select * from users"
ID                                   NAME                                     MAIL                                                                                                 PASSWORD                       CREATED_AT                 UPDATED_AT
------------------------------------ ---------------------------------------- ---------------------------------------------------------------------------------------------------- ------------------------------ -------------------------- --------------------------
hogeeee                              hoge                                     hoge@hoge.com                                                                                        hogehoge                       2020-10-09-12.00.00.000000 2020-10-09-12.00.00.000000
fugaaaa                              fuga                                     fuga@fuga.com                                                                                        fugafuga                       2020-10-10-12.00.00.000000 2020-10-10-12.00.00.000000

  2 record(s) selected.

[db2inst1@99a855c216d7 ~]$ 
```

ちゃんとcsvファイルで書いたデータが挿入されていますね。
コンテナから出るときは`exit`と打つと出れます。

## まとめ
さぁ、いかがでしたでしょうか？
簡単にではございますが、Db2にデータを挿入した状態でのセットアップ方法を紹介しました。
このコンテナイメージは使いまわせるので、他のプロジェクトで別のDockerfileを書いて別のカスタムイメージもすぐに作れますし、
簡単にそのプロジェクトの仕様に沿ったデータベースを構築することができます。

また、揮発性なので、データを挿入しまくっても、`docker rm -f`で消してしまえば、次の立ち上げの時には簡単に初期化できるので、テストも大変しやすくなります。

このデータベースをGUIで操作する方法など、今後調べていこうと思います。

また、実際にこのデータベースに接続して、データを操作する方法をGo言語を使って説明する記事も書こうと思います。

## 参考文献
