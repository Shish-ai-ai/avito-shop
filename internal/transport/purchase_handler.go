package transport

import (
	"avito-shop/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BuyItemHandler(purchaseService *services.PurchaseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized"})
			return
		}

		item := c.Param("item")
		if item == "" {
			c.JSON(http.StatusBadRequest, gin.H{"errors": "Item name is required"})
			return
		}

		err := purchaseService.BuyItem(userID.(uint), item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Item purchased successfully"})
	}
}
