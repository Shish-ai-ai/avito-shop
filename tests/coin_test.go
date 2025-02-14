package tests

import (
	"avito-shop/internal/models"
	"avito-shop/internal/services"
	"testing"
)

func TestCoinService_SendCoins(t *testing.T) {
	db := setupTestDB()
	coinService := services.NewCoinService(db)

	db.Create(&models.User{ID: 1, Name: "user1", Balance: 1000})
	db.Create(&models.User{ID: 2, Name: "user2", Balance: 1000})

	err := coinService.SendCoins(1, 2, 500)
	if err != nil {
		t.Errorf("Ошибка перевода монет: %v", err)
	}
}
