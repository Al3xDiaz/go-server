package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/al3xdiaz/go-server/services"
	request "github.com/al3xdiaz/go-server/utils"
	"github.com/gorilla/mux"
)

func ListGaleries(w http.ResponseWriter, r *http.Request) {
	// ...
	username := r.URL.Query().Get("username")
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = -1
	}
	service := services.GaleryService{}
	response, err := service.ListGaleries(username, limit)
	if err != nil {
		request.InternalServerError(w, "Internal server Error")
		return
	}
	request.Ok(w, response)
}
func CreateGalery(w http.ResponseWriter, r *http.Request) {
	// ...
	data, _ := request.ValidateJWT(w, r)
	username := data["username"].(string)
	service := services.GaleryService{}
	response, err := service.CreateGalery(username, &r.Body)
	if err != nil {
		log.Print(err.Error())
		request.InternalServerError(w, "internal server error")
		return
	}
	request.Ok(w, response)
}
func UpdateGalery(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	data, _ := request.ValidateJWT(w, r)
	username := data["username"].(string)
	service := services.GaleryService{}
	response, err := service.UpdateGalery(id, username, &r.Body)
	if err != nil {
		request.NotFound(w, err.Error())
		return
	}
	request.Ok(w, response)
}
func DeleteGalery(w http.ResponseWriter, r *http.Request) {
	// ...
	params := mux.Vars(r)
	id := params["id"]
	data, _ := request.ValidateJWT(w, r)
	username := data["username"].(string)
	service := services.GaleryService{}
	_, err := service.DeleteGalery(id, username)
	if err != nil {
		request.NotFound(w, err.Error())
		return
	}
	request.NoContend(w)
}
