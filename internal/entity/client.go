package entity

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name  string `gorm:"type:text"`
	Email string `gorm:"type:text,unique"`
	Books []Book
}

type CreateClientReq struct {
	Name     string
	Email    string `gorm:"unique"`
	Password string
}
