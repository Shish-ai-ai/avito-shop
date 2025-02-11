package routes

import (
	"avito-shop/internal/services"
	"avito-shop/internal/transport"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authService *services.AuthService) *gin.Engine {
	router := gin.Default()

	router.POST("/api/auth", transport.AuthHandler(authService))

	return router
}
