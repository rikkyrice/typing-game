package usecase

import (
	"api/internal/domain/model"
	"api/internal/domain/repository"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// UserUseCase ユーザーのサービスインターフェース
type UserUseCase interface {
	Signup(user model.User) (*model.Token, error)
	Login(userID string, password string) (*model.Token, error)
}

// NewUserUseCase ユーザー用サービス生成
func NewUserUseCase(uR repository.UserRepository, authUC AuthUseCase) UserUseCase {
	return &userUseCase{
		UserRepository: uR,
		AuthUseCase:    authUC,
	}
}

type userUseCase struct {
	UserRepository repository.UserRepository
	AuthUseCase    AuthUseCase
}

func (u *userUseCase) Signup(user model.User) (*model.Token, error) {
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

func (u *userUseCase) Login(userID string, password string) (*model.Token, error) {
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

func cryptPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
