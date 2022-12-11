package main

import (
	"Wave/database"
	"Wave/router"
	"context"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	database.MysqlConnection(ctx)
	database.RedisConnection(ctx)
	// time.Sleep(time.Second * 1)
	//migrator.Migrate()
	router.NewRouter()

}
