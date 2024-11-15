package handlers

import (
	"github.com/VaheMuradyan/BookRentalSystem/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type OrderHandler struct {
	db *gorm.DB
}

func NewOrderHandler(db *gorm.DB) *OrderHandler {
	return &OrderHandler{
		db: db,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	userID := c.Param("userID")
	var cart db.Cart
	if result := h.db.Preload("Books").Where("user_id = ?", userID).First(&cart); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
		return
	}

	for _, book := range cart.Books {
		order := db.Order{
			UserId:    cart.UserId,
			BookId:    book.ID,
			OrderDate: time.Now(),
		}

		h.db.Create(&order)
	}

	if err := h.db.Model(&cart).Association("Books").Clear(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Cant create order or cant cleared card"})
		return
	}
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	orderID := c.Param("orderID")

	var order db.Order
	if result := h.db.First(&order, orderID); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"order": order})
}

func (h *OrderHandler) GetOrdersFromUser(c *gin.Context) {
	userID := c.Param("userID")

	var orders []db.Order
	if result := h.db.Where("user_id = ?", userID).Find(&orders); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"orders": orders})
}

func (h *OrderHandler) ReturnBook(c *gin.Context) {
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
