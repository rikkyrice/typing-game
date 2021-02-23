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
	ID           string    `json:"id"`
	Word         string    `json:"word"`
	Yomi         string    `json:"yomi"`
	Meaning      string    `json:"meaning"`
	MYomi        string    `json:"m_yomi"`
	Explanation  string    `json:"explanation"`
	IsRemembered bool      `json:"is_remembered"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func toWordResponse(word *model.Word) *wordResponse {
	return &wordResponse{
		ID:           word.ID,
		Word:         word.Word,
		Yomi:         word.Yomi,
		Meaning:      word.Meaning,
		MYomi:        word.MYomi,
		Explanation:  word.Explanation,
		IsRemembered: word.IsRemembered,
		CreatedAt:    word.CreatedAt,
		UpdatedAt:    word.UpdatedAt,
	}
}

func (w *wordHandler) GETWord() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := usecase.Authenticate(c)
		if err != nil {
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
		res := toWordResponse(word)
		return c.JSON(http.StatusOK, res)
	}
}

type wordsResponse struct {
	Matched int             `json:"matched"`
	Words   []*wordResponse `json:"words"`
}

func (w *wordHandler) GETWords() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := usecase.Authenticate(c)
		if err != nil {
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
		ws := []*wordResponse{}
		for _, word := range words {
			ws = append(ws, toWordResponse(word))
		}
		res := &wordsResponse{
			Matched: len(words),
			Words:   ws,
		}
		return c.JSON(http.StatusOK, res)
	}
}

type wordQueryRequest struct {
	WordListID   string `json:"word_list_id" validate:"required"`
	Word         string `json:"word" validate:"required"`
	Meaning      string `json:"meaning" validate:"required"`
	Explanation  string `json:"explanation" validate:"required"`
	IsRemembered bool   `json:"is_remembered" validate:"required"`
	CreatedAt    time.Time
}

func (wR *wordQueryRequest) toWord() model.Word {
	return model.Word{
		ID:           "",
		WordListID:   wR.WordListID,
		Word:         wR.Word,
		Yomi:         "",
		Meaning:      wR.Meaning,
		MYomi:        "",
		Explanation:  wR.Explanation,
		IsRemembered: wR.IsRemembered,
		CreatedAt:    wR.CreatedAt,
		UpdatedAt:    time.Now(),
	}
}

func (w *wordHandler) POSTWord() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := usecase.Authenticate(c)
		if err != nil {
			return c.JSON(err.StatusCode, err)
		}
		c.Echo().Logger.Info("認証OK")
		var queryParams wordQueryRequest
		if err := getQueryParams(c, &queryParams); err != nil {
			return c.JSON(err.StatusCode, err)
		}
		queryParams.CreatedAt = time.Now()

		word, err := w.WordUseCase.PostWord(queryParams.toWord())
		if err != nil {
			c.Echo().Logger.Errorf("単語の作成に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		res := toWordResponse(word)
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
		_, err := usecase.Authenticate(c)
		if err != nil {
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
		_, err := usecase.Authenticate(c)
		if err != nil {
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
		res := toWordResponse(word)
		return c.JSON(http.StatusCreated, res)
	}
}

func (w *wordHandler) DELETEWord() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := usecase.Authenticate(c)
		if err != nil {
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
		_, err := usecase.Authenticate(c)
		if err != nil {
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
