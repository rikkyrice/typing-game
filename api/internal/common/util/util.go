package util

import (
	"github.com/google/uuid"
)

// GenerateUUID UUID(v4)文字列の生成
func GenerateUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	uu := u.String()
	return uu, nil
}
