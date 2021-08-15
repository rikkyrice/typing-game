package repository

import (
	"api/internal/common/apierror"
	"api/internal/domain/model"
)

// WordListRepository 単語帳リポジトリのインターフェース
type WordListRepository interface {
	// FIndWordListByID
	FindWordListByID(id string) (*model.WordListSummary, *apierror.Error)
	// FindWordListByUserID ユーザーIDに紐づく単語帳を取得
	FindWordListByUserID(userID string) ([]*model.WordListSummary, *apierror.Error)
	// CreateWordList 単語帳作成
	CreateWordList(wl model.WordList) (*model.WordList, *apierror.Error)
	// UpdateWordListByID 単語帳編集
	UpdateWordListByID(id string, wl model.WordList) (*model.WordList, *apierror.Error)
	// RemoveWordListByID 単語帳削除
	RemoveWordListByID(id string) *apierror.Error
}
