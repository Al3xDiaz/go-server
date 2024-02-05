package services

import (
	"errors"
	"net/http"

	"github.com/al3xdiaz/go-server/db"
)

type VCardService struct {
}

func (o VCardService) CreateVCard(username string, w *http.ResponseWriter) error {
	service := ProfileService{
		DB: db.DB,
	}
	user, err := service.GetData(username)
	if err != nil {
		return errors.New("user not found")
	}
	if user.Profile.Telephone == "" {
		return errors.New("the user not has telephone")
	}

	(*w).Write([]byte("BEGIN:VCARD\r\n"))
	(*w).Write([]byte("VERSION:2.1\r\n"))
	(*w).Write([]byte("N:" + user.Profile.FirstName + "\r\n"))
	(*w).Write([]byte("FN:" + user.Profile.FirstName + " " + user.Profile.LastName + "\r\n"))
	(*w).Write([]byte("NICKNAME:" + user.UserName + "\r\n"))
	(*w).Write([]byte("TEL;CELL:" + user.Profile.Telephone + "\r\n"))
	(*w).Write([]byte("EMAIL:" + user.Email + "\r\n"))
	if user.Profile.Photo != "" {
		(*w).Write([]byte("PHOTO;VALUE=URI:" + user.Profile.Photo + "\r\n"))
	}
	if user.Profile.Facebook != "" {
		(*w).Write([]byte("X-SOCIALPROFILE;TYPE=facebook:" + user.Profile.Facebook + "\r\n"))
	}
	if user.Profile.Twitter != "" {
		(*w).Write([]byte("X-SOCIALPROFILE;TYPE=twitter:" + user.Profile.Twitter + "\r\n"))
	}
	if user.Profile.Linkedin != "" {
		(*w).Write([]byte("X-SOCIALPROFILE;TYPE=linkedin:" + user.Profile.Linkedin + "\r\n"))
	}
	if user.Profile.Github != "" {
		(*w).Write([]byte("X-SOCIALPROFILE;TYPE=github:" + user.Profile.Github + "\r\n"))
	}
	if user.Profile.Instagram != "" {
		(*w).Write([]byte("X-SOCIALPROFILE;TYPE=instagram:" + user.Profile.Instagram + "\r\n"))
	}
	if user.Profile.Youtube != "" {
		(*w).Write([]byte("X-SOCIALPROFILE;TYPE=youtube:" + user.Profile.Youtube + "\r\n"))
	}
	if user.Profile.Website != "" {
		(*w).Write([]byte("URL:" + user.Profile.Website + "\r\n"))
	}

	(*w).Write([]byte("END:VCARD\r\n"))
	return nil
}
