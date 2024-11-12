package Database

import (
	"github.com/VaheMuradyan/BookRentalSystem/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func ConnectDatabase() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
		},
	)

	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/books?charset=utf8&parseTime=true"), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic("Failed to connect to databse!")
	}

	DB = database
}

func DBMigrate() {
	if err := DB.AutoMigrate(&entity.Book{}, &entity.Client{}, &entity.Genre{}); err != nil {
		log.Fatal(err)
	}
}
