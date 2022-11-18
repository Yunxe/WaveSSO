package main

import (
	"Wave/database"
	"Wave/migrator"
	"Wave/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	database.GetConnection()
	migrator.Migrate()
	router.NewRouter()

}
