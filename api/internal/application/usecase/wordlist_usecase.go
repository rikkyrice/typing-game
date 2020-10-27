package usecase

import (
	"api/internal/domain/model"
	"api/internal/domain/repository"

	"github.com/pkg/errors"
)

// WordListUseCase 単語帳のサービスインターフェース
type WordListUseCase interface {
	GetWordList(userID string) ([]*model.WordList, error)
	PostWordList(wordlist model.WordList) (*model.WordList, error)
	PutWordList(id string, wordlist model.WordList) (*model.WordList, error)
	DeleteWordList(id string) error
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

func (wl *wordlistUseCase) GetWordList(userID string) ([]*model.WordList, error) {
	wordlists := []*model.WordList{}
	wordlists, err := wl.WordListRepository.FindWordListByUserID(userID)
	if err != nil {
		return wordlists, errors.Wrap(err, "単語帳が一つもありません。")
	}
	return wordlists, nil
}

func (wl *wordlistUseCase) PostWordList(wordlist model.WordList) (*model.WordList, error) {
	createdWL, err := wl.WordListRepository.CreateWordList(wordlist)
	if err != nil {
		return nil, errors.Wrap(err, "単語帳の作成に失敗しました。")
	}
	return createdWL, nil
}

func (wl *wordlistUseCase) PutWordList(id string, wordlist model.WordList) (*model.WordList, error) {
	updatedWL, err := wl.WordListRepository.UpdateWordListByID(id, wordlist)
	if err != nil {
		return nil, errors.Wrap(err, "単語帳の更新に失敗しました。")
	}
	return updatedWL, nil
}

func (wl *wordlistUseCase) DeleteWordList(id string) error {
	if err := wl.WordListRepository.RemoveWordListByID(id); err != nil {
		return errors.Wrap(err, "単語帳の削除に失敗しました。")
	}
	return nil
}
