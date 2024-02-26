package services

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
)

type AchievementsHistoryService struct{}

func (service AchievementsHistoryService) ListAchievements(username string) ([]models.AchievementsHistory, error) {
	var model []models.AchievementsHistory
	err := db.DB.
		Order("year desc").
		Joins("INNER JOIN users u ON u.id = achievements_histories.user_id").
		Where("u.user_name = ?", username).
		Find(&model).Error
	if err != nil {
		return model, errors.New("has ocurrent problem get data from database")
	}
	return model, nil
}
func (service AchievementsHistoryService) CreateAchievement(username string, body *io.ReadCloser) (models.AchievementsHistory, error) {
	var model models.AchievementsHistory
	json.NewDecoder(*body).Decode(&model)

	var user models.User
	db.DB.Where("user_name = ?", username).First(&user)
	model.UserID = user.ID
	db.DB.Create(&model)
	return model, nil
}
func (service AchievementsHistoryService) CreateAchievements(username string, body *io.ReadCloser) error {
	var model []*models.AchievementsHistory
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
func (service AchievementsHistoryService) UpdateAchievement(id string, username string, body *io.ReadCloser) (models.AchievementsHistory, error) {
	var model models.AchievementsHistory
	err := db.DB.
		Joins("INNER JOIN users u ON u.id = achievements_histories.user_id").
		Where("achievements_histories.id = ? and u.user_name = ?", id, username).
		First(&model).Error
	if err != nil {
		return model, errors.New("the user: " + username + " doesn'n have this achievements")
	}
	json.NewDecoder(*body).Decode(&model)
	db.DB.Save(&model)
	return model, nil
}
func (service AchievementsHistoryService) DeleteAchievement(id string, username string) (models.AchievementsHistory, error) {
	var model models.AchievementsHistory
	err := db.DB.
		Joins("INNER JOIN users u ON u.id = achievements_histories.user_id").
		Where("achievements_histories.id = ? and u.user_name = ?", id, username).
		First(&model).Error
	if err != nil {
		return model, errors.New("the user: " + username + " doesn'n have this achievements")
	}
	db.DB.Delete(&model)
	return model, nil
}
