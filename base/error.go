package base

import (
	"errors"
)

var (
	ErrLoginUserNotFound = errors.New("用户不存在")
	ErrLoginPasswdError  = errors.New("密码错误")
)
