package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Uid       uint           `json:"uid" gorm:"primaryKey" `
	UserName  string         `json:"userName" `
	Password  string         `json:"password"`
	Email     string         `json:"email"`
	Role      int            `json:"role"`
	Gender    int            `json:"gender"`
	Status    int            `json:"status"`
	AvatarUrl string         `json:"avatarUrl"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt ` json:"deletedAt" gorm:"index"`
}

var NilUser = &User{}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var user User
	tx.Where("email = ?", u.Email).First(&user)
	if user != *NilUser {
		return errors.New("邮箱已存在")
	}
	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.Uid == 1 {
		tx.Model(u).Update("role", 2)
	}
	return
}
