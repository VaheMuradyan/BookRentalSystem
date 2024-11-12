package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey"`
	Name     string `gorm:"type:text"`
	Author   string `gorm:"type:text"`
	ClientID uint64
	Client   Client  `gorm:"foreignKey:ClientEmail"`
	Genres   []Genre `gorm:"many2many:book_genre;"`
}
