package repositories

import (
	"go-api-server/config"
	"go-api-server/models"
)

func CreateOrganization(org *models.Organization) error {
    return config.DB.Create(org).Error
}
