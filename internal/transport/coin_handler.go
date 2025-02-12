package transport

import (
	"avito-shop/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendCoinRequest struct {
	ToUser uint `json:"to_user" binding:"required"`
	Amount int  `json:"amount" binding:"required,min=1"`
}

func SendCoinHandler(coinService *services.CoinService) gin.HandlerFunc {
	return func(c *gin.Context) {
		fromUserID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized"})
			return
		}
		fromID, ok := fromUserID.(uint)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Invalid user ID"})
			return
		}

		var req SendCoinRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid request"})
			return
		}

		if err := coinService.SendCoins(fromID, req.ToUser, req.Amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Coins sent successfully"})
	}
}
