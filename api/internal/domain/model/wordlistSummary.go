package model

import (
	"time"

	"github.com/lib/pq"
)

// WordListSummary defines wordList table
type WordListSummary struct {
	// ID autoincrementされる値
	ID string `json:"id" validate:"required"`
	// UserID ユーザーID
	UserID string `json:"userID" validate:"required"`
	// Title 単語帳名
	Title string `json:"title" validate:"required"`
	// Explanation 説明
	Explanation string `json:"explanation" validate:"required"`
	// WordCount 単語数
	WordCount int `json:"wordCount" validate:"required"`
	// PlayCount プレイ回数
	PlayCount int `json:"playCount" validate:"required"`
	// PlayedAt プレイ日時
	PlayedAt pq.NullTime `json:"playedAt"`
	// CreatedAt 作成された日付
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	// UpdatedAt 更新された日付
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}
