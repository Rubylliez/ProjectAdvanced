package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JSONData struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handleJSONRequest)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleJSONRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var jsonData JSONData
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if jsonData.Message != "" {
		fmt.Println("Received message:", jsonData.Message)

		response := map[string]string{
			"status":  "success",
			"message": "Data successfully received",
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	} else {
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}
}
