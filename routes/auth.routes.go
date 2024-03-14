package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
	"github.com/al3xdiaz/go-server/services"
	"github.com/al3xdiaz/go-server/utils"
)

type ILogin struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var login ILogin
	var err error
	json.NewDecoder(r.Body).Decode(&login)
	login.Password = models.MakePassword(login.Password)

	var user models.User
	db.DB.First(&user, login)
	if user.UserName == "" {
		utils.Unauthorized(w, nil)
		return
	}
	db.DB.Model(&user).Association("Profile").Find(&user.Profile)
	db.DB.Model(&user).Association("Permisions").Find(&user.Permisions)

	token, err := utils.CreateJWT(map[string]any{
		"username":   user.UserName,
		"permisions": user.MakeMapString(),
	})
	if err != nil {
		utils.InternalServerError(w, "error create token")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	utils.Ok(w, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}
func LogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})
	utils.NoContend(w)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		utils.BadRequest(w, err.Error())
		return
	}
	token, err := utils.CreateJWT(map[string]any{
		"username":   user.UserName,
		"permisions": user.MakeMapString(),
	})
	if err != nil {
		utils.InternalServerError(w, "error create token")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		// SameSite: http.SameSiteLaxMode,
		SameSite: http.SameSiteNoneMode,
	})
	utils.Ok(w, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

func UserData(w http.ResponseWriter, r *http.Request) {
	data, _ := utils.ValidateJWT(w, r)
	username := data["username"]

	service := services.ProfileService{}
	user, err := service.GetData(username.(string))
	if err != nil {
		utils.InternalServerError(w, "User Not exist")
		return
	}
	utils.Ok(w, user)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// var profile models.Profile

	data, _ := utils.ValidateJWT(w, r)
	username := data["username"]

	service := services.ProfileService{}
	user, err := service.UpdateProfile(username.(string), r.Body)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}
	utils.Ok(w, user)
}
func GetProfile(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	service := services.ProfileService{}
	user, err := service.GetProfile(username)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}
	utils.Ok(w, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	service := services.UserService{
		DB: db.DB,
	}
	users, err := service.GetUsers()
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}
	utils.Ok(w, users)
}
