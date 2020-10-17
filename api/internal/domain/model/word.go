package model

import "time"

// Word defines word table
type Word struct {
	// ID autoincrementされる値
	ID int
	// WordListID 単語帳ID
	WordListID int
	// word 単語名
	Word string
	// Meaning 意味
	Meaning string
	// Explanation 説明
	Explanation string
	// CreatedAt 作成された日付
	CreatedAt time.Time
	// UpdatedAt 更新された日付
	UpdatedAt time.Time
}
