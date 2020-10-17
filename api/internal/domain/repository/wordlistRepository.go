package repository

import (
	"api/internal/domain/model"
	"context"
)

// WordListRepository 単語帳リポジトリのインターフェース
type WordListRepository interface {
	// FindWordListByUserID ユーザーIDに紐づく単語帳を取得
	FindWordListByUserID(ctx context.Context, userID string) ([]*model.WordList, error)
	// FIndWordListByID
	FindWordListByID(ctx context.Context, id int) (*model.WordList, error)
	// CreateWordList 単語帳作成
	CreateWordList(ctx context.Context, wl model.WordList) (*model.WordList, error)
	// UpdateWordListByID 単語帳編集
	UpdateWordListByID(ctx context.Context, wl model.WordList) (*model.WordList, error)
	// RemoveWordListByID 単語帳削除
	RemoveWOrdListByID(ctx context.Context, id int) error
}
