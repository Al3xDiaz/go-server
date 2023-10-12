package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

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
	origin, _ := url.Parse(r.Header.Get("origin"))
	var commentaries []models.Commentary
	data := db.DB.
		Joins("INNER JOIN sites s ON s.id = commentaries.site_id").
		Where("s.url = ?", origin.Hostname()).
		Find(&commentaries)
	if data.Error != nil {
		log.Fatal(data.Error)
		request.InternalServerError(w, "Error getting commentaries")
		return
	}
	request.Ok(w, commentaries)
}
func CreateCommentary(w http.ResponseWriter, r *http.Request) {
	// ...
	origin, _ := url.Parse(r.Header.Get("origin"))
	_, data := request.ValidateJWT(w, r)
	username := data["username"]
	var site models.Site
	db.DB.FirstOrCreate(&site, models.Site{
		Url: origin.Hostname(),
	})

	var commentary models.Commentary
	json.NewDecoder(r.Body).Decode(&commentary)
	var user models.User
	db.DB.First(&user, "user_name = ?", username)
	commentary.UserID = user.ID
	commentary.SiteId = site.ID
	createCommentary := db.DB.Create(&commentary)
	if createCommentary.Error != nil {
		request.InternalServerError(w, "Error creating commentary")
		return
	}
	request.Ok(w, commentary)
}
