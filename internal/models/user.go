package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"size:100;not null;unique"`
	Email string `gorm:"size:100;not null;unique"`
}
