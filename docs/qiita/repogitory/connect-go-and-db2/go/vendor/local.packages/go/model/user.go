package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
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
		"ユーザー名:%s",
		u.name,
	)
}

// GetID returns user's id
func (u *User) GetID() string {
	return u.id
}

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
