package handler

import (
	"api/internal/application/usecase"
	"api/internal/domain/model"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// WordListHandler 単語帳ハンドラインターフェース
type WordListHandler interface {
	GETWordList() echo.HandlerFunc
	POSTWordList() echo.HandlerFunc
	PUTWordList() echo.HandlerFunc
	DELETEWordList() echo.HandlerFunc
}

// NewWordListHandler 単語帳ハンドラ生成
func NewWordListHandler(wlUC usecase.WordListUseCase) WordListHandler {
	return &wordlistHandler{
		WordListUseCase: wlUC,
	}
}

type wordlistHandler struct {
	WordListUseCase usecase.WordListUseCase
}

type wordlistsResponse struct {
	Matched   int               `json:"matched"`
	WordLists []*model.WordList `json:"wordlists"`
}

func (wl *wordlistHandler) GETWordList() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := usecase.Authenticate(c); err != nil {
			return c.JSON(err.StatusCode, err)
		}
		c.Echo().Logger.Info("認証OK")
		headerParams, err := getHeaderParams(c)
		if err != nil {
			c.Echo().Logger.Errorf("ヘッダーの読み込みに失敗しました。%+v", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		wordlists, err := wl.WordListUseCase.GetWordList(headerParams.UserID)
		if err != nil {
			return c.JSON(err.StatusCode, err)
		}
		res := &wordlistsResponse{
			Matched:   len(wordlists),
			WordLists: wordlists,
		}
		return c.JSON(http.StatusOK, res)
	}
}

type wordlistQueryRequest struct {
	UserID      string    `json:"userID" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Explanation string    `json:"explanation" validate:"required"`
	CreatedAt   time.Time `json:"createdAt" validate:"required"`
	UpdatedAt   time.Time `json:"updatedAt" validate:"required"`
}

func (wlR *wordlistQueryRequest) toWordList() model.WordList {
	return model.WordList{
		ID:          "",
		UserID:      wlR.UserID,
		Title:       wlR.Title,
		Explanation: wlR.Explanation,
		CreatedAt:   wlR.CreatedAt,
		UpdatedAt:   wlR.UpdatedAt,
	}
}

type wordlistResponse struct {
	WordList *model.WordList `json:"wordlist"`
}

func (wl *wordlistHandler) POSTWordList() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := usecase.Authenticate(c); err != nil {
			return c.JSON(err.StatusCode, err)
		}
		c.Echo().Logger.Info("認証OK")
		var queryParams wordlistQueryRequest
		if err := getQueryParams(c, &queryParams); err != nil {
			return c.JSON(err.StatusCode, err)
		}

		wordlist, err := wl.WordListUseCase.PostWordList(queryParams.toWordList())
		if err != nil {
			c.Echo().Logger.Errorf("単語帳の作成に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		res := &wordlistResponse{
			WordList: wordlist,
		}
		return c.JSON(http.StatusCreated, res)
	}
}

func (wl *wordlistHandler) PUTWordList() echo.HandlerFunc {
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
		var queryParams wordlistQueryRequest
		if err := getQueryParams(c, &queryParams); err != nil {
			return c.JSON(err.StatusCode, err)
		}

		wordlist, err := wl.WordListUseCase.PutWordList(pathParams.ID, queryParams.toWordList())
		if err != nil {
			c.Echo().Logger.Errorf("単語帳の更新に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		res := &wordlistResponse{
			WordList: wordlist,
		}
		return c.JSON(http.StatusCreated, res)
	}
}

func (wl *wordlistHandler) DELETEWordList() echo.HandlerFunc {
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

		err = wl.WordListUseCase.DeleteWordList(pathParams.ID)
		if err != nil {
			c.Echo().Logger.Errorf("単語帳の削除に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		return c.NoContent(http.StatusNoContent)
	}
}
