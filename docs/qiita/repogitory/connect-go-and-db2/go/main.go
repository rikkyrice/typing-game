package main

import (
	"fmt"
	"time"

	"github.com/pkg/errors"

	"local.packages/config"
	"local.packages/db"
)

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
		"ユーザーID:%s, ユーザー名:%s, メールアドレス:%s, パスワード:%s, 作成日:%s, 更新日:%s",
		u.id,
		u.name,
		u.mail,
		u.password,
		u.createdAt,
		u.updatedAt,
	)
}

// Tweet is tweets entity
type Tweet struct {
	id        string
	userID    string
	body      string
	createdAt time.Time
	updatedAt time.Time
}

func (t *Tweet) String() string {
	return fmt.Sprintf(
		"ツイートID:%s, ユーザーID:%s, 本文:%s, 作成日:%s, 更新日:%s",
		t.id,
		t.userID,
		t.body,
		t.createdAt,
		t.updatedAt,
	)
}

// Reply is replys entity
type Reply struct {
	id        string
	tweetID   string
	userID    string
	body      string
	createdAt time.Time
}

func (r *Reply) String() string {
	return fmt.Sprintf(
		"リプライID:%s, ツイートID:%s, ユーザーID:%s, 本文:%s, 作成日:%s",
		r.id,
		r.tweetID,
		r.userID,
		r.body,
		r.createdAt,
	)
}

func main() {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}

	conn, err := db.NewDBConn(c.DB)
	if err != nil {
		fmt.Printf("DBとの接続に失敗しました。%+v", err)
	}
	defer conn.Close()

	users, err := getAllUser(conn)
	if err != nil {
		fmt.Printf("取得に失敗 %+v", err)
	}

	// 件数少ないので3重for文で。
	for _, user := range users {
		fmt.Println(user.String())
		tweets, err := getAllTweets(conn, user.id)
		if err != nil {
			fmt.Printf("取得に失敗 %+v", err)
		}
		for _, tweet := range tweets {
			fmt.Println(tweet.String())
			replys, err := getAllReplys(conn, tweet.id)
			if err != nil {
				fmt.Printf("取得に失敗", err)
			}
			for _, reply := range replys {
				fmt.Println(reply.String())
			}
		}
	}
}

func getAllUser(conn *db.DBConn) ([]User, error) {
	selectAllUserQuery := `SELECT * FROM users`

	selectAllUserPstmt, err := conn.GetPstmt(selectAllUserQuery)
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

func getAllTweets(conn *db.DBConn, userID string) ([]Tweet, error) {
	selectAllTweetQuery := `SELECT * FROM Tweets WHERE user_id = ?`

	selectAllTweetPstmt, err := conn.GetPstmt(selectAllTweetQuery)
	if err != nil {
		return []Tweet{}, errors.Wrapf(err, "ステートメントの作成に失敗しました")
	}

	var tweets []Tweet

	rows, err := selectAllTweetPstmt.Query(userID)
	if err != nil {
		return []Tweet{}, errors.Wrap(err, "クエリ実行に失敗")
	}
	for rows.Next() {
		var tweet Tweet
		if err := rows.Scan(
			&tweet.id,
			&tweet.userID,
			&tweet.body,
			&tweet.createdAt,
			&tweet.updatedAt,
		); err != nil {
			return []Tweet{}, errors.Wrap(err, "結果読み込み失敗")
		}
		tweets = append(tweets, tweet)
	}
	return tweets, nil
}

func getAllReplys(conn *db.DBConn, tweetID string) ([]Reply, error) {
	selectAllReplyQuery := `SELECT * FROM Replys WHERE tweet_id = ?`

	selectAllReplyPstmt, err := conn.GetPstmt(selectAllReplyQuery)
	if err != nil {
		return []Reply{}, errors.Wrapf(err, "ステートメントの作成に失敗しました")
	}

	var replys []Reply

	rows, err := selectAllReplyPstmt.Query(tweetID)
	if err != nil {
		return []Reply{}, errors.Wrap(err, "クエリ実行に失敗")
	}
	for rows.Next() {
		var reply Reply
		if err := rows.Scan(
			&reply.id,
			&reply.tweetID,
			&reply.userID,
			&reply.body,
			&reply.createdAt,
		); err != nil {
			return []Reply{}, errors.Wrap(err, "結果読み込み失敗")
		}
		replys = append(replys, reply)
	}
	return replys, nil
}
