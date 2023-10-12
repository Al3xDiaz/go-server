package utils

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
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

func ValidateJWT(w http.ResponseWriter, r *http.Request) (bool, map[string]interface{}) {
	if r.Header["Authorization"] != nil {
		header := r.Header["Authorization"][0]
		header = strings.Split(header, "Bearer ")[1]
		if len(header) == 0 {
			Unauthorized(w, map[string]string{"msg": "token not exist"})
		}
		token, err := jwt.Parse(header, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return false, nil
			}
			return SECRET, nil
		})

		if err != nil {
			return false, nil
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			resp := claims["data"].(map[string]interface{})
			var user models.User
			search := db.DB.Find(&user, "user_name = ?", resp["username"])
			if search.Error != nil || user.ID == 0 {
				return false, nil
			}
			return true, resp
		}
	}
	return false, nil
}

func RequireAuth(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid, _ := ValidateJWT(w, r)
		if valid {
			next(w, r)
			return
		}
		Unauthorized(w, map[string]interface{}{
			"msg": "invalid token",
		})
	})
}
func RequirePermision(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid, data := ValidateJWT(w, r)
		if valid {
			var user models.User
			db.DB.First(&user, "user_name = ?", data["username"])
			db.DB.Model(&user).Association("Permisions")

			for _, v := range user.Permisions {
				reg, err := regexp.Compile(v.Path)
				if err != nil {
					break
				}
				if reg.FindString(r.URL.Path) != "" {
					methodsreg, _ := regexp.Compile(v.Methods)
					if methodsreg.FindString(r.Method) != "" {
						next(w, r)
						return
					}
				}
			}
			Forbidden(w, map[string]string{"msg": "not authorized: " + r.URL.Path})
		} else {
			Unauthorized(w, map[string]string{"msg": "unauthorized"})
		}
	})
}
