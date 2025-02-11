package main

import (
	"avito-shop/internal/database"
	"avito-shop/internal/routes"
	"avito-shop/internal/services"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	godotenv.Load(".env")

	if err := database.InitDB(); err != nil {
		log.Fatal("failed to connect to the database: ", err)
	}

	authService := services.NewAuthService(database.DB)
	infoService := services.NewInfoService(database.DB)
	purchaseService := services.NewPurchaseService(database.DB)

	router := routes.SetupRouter(authService, infoService, purchaseService)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
