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
	UserID    string `json:"user_id" validate:"required"`
	Mail      string `json:"mail" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	CreatedAt time.Time
}

func (sR *signupQueryRequest) toUser() model.User {
	return model.User{
		ID:        sR.UserID,
		Mail:      sR.Mail,
		Password:  sR.Password,
		CreatedAt: sR.CreatedAt,
	}
}

type tokenResponse struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

// SignUp /signup
func (u *userHandler) Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		var queryParams signupQueryRequest
		if err := getQueryParams(c, &queryParams); err != nil {
			return c.JSON(err.StatusCode, err)
		}
		queryParams.CreatedAt = time.Now()

		token, err := u.userUC.Signup(queryParams.toUser())
		if err != nil {
			c.Echo().Logger.Errorf("ユーザー登録に失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		res := &tokenResponse{
			Token:  token.Token,
			UserID: queryParams.UserID,
		}
		return c.JSON(http.StatusCreated, res)
	}
}

type loginQueryRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Login /login
func (u *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var queryParams loginQueryRequest
		if err := getQueryParams(c, &queryParams); err != nil {
			return c.JSON(err.StatusCode, err)
		}

		token, err := u.userUC.Login(queryParams.UserID, queryParams.Password)
		if err != nil {
			c.Echo().Logger.Errorf("ログインに失敗しました。%+v", err)
			return c.JSON(err.StatusCode, err)
		}
		res := &tokenResponse{
			Token:  token.Token,
			UserID: queryParams.UserID,
		}
		return c.JSON(http.StatusOK, res)
	}
}
