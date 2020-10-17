package model

import "time"

// WordList defines wordList table
type WordList struct {
	// ID autoincrementされる値
	ID int
	// UserID ユーザーID
	UserID string
	// Title 単語帳名
	Title string
	// Explanation 説明
	Explanation string
	// CreatedAt 作成された日付
	CreatedAt time.Time
	// UpdatedAt 更新された日付
	UpdatedAt time.Time
}
