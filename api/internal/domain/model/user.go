package model

import "time"

// User defines user table
type User struct {
	// ID ユーザーが登録時に一意の値を入力する
	ID string `json:"id" validate:"required"`
	// Mail ユーザーメールアドレス
	Mail string `json:"mail" validate:"required,email"`
	// Password ユーザーパスワード
	Password string `json:"password" validate:"required"`
	// CreatedAt 作成された日付
	CreatedAt time.Time `json:"createdAt" validate:"required"`
}
