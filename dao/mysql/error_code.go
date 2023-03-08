package mysql

import "errors"

var(
	ErrorUserExist = errors.New("用户已存在")
	ErrorUserNotExist = errors.New("用户不存在")
	ErrorInvalidPassword= errors.New("密码错误")
	ErrorUserNotLogin= errors.New("用户未登录")
	ErrorInvalidId = errors.New("无效的ID")

)