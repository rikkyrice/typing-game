package model

import "time"

// Word defines word table
type Word struct {
	// ID autoincrementされる値
	ID string `json:"id" validate:"required"`
	// WordListID 単語帳ID
	WordListID string `json:"wordlistID" validate:"required"`
	// word 単語名
	Word string `json:"word" validate:"required"`
	// Meaning 意味
	Meaning string `json:"meaning" validate:"required"`
	// Explanation 説明
	Explanation string `json:"explanation" validate:"required"`
	// CreatedAt 作成された日付
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	// UpdatedAt 更新された日付
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}
