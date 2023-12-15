package api

import "errors"

// ErrorResponse 响应错误结构体
type ErrorResponse struct {
	Message string `json:"msg"`
}

var (
	ErrInvalidBody           = errors.New("invalid body")
	ErrInternalServer        = errors.New("internal server error")
	ErrStatusForbidden       = errors.New("访问被拒绝")
	ErrStatusTooManyRequests = errors.New("请求太频繁")
	ErrJwtGenerate           = errors.New("JWT生成错误")
)
