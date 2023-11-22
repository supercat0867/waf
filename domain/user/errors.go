package user

import "errors"

var (
	ErrMismatchedPassword = errors.New("两次密码不相同")
	ErrUserIsExist        = errors.New("已存在此用户")
	ErrUserNotExist       = errors.New("不存在此用户")
	ErrPasswordNotCorrect = errors.New("密码错误")
)
