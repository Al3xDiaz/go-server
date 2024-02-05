package routes

import (
	"net/http"

	"github.com/al3xdiaz/go-server/models"
	"github.com/al3xdiaz/go-server/services"
	request "github.com/al3xdiaz/go-server/utils"
	"github.com/gorilla/mux"
)

func VCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	service := services.VCardService{}
	err := service.CreateVCard(username, func(user models.User) {
		w.Header().Set("Content-Disposition", "attachment; filename="+username+".vcf")
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		w.Write([]byte("BEGIN:VCARD\r\n"))
		w.Write([]byte("VERSION:2.1\r\n"))
		w.Write([]byte("N:" + user.Profile.FirstName + "\r\n"))
		w.Write([]byte("FN:" + user.Profile.FirstName + " " + user.Profile.LastName + "\r\n"))
		w.Write([]byte("NICKNAME:" + user.UserName + "\r\n"))
		w.Write([]byte("TEL;CELL:" + user.Profile.Telephone + "\r\n"))
		w.Write([]byte("EMAIL:" + user.Email + "\r\n"))
		if user.Profile.Photo != "" {
			w.Write([]byte("PHOTO;VALUE=URI:" + user.Profile.Photo + "\r\n"))
		}
		if user.Profile.Facebook != "" {
			w.Write([]byte("X-SOCIALPROFILE;TYPE=facebook:" + user.Profile.Facebook + "\r\n"))
		}
		if user.Profile.Twitter != "" {
			w.Write([]byte("X-SOCIALPROFILE;TYPE=twitter:" + user.Profile.Twitter + "\r\n"))
		}
		if user.Profile.Linkedin != "" {
			w.Write([]byte("X-SOCIALPROFILE;TYPE=linkedin:" + user.Profile.Linkedin + "\r\n"))
		}
		if user.Profile.Github != "" {
			w.Write([]byte("X-SOCIALPROFILE;TYPE=github:" + user.Profile.Github + "\r\n"))
		}
		if user.Profile.Instagram != "" {
			w.Write([]byte("X-SOCIALPROFILE;TYPE=instagram:" + user.Profile.Instagram + "\r\n"))
		}
		if user.Profile.Youtube != "" {
			w.Write([]byte("X-SOCIALPROFILE;TYPE=youtube:" + user.Profile.Youtube + "\r\n"))
		}
		if user.Profile.Website != "" {
			w.Write([]byte("URL:" + user.Profile.Website + "\r\n"))
		}

		w.Write([]byte("END:VCARD\r\n"))
	})
	if err != nil {
		request.NotFound(w, err.Error())
		return
	}
}
