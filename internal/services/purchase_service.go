package services

import (
	"avito-shop/internal/models"
	"fmt"
	"gorm.io/gorm"
)

type PurchaseService struct {
	DB *gorm.DB
}

func NewPurchaseService(db *gorm.DB) *PurchaseService {
	return &PurchaseService{DB: db}
}

func (s *PurchaseService) BuyItem(userID uint, itemName string) error {
	var merch models.Merch
	if err := s.DB.Where("type = ?", itemName).First(&merch).Error; err != nil {
		return fmt.Errorf("item not found: %v", err)
	}

	var user models.User
	if err := s.DB.First(&user, userID).Error; err != nil {
		return fmt.Errorf("failed to fetch user balance: %v", err)
	}

	if user.Balance < merch.Price {
		return fmt.Errorf("insufficient balance")
	}

	if err := s.DB.Model(&user).Update("balance", gorm.Expr("balance - ?", merch.Price)).Error; err != nil {
		return fmt.Errorf("failed to update user balance: %v", err)
	}

	purchase := models.Purchase{
		UserID:  userID,
		MerchID: merch.ID,
		Amount:  1,
	}
	if err := s.DB.Create(&purchase).Error; err != nil {
		return fmt.Errorf("failed to record purchase: %v", err)
	}

	return nil
}
