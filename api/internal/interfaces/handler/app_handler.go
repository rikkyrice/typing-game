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
