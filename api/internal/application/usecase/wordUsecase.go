package usecase

import (
	"api/internal/common/apierror"
	"api/internal/domain/model"
	"api/internal/domain/repository"
)

// WordUseCase 単語帳のサービスインターフェース
type WordUseCase interface {
	GetWord(id string) (*model.Word, *apierror.Error)
	GetWordByWordListID(id string) ([]*model.Word, *apierror.Error)
	PostWord(word model.Word) (*model.Word, *apierror.Error)
	PostAllWord(ws []model.Word) ([]*model.Word, *apierror.Error)
	PutWord(id string, word model.Word) (*model.Word, *apierror.Error)
	DeleteWord(id string) *apierror.Error
	DeleteAllWord(id string) *apierror.Error
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

func (w *wordUseCase) GetWord(id string) (*model.Word, *apierror.Error) {
	word, err := w.WordRepository.FindWordByID(id)
	if err != nil {
		return nil, err
	}
	return word, nil
}

func (w *wordUseCase) GetWordByWordListID(wlID string) ([]*model.Word, *apierror.Error) {
	words := []*model.Word{}
	words, err := w.WordRepository.FindWordByWordListID(wlID)
	if err != nil {
		return words, err
	}
	return words, nil
}

func (w *wordUseCase) PostWord(word model.Word) (*model.Word, *apierror.Error) {
	createdW, err := w.WordRepository.CreateWord(word)
	if err != nil {
		return nil, err
	}
	return createdW, nil
}

func (w *wordUseCase) PostAllWord(words []model.Word) ([]*model.Word, *apierror.Error) {
	createdWs := []*model.Word{}
	createdWs, err := w.WordRepository.CreateAllWord(words)
	if err != nil {
		return nil, err
	}
	return createdWs, nil
}

func (w *wordUseCase) PutWord(id string, word model.Word) (*model.Word, *apierror.Error) {
	updatedW, err := w.WordRepository.UpdateWordByID(id, word)
	if err != nil {
		return nil, err
	}
	return updatedW, nil
}

func (w *wordUseCase) DeleteWord(id string) *apierror.Error {
	if err := w.WordRepository.RemoveWordByID(id); err != nil {
		return err
	}
	return nil
}

func (w *wordUseCase) DeleteAllWord(wlID string) *apierror.Error {
	if err := w.WordRepository.RemoveAllWordByWordListID(wlID); err != nil {
		return err
	}
	return nil
}
