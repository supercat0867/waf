package api

import "errors"

// ErrorResponse 响应错误结构体
type ErrorResponse struct {
	Message string `json:"msg"`
}

var (
	ErrInvalidBody = errors.New("参数不合法")
)
