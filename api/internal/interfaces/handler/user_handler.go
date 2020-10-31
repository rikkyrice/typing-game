package handler

import (
	"net/http"
	"time"

	"api/internal/application/usecase"
	"api/internal/domain/model"

	"github.com/labstack/echo"
)

// UserHandler ユーザーを操作するAPI
// /user
type UserHandler interface {
	// /signup
	Signup() echo.HandlerFunc

	// /login
	Login() echo.HandlerFunc
}

// NewUserHandler ユーザーハンドラー定義
func NewUserHandler(userUC usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUC: userUC,
	}
}

type userHandler struct {
	userUC usecase.UserUseCase
}

type signupQueryRequest struct {
	ID        string    `json:"id" validate:"required"`
	Mail      string    `json:"mail" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
}

func (sR *signupQueryRequest) toUser() model.User {
	return model.User{
		ID:        sR.ID,
		Mail:      sR.Mail,
		Password:  sR.Password,
		CreatedAt: sR.CreatedAt,
	}
}

type tokenResponse struct {
	Token string `json:"token"`
}

// SignUp /signup
func (u *userHandler) Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		var queryParams signupQueryRequest
		if err := getQueryParams(c, &queryParams); err != nil {
			return c.JSON(err.StatusCode, err)
		}

		token, err := u.userUC.Signup(queryParams.toUser())
		if err != nil {
			c.Echo().Logger.Errorf("ユーザー登録に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		res := &tokenResponse{
			Token: token.Token,
		}
		return c.JSON(http.StatusCreated, res)
	}
}

type loginQueryRequest struct {
	Password string `json:"password" validate:"required"`
}

// Login /login
func (u *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		headerParams, err := getHeaderParams(c)
		if err != nil {
			c.Echo().Logger.Errorf("ヘッダーの読み込みに失敗しました。%+v", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		var queryParams loginQueryRequest
		if err := getQueryParams(c, &queryParams); err != nil {
			return c.JSON(err.StatusCode, err)
		}

		token, err := u.userUC.Login(headerParams.UserID, queryParams.Password)
		if err != nil {
			c.Echo().Logger.Errorf("ログインに失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		res := &tokenResponse{
			Token: token.Token,
		}
		return c.JSON(http.StatusOK, res)
	}
}
