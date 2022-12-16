package migrator

import (
	"Wave/database"
	"Wave/model"
	"fmt"
)

func Migrate() {
	migrator := database.DB.Migrator()
	if err := migrator.DropTable("users"); err != nil {
		fmt.Println("drop table error" + err.Error())
		return
	}
	if err := migrator.CreateTable(&model.User{}); err != nil {
		fmt.Println("create table error" + err.Error())
		return
	}
	fmt.Println("migrator running")
}
