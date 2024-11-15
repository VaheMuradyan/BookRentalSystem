package handlers

import (
	"github.com/VaheMuradyan/BookRentalSystem/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type AuthorHandler struct {
	db *gorm.DB
}

func NewAuthorHandler(db *gorm.DB) *AuthorHandler {
	return &AuthorHandler{
		db: db,
	}
}

func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var author db.Author
	if err := c.BindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Create(&author).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	c.IndentedJSON(http.StatusCreated, author)
}

func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	var author db.Author
	id := c.Param("id")
	if result := h.db.First(&author, id).Error; result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": result.Error()})
		return
	}

	if err := c.BindJSON(&author); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	h.db.Save(&author)
	c.IndentedJSON(http.StatusOK, author)

}

func (h *AuthorHandler) GetAuthors(c *gin.Context) {
	var authors []db.Author
	if err := h.db.Find(&authors).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, authors)
}

func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	var author db.Author
	if err := h.db.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}
	if err := h.db.Delete(&author).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, author)
}

func (h *AuthorHandler) SearchAuthor(c *gin.Context) {
	var author db.Author
	if err := h.db.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, author)
}
