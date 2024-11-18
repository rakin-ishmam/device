package web

import (
	"encoding/json"
	"net/http"
)

// Send sends a JSON response
func Send(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// SendCreated sends a JSON response with status code 201
func SendCreated(w http.ResponseWriter, data interface{}) {
	Send(w, http.StatusCreated, data)
}

// SendOk sends a JSON response with status code 200
func SendOk(w http.ResponseWriter, data interface{}) {
	Send(w, http.StatusOK, data)
}

// SendError sends a JSON error response
func SendError(w http.ResponseWriter, statusCode int, err error) {
	Send(w, statusCode, map[string]string{"error": err.Error()})
}
