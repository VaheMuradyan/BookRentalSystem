package handlers

import (
	"github.com/VaheMuradyan/BookRentalSystem/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type UserHandler struct {
	db *gorm.DB
}

func NewHandler(d *gorm.DB) *UserHandler {
	return &UserHandler{
		db: d,
	}
}

func (h *UserHandler) Signup(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := db.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hash),
	}

	result := h.db.Create(&user)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	cart := db.Cart{
		UserId: user.ID,
	}

	if result := h.db.Create(&cart); result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"user": user})
}

func (h *UserHandler) Login(c *gin.Context) {

}

func (h *UserHandler) PlaceOrder(c *gin.Context) {

}
