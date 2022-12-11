package util

import (
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
	// 0 通用成功
	SUCCESS = &Response{HttpCode: http.StatusOK, Status: &Status{Code: 0, Message: "成功"}}
)

var ErrorMap = map[error]*Response{

	// 99 通用错误
	COMMON_ERR: {HttpCode: http.StatusInternalServerError, Status: &Status{Code: 99, Message: "服务器内部错误"}},

	// 100xx 基本常用错误码
	REQ_PARAM_INVALID_ERR: {HttpCode: http.StatusBadRequest, Status: &Status{Code: 10001, Message: "请求参数无效"}},

	// 200xx 鉴定权限相关错误
	AUTH_REQUIRE:              {HttpCode: http.StatusUnauthorized, Status: &Status{Code: 20001, Message: "需要认证"}},
	AUTH_TYPE_ERR:             {HttpCode: http.StatusForbidden, Status: &Status{Code: 20002, Message: "认证类型错误"}},
	AUTH_PARSE_TOKEN_ERR:      {HttpCode: http.StatusInternalServerError, Status: &Status{Code: 20003, Message: "解析令牌错误"}},
	AUTH_TOKEN_EXPIRED:        {HttpCode: http.StatusForbidden, Status: &Status{Code: 20004, Message: "令牌已过期"}},
	AUTH_TOKEN_INVALID_ISSUER: {HttpCode: http.StatusForbidden, Status: &Status{Code: 20005, Message: "令牌签发者无效"}},

	// 310xx SSO业务平台的User模块错误
	USER_NOT_FOUND:          {HttpCode: http.StatusUnauthorized, Status: &Status{Code: 31001, Message: "用户不存在"}},
	USER_EMAIL_EXIST:        {HttpCode: http.StatusOK, Status: &Status{Code: 31002, Message: "邮箱已存在"}},
	USER_LOGIN_PASSWORD_ERR: {HttpCode: http.StatusUnauthorized, Status: &Status{Code: 31003, Message: "密码错误"}},
}
