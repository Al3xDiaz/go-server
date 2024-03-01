package services

import (
	"encoding/json"
	"io"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
)

type TelephoneService struct {
}

func (t *TelephoneService) CreateTelephone(username string, Body io.ReadCloser) (model models.Telephone, err error) {
	var user models.User
	err = db.DB.Where("user_name=?", username).First(&user).Error
	if err != nil {
		return models.Telephone{}, err
	}
	err = db.DB.Model(&user).Association("Profile").Find(&user.Profile)
	if err != nil {
		return models.Telephone{}, err
	}
	var telephone models.Telephone
	json.NewDecoder(Body).Decode(&telephone)
	telephone.ProfileID = user.Profile.ID
	err = db.DB.Create(&telephone).Error
	return telephone, err
}
