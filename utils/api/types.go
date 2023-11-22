package api

import "errors"

// ErrorResponse 响应错误结构体
type ErrorResponse struct {
	Message string `json:"msg"`
}

var (
	ErrInvalidBody           = errors.New("invalid body error")
	ErrInternalServer        = errors.New("internal server error")
	ErrStatusForbidden       = errors.New("access denied")
	ErrStatusTooManyRequests = errors.New("Too Many Request")
)
