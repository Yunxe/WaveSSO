package main

import (
	"Wave/database"
	"Wave/migrator"
	"Wave/router"
	"context"
	"flag"

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
	ifMigrate := flag.Bool("m", false, "recreate table")
	flag.Parse()
	if *ifMigrate {
		migrator.Migrate()
	}
	router.NewRouter()

}
