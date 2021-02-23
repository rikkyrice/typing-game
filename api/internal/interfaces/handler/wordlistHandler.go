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

type wordlistResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"word_list_title"`
	Explanation string    `json:"explanation"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func toWordlistResponse(wordlist *model.WordList) *wordlistResponse {
	return &wordlistResponse{
		ID:          wordlist.ID,
		Title:       wordlist.Title,
		Explanation: wordlist.Explanation,
		CreatedAt:   wordlist.CreatedAt,
		UpdatedAt:   wordlist.UpdatedAt,
	}
}

type wordlistsResponse struct {
	Matched   int                 `json:"matched"`
	WordLists []*wordlistResponse `json:"wordlists"`
}

func (wl *wordlistHandler) GETWordList() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := usecase.Authenticate(c)
		if err != nil {
			return c.JSON(err.StatusCode, err)
		}
		c.Echo().Logger.Info("認証OK")

		wordlists, err := wl.WordListUseCase.GetWordList(userID)
		if err != nil {
			return c.JSON(err.StatusCode, err)
		}
		wls := []*wordlistResponse{}
		for _, wordlist := range wordlists {
			wls = append(wls, toWordlistResponse(wordlist))
		}
		res := &wordlistsResponse{
			Matched:   len(wordlists),
			WordLists: wls,
		}
		return c.JSON(http.StatusOK, res)
	}
}

type wordlistQueryRequest struct {
	Title       string `json:"word_list_title" validate:"required"`
	Explanation string `json:"explanation" validate:"required"`
	CreatedAt   time.Time
}

func (wlR *wordlistQueryRequest) toWordList(userID string) model.WordList {
	return model.WordList{
		ID:          "",
		UserID:      userID,
		Title:       wlR.Title,
		Explanation: wlR.Explanation,
		CreatedAt:   wlR.CreatedAt,
		UpdatedAt:   time.Now(),
	}
}

func (wl *wordlistHandler) POSTWordList() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := usecase.Authenticate(c)
		if err != nil {
			return c.JSON(err.StatusCode, err)
		}
		c.Echo().Logger.Info("認証OK")
		var queryParams wordlistQueryRequest
		if err := getQueryParams(c, &queryParams); err != nil {
			return c.JSON(err.StatusCode, err)
		}
		queryParams.CreatedAt = time.Now()

		wordlist, err := wl.WordListUseCase.PostWordList(queryParams.toWordList(userID))
		if err != nil {
			c.Echo().Logger.Errorf("単語帳の作成に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		res := toWordlistResponse(wordlist)
		return c.JSON(http.StatusCreated, res)
	}
}

func (wl *wordlistHandler) PUTWordList() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := usecase.Authenticate(c)
		if err != nil {
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

		wordlist, err := wl.WordListUseCase.PutWordList(pathParams.ID, queryParams.toWordList(userID))
		if err != nil {
			c.Echo().Logger.Errorf("単語帳の更新に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		res := toWordlistResponse(wordlist)
		return c.JSON(http.StatusCreated, res)
	}
}

func (wl *wordlistHandler) DELETEWordList() echo.HandlerFunc {
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

		err = wl.WordListUseCase.DeleteWordList(pathParams.ID)
		if err != nil {
			c.Echo().Logger.Errorf("単語帳の削除に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		return c.NoContent(http.StatusNoContent)
	}
}
