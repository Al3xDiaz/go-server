package routes

import (
	"net/http"

	"github.com/al3xdiaz/go-server/services"
	request "github.com/al3xdiaz/go-server/utils"
	"github.com/gorilla/mux"
)

func VCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	service := services.VCardService{}
	err := service.CreateVCard(username, &w)
	if err != nil {
		request.NotFound(w, err.Error())
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+username+".vcf")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
}
