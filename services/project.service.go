package services

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
)

type ProjectService struct{}

func (service ProjectService) ListProjects(username string, limit int) ([]models.Project, error) {
	var model []models.Project
	err := db.DB.
		Order("start_date desc").
		Joins("INNER JOIN users u ON u.id = projects.user_id").
		Where("u.user_name = ?", username).
		Find(&model).Error
	if err != nil {
		return model, errors.New("has ocurrent problem get data from database")
	}
	return model, nil
}
func (service ProjectService) CreateProject(username string, body *io.ReadCloser) (models.Project, error) {
	var model models.Project
	json.NewDecoder(*body).Decode(&model)

	var user models.User
	db.DB.Where("user_name = ?", username).First(&user)
	model.UserID = user.ID
	db.DB.Create(&model)
	return model, nil
}
func (service ProjectService) UpdateProject(id string, username string, body *io.ReadCloser) (models.Project, error) {
	var model models.Project
	err := db.DB.
		Order("start_date").
		Joins("INNER JOIN users u ON u.id = projects.user_id").
		Where("u.user_name = ?", username).
		Find(&model).Error
	if err != nil {
		return model, errors.New("the user: " + username + " doesn'n have this project")
	}
	json.NewDecoder(*body).Decode(&model)
	db.DB.Save(&model)
	return model, nil
}
func (service ProjectService) DeleteProject(id string, username string) (models.Project, error) {
	var model models.Project
	err := db.DB.
		Order("start_date").
		Joins("INNER JOIN users u ON u.id = projects.user_id").
		Where("u.user_name = ?", username).
		Find(&model).Error
	if err != nil {
		return model, errors.New("the user: " + username + " doesn'n have this course")
	}
	db.DB.Delete(&model)
	return model, nil
}
