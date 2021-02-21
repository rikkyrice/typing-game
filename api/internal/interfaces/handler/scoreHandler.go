package handler

import (
	"api/internal/application/usecase"
	"api/internal/domain/model"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// ScoreHandler スコアハンドラインターフェース
type ScoreHandler interface {
	GETScores() echo.HandlerFunc
	GETLatestScore() echo.HandlerFunc
	POSTScore() echo.HandlerFunc
	DELETEScores() echo.HandlerFunc
}

// NewScoreHandler スコアハンドラ生成
func NewScoreHandler(sUC usecase.ScoreUseCase) ScoreHandler {
	return &scoreHandler{
		ScoreUseCase: sUC,
	}
}

type scoreHandler struct {
	ScoreUseCase usecase.ScoreUseCase
}

type scoresResponse struct {
	Matched int            `json:"matched"`
	Scores  []*model.Score `json:"scores"`
}

func (s *scoreHandler) GETScores() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := usecase.Authenticate(c); err != nil {
			return c.JSON(err.StatusCode, err)
		}
		c.Echo().Logger.Info("認証OK")
		pathParams, err := getPathParams(c)
		if err != nil {
			c.Echo().Logger.Errorf("パスパラメータの読み込みに失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}

		scores, err := s.ScoreUseCase.GetScores(pathParams.ID)
		if err != nil {
			return c.JSON(err.StatusCode, err)
		}
		res := &scoresResponse{
			Matched: len(scores),
			Scores:  scores,
		}
		return c.JSON(http.StatusOK, res)
	}
}

type scoreResponse struct {
	Score *model.Score `json:"score"`
}

func (s *scoreHandler) GETLatestScore() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := usecase.Authenticate(c); err != nil {
			return c.JSON(err.StatusCode, err)
		}
		c.Echo().Logger.Info("認証OK")
		pathParams, err := getPathParams(c)
		if err != nil {
			c.Echo().Logger.Errorf("パスパラメータの読み込みに失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}

		score, err := s.ScoreUseCase.GetLatestScore(pathParams.ID)
		if err != nil {
			return c.JSON(err.StatusCode, err)
		}
		res := &scoreResponse{
			Score: score,
		}
		return c.JSON(http.StatusOK, res)
	}
}

type scoreQueryRequest struct {
	WordListID     string    `json:"wordlistID" validate:"required"`
	PlayCount      int       `json:"playCount" validate:"required"`
	ClearTypeCount int       `json:"clearTypeCount" validate:"required"`
	MissTypeCount  int       `json:"missTypeCount" validate:"required"`
	PlayedAt       time.Time `json:"playedAt" validate:"required"`
}

func (sR *scoreQueryRequest) toScore() model.Score {
	return model.Score{
		ID:             "",
		WordListID:     sR.WordListID,
		PlayCount:      sR.PlayCount,
		ClearTypeCount: sR.ClearTypeCount,
		MissTypeCount:  sR.MissTypeCount,
		PlayedAt:       sR.PlayedAt,
	}
}

func (s *scoreHandler) POSTScore() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := usecase.Authenticate(c); err != nil {
			return c.JSON(err.StatusCode, err)
		}
		c.Echo().Logger.Info("認証OK")
		var queryParams scoreQueryRequest
		if err := getQueryParams(c, &queryParams); err != nil {
			return c.JSON(err.StatusCode, err)
		}

		score, err := s.ScoreUseCase.PostScore(queryParams.toScore())
		if err != nil {
			c.Echo().Logger.Errorf("スコアの作成に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		res := &scoreResponse{
			Score: score,
		}
		return c.JSON(http.StatusCreated, res)
	}
}

func (s *scoreHandler) DELETEScores() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := usecase.Authenticate(c); err != nil {
			return c.JSON(err.StatusCode, err)
		}
		c.Echo().Logger.Info("認証OK")
		pathParams, err := getPathParams(c)
		if err != nil {
			c.Echo().Logger.Errorf("パスパラメータの読み込みに失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}

		err = s.ScoreUseCase.DeleteAllScore(pathParams.ID)
		if err != nil {
			c.Echo().Logger.Errorf("スコアの削除に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		return c.NoContent(http.StatusNoContent)
	}
}
