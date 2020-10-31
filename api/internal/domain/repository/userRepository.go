package repository

import (
	"api/internal/common/apierror"
	"api/internal/domain/model"
)

// UserRepository ユーザーのリポジトリインターフェース
type UserRepository interface {
	// FindUserByID login用のメソッド
	FindUserByID(userID string) (*model.User, *apierror.Error)
	// CreateUser ユーザー登録用メソッド
	CreateUser(user model.User) (string, *apierror.Error)
}
