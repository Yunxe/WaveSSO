package service

//type user struct {
//	model.User
//}

type UserRegisterInfo struct {
	UserName string `json:"userName" form:"userName" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Email    string `json:"email" form:"email" binding:"email,required"`
}

