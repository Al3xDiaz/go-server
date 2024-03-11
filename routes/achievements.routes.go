package routes

import (
	"log"
	"net/http"

	"github.com/al3xdiaz/go-server/services"
	request "github.com/al3xdiaz/go-server/utils"
	"github.com/gorilla/mux"
)

func ListAchievements(w http.ResponseWriter, r *http.Request) {
	// ...
	username := r.URL.Query().Get("username")
	service := services.AchievementsHistoryService{}
	response, err := service.ListAchievements(username)
	if err != nil {
		request.InternalServerError(w, "Internal server Error")
		return
	}
	request.Ok(w, response)
}
func CreateAchievement(w http.ResponseWriter, r *http.Request) {
	// ...
	typeInsert := r.URL.Query().Get("type")
	data, _ := request.ValidateJWT(w, r)
	username := data["username"].(string)
	service := services.AchievementsHistoryService{}
	if typeInsert == "bulk" {
		err := service.CreateAchievements(username, &r.Body)
		if err != nil {
			request.InternalServerError(w, "internal server error")
			return
		}
		request.NoContend(w)
		return
	}
	response, err := service.CreateAchievement(username, &r.Body)
	if err != nil {
		log.Print(err.Error())
		request.InternalServerError(w, "internal server error")
		return
	}
	request.Ok(w, response)
}
func UpdateAchievement(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	data, _ := request.ValidateJWT(w, r)
	username := data["username"].(string)
	service := services.AchievementsHistoryService{}
	response, err := service.UpdateAchievement(id, username, &r.Body)
	if err != nil {
		request.NotFound(w, err.Error())
		return
	}
	request.Ok(w, response)
}
func DeleteAchievement(w http.ResponseWriter, r *http.Request) {
	// ...
	params := mux.Vars(r)
	id := params["id"]
	data, _ := request.ValidateJWT(w, r)
	username := data["username"].(string)
	service := services.AchievementsHistoryService{}
	_, err := service.DeleteAchievement(id, username)
	if err != nil {
		request.NotFound(w, err.Error())
		return
	}
	request.NoContend(w)
}
