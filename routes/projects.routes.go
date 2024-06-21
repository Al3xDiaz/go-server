package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/al3xdiaz/go-server/services"
	request "github.com/al3xdiaz/go-server/utils"
	"github.com/gorilla/mux"
)

func ListProjects(w http.ResponseWriter, r *http.Request) {
	// ...
	username := r.URL.Query().Get("username")
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = -1
	}
	service := services.ProjectService{}
	response, err := service.ListProjects(username, limit)
	if err != nil {
		request.InternalServerError(w, "Internal server Error")
		return
	}
	request.Ok(w, response)
}
func CreateProject(w http.ResponseWriter, r *http.Request) {
	// ...
	data, _ := request.ValidateJWT(w, r)
	username := data["username"].(string)
	service := services.ProjectService{}
	response, err := service.CreateProject(username, &r.Body)
	if err != nil {
		log.Print(err.Error())
		request.InternalServerError(w, "internal server error")
		return
	}
	request.Ok(w, response)
}
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	data, _ := request.ValidateJWT(w, r)
	username := data["username"].(string)
	service := services.ProjectService{}
	response, err := service.UpdateProject(id, username, &r.Body)
	if err != nil {
		request.NotFound(w, err.Error())
		return
	}
	request.Ok(w, response)
}
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	// ...
	params := mux.Vars(r)
	id := params["id"]
	data, _ := request.ValidateJWT(w, r)
	username := data["username"].(string)
	service := services.ProjectService{}
	_, err := service.DeleteProject(id, username)
	if err != nil {
		request.NotFound(w, err.Error())
		return
	}
	request.NoContend(w)
}
