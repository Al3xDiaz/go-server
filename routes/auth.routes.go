package routes

import (
	"encoding/json"
	"net/http"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
	"github.com/al3xdiaz/go-server/services"
	"github.com/al3xdiaz/go-server/utils"
	request "github.com/al3xdiaz/go-server/utils"
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
	utils.Ok(w, map[string]interface{}{
		"user":  user,
		"token": token,
	})
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
	utils.Ok(w, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}
func UserData(w http.ResponseWriter, r *http.Request) {
	_, data := utils.ValidateJWT(w, r)
	username := data["username"]

	service := services.ProfileService{
		DB: db.DB,
	}
	user, err := service.GetData(username.(string))
	if err != nil {
		utils.InternalServerError(w, "User Not exist")
		return
	}
	db.DB.Model(&user).Association("Permisions").Find(&user.Permisions)

	utils.Ok(w, user)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// var profile models.Profile

	_, data := utils.ValidateJWT(w, r)
	username := data["username"]

	service := services.ProfileService{
		DB: db.DB,
	}
	user, err := service.UpdateProfile(username.(string), r.Body)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}
	utils.Ok(w, user)
}
func GetProfile(w http.ResponseWriter, r *http.Request) {
	_, data := request.ValidateJWT(w, r)
	username := data["username"]
	service := services.ProfileService{
		DB: db.DB,
	}
	profile, err := service.GetProfile(username.(string))
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}
	utils.Ok(w, profile)
}
func PostTelephone(w http.ResponseWriter, r *http.Request) {
	_, data := request.ValidateJWT(w, r)
	username := data["username"]
	service := services.TelephoneService{}
	profile, err := service.CreateTelephone(username.(string), r.Body)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}
	utils.Ok(w, profile)
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
