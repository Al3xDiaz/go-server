package services

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
)

type GalleryService struct{}

func (service GalleryService) ListGalleries(username string, limit int) ([]models.Gallery, error) {
	var model []models.Gallery
	err := db.DB.
		Joins("INNER JOIN users u ON u.id = galleries.user_id").
		Where("u.user_name = ?", username).
		Find(&model).Error
	if err != nil {
		return model, errors.New("has ocurrent problem get data from database")
	}
	return model, nil
}
func (service GalleryService) CreateGallery(username string, body *io.ReadCloser) (models.Gallery, error) {
	var model models.Gallery
	json.NewDecoder(*body).Decode(&model)

	var user models.User
	db.DB.Where("user_name = ?", username).First(&user)
	model.UserID = user.ID
	db.DB.Create(&model)
	return model, nil
}
func (service GalleryService) UpdateGallery(id string, username string, body *io.ReadCloser) (models.Gallery, error) {
	var model models.Gallery
	err := db.DB.
		Joins("INNER JOIN users u ON u.id = galleries.user_id").
		Where("u.user_name = ?", username).
		Find(&model).Error
	if err != nil {
		return model, errors.New("the user: " + username + " doesn'n have this project")
	}
	json.NewDecoder(*body).Decode(&model)
	db.DB.Save(&model)
	return model, nil
}
func (service GalleryService) DeleteGallery(id string, username string) (models.Gallery, error) {
	var model models.Gallery
	err := db.DB.
		Joins("INNER JOIN users u ON u.id = galleries.user_id").
		Where("u.user_name = ?", username).
		Find(&model).Error
	if err != nil {
		return model, errors.New("the user: " + username + " doesn'n have this course")
	}
	db.DB.Delete(&model)
	return model, nil
}
