package handlers

import (
	"github.com/VaheMuradyan/BookRentalSystem/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type orderHandler struct {
	db *gorm.DB
}

func NewOrderHandler(db *gorm.DB) *orderHandler {
	return &orderHandler{
		db: db,
	}
}

func (h *orderHandler) CreateOrder(c *gin.Context) {
	userID := c.Param("userID")
	id, err := strconv.ParseUint(userID, 10, 32)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var cart db.Cart
	if result := h.db.Preload("Books").Where("user_id = ?", uint(id)).First(&cart); result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Cart not found"})
		return
	}

	for _, book := range cart.Books {
		order := db.Order{
			UserId:    uint(id),
			BookId:    book.ID,
			OrderDate: time.Now(),
		}

		if err := h.db.Create(&order).Error; err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Faild to create order"})
			return
		}
	}

	if err := h.db.Model(&cart).Association("Books").Clear(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Faild to clear card"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Order created successfully"})
}

func (h *orderHandler) GetOrder(c *gin.Context) {
	orderID := c.Param("orderID")

	var order db.Order
	if result := h.db.First(&order, orderID); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"order": order})
}

func (h *orderHandler) GetOrdersFromUser(c *gin.Context) {
	userID := c.Param("userID")

	var orders []db.Order
	if result := h.db.Where("user_id = ?", userID).Find(&orders); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"orders": orders})
}

func (h *orderHandler) ReturnBook(c *gin.Context) {
	orderID := c.Param("orderID")

	var order db.Order
	if result := h.db.First(&order, orderID); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		return
	}

	order.ReturnData = time.Now()
	h.db.Save(&order)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book returnd", "order": order})
}
