package transport

import (
	"avito-shop/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthHandler(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var authRequest struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&authRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid request"})
			return
		}

		token, err := authService.Authenticate(authRequest.Name, authRequest.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
