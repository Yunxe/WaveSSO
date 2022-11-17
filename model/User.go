package model

import (
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
