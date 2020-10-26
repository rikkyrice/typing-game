package validation

import (
	"gopkg.in/go-playground/validator.v9"
)

// CustomValidator カスタムバリデーター
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate バリデートを実施
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
