package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prastavna/form-to-sheet/internals/services"
)

type User struct {
	Email  string `json:"email"`
	Gender string `json:"gender"`
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding the request body", http.StatusBadRequest)
		log.Println("Error decoding the request body")
	}

	w.Header().Set("Content-Type", "application/json")
	response := services.Submit(user.Email, user.Gender)
	if response.Status != 200 {
		http.Error(w, response.Message, response.Status)
		return
	}
	json.NewEncoder(w).Encode(response)
}
