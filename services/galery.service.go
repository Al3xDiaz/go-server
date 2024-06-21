package services

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
)

type GaleryService struct{}

func (service GaleryService) ListGaleries(username string, limit int) ([]models.Galery, error) {
	var model []models.Galery
	err := db.DB.
		Joins("INNER JOIN users u ON u.id = galeries.user_id").
		Where("u.user_name = ?", username).
		Find(&model).Error
	if err != nil {
		return model, errors.New("has ocurrent problem get data from database")
	}
	return model, nil
}
func (service GaleryService) CreateGalery(username string, body *io.ReadCloser) (models.Galery, error) {
	var model models.Galery
	json.NewDecoder(*body).Decode(&model)

	var user models.User
	db.DB.Where("user_name = ?", username).First(&user)
	model.UserID = user.ID
	db.DB.Create(&model)
	return model, nil
}
func (service GaleryService) UpdateGalery(id string, username string, body *io.ReadCloser) (models.Galery, error) {
	var model models.Galery
	err := db.DB.
		Joins("INNER JOIN users u ON u.id = galeries.user_id").
		Where("u.user_name = ?", username).
		Find(&model).Error
	if err != nil {
		return model, errors.New("the user: " + username + " doesn'n have this project")
	}
	json.NewDecoder(*body).Decode(&model)
	db.DB.Save(&model)
	return model, nil
}
func (service GaleryService) DeleteGalery(id string, username string) (models.Galery, error) {
	var model models.Galery
	err := db.DB.
		Joins("INNER JOIN users u ON u.id = galeries.user_id").
		Where("u.user_name = ?", username).
		Find(&model).Error
	if err != nil {
		return model, errors.New("the user: " + username + " doesn'n have this course")
	}
	db.DB.Delete(&model)
	return model, nil
}
