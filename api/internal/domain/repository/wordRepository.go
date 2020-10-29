package repository

import (
	"api/internal/domain/model"
)

// WordRepository 単語リポジトリのインターフェース
type WordRepository interface {
	// FIndWordByID IDから単語を取得
	FindWordByID(id string) (*model.Word, error)
	// FindWordByWordListID 単語張IDに紐づく単語リストを取得
	FindWordByWordListID(wlID string) ([]*model.Word, error)
	// CreateWord 単語作成
	CreateWord(w model.Word) (*model.Word, error)
	// CreateAllWord 複数の単語一度に作成
	CreateAllWord(ws []model.Word) ([]*model.Word, error)
	// UpdateWordByID 単語編集
	UpdateWordByID(id string, w model.Word) (*model.Word, error)
	// RemoveWordByID 単語削除
	RemoveWordByID(id string) error
	// RemoveAllWordByWordListID 単語張IDに紐づく単語をすべて削除
	RemoveAllWordByWordListID(wlID string) error
}
