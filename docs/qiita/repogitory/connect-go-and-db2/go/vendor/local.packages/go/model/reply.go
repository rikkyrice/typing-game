package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

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

// GetAllReplys returns all reply instances
func GetAllReplys(conn *sql.DB, tweetID string) ([]Reply, error) {
	selectAllReplyQuery := `SELECT * FROM Replys WHERE tweet_id = ?`

	selectAllReplyPstmt, err := conn.Prepare(selectAllReplyQuery)
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
