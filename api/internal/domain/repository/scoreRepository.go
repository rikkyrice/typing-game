package repository

import (
	"api/internal/domain/model"
	"context"
)

// ScoreRepository スコアリポジトリのインターフェース
type ScoreRepository interface {
	// FindScoreByWordListID 単語張IDに紐づくスコアリストを取得
	FindScoreByWordListID(ctx context.Context, wlID int) ([]*model.Score, error)
	// FIndLatestScoreByWordListID 単語帳DIに紐づくスコアのうち最新のものを取得
	FIndLatestScoreByWordListID(ctx context.Context, wlID int) (*model.Score, error)
	// CreateScore スコア作成
	CreateScore(ctx context.Context, s model.Score) (*model.Score, error)
	// RemoveScoreByID スコア削除
	RemoveScoreByID(ctx context.Context, id int) error
	// RemoveAllScoreByWordListID 単語張IDに紐づくスコアをすべて削除
	RemoveAllScoreByWordListID(ctx context.Context, wlID int) error
}
