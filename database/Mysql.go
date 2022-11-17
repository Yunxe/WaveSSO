package database

import (
	"Wave/model"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func GetConnection() {
	if err:=godotenv.Load();err!=nil{
		panic("Error loading .env file")
	}
	DSN:=os.Getenv("DSN")
	db,err := gorm.Open(mysql.New(mysql.Config{
		DSN: DSN,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection error")
		return
	}
	//db.AutoMigrate(&model.User{})
	migrator:=db.Migrator()
	if err:=migrator.DropTable("users");err!=nil{
		fmt.Println("drop table error"+err.Error())
		return
	}
	if err := migrator.CreateTable(&model.User{});err!=nil{
		fmt.Println("create table error"+err.Error())
		return
	}


	DB = db
}
