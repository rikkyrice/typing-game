package model

import "time"

// Score defines score table
type Score struct {
	// ID autoincrementされる値
	ID string `json:"id" validate:"required"`
	// WordListID 単語帳ID
	WordListID string `json:"wordlistID" validate:"required"`
	// ClearTypeCount クリアタイプ数
	ClearTypeCount int `json:"clearTypeCount" validate:"required"`
	// MissTypeCount ミスタイプ数
	MissTypeCount int `json:"missTypeCount" validate:"required"`
	// PlayedAt プレイ日時
	PlayedAt time.Time `json:"playedAt" validate:"required"`
}
