package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET = []byte(getEnv("SECRET", "super-secret-auth-key"))

func CreateJWT(data interface{}) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["data"] = data

	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] != nil {
			header := r.Header["Authorization"][0]
			header = strings.Split(header, "Bearer ")[1]
			if len(header) == 0 {
				Unauthorized(w, map[string]string{"msg": "token not exist"})
			}
			token, err := jwt.Parse(header, func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					Unauthorized(w, map[string]string{"msg": "not authorized"})
				}
				return SECRET, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				Unauthorized(w, map[string]string{"msg": "not authorized: " + err.Error()})
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			Unauthorized(w, map[string]string{"msg": "not authorized"})
		}
	})
}
