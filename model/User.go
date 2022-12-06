package model

import (
	"Wave/util"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Uid       uint           `json:"uid" gorm:"primaryKey" `
	UserName  string         `json:"userName" `
	Password  string         `json:"password"`
	Email     string         `json:"email"`
	Role      int8           `json:"role"`
	Gender    int8           `json:"gender"`
	Status    int8           `json:"status"`
	AvatarUrl string         `json:"avatarUrl"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt ` json:"deletedAt" gorm:"index"`
}

//var NilUser = &User{}

func NewUser() *User {
	return &User{}
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	user := NewUser()
	tx.Where("email = ?", u.Email).First(&user)
	if *user != *NewUser() {
		return util.USER_EMAIL_EXIST
	}
	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.Uid == 1 {
		tx.Model(u).Update("role", 2)
	}
	return
}
