package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/al3xdiaz/go-server/services"
	request "github.com/al3xdiaz/go-server/utils"
	"github.com/gorilla/mux"
)

func ListCourses(w http.ResponseWriter, r *http.Request) {
	// ...
	username := r.URL.Query().Get("username")
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = -1
	}
	service := services.CourseService{}
	response, err := service.ListCourses(username, limit)
	if err != nil {
		request.InternalServerError(w, "Internal server Error")
	}
	request.Ok(w, response)
}
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	// ...
	typeInsert := r.URL.Query().Get("type")
	data, _ := request.ValidateJWT(w, r)
	username := data["username"].(string)
	service := services.CourseService{}
	if typeInsert == "bulk" {
		err := service.CreateCourses(username, &r.Body)
		if err != nil {
			log.Print(err.Error())
			request.InternalServerError(w, "internal server error")
			return
		}
		request.NoContend(w)
		return
	}
	response, err := service.CreateCourse(username, &r.Body)
	if err != nil {
		request.InternalServerError(w, "internal server error")
		return
	}
	request.Ok(w, response)
}
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	data, _ := request.ValidateJWT(w, r)
	username := data["username"].(string)
	service := services.CourseService{}
	response, err := service.UpdateCourse(id, username, &r.Body)
	if err != nil {
		request.NotFound(w, err.Error())
		return
	}
	request.Ok(w, response)
}
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	// ...
	params := mux.Vars(r)
	id := params["id"]
	data, _ := request.ValidateJWT(w, r)
	username := data["username"].(string)
	service := services.CourseService{}
	_, err := service.DeleteCourse(id, username)
	if err != nil {
		request.NotFound(w, err.Error())
		return
	}
	request.NoContend(w)
}
