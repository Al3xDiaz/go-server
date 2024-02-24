package routes

import (
	"net/http"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
	"github.com/al3xdiaz/go-server/services"
	request "github.com/al3xdiaz/go-server/utils"
	"github.com/gorilla/mux"
)

func ListCourses(w http.ResponseWriter, r *http.Request) {
	// ...
	username := r.URL.Query().Get("username")
	response, err := services.ListCourses(username)
	if err != nil {
		request.InternalServerError(w, "Internal server Error")
	}
	request.Ok(w, response)
}
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	// ...
	_, data := request.ValidateJWT(w, r)
	username := data["username"].(string)
	response, err := services.CreateCourse(username, &r.Body)
	if err != nil {
		request.InternalServerError(w, "internal server error")
	}
	request.Ok(w, response)
}
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	_, data := request.ValidateJWT(w, r)
	username := data["username"].(string)
	response, err := services.UpdateCourse(id, username, &r.Body)
	if err != nil {
		request.NotFound(w, err.Error())
		return
	}
	request.Ok(w, response)
}
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	// ...
	var commentary models.Commentary

	// Get params
	params := mux.Vars(r)
	id := params["id"]
	_, data := request.ValidateJWT(w, r)
	username := data["username"]
	response := db.DB.
		Joins("INNER JOIN users u ON u.id = commentaries.user_id").
		Where("commentaries.id = ?", id).
		Where("u.user_name = ?", username).
		First(&commentary)
	if response.Error != nil {
		request.NotFound(w, "Commentary not found")
		return
	}
	db.DB.Delete(&commentary)
	request.NoContend(w)
}
