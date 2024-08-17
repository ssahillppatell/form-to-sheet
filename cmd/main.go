package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/prastavna/form-to-sheet/internals/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api.Routes()

	port := os.Getenv("PORT")

	log.Print("Server started at port: " + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting the server")
	}

}
