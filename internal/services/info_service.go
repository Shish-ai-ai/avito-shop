package services

import (
	"avito-shop/internal/models"
	"gorm.io/gorm"
)

type InfoService struct {
	DB *gorm.DB
}

func NewInfoService(db *gorm.DB) *InfoService {
	return &InfoService{DB: db}
}

func (s *InfoService) GetUserInfo(userID uint) (map[string]interface{}, error) {
	var user models.User
	if err := s.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}

	var purchases []models.Purchase
	if err := s.DB.Where("user_id = ?", userID).Preload("Merch").Find(&purchases).Error; err != nil {
		return nil, err
	}

	var sentOperations []models.Operation
	if err := s.DB.Where("from_user = ?", userID).Find(&sentOperations).Error; err != nil {
		return nil, err
	}
	var receivedOperations []models.Operation
	if err := s.DB.Where("to_user = ?", userID).Find(&receivedOperations).Error; err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"balance":             user.Balance,
		"purchases":           purchases,
		"sent_operations":     sentOperations,
		"received_operations": receivedOperations,
	}

	return response, nil
}
