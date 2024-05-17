package repositories

import (
	"go-api-server/config"
	"go-api-server/models"
)

func CreateUser(user *models.User) error {
    return config.DB.Create(user).Error
}

func FindUserByEmail(email string) (models.User, error) {
    var user models.User
    err := config.DB.Where("email = ?", email).First(&user).Error
    return user, err
}

func UpdateUser(user *models.User) error {
    return config.DB.Save(user).Error
}
