package routes

import (
	"encoding/json"
	"net/http"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
	request "github.com/al3xdiaz/go-server/utils"
	"github.com/gorilla/mux"
)

func GetCommentary(w http.ResponseWriter, r *http.Request) {
	// ...
	var commentary models.Commentary

	// Get params
	params := mux.Vars(r)
	data := db.DB.First(&commentary, params["id"])
	if data.Error != nil {
		request.NotFound(w, "Commentary not found")
		return
	}
	request.Ok(w, commentary)
}

func GetCommentaries(w http.ResponseWriter, r *http.Request) {
	// ...
	var commentaries []models.Commentary
	data := db.DB.Find(&commentaries)
	if data.Error != nil {
		request.InternalServerError(w, "Error getting commentaries")
		return
	}
	request.Ok(w, commentaries)
}
func CreateCommentary(w http.ResponseWriter, r *http.Request) {
	// ...
	_, data := request.ValidateJWT(w, r)
	username := data["username"]

	var commentary models.Commentary
	json.NewDecoder(r.Body).Decode(&commentary)
	var user models.User
	db.DB.First(&user, "user_name = ?", username)
	commentary.UserID = user.ID
	createCommentary := db.DB.Create(&commentary)
	if createCommentary.Error != nil {
		request.InternalServerError(w, "Error creating commentary")
		return
	}
	request.Ok(w, commentary)
}
