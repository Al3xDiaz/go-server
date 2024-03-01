package services

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
)

type TelephoneService struct{}

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
func (t *TelephoneService) DeleteTelephone(username string, id string) error {
	var model models.Telephone
	err := db.DB.
		Joins("INNER JOIN profiles p ON p.id = telephones.profile_id").
		Joins("INNER JOIN users u ON u.id = p.user_id").
		Where("telephones.id = ? and u.user_name = ?", id, username).
		First(&model).Error
	if err != nil {
		return errors.New("the user: " + username + " doesn'n have this course")
	}
	db.DB.Delete(&model)
	return nil
}
