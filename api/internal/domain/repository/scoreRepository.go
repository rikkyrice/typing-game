package repository

import (
	"api/internal/domain/model"
)

// ScoreRepository スコアリポジトリのインターフェース
type ScoreRepository interface {
	// FindScoreByWordListID 単語張IDに紐づくスコアリストを取得
	FindScoreByWordListID(wlID string) ([]*model.Score, error)
	// FIndLatestScoreByWordListID 単語帳DIに紐づくスコアのうち最新のものを取得
	FIndLatestScoreByWordListID(wlID string) (*model.Score, error)
	// CreateScore スコア作成
	CreateScore(s model.Score) (*model.Score, error)
	// RemoveScoreByID スコア削除
	RemoveScoreByID(id string) error
	// RemoveAllScoreByWordListID 単語張IDに紐づくスコアをすべて削除
	RemoveAllScoreByWordListID(wlID string) error
}
