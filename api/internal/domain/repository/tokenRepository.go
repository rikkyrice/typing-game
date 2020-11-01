package repository

import (
	"api/internal/common/apierror"
	"api/internal/domain/model"
)

// TokenRepository トークンのリポジトリインターフェース
type TokenRepository interface {
	// FindTokenByUserID トークン取得
	FindLatestTokenByUserID(userID string) (*model.Token, *apierror.Error)
	// StoreToken トークン作成
	StoreToken(t *model.Token) *apierror.Error
	// RemoveTokenByUserID トークン削除
	RemoveTokenByUserID(userID string) *apierror.Error
}
