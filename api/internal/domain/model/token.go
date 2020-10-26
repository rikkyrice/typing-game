package model

import "time"

// Token アクセストークン
type Token struct {
	// Token アクセス用トークン文字列
	Token string `json:"token"`
	// UserID どのユーザーのアクセストークンかを保持
	UserID string `json:"userID"`
	// CreatedAt トークン作成日時
	CreatedAt time.Time `json:"createdAt"`
	// ExpiredAt トークン有効期限
	ExpiredAt time.Time `json:"expiredAt"`
}
