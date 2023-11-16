package utils

import (
	"encoding/json"
	"net/http"
)

func BadRequest(w http.ResponseWriter, msg string) {
	Response(&w, http.StatusBadRequest, map[string]string{"error": msg})
}

func InternalServerError(w http.ResponseWriter, msg string) {
	Response(&w, http.StatusInternalServerError, map[string]string{"error": msg})
}

func NotFound(w http.ResponseWriter, msg string) {
	Response(&w, http.StatusNotFound, map[string]string{"error": msg})
}

func Ok(w http.ResponseWriter, data interface{}) {
	Response(&w, http.StatusOK, data)
}
func Unauthorized(w http.ResponseWriter, data interface{}) {
	Response(&w, http.StatusUnauthorized, data)
}

func Forbidden(w http.ResponseWriter, data interface{}) {
	Response(&w, http.StatusForbidden, data)
}
func Response(w *http.ResponseWriter, code int, data interface{}) {
	// (*w).Header().Set("Access-Control-Allow-Origin", "*")
	// (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// (*w).Header().Add("Content-Type", "application/json")
	(*w).Header().Add("Access-Control-Allow-Headers", "Authorization, content-type")
	(*w).WriteHeader(code)
	json.NewEncoder(*w).Encode(data)
}
