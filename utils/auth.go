package utils

import (
	"errors"
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

func ValidateJWT(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	var access_token string
	if r.Header["Authorization"] != nil {
		header := r.Header["Authorization"][0]
		access_token = strings.Split(header, "Bearer ")[1]
	}
	cookie, err := r.Cookie("access_token")
	if err == nil {
		access_token = cookie.Value
	}

	token, err := jwt.Parse(access_token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("token not valid")
		}
		return SECRET, nil
	})
	if err != nil {
		return nil, errors.New("token cant parse")
	}

	claims := token.Claims.(jwt.MapClaims)
	resp := claims["data"].(map[string]interface{})
	var user models.User
	search := db.DB.Find(&user, "user_name = ?", resp["username"])
	if search.Error != nil || user.ID == 0 {
		return nil, errors.New("user not exist")
	}
	return resp, nil
}

func RequireAuth(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := ValidateJWT(w, r)
		if err != nil {
			Unauthorized(w, map[string]interface{}{
				"msg": err.Error(),
			})
			return
		}
		next(w, r)
	})
}
func RequireStaff(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := ValidateJWT(w, r)
		if err != nil {
			Unauthorized(w, map[string]string{"msg": err.Error()})
			return
		}
		var user models.User
		db.DB.First(&user, "user_name=? and staff", data["username"])
		if user.ID == 0 {
			Forbidden(w, map[string]string{"msg": "not authorized"})
			return
		}
		next(w, r)
	})
}
func RequirePermision(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := ValidateJWT(w, r)
		if err != nil {
			Unauthorized(w, map[string]string{"msg": err.Error()})
			return
		}
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
	})
}
