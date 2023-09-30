package routes

import (
	"encoding/json"
	"log"
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

	var user models.Users
	db.DB.First(&user, login)
	log.Output(0, user.Password)
	if user.UserName == "" {
		utils.Unauthorized(w, nil)
		return
	}

	token, err := utils.CreateJWT()
	if err != nil {
		utils.InternalServerError(w, "error create token")
		return
	}
	utils.Ok(w, token)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		utils.BadRequest(w, err.Error())
	}
	json.NewEncoder(w).Encode(user)
}
