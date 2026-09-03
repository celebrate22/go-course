package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// RequestBody defines the expected incoming JSON structure
type RequestBody struct {
	Message *string `json:"message"` // Used a pointer to distinguish between an empty string "" and a missing field
}

// ResponseBody defines the successful JSON output structure
type ResponseBody struct {
	Echo   string `json:"echo"`
	Length int    `json:"length"`
}

// ErrorResponse defines the standard error JSON structure
type ErrorResponse struct {
	Error string `json:"error"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Enforce POST method
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Method not allowed"})
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// 2. Decode the incoming JSON body
	var req RequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid JSON format"})
		return
	}

	// 3. Validate that the "message" field is present
	if req.Message == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Missing required field: message"})
		return
	}

	// 4. Build and send the successful response
	resp := ResponseBody{
		Echo:   *req.Message,
		Length: len(*req.Message), // Character count of the string
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", echoHandler)

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
