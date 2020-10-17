package repository

import (
	"api/internal/domain/model"
	"context"
)

// WordRepository 単語リポジトリのインターフェース
type WordRepository interface {
	// FindWordByWordListID 単語張IDに紐づく単語リストを取得
	FindWordByWordListID(ctx context.Context, wlID int) ([]*model.Word, error)
	// FIndWordByID IDから単語を取得
	FindWordByID(ctx context.Context, id int) (*model.Word, error)
	// CreateWord 単語作成
	CreateWord(ctx context.Context, w model.Word) (*model.Word, error)
	// CreateAllWord 複数の単語一度に作成
	CreateAllWord(ctx context.Context, ws []model.Word) ([]*model.Word, error)
	// UpdateWordByID 単語編集
	UpdateWordByID(ctx context.Context, w model.Word) (*model.Word, error)
	// RemoveWordByID 単語削除
	RemoveWordByID(ctx context.Context, id int) error
	// RemoveAllWordByWordListID 単語張IDに紐づく単語をすべて削除
	RemoveAllWordByWordListID(ctx context.Context, wlID int) error
}
