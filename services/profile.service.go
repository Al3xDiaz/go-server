package services

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
)

type ProfileService struct{}

func (o ProfileService) GetData(username string) (model models.User, err error) {
	var user models.User
	err = db.DB.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	err = db.DB.Model(&user).Association("Profile").Find(&user.Profile)
	if err != nil {
		return user, err
	}
	return user, err
}

func (o ProfileService) UpdateProfile(username string, Body io.ReadCloser) (model models.User, err error) {
	user, err := o.GetData(username)
	if err != nil {
		return user, errors.New("user not exist")
	}
	json.NewDecoder(Body).Decode(&user.Profile)
	db.DB.Save(&user.Profile)

	return user, nil
}
func (o ProfileService) GetProfile(username string) (models.User, error) {
	var user models.User
	err := db.DB.Where("user_name=?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	err = db.DB.Model(&user).Association("Profile").Find(&user.Profile)
	if err != nil {
		return user, err
	}
	err = db.DB.Model(&user.Profile).Association("Telephone").Find(&user.Profile.Telephone)
	if err != nil {
		return user, err
	}
	return user, err
}
