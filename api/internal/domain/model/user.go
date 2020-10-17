package model

import "time"

// User defines user table
type User struct {
	// ID ユーザーが登録時に一意の値を入力する
	ID string
	// Mail ユーザーメールアドレス
	Mail string
	// Password ユーザーパスワード
	Password string
	// CreatedAt 作成された日付
	CreatedAt time.Time
}
