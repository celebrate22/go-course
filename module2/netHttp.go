package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/health", healthHandler)

    log.Println("listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}