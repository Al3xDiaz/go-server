package services

import (
	"encoding/json"
	"io"

	"github.com/al3xdiaz/go-server/models"
	"gorm.io/gorm"
)

type TelephoneService struct {
	DB *gorm.DB
}

func (t *TelephoneService) CreateTelephone(username string, Body io.ReadCloser) (model models.Telephone, err error) {
	var user models.User
	err = t.DB.Where("user_name=? and staff", username).First(&user).Error
	if err != nil {
		return models.Telephone{}, err
	}
	err = t.DB.Model(&user).Association("Profile").Find(&user.Profile)
	if err != nil {
		return models.Telephone{}, err
	}
	var telephone models.Telephone
	json.NewDecoder(Body).Decode(&telephone)
	telephone.ProfileID = user.Profile.ID
	err = t.DB.Create(&telephone).Error
	return telephone, err
}
