package handlers

import (
	"github.com/VaheMuradyan/BookRentalSystem/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
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
	var body struct {
		Email    string
		Password string
	}

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Faild tou read body"})
		return
	}

	var user db.User
	h.db.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Invalid tou create token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", tokenString, 3600*24*30, "", "", false, true)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "all done"})
}

func (h *UserHandler) PlaceOrder(c *gin.Context) {

}
