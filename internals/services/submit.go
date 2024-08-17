package services

import (
	"context"
	"log"
	"net/mail"
	"os"
	"time"

	"google.golang.org/api/sheets/v4"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func Submit(email string, gender string) Response {
	errResponse := Response{Message: "Something went wrong!", Status: 500}
	successResponse := Response{Message: "success", Status: 200}

	mailAddress, err := mail.ParseAddress(email)
	if err != nil {
		log.Println("Error parsing email address")
		log.Println(err)
		return errResponse
	}

	if gender != "Male" && gender != "Female" {
		log.Println("Incorrect Gender!")
		return errResponse
	}

	sheetId := os.Getenv("SPREADSHEET_ID")
	writeRange := "Sheet1!A1:C1"

	ctx := context.Background()

	sheetsService, err := sheets.NewService(ctx)
	if err != nil {
		log.Println("Error creating sheets service")
		log.Println(err)
		return errResponse
	}

	currentTime := time.Now().Format(time.RFC3339)

	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{{mailAddress.Address, gender, currentTime}},
	}

	_, err = sheetsService.Spreadsheets.Values.Append(sheetId, writeRange, valueRange).ValueInputOption("USER_ENTERED").Context(ctx).Do()
	if err != nil {
		log.Println("Error appending value to sheet")
		log.Println(err)
		return errResponse
	}

	return successResponse
}
