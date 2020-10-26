package user

import (
	"api/internal/application/usecase/auth"
	"api/internal/domain/model"
	"api/internal/domain/repository"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// UserSignupUseCase ユーザー登録用のサービスインターフェース
type UserSignupUseCase interface {
	Signup(user model.User) (*model.Token, error)
}

// NewUserSignupUseCase ユーザー登録用サービス生成
func NewUserSignupUseCase(uR repository.UserRepository, authUC auth.AuthUseCase) UserSignupUseCase {
	return &userSignupUseCase{
		UserRepository: uR,
		AuthUseCase:    authUC,
	}
}

type userSignupUseCase struct {
	UserRepository repository.UserRepository
	AuthUseCase    auth.AuthUseCase
}

func (u *userSignupUseCase) Signup(user model.User) (*model.Token, error) {
	hashedPassword, err := cryptPassword(user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "パスワードの暗号化に失敗しました。")
	}

	user.Password = string(hashedPassword)

	userID, err := u.UserRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	token, err := u.AuthUseCase.PostToken(userID)
	if err != nil {
		return nil, errors.Wrap(err, "トークン発行に失敗しました。")
	}

	return token, nil
}

func cryptPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
