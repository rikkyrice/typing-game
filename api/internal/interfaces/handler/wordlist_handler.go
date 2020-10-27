package handler

import (
	"api/internal/application/usecase"
	"api/internal/common/apierror"
	"api/internal/domain/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// WordListHandler 単語帳ハンドラインターフェース
type WordListHandler interface {
	GETWordList() echo.HandlerFunc
	// POSTWordList() echo.HandlerFunc
	// PUTWordList() echo.HandlerFunc
	// DELETEWordList() echo.HandlerFunc
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

type WordListsResponse struct {
	Matched   int               `json:"matched"`
	WordLists []*model.WordList `json:"wordlists"`
}

func (wl *wordlistHandler) GETWordList() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := usecase.Authenticate(c); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		} else {
			fmt.Println("認証OK")
		}
		headerParams, err := getHeaderParams(c)
		if err != nil {
			c.Echo().Logger.Errorf("ヘッダーの読み込みに失敗しました。%+v", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		wordlists, err := wl.WordListUseCase.GetWordList(headerParams.UserID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, apierror.NewError(http.StatusInternalServerError, err))
		}
		res := &WordListsResponse{
			Matched:   len(wordlists),
			WordLists: wordlists,
		}
		return c.JSON(http.StatusOK, res)
	}
}

// func (wl *wordlistHandler) POSTWordList() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		headerParams, err := getHeaderParams(c)
// 		if err != nil {
// 			c.Echo().Logger.Errorf("ヘッダーの読み込みに失敗しました。%+v", err)
// 			return c.JSON(http.StatusBadRequest, err)
// 		}

// 	}
// }

// func (wl *wordlistHandler) PUTWordList() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		headerParams, err := getHeaderParams(c)
// 		if err != nil {
// 			c.Echo().Logger.Errorf("ヘッダーの読み込みに失敗しました。%+v", err)
// 			return c.JSON(http.StatusBadRequest, err)
// 		}

// 	}
// }

// func (wl *wordlistHandler) DELETEWordList() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		headerParams, err := getHeaderParams(c)
// 		if err != nil {
// 			c.Echo().Logger.Errorf("ヘッダーの読み込みに失敗しました。%+v", err)
// 			return c.JSON(http.StatusBadRequest, err)
// 		}

// 	}
// }
