package main

import (
	"Wave/database"
	"Wave/router"
)

func main() {
	database.GetConnection()
	router.NewRouter()

}
