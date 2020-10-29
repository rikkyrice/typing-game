package usecase

import (
	"api/internal/domain/model"
	"api/internal/domain/repository"

	"github.com/pkg/errors"
)

// ScoreUseCase スコアのサービスインターフェース
type ScoreUseCase interface {
	GetScores(wlID string) ([]*model.Score, error)
	GetLatestScore(wlID string) (*model.Score, error)
	PostScore(score model.Score) (*model.Score, error)
	DeleteAllScore(wlID string) error
}

// NewScoreUseCase スコア用サービス生成
func NewScoreUseCase(sR repository.ScoreRepository) ScoreUseCase {
	return &scoreUseCase{
		ScoreRepository: sR,
	}
}

type scoreUseCase struct {
	ScoreRepository repository.ScoreRepository
}

func (s *scoreUseCase) GetScores(wlID string) ([]*model.Score, error) {
	scores := []*model.Score{}
	scores, err := s.ScoreRepository.FindScoreByWordListID(wlID)
	if err != nil {
		return scores, errors.Wrap(err, "スコアの全件取得に失敗しました。")
	}
	return scores, nil
}

func (s *scoreUseCase) GetLatestScore(wlID string) (*model.Score, error) {
	score, err := s.ScoreRepository.FIndLatestScoreByWordListID(wlID)
	if err != nil {
		return nil, errors.Wrap(err, "最新のスコア取得に失敗しました。")
	}
	return score, nil
}

func (s *scoreUseCase) PostScore(score model.Score) (*model.Score, error) {
	createdS, err := s.ScoreRepository.CreateScore(score)
	if err != nil {
		return nil, errors.Wrap(err, "スコアの作成に失敗しました。")
	}
	return createdS, nil
}

func (s *scoreUseCase) DeleteAllScore(wlID string) error {
	if err := s.ScoreRepository.RemoveAllScoreByWordListID(wlID); err != nil {
		return errors.Wrap(err, "スコアの一括削除に失敗しました。")
	}
	return nil
}
