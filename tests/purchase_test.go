package tests

import (
	"avito-shop/internal/models"
	"avito-shop/internal/services"
	"testing"
)

func TestPurchaseService_BuyItem(t *testing.T) {
	db := setupTestDB()
	purchaseService := services.NewPurchaseService(db)

	db.Create(&models.User{ID: 1, Name: "buyer", Balance: 1000})
	db.Create(&models.Merch{ID: 1, Type: "item1", Price: 500})

	err := purchaseService.BuyItem(1, "item1")
	if err != nil {
		t.Errorf("Ошибка покупки товара: %v", err)
	}
}
