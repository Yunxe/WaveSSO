package database

import (
	"context"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MysqlConnection(ctx context.Context) {
	DSN := os.Getenv("DSN")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: DSN,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection error")
		return
	}
	DB = db
	fmt.Println("mysql pong")
}
