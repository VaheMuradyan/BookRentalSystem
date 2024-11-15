package handlers

import (
	"github.com/VaheMuradyan/BookRentalSystem/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type BookHandler struct {
	db *gorm.DB
}

func NewBookHandler(db *gorm.DB) *BookHandler {
	return &BookHandler{
		db: db,
	}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book db.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Create(&book).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, book)
}

func (h *BookHandler) GetBook(c *gin.Context) {
	var book db.Book
	id := c.Param("id")
	if result := h.db.First(&book, id); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	var books []db.Book
	if result := h.db.Find(&books); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	var book db.Book
	id := c.Param("id")
	if result := h.db.First(&book, id); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.db.Save(&book)
	c.IndentedJSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if result := h.db.Delete(&db.Book{}, id); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
