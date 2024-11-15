package handlers

import (
	"github.com/VaheMuradyan/BookRentalSystem/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type CartHandler struct {
	db *gorm.DB
}

func NewCartHandler(db *gorm.DB) *CartHandler {
	return &CartHandler{
		db: db,
	}
}

func (h *CartHandler) AddBookToCart(c *gin.Context) {
	var body struct {
		UserId uint
		BookId uint
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cart db.Cart
	if result := h.db.Where("user_id = ?", body.UserId).First(&cart); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
	}

	var book db.Book
	if result := h.db.First(&book, body.BookId); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
		return
	}

	if err := h.db.Model(&cart).Association("Books").Append(&book); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Can't add book to cart"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Book added to cart"})
}

func (h *CartHandler) RemoveBookFromCart(c *gin.Context) {
	var body struct {
		UserId uint
		BookId uint
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cart db.Cart
	if result := h.db.Where("user_id = ?", body.UserId).First(&cart); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
		return
	}

	var book db.Book
	if result := h.db.First(&book, body.BookId); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
		return
	}

	if err := h.db.Model(&cart).Association("Books").Delete(&book); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Can't remove book from cart"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book removed from cart"})
}

func (h *CartHandler) ClearCart(c *gin.Context) {
	userID := c.Param("userID")

	var cart db.Cart
	if result := h.db.Where("user_id = ?", userID).First(&cart); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
		return
	}

	if err := h.db.Model(&cart).Association("Carts").Clear(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Can't clear cart"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Carts cleared"})
}
