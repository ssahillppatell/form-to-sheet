package services

import (
	"context"
	"log"
	"net/mail"
	"os"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/api/sheets/v4"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func Submit(email string) Response {
	errResponse := Response{Message: "Something went wrong!", Status: 500}
	successResponse := Response{Message: "success", Status: 200}

	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
		log.Println(err)
		return errResponse
	}

	mailAddress, err := mail.ParseAddress(email)
	if err != nil {
		log.Println("Error parsing email address")
		log.Println(err)
		return errResponse
	}

	sheetID := os.Getenv("SPREADSHEET_ID")
	writeRange := "Sheet1!A1:B1"

	ctx := context.Background()

	sheetsService, err := sheets.NewService(ctx)
	if err != nil {
		log.Println("Error creating sheets service")
		log.Println(err)
		return errResponse
	}

	currentTime := time.Now().Format(time.RFC3339)

	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{{mailAddress.Address, currentTime}},
	}

	_, err = sheetsService.Spreadsheets.Values.Append(sheetID, writeRange, valueRange).ValueInputOption("USER_ENTERED").Context(ctx).Do()
	if err != nil {
		log.Println("Error appending value to sheet")
		log.Println(err)
		return errResponse
	}

	return successResponse
}
