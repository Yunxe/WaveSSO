package util

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

var (
	COMMON_ERR = errors.New("common_error")

	REQ_PARAM_INVALID_ERR = errors.New("request_param_invalid_error")

	AUTH_REQUIRE                     = errors.New("authorization_required")
	AUTH_TYPE_ERR                    = errors.New("authorization_type_error")
	AUTH_PARSE_TOKEN_ERR             = errors.New("authorization_parse_token_error")
	AUTH_TOKEN_EXPIRED               = jwt.ErrTokenExpired
	AUTH_TOKEN_INVALID_ISSUER        = jwt.ErrTokenInvalidIssuer
	AUTH_TOKEN_INVALID_IN_BLACK_LIST = errors.New("authorization_token_invalid_in_black_list")

	USER_NOT_FOUND          = errors.New("user_not_found")
	USER_EMAIL_EXIST        = errors.New("user_email_exist")
	USER_LOGIN_PASSWORD_ERR = errors.New("user_login_password_error")
)
