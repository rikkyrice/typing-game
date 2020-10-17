package repository

import (
	"api/internal/domain/model"
	"context"
)

// UserRepository ユーザーのリポジトリインターフェース
type UserRepository interface {
	// FindUserByID login用のメソッド
	FindUserByID(ctx context.Context, userID string) (*model.User, error)
	// CreateUser ユーザー登録用メソッド
	CreateUser(ctx context.Context, user model.User) (*model.User, error)
}
