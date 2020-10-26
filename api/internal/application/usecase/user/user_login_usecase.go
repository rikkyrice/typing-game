package user

import (
	"api/internal/application/usecase/auth"
	"api/internal/domain/model"
	"api/internal/domain/repository"

	"golang.org/x/crypto/bcrypt"

	"github.com/pkg/errors"
)

// UserLoginUseCase ログイン用のサービス
type UserLoginUseCase interface {
	Login(userID string, password string) (*model.Token, error)
}

// NewUserLoginUseCase ログイン用サービス
func NewUserLoginUseCase(uR repository.UserRepository, authUC auth.AuthUseCase) UserLoginUseCase {
	return &userLoginUseCase{
		UserRepository: uR,
		AuthUseCase:    authUC,
	}
}

type userLoginUseCase struct {
	UserRepository repository.UserRepository
	AuthUseCase    auth.AuthUseCase
}

func (u *userLoginUseCase) Login(userID string, password string) (*model.Token, error) {
	user, err := u.UserRepository.FindUserByID(userID)
	if err != nil {
		return nil, errors.Wrap(err, "ユーザーが存在しません。")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.Wrap(err, "パスワードが違います。")
	}

	token, err := u.AuthUseCase.PostToken(user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "トークン発行に失敗しました。")
	}

	return token, nil
}
