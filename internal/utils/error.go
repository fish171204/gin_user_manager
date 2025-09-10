package utils

import (
	"net/http"
)

type ErrorCode string

const (
	ErrCodeBadeRequest ErrorCode = "BAD_REQUEST"
	ErrCodeNotFound    ErrorCode = "NOT_FOUND"
	ErrCodeConflict    ErrorCode = "CONFLICT"
	ErrCodeInternal    ErrorCode = "INTERNAL_SERVER_ERR"
)

type AppError struct {
	Message string
	Code    ErrorCode
	Err     error
}

func (ae *AppError) Error() string {
	return ""
}

func NewError(message string, code ErrorCode) error {
	return &AppError{
		Message: message,
		Code:    code,
	}
}

func WrapError(message string, code ErrorCode, err error) error {
	return &AppError{
		Message: message,
		Code:    code,
		Err:     err,
	}
}

func httpStatusFromCode(code ErrorCode) int {
	switch code {
	case ErrCodeBadeRequest:
		return http.StatusBadRequest
	case ErrCodeNotFound:
		return http.StatusNotFound
	case ErrCodeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
