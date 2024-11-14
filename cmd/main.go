package main

import (
	"github.com/VaheMuradyan/BookRentalSystem/db"
	"github.com/VaheMuradyan/BookRentalSystem/internal/handlers"
	"github.com/gin-gonic/gin"
)

func init() {
	db.ConnectDatabase()
	db.DBMigrate()
}

func main() {
	r := gin.Default()

	userHandler := handlers.NewHandler(db.DB)

	r.POST("/signup", userHandler.Signup)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
