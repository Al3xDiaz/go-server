package routes

import (
	"net/http"

	"github.com/al3xdiaz/go-server/services"
	"github.com/al3xdiaz/go-server/utils"
	"github.com/gorilla/mux"
)

func PostTelephone(w http.ResponseWriter, r *http.Request) {
	_, data := utils.ValidateJWT(w, r)
	username := data["username"]
	service := services.TelephoneService{}
	profile, err := service.CreateTelephone(username.(string), r.Body)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}
	utils.Ok(w, profile)
}
func DeleteTelephone(w http.ResponseWriter, r *http.Request) {
	_, data := utils.ValidateJWT(w, r)
	username := data["username"]
	params := mux.Vars(r)
	id := params["id"]
	service := services.TelephoneService{}
	err := service.DeleteTelephone(username.(string), id)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}
	utils.NoContend(w)
}
