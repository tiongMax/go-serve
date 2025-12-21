package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func extractID(r *http.Request, paramName string) (int, error) {
	idStr := r.PathValue(paramName)
	if idStr == "" {
		return 0, fmt.Errorf("%s is required", paramName)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid %s", paramName)
	}
	return id, nil
}

// writeJSON sends a JSON response with the given status code and data
func writeJSON(w http.ResponseWriter, status int, data any) {
	// 1. Marshal to memory first to catch encoding errors before touching the network
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// 2. Now that we know the data is valid, set headers and write
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

// readJSON decodes the request body into the provided destination
func readJSON(r *http.Request, dest any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Strict checking
	return decoder.Decode(dest)
}

// writeError sends a JSON error response
func writeError(w http.ResponseWriter, status int, message string) {
	type Envelope struct {
		Error string `json:"error"`
	}
	writeJSON(w, status, &Envelope{Error: message})
}
