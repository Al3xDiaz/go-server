package services

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
)

type CourseService struct{}

func (service CourseService) ListCourses(username string, limit int) ([]models.Course, error) {
	var user models.User
	err := db.DB.Where("user_name = ?", username).First(&user).Error
	if err != nil || user.ID == 0 {
		return user.Courses, errors.New("user not found")
	}
	db.DB.Limit(limit).Model(&user).Association("Courses").Find(&user.Courses)
	return user.Courses, nil
}
func (service CourseService) CreateCourse(username string, body *io.ReadCloser) (models.Course, error) {
	var model models.Course
	json.NewDecoder(*body).Decode(&model)

	var user models.User
	db.DB.Where("user_name = ?", username).First(&user)
	model.UserID = user.ID
	db.DB.Create(&model)
	return model, nil
}
func (service CourseService) CreateCourses(username string, body *io.ReadCloser) error {
	var model []*models.Course
	json.NewDecoder(*body).Decode(&model)

	var user models.User
	db.DB.Where("user_name = ?", username).First(&user)
	for _, achievements := range model {
		achievements.UserID = user.ID
	}
	err := db.DB.Create(&model)
	if err != nil {
		return err.Error
	}
	return nil
}
func (service CourseService) UpdateCourse(id string, username string, body *io.ReadCloser) (models.Course, error) {
	var model models.Course
	err := db.DB.
		Joins("INNER JOIN users u ON u.id = courses.user_id").
		Where("courses.id = ? and u.user_name = ?", id, username).
		First(&model).Error
	if err != nil {
		return model, errors.New("the user: " + username + " doesn'n have this course")
	}
	json.NewDecoder(*body).Decode(&model)
	db.DB.Save(&model)
	return model, nil
}
func (service CourseService) DeleteCourse(id string, username string) (models.Course, error) {
	var model models.Course
	err := db.DB.
		Joins("INNER JOIN users u ON u.id = courses.user_id").
		Where("courses.id = ? and u.user_name = ?", id, username).
		First(&model).Error
	if err != nil {
		return model, errors.New("the user: " + username + " doesn'n have this course")
	}
	db.DB.Delete(&model)
	return model, nil
}
