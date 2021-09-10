package utils

import "gorm.io/gorm"

type returnError struct {
	code string
	msg  string
}

var errorMap = map[error]returnError{
	ErrBodyCanNotParser: {
		code: "1001",
		msg:  "Request Body 解析失败",
	},
	ErrParamsMiss: {
		code: "1002",
		msg:  "参数错误，请检查",
	},
	gorm.ErrRecordNotFound: {
		code: "1011",
		msg:  "记录未找到",
	},
	ErrTypeMismatch: {
		code: "1021",
		msg:  "服务转型失败",
	},
	ErrInfoGetNull: {
		code: "1031",
		msg:  "info get failed",
	},
	ErrInfoUpdateError: {
		code: "1041",
		msg:  "info update error",
	},
	ErrSetOptionFailed: {
		code: "1051",
		msg:  "option set failed",
	},
	ErrGetOptionFailed: {
		code: "1061",
		msg:  "option get failed",
	},
}
