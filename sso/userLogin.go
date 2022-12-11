package sso

import (
	"Wave/database"
	"Wave/model"
	"Wave/util"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserLoginInfo struct {
	Email    string `json:"email" form:"email" binding:"email,required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func UserLogin(c *gin.Context) (err error, data any) {
	var (
		u    UserLoginInfo
		user model.User
	)
	if err := c.ShouldBind(&u); err != nil {
		return util.REQ_PARAM_INVALID_ERR, nil
	}

	database.DB.Where("email = ?", u.Email).First(&user)

	if user == *model.NewUser() {
		return util.USER_NOT_FOUND, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		return util.USER_LOGIN_PASSWORD_ERR, nil
	}

	userClaims := &util.UserClaims{
		Uid:    user.Uid,
		Role:   user.Role,
		Status: user.Status,
	}

	token, err := userClaims.CreateToken()
	if err != nil {
		return err, nil
	}

	return nil, &util.StatusWithData{
		Code:    0,
		Message: "成功",
		Data: &util.TokenInfo{
			Token:     token,
			TokenType: "bearer",
			ExpiresIn: time.Time{}.Second() * 10,
		},
	}
}
