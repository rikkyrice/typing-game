package main

import (
	"fmt"
	"time"

	"api/db"
	"api/internal/config"

	"github.com/pkg/errors"
)

// User is user entity
type User struct {
	userID    string
	mail      string
	password  string
	createdAt time.Time
}

func (u *User) String() string {
	return fmt.Sprintf(
		"ユーザーID:%s, メールアドレス:%s, パスワード:%s, 作成日:%s",
		u.userID,
		u.mail,
		u.password,
		u.createdAt,
	)
}

// WordList is wordlists entity
type WordList struct {
	id            int
	userID        string
	wordListTitle string
	explanation   string
	createdAt     time.Time
	updatedAt     time.Time
}

func (wl *WordList) String() string {
	return fmt.Sprintf(
		"単語帳ID:%v, ユーザーID:%s, 単語帳名:%s, 説明:%s, 作成日:%s, 更新日:%s",
		wl.id,
		wl.userID,
		wl.wordListTitle,
		wl.explanation,
		wl.createdAt,
		wl.updatedAt,
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
	for _, user := range users {
		fmt.Println(user.String())
		wordLists, err := getAllWordLists(conn, user.userID)
		if err != nil {
			fmt.Printf("取得に失敗 %+v", err)
		}
		for _, wordList := range wordLists {
			fmt.Println(wordList.String())
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
			&user.userID,
			&user.mail,
			&user.password,
			&user.createdAt,
		); err != nil {
			return []User{}, errors.Wrap(err, "結果読み込み失敗")
		}
		users = append(users, user)
	}
	return users, nil
}

func getAllWordLists(conn *db.DBConn, userID string) ([]WordList, error) {
	selectAllWordListQuery := `SELECT * FROM wordlists WHERE user_id = ?`

	selectAllWordListPstmt, err := conn.GetPstmt(selectAllWordListQuery)
	if err != nil {
		return []WordList{}, errors.Wrapf(err, "ステートメントの作成に失敗しました")
	}

	var wordLists []WordList

	rows, err := selectAllWordListPstmt.Query(userID)
	if err != nil {
		return []WordList{}, errors.Wrap(err, "クエリ実行に失敗")
	}
	for rows.Next() {
		var wordList WordList
		if err := rows.Scan(
			&wordList.id,
			&wordList.userID,
			&wordList.wordListTitle,
			&wordList.explanation,
			&wordList.createdAt,
			&wordList.updatedAt,
		); err != nil {
			return []WordList{}, errors.Wrap(err, "結果読み込み失敗")
		}
		wordLists = append(wordLists, wordList)
	}
	return wordLists, nil
}
