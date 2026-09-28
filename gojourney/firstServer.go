package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// HealthResponse is the JSON body returned by the /health endpoint.
type HealthResponse struct {
	Status string `json:"status"`

	Message string `json:"message,omitempty"`
}

// health handles GET /health and GET /health?name=<name>.
func health(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{Status: "ok"}

	// The name query parameter is optional; Get returns "" if it's absent.
	if name := r.URL.Query().Get("name"); name != "" {
		resp.Message = "Hello, " + name + "!"
	}

	// Headers must be set before the first write to the body.
	w.Header().Set("Content-Type", "application/json")

	// Encode straight to the response writer. At this point the status line
	// is already sent, so on failure we can only log the error.
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("encode error:", err)
	}
}

func main() {
	mux := http.NewServeMux()

	// Method-qualified patterns require Go 1.22+.
	mux.HandleFunc("GET /health", health)

	log.Println("listening on :8080")
	// ListenAndServe blocks; it only returns on error, so log.Fatal exits with it.
	log.Fatal(http.ListenAndServe(":8080", mux))
}
