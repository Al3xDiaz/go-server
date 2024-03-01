package services

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/al3xdiaz/go-server/models"
	"gorm.io/gorm"
)

type ProfileService struct {
	DB *gorm.DB
}

func (o ProfileService) GetData(username string) (model models.User, err error) {
	var user models.User
	err = o.DB.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	err = o.DB.Model(&user).Association("Profile").Find(&user.Profile)
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
	o.DB.Save(&user.Profile)

	return user, nil
}
func (o ProfileService) GetProfile(username string) (model models.Profile, err error) {
	var user models.User
	err = o.DB.Where("user_name=?", username).First(&user).Error
	if err != nil {
		return user.Profile, err
	}
	err = o.DB.Model(&user).Association("Profile").Find(&user.Profile)
	if err != nil {
		return user.Profile, err
	}
	err = o.DB.Model(&user.Profile).Association("Telephone").Find(&user.Profile.Telephone)
	if err != nil {
		return user.Profile, err
	}
	return user.Profile, err
}
