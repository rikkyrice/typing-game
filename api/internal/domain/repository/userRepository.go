package repository

import (
	"api/internal/domain/model"
)

// UserRepository ユーザーのリポジトリインターフェース
type UserRepository interface {
	// FindUserByID login用のメソッド
	FindUserByID(userID string) (model.User, error)
	// CreateUser ユーザー登録用メソッド
	CreateUser(user model.User) (string, error)
}
