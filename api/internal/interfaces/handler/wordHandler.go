package handler

import (
	"api/internal/application/usecase"
	"api/internal/domain/model"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// WordHandler 単語ハンドラインターフェース
type WordHandler interface {
	GETWord() echo.HandlerFunc
	GETWords() echo.HandlerFunc
	POSTWord() echo.HandlerFunc
	POSTWords() echo.HandlerFunc
	PUTWord() echo.HandlerFunc
	DELETEWord() echo.HandlerFunc
	DELETEWords() echo.HandlerFunc
}

// NewWordHandler 単語ハンドラ生成
func NewWordHandler(wUC usecase.WordUseCase) WordHandler {
	return &wordHandler{
		WordUseCase: wUC,
	}
}

type wordHandler struct {
	WordUseCase usecase.WordUseCase
}

type wordResponse struct {
	Word *model.Word `json:"word"`
}

func (w *wordHandler) GETWord() echo.HandlerFunc {
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

		word, err := w.WordUseCase.GetWord(pathParams.ID)
		if err != nil {
			return c.JSON(err.StatusCode, err)
		}
		res := &wordResponse{
			Word: word,
		}
		return c.JSON(http.StatusOK, res)
	}
}

type wordsResponse struct {
	Matched int           `json:"matched"`
	Words   []*model.Word `json:"words"`
}

func (w *wordHandler) GETWords() echo.HandlerFunc {
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

		words, err := w.WordUseCase.GetWordByWordListID(pathParams.ID)
		if err != nil {
			return c.JSON(err.StatusCode, err)
		}
		res := &wordsResponse{
			Matched: len(words),
			Words:   words,
		}
		return c.JSON(http.StatusOK, res)
	}
}

type wordQueryRequest struct {
	WordListID  string    `json:"wordlistID" validate:"required"`
	Word        string    `json:"word" validate:"required"`
	Meaning     string    `json:"meaning" validate:"required"`
	Explanation string    `json:"explanation" validate:"required"`
	CreatedAt   time.Time `json:"createdAt" validate:"required"`
	UpdatedAt   time.Time `json:"updatedAt" validate:"required"`
}

func (wR *wordQueryRequest) toWord() model.Word {
	return model.Word{
		ID:          "",
		WordListID:  wR.WordListID,
		Word:        wR.Word,
		Meaning:     wR.Meaning,
		Explanation: wR.Explanation,
		CreatedAt:   wR.CreatedAt,
		UpdatedAt:   wR.UpdatedAt,
	}
}

func (w *wordHandler) POSTWord() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := usecase.Authenticate(c); err != nil {
			return c.JSON(err.StatusCode, err)
		}
		c.Echo().Logger.Info("認証OK")
		var queryParams wordQueryRequest
		if err := getQueryParams(c, &queryParams); err != nil {
			return c.JSON(err.StatusCode, err)
		}

		word, err := w.WordUseCase.PostWord(queryParams.toWord())
		if err != nil {
			c.Echo().Logger.Errorf("単語の作成に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		res := &wordResponse{
			Word: word,
		}
		return c.JSON(http.StatusCreated, res)
	}
}

type wordsQueryRequest struct {
	Words []*wordQueryRequest `json:"words" validate:"required"`
}

type wordsPostResponse struct {
	Words []*model.Word `json:"words"`
}

func (w *wordHandler) POSTWords() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := usecase.Authenticate(c); err != nil {
			return c.JSON(err.StatusCode, err)
		}
		c.Echo().Logger.Info("認証OK")
		var queryParams wordsQueryRequest
		if err := getQueryParams(c, &queryParams); err != nil {
			return c.JSON(err.StatusCode, err)
		}
		queryWords := []model.Word{}
		for _, word := range queryParams.Words {
			queryWords = append(queryWords, word.toWord())
		}

		words, err := w.WordUseCase.PostAllWord(queryWords)
		if err != nil {
			c.Echo().Logger.Errorf("単語帳の作成に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		res := &wordsPostResponse{
			Words: words,
		}
		return c.JSON(http.StatusCreated, res)
	}
}

func (w *wordHandler) PUTWord() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := usecase.Authenticate(c); err != nil {
			return c.JSON(err.StatusCode, err)
		}
		c.Echo().Logger.Info("認証OK")
		pathParams, err := getPathParams(c)
		if err != nil {
			c.Echo().Logger.Errorf("パスパラメータの読み込みに失敗しました。%+v", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		var queryParams wordQueryRequest
		if err := getQueryParams(c, &queryParams); err != nil {
			return c.JSON(err.StatusCode, err)
		}

		word, err := w.WordUseCase.PutWord(pathParams.ID, queryParams.toWord())
		if err != nil {
			c.Echo().Logger.Errorf("単語帳の更新に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		res := &wordResponse{
			Word: word,
		}
		return c.JSON(http.StatusCreated, res)
	}
}

func (w *wordHandler) DELETEWord() echo.HandlerFunc {
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

		err = w.WordUseCase.DeleteWord(pathParams.ID)
		if err != nil {
			c.Echo().Logger.Errorf("単語帳の削除に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		return c.NoContent(http.StatusNoContent)
	}
}

func (w *wordHandler) DELETEWords() echo.HandlerFunc {
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

		err = w.WordUseCase.DeleteAllWord(pathParams.ID)
		if err != nil {
			c.Echo().Logger.Errorf("単語帳の削除に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		return c.NoContent(http.StatusNoContent)
	}
}
