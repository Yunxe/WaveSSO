package util

import (
	"errors"
	"net/http"
)

type Status struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type StatusWithData struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Response struct {
	HttpCode int
	Status   *Status
}

var (
	COMMON_ERR              = errors.New("common_error")
	REQ_PARAM_INVALID_ERR   = errors.New("request_param_invalid_error")
	USER_NOT_FOUND          = errors.New("user_not_found")
	USER_EMAIL_EXIST        = errors.New("user_email_exist")
	USER_LOGIN_PASSWORD_ERR = errors.New("user_login_password_error")
)

var ErrorMap = map[error]*Response{
	// 0 通用成功

	// 99 通用错误
	COMMON_ERR: {HttpCode: http.StatusInternalServerError, Status: &Status{Code: 99, Message: "服务器内部错误"}},

	// 100xx 基本常用错误码
	REQ_PARAM_INVALID_ERR: {HttpCode: http.StatusBadRequest, Status: &Status{Code: 10002, Message: "请求参数无效"}},

	// 200xx 鉴定权限相关错误

	// 310xx SSO业务平台的User模块错误
	USER_NOT_FOUND:          {HttpCode: http.StatusUnauthorized, Status: &Status{Code: 31001, Message: "用户不存在"}},
	USER_EMAIL_EXIST:        {HttpCode: http.StatusOK, Status: &Status{Code: 31002, Message: "邮箱已存在"}},
	USER_LOGIN_PASSWORD_ERR: {HttpCode: http.StatusUnauthorized, Status: &Status{Code: 31003, Message: "密码错误"}},
}
