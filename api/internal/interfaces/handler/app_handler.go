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

func getHeaderParams(c echo.Context) (*headerRequest, error) {
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

func getPathParams(c echo.Context) (*pathRequest, error) {
	id := c.Param("id")
	if id == "" {
		return nil, apierror.NewError(http.StatusBadRequest, errors.New("パスパラメータに値が含まれていません。"))
	}
	return &pathRequest{
		ID: id,
	}, nil
}
