package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/al3xdiaz/go-server/services"
	request "github.com/al3xdiaz/go-server/utils"
	"github.com/gorilla/mux"
)

func ListGalleries(w http.ResponseWriter, r *http.Request) {
	// ...
	username := r.URL.Query().Get("username")
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = -1
	}
	service := services.GalleryService{}
	response, err := service.ListGalleries(username, limit)
	if err != nil {
		request.InternalServerError(w, "Internal server Error")
		return
	}
	request.Ok(w, response)
}
func CreateGallery(w http.ResponseWriter, r *http.Request) {
	// ...
	data, _ := request.ValidateJWT(w, r)
	username := data["username"].(string)
	service := services.GalleryService{}
	response, err := service.CreateGallery(username, &r.Body)
	if err != nil {
		log.Print(err.Error())
		request.InternalServerError(w, "internal server error")
		return
	}
	request.Ok(w, response)
}
func UpdateGallery(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	data, _ := request.ValidateJWT(w, r)
	username := data["username"].(string)
	service := services.GalleryService{}
	response, err := service.UpdateGallery(id, username, &r.Body)
	if err != nil {
		request.NotFound(w, err.Error())
		return
	}
	request.Ok(w, response)
}
func DeleteGallery(w http.ResponseWriter, r *http.Request) {
	// ...
	params := mux.Vars(r)
	id := params["id"]
	data, _ := request.ValidateJWT(w, r)
	username := data["username"].(string)
	service := services.GalleryService{}
	_, err := service.DeleteGallery(id, username)
	if err != nil {
		request.NotFound(w, err.Error())
		return
	}
	request.NoContend(w)
}
