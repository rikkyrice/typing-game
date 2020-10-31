package repository

import (
	"api/internal/common/apierror"
	"api/internal/domain/model"
)

// ScoreRepository スコアリポジトリのインターフェース
type ScoreRepository interface {
	// FindScoreByWordListID 単語張IDに紐づくスコアリストを取得
	FindScoreByWordListID(wlID string) ([]*model.Score, *apierror.Error)
	// FIndLatestScoreByWordListID 単語帳DIに紐づくスコアのうち最新のものを取得
	FIndLatestScoreByWordListID(wlID string) (*model.Score, *apierror.Error)
	// CreateScore スコア作成
	CreateScore(s model.Score) (*model.Score, *apierror.Error)
	// RemoveScoreByID スコア削除
	RemoveScoreByID(id string) *apierror.Error
	// RemoveAllScoreByWordListID 単語張IDに紐づくスコアをすべて削除
	RemoveAllScoreByWordListID(wlID string) *apierror.Error
}
