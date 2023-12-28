package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type RequestData struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/handlepost", handlePost)
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	if requestData.Message == "" {
		errorResponse := ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid JSON message",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	fmt.Printf("Received message: %s\n", requestData.Message)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Data received successfully")
}
