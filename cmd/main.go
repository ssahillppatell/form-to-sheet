package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/prastavna/form-to-sheet/internals/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api.Routes()

	log.Print("Server started at :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server")
	}

}
