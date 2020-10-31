package usecase

import (
	"api/internal/common/apierror"
	"api/internal/domain/model"
	"api/internal/domain/repository"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// UserUseCase ユーザーのサービスインターフェース
type UserUseCase interface {
	Signup(user model.User) (*model.Token, *apierror.Error)
	Login(userID string, password string) (*model.Token, *apierror.Error)
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

func (u *userUseCase) Signup(user model.User) (*model.Token, *apierror.Error) {
	hashedPassword, internalErr := cryptPassword(user.Password)
	if internalErr != nil {
		return nil, apierror.NewError(http.StatusInternalServerError, errors.Wrap(internalErr, "パスワードの暗号化に失敗しました。"))
	}

	user.Password = string(hashedPassword)

	userID, err := u.UserRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	token, err := u.AuthUseCase.PostToken(userID)
	if err != nil {
		return nil, apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "トークン発行に失敗しました。"))
	}

	return token, nil
}

func (u *userUseCase) Login(userID string, password string) (*model.Token, *apierror.Error) {
	user, err := u.UserRepository.FindUserByID(userID)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, apierror.NewError(http.StatusBadRequest, errors.Wrap(err, "パスワードが違います。"))
	}

	token, err := u.AuthUseCase.PostToken(user.ID)
	if err != nil {
		return nil, apierror.NewError(http.StatusInternalServerError, errors.Wrap(err, "トークン発行に失敗しました。"))
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
