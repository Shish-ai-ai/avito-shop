package routes

import (
	"avito-shop/internal/middleware"
	"avito-shop/internal/services"
	"avito-shop/internal/transport"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authService *services.AuthService, infoService *services.InfoService, purchaseService *services.PurchaseService) *gin.Engine {
	router := gin.Default()

	router.POST("/api/auth", transport.AuthHandler(authService))

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/info", transport.InfoHandler(infoService))

		protected.GET("/buy/:item", transport.BuyItemHandler(purchaseService))
	}

	return router
}
