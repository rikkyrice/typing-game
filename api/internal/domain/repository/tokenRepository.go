package repository

import (
	"api/internal/domain/model"
	"context"
)

// TokenRepository トークンのリポジトリインターフェース
type TokenRepository interface {
	// FindTokenByUserID トークン取得
	FindTokenByUserID(ctx context.Context, userID string) (model.Token, error)
	// StoreToken トークン作成
	StoreToken(t *model.Token) error
	// RemoveTokenByUserID トークン削除
	RemoveTokenByUserID(ctx context.Context, userID string) error
}
