package usecase

import (
	"api/internal/common/apierror"
	"api/internal/common/util"
	"api/internal/domain/model"
	"api/internal/domain/repository"
	"net/http"

	"github.com/pkg/errors"
)

// ScoreUseCase スコアのサービスインターフェース
type ScoreUseCase interface {
	GetScores(wlID string) ([]*model.Score, *apierror.Error)
	GetLatestScore(wlID string) (*model.Score, *apierror.Error)
	PostScore(score model.Score) (*model.Score, *apierror.Error)
	DeleteAllScore(wlID string) *apierror.Error
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

func (s *scoreUseCase) GetScores(wlID string) ([]*model.Score, *apierror.Error) {
	scores := []*model.Score{}
	scores, err := s.ScoreRepository.FindScoreByWordListID(wlID)
	if err != nil {
		return scores, err
	}
	return scores, nil
}

func (s *scoreUseCase) GetLatestScore(wlID string) (*model.Score, *apierror.Error) {
	score, err := s.ScoreRepository.FIndLatestScoreByWordListID(wlID)
	if err != nil {
		return nil, err
	}
	return score, nil
}

func (s *scoreUseCase) PostScore(score model.Score) (*model.Score, *apierror.Error) {
	id, utilerr := util.GenerateUUID()
	if utilerr != nil {
		return nil, apierror.NewError(http.StatusInternalServerError, errors.Wrap(utilerr, "UUIDの生成に失敗しました。"))
	}
	score.ID = id
	createdS, err := s.ScoreRepository.CreateScore(score)
	if err != nil {
		return nil, err
	}
	return createdS, nil
}

func (s *scoreUseCase) DeleteAllScore(wlID string) *apierror.Error {
	if err := s.ScoreRepository.RemoveAllScoreByWordListID(wlID); err != nil {
		return err
	}
	return nil
}
