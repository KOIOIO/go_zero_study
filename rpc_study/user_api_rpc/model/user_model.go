package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:32" json:"username"`
	Password string `gorm:"size:32" json:"password"`
}
