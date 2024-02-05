package services

import (
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
