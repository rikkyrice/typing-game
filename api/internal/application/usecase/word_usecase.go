package usecase

import (
	"api/internal/domain/model"
	"api/internal/domain/repository"

	"github.com/pkg/errors"
)

// WordUseCase 単語帳のサービスインターフェース
type WordUseCase interface {
	GetWord(id string) (*model.Word, error)
	GetWordByWordListID(id string) ([]*model.Word, error)
	PostWord(word model.Word) (*model.Word, error)
	PostAllWord(ws []model.Word) ([]*model.Word, error)
	PutWord(id string, word model.Word) (*model.Word, error)
	DeleteWord(id string) error
	DeleteAllWord(id string) error
}

// NewWordUseCase 単語帳用サービス生成
func NewWordUseCase(wR repository.WordRepository) WordUseCase {
	return &wordUseCase{
		WordRepository: wR,
	}
}

type wordUseCase struct {
	WordRepository repository.WordRepository
}

func (w *wordUseCase) GetWord(id string) (*model.Word, error) {
	word, err := w.WordRepository.FindWordByID(id)
	if err != nil {
		return nil, errors.Wrapf(err, "ID[%s]の単語がございません。", id)
	}
	return word, nil
}

func (w *wordUseCase) GetWordByWordListID(wlID string) ([]*model.Word, error) {
	words := []*model.Word{}
	words, err := w.WordRepository.FindWordByWordListID(wlID)
	if err != nil {
		return words, errors.Wrap(err, "単語が一つもありません。")
	}
	return words, nil
}

func (w *wordUseCase) PostWord(word model.Word) (*model.Word, error) {
	createdW, err := w.WordRepository.CreateWord(word)
	if err != nil {
		return nil, errors.Wrap(err, "単語の作成に失敗しました。")
	}
	return createdW, nil
}

func (w *wordUseCase) PostAllWord(words []model.Word) ([]*model.Word, error) {
	createdWs := []*model.Word{}
	createdWs, err := w.WordRepository.CreateAllWord(words)
	if err != nil {
		return nil, errors.Wrap(err, "単語の一括作成に失敗しました。")
	}
	return createdWs, nil
}

func (w *wordUseCase) PutWord(id string, word model.Word) (*model.Word, error) {
	updatedW, err := w.WordRepository.UpdateWordByID(id, word)
	if err != nil {
		return nil, errors.Wrap(err, "単語の更新に失敗しました。")
	}
	return updatedW, nil
}

func (w *wordUseCase) DeleteWord(id string) error {
	if err := w.WordRepository.RemoveWordByID(id); err != nil {
		return errors.Wrap(err, "単語の削除に失敗しました。")
	}
	return nil
}

func (w *wordUseCase) DeleteAllWord(wlID string) error {
	if err := w.WordRepository.RemoveAllWordByWordListID(wlID); err != nil {
		return errors.Wrap(err, "単語の一括削除に失敗しました。")
	}
	return nil
}
