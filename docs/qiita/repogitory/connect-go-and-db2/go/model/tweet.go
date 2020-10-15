package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

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
		"ツイート本文:%s, 作成日:%s",
		t.body,
		t.createdAt,
	)
}

// GetID returns tweet's id
func (t *Tweet) GetID() string {
	return t.id
}

// GetAllTweets returns all tweet instances
func GetAllTweets(conn *sql.DB, userID string) ([]Tweet, error) {
	selectAllTweetQuery := `SELECT * FROM Tweets WHERE user_id = ?`

	selectAllTweetPstmt, err := conn.Prepare(selectAllTweetQuery)
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
