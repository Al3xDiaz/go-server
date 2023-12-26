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
func NoContend(w http.ResponseWriter) {
	Response(&w, http.StatusNoContent, nil)
}
func Unauthorized(w http.ResponseWriter, data interface{}) {
	Response(&w, http.StatusUnauthorized, data)
}

func Forbidden(w http.ResponseWriter, data interface{}) {
	Response(&w, http.StatusForbidden, data)
}
func Response(w *http.ResponseWriter, code int, data interface{}) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST,OPTIONS, PUT, DELETE")
	(*w).Header().Add("Content-Type", "application/json")
	(*w).Header().Add("Access-Control-Allow-Headers", "Authorization, content-type")
	(*w).WriteHeader(code)
	json.NewEncoder(*w).Encode(data)
}
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func HandlerCors(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Methods", "GET, POST,OPTIONS, PUT, DELETE")
		(w).Header().Add("Content-Type", "application/json")
		(w).Header().Add("Access-Control-Allow-Headers", "Authorization, content-type")
		if r.Method == http.MethodOptions {
			NoContend(w)
			return
		}
		next(w, r)
	})
}
