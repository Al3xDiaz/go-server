package routes

import (
	"encoding/json"
	"net/http"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
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
	db.DB.Model(&user).Association("Permisions").Find(&user.Permisions)

	token, err := utils.CreateJWT(map[string]any{
		"username":   user.UserName,
		"permisions": models.MakeMapString(user.Permisions),
	})
	if err != nil {
		utils.InternalServerError(w, "error create token")
		return
	}
	utils.Ok(w, map[string]string{"token": token})
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		utils.BadRequest(w, err.Error())
	}
	utils.Ok(w, user)
}
func UserData(w http.ResponseWriter, r *http.Request) {
	_, data := utils.ValidateJWT(w, r)
	username := data["username"]
	var user models.User
	db.DB.First(&user, "user_name = ?", username)
	db.DB.Model(&user).Association("Permisions").Find(&user.Permisions)
	utils.Ok(w, user)
}
