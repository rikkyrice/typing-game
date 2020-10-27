package model

import "time"

// WordList defines wordList table
type WordList struct {
	// ID autoincrementされる値
	ID int `json:"id" validate:"required"`
	// UserID ユーザーID
	UserID string `json:"userID" validate:"required"`
	// Title 単語帳名
	Title string `json:"title" validate:"required"`
	// Explanation 説明
	Explanation string `json:"explanation" validate:"required"`
	// CreatedAt 作成された日付
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	// UpdatedAt 更新された日付
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}
