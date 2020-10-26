package auth

import (
	"api/internal/common/apierror"
	"api/internal/config"
	"api/internal/domain/model"
	"api/internal/domain/repository"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

const (
	secret = "JTtwVvQXPRFZXFTeMC6AB2VNJ+x70wDc1HpBObmRAZL0HgMHpqklXQ=="

	userIDKey = "user_id"

	iatKey = "iat"
	expKey = "exp"
)

// Auth 認証用の構造体
type Auth struct {
	UserID string
	Iat    int64
}

// AuthUseCase auth用のサービス
type AuthUseCase interface {
	GetUserByID(id string) (model.User, error)
	PostToken(userID string) (*model.Token, error)
}

// NewAuthUseCase 認証ユースケースを生成
func NewAuthUseCase(tR repository.TokenRepository, uR repository.UserRepository) AuthUseCase {
	return &authUseCase{
		TokenRepository: tR,
		UserRepository:  uR,
	}
}

type authUseCase struct {
	TokenRepository repository.TokenRepository
	UserRepository  repository.UserRepository
}

func (a *authUseCase) GetUserByID(id string) (model.User, error) {
	user, err := a.UserRepository.FindUserByID(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (a *authUseCase) PostToken(userID string) (*model.Token, error) {
	c, err := config.Init("config/env.yaml")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みに失敗しました。%+v", err)
	}
	iat := time.Now()
	exp := iat.Add(c.TC.Expired * time.Second)
	tokenString, err := createToken(userID, iat, exp)
	if err != nil {
		return nil, err
	}

	token := &model.Token{
		Token:     tokenString,
		UserID:    userID,
		CreatedAt: iat,
		ExpiredAt: exp,
	}
	if err := a.TokenRepository.StoreToken(token); err != nil {
		return nil, err
	}
	return token, nil
}

func createToken(userID string, i time.Time, e time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		userIDKey: userID,
		iatKey:    i.Unix(),
		expKey:    e.Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// TODO: エラーがちゃんと帰ってない気がするので修正
func parse(signedString string) (*Auth, error) {
	token, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, apierror.NewError(http.StatusUnauthorized, errors.Wrapf(err, "%s is expired", signedString))
			} else {
				return nil, apierror.NewError(http.StatusUnauthorized, errors.Wrapf(err, "%s is invalid", signedString))
			}
		} else {
			return nil, apierror.NewError(http.StatusUnauthorized, errors.Wrapf(err, "%s is invalid", signedString))
		}
	}

	if token == nil {
		return nil, apierror.NewError(http.StatusUnauthorized, errors.Errorf("not found token in %s:", signedString))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, apierror.NewError(http.StatusUnauthorized, errors.Errorf("not found claims in %s", signedString))
	}
	userID, ok := claims[userIDKey].(string)
	if !ok {
		return nil, apierror.NewError(http.StatusUnauthorized, errors.Errorf("not found %s in %s", userIDKey, signedString))
	}
	iat, ok := claims[iatKey].(float64)
	if !ok {
		return nil, apierror.NewError(http.StatusUnauthorized, errors.Errorf("not found %s in %s", iatKey, signedString))
	}

	return &Auth{
		UserID: userID,
		Iat:    int64(iat),
	}, nil
}

// Authenticate 認証を実施
func Authenticate(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return apierror.NewError(http.StatusUnauthorized, errors.New("Authorizationヘッダーに値がありません。"))
	}
	userID := c.Request().Header.Get("X-User-ID")
	if userID == "" {
		return apierror.NewError(http.StatusUnauthorized, errors.New("X-User-IDヘッダーに値がありません。"))
	}

	if err := validateToken(strings.Replace(tokenString, "Bearer ", "", 1), userID); err != nil {
		return err
	}
	return nil
}

func validateToken(tokenString string, userID string) error {
	auth, err := parse(tokenString)
	if err != nil {
		return apierror.NewError(http.StatusUnauthorized, err)
	}
	if userID != auth.UserID {
		return apierror.NewError(http.StatusUnauthorized, errors.Wrapf(err, "failed authentication, userID=%s, userIDPayload=%s", userID, auth.UserID))
	}
	return nil
}
