package main

import (
	"github.com/BennoAlif/ps-cats-social/src/drivers/db"
	"github.com/BennoAlif/ps-cats-social/src/http"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	dbConnection := db.CreateConnection()

	h := http.New(&http.Http{
		DB: dbConnection,
	})

	h.Launch()
}
