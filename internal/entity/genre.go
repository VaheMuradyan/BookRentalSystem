package entity

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Name  string `gorm:"type:text"`
	Books []Book `gorm:"many2many:book_genre;"`
}
