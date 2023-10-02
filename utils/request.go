package utils

import (
	"encoding/json"
	"net/http"
)

func BadRequest(w http.ResponseWriter, msg string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func InternalServerError(w http.ResponseWriter, msg string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func NotFound(w http.ResponseWriter, msg string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func Ok(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
func Unauthorized(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(data)
}

func Forbidden(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(data)
}
