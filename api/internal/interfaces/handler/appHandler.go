package handler

import (
	"api/internal/common/apierror"
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

// AppHandler すべてのハンドラを内包するインターフェース
type AppHandler interface {
	HealthCheckHandler
	UserHandler
	WordListHandler
	WordHandler
}

type headerRequest struct {
	UserID string
}

func getHeaderParams(c echo.Context) (*headerRequest, *apierror.Error) {
	userID := c.Request().Header.Get("X-User-ID")
	if userID == "" {
		return nil, apierror.NewError(http.StatusBadRequest, errors.New("ヘッダーにユーザーIDが含まれていません。"))
	}
	return &headerRequest{
		UserID: userID,
	}, nil
}

type pathRequest struct {
	ID string
}

func getPathParams(c echo.Context) (*pathRequest, *apierror.Error) {
	id := c.Param("id")
	if id == "" {
		return nil, apierror.NewError(http.StatusBadRequest, errors.New("パスパラメータに値が含まれていません。"))
	}
	return &pathRequest{
		ID: id,
	}, nil
}

type queryRequest interface{}

func getQueryParams(c echo.Context, req queryRequest) *apierror.Error {
	if err := c.Bind(req); err != nil {
		c.Echo().Logger.Errorf("リクエストボディの読み込みに失敗しました。%+v", err)
		return apierror.NewError(http.StatusBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return apierror.NewError(http.StatusBadRequest, err)
	}
	return nil
}
