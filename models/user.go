package models

import (
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email    string `gorm:"unique"`
    Password string
    OTP      string
    Verified bool
    Organizations []Organization `gorm:"many2many:user_organizations;"`
}
