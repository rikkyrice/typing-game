package usecase

import (
	"api/internal/common/apierror"
	"api/internal/domain/model"
	"api/internal/domain/repository"
)

// WordListUseCase 単語帳のサービスインターフェース
type WordListUseCase interface {
	GetWordList(userID string) ([]*model.WordList, *apierror.Error)
	PostWordList(wordlist model.WordList) (*model.WordList, *apierror.Error)
	PutWordList(id string, wordlist model.WordList) (*model.WordList, *apierror.Error)
	DeleteWordList(id string) *apierror.Error
}

// NewWordListUseCase 単語帳用サービス生成
func NewWordListUseCase(wlR repository.WordListRepository) WordListUseCase {
	return &wordlistUseCase{
		WordListRepository: wlR,
	}
}

type wordlistUseCase struct {
	WordListRepository repository.WordListRepository
}

func (wl *wordlistUseCase) GetWordList(userID string) ([]*model.WordList, *apierror.Error) {
	wordlists := []*model.WordList{}
	wordlists, err := wl.WordListRepository.FindWordListByUserID(userID)
	if err != nil {
		return wordlists, err
	}
	return wordlists, nil
}

func (wl *wordlistUseCase) PostWordList(wordlist model.WordList) (*model.WordList, *apierror.Error) {
	createdWL, err := wl.WordListRepository.CreateWordList(wordlist)
	if err != nil {
		return nil, err
	}
	return createdWL, nil
}

func (wl *wordlistUseCase) PutWordList(id string, wordlist model.WordList) (*model.WordList, *apierror.Error) {
	updatedWL, err := wl.WordListRepository.UpdateWordListByID(id, wordlist)
	if err != nil {
		return nil, err
	}
	return updatedWL, nil
}

func (wl *wordlistUseCase) DeleteWordList(id string) *apierror.Error {
	if err := wl.WordListRepository.RemoveWordListByID(id); err != nil {
		return err
	}
	return nil
}
