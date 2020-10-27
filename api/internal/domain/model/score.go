package model

import "time"

// Score defines score table
type Score struct {
	// ID autoincrementされる値
	ID string
	// WordListID 単語帳ID
	WordListID string
	// PlayCount プレイ回数
	PlayCount int
	// ClearTypeCount クリアタイプ数
	ClearTypeCount int
	// MissTypeCount ミスタイプ数
	MissTypeCount int
	// CorrectRate
	CorrectRate float64
	// PlayedAt プレイ日時
	PlayedAt time.Time
}
