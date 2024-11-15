package db

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Title     string `gorm:"unique"`
	AuthorID  string
	Available bool
	Price     float64
	Carts     []Cart `gorm:"many2many:cart_books;"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Cart     Cart
}

type Author struct {
	gorm.Model
	Name string
}

type Cart struct {
	gorm.Model
	UserId uint
	Books  []Book `gorm:"many2many:cart_books;"`
}

type Order struct {
	gorm.Model
	UserId     uint
	BookId     uint
	OrderDate  time.Time
	ReturnData time.Time
}
