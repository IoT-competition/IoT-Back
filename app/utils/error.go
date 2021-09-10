package utils

import "errors"

var (
	ErrBodyCanNotParser = errors.New("请求体解析失败")
	ErrTypeMismatch     = errors.New("查询类型错误")
	ErrParamsMiss       = errors.New("参数不匹配请进行检查")
	ErrInfoGetNull      = errors.New("info Get Null")
	ErrInfoUpdateError  = errors.New("info Update Error")
	ErrSetOptionFailed  = errors.New("set option failed")
	ErrGetOptionFailed  = errors.New("get option failed")
)
