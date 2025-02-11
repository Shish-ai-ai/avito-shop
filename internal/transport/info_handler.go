package transport

import (
	"avito-shop/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InfoHandler(infoService *services.InfoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized"})
			return
		}

		uid, ok := userID.(uint)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Invalid user ID"})
			return
		}

		info, err := infoService.GetUserInfo(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
			return
		}

		c.JSON(http.StatusOK, info)
	}
}
