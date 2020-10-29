package apierror

import "errors"

// Error APIエラー
type Error struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

// NewError errorの定義
func NewError(statuscode int, originalError error) *Error {
	if originalError == nil {
		originalError = errors.New("")
	}
	return &Error{
		StatusCode: statuscode,
		Message:    originalError.Error(),
	}
}

func (e *Error) Error() string {
	return e.Message
}
