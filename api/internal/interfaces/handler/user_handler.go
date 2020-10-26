package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"api/internal/application/usecase/auth"
	"api/internal/application/usecase/user"
	"api/internal/common/apierror"
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

type userHandler struct {
	userSignupUC user.UserSignupUseCase
	userLoginUC  user.UserLoginUseCase
}

// NewUserHandler ユーザーハンドラー定義
func NewUserHandler(userSignupUC user.UserSignupUseCase,
	userLoginUC user.UserLoginUseCase) UserHandler {
	return &userHandler{
		userSignupUC: userSignupUC,
		userLoginUC:  userLoginUC,
	}
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
		if err := c.Bind(&queryParams); err != nil {
			c.Echo().Logger.Errorf("リクエストボディの読み込みに失敗しました。%+v", err)
			return c.JSON(http.StatusBadRequest, apierror.NewError(http.StatusBadRequest, err))
		}
		if err := c.Validate(&queryParams); err != nil {
			return c.JSON(http.StatusBadRequest, apierror.NewError(http.StatusBadRequest, err))
		}

		token, err := u.userSignupUC.Signup(queryParams.toUser())
		if err != nil {
			c.Echo().Logger.Errorf("ユーザー登録に失敗しました。%+v", err)
			return c.JSON(http.StatusInternalServerError, apierror.NewError(http.StatusInternalServerError, err))
		}
		res := &tokenResponse{
			Token: token.Token,
		}
		return c.JSON(http.StatusOK, res)
	}
}

type loginHeaderRequest struct {
	UserID string
}

func getHeaderParams(c echo.Context) (*loginHeaderRequest, error) {
	userID := c.Request().Header.Get("X-User-ID")
	if userID == "" {
		return nil, apierror.NewError(http.StatusBadRequest, errors.New("ヘッダーにユーザーIDが含まれていません。"))
	}
	return &loginHeaderRequest{
		UserID: userID,
	}, nil
}

type loginQueryRequest struct {
	Password string `json:"password" validate:"required"`
}

// Login /login
func (u *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		headerParams, err := getHeaderParams(c)
		if err != nil {
			c.Echo().Logger.Errorf("リクエストボディの読み込みに失敗しました。%+v", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		var queryParams loginQueryRequest
		if err := c.Bind(&queryParams); err != nil {
			c.Echo().Logger.Errorf("リクエストボディの読み込みに失敗しました。%+v", err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		if err := c.Validate(&queryParams); err != nil {
			return c.JSON(http.StatusBadRequest, apierror.NewError(http.StatusBadRequest, err))
		}

		// すでにjwtトークンがあれば上書き、あ。別にしなくていいかも
		if err := auth.Authenticate(c); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("認証OK")
		}

		token, err := u.userLoginUC.Login(headerParams.UserID, queryParams.Password)
		if err != nil {
			c.Echo().Logger.Errorf("ログインに失敗しました。%+v", err)
			return c.JSON(http.StatusInternalServerError, apierror.NewError(http.StatusInternalServerError, err))
		}
		res := &tokenResponse{
			Token: token.Token,
		}
		return c.JSON(http.StatusOK, res)
	}
}
