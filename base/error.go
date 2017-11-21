package base

import (
	"errors"
)

type PiccError struct{
	Err error
	Code int
	Desc string
}

func (e *PiccError) Error() string{ 
	return e.Err.Error()
}

func New(err error, code int, desc string) *PiccError{ 
	return &PiccError{
		Err:err,
		Code:code,
		Desc:desc,
	}
}

const(
	ERR_DB_BEGIN_TRANSCATION_CODE int = 100
	ERR_DB_BEGIN_TRANSCATION_DESC string = "[DB]开启事务失败"

	ERR_DB_COMMIT_TRANSCATION_CODE int = 101
	ERR_DB_COMMIT_TRANSCATION_DESC string = "[DB]提交事务失败"

	ERR_USER_ID_OR_TEL_EXIST_CODE int = 200
	ERR_USER_ID_OR_TEL_EXIST_DESC string = "用户身份证号或手机号码已被占用"

	ERR_USER_ADD_CODE int = 201
	ERR_USER_ADD_DESC string = "添加用户失败"

	ERR_USER_ROLE_ADD_CODE int = 202
	ERR_USER_ROLE_ADD_DESC string = "添加角色失败"

	ERR_INSURANCE_SUBJECT_ADD_NO_ATTR_CODE int = 301
	ERR_INSURANCE_SUBJECT_ADD_NO_ATTR_DESC string = "没有属性，请为标的添加属性"

	ERR_INSURANCE_SUBJECT_ADD_EXIST_CODE int = 302
	ERR_INSURANCE_SUBJECT_ADD_EXIST_DESC string = "标的已经存在"

	ERR_INSURANCE_SUBJECT_ADD_CODE int = 303
	ERR_INSURANCE_SUBJECT_ADD_DESC string = "添加标的失败"

)


var (
	ErrLoginUserNotFound = errors.New("用户不存在")
	ErrLoginPasswdError  = errors.New("密码错误")
)
