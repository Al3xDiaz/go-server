package services

import (
	"errors"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
)

type VCardService struct {
}

func (o VCardService) CreateVCard(username string, next func(user models.User)) error {
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
	next(user)
	return nil
}
