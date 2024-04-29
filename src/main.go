package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BennoAlif/ps-cats-social/src/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := http.NewServeMux()

	routes.LoadAuthRoutes(router)

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	server := http.Server{
		Addr:    ":8080",
		Handler: v1,
	}

	fmt.Println("Server is breathing on port :8080")
	server.ListenAndServe()

}
