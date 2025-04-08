package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `gorm:"size:32" json:"title"`
	Content string `gorm:"size:256" json:"content"`
	Writer  string `gorm:"size:32" json:"writer"`
}
