package services

import (
	"github.com/al3xdiaz/go-server/models"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (o UserService) GetUsers() (model []string, err error) {
	var users []string
	err = o.DB.Model(&models.User{}).Pluck("user_name", &users).Error
	return users, err
}
