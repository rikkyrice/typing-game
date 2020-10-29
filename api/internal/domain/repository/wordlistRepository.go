package repository

import (
	"api/internal/domain/model"
)

// WordListRepository 単語帳リポジトリのインターフェース
type WordListRepository interface {
	// FIndWordListByID
	FindWordListByID(id string) (*model.WordList, error)
	// FindWordListByUserID ユーザーIDに紐づく単語帳を取得
	FindWordListByUserID(userID string) ([]*model.WordList, error)
	// CreateWordList 単語帳作成
	CreateWordList(wl model.WordList) (*model.WordList, error)
	// UpdateWordListByID 単語帳編集
	UpdateWordListByID(id string, wl model.WordList) (*model.WordList, error)
	// RemoveWordListByID 単語帳削除
	RemoveWordListByID(id string) error
}
