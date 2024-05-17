package models

import (
	"gorm.io/gorm"
)

type Organization struct {
    gorm.Model
    Name  string
    Users []User `gorm:"many2many:user_organizations;"`
}
