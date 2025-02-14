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

func (s *InfoService) GetUserInfo(userID uint) (*models.InfoResponse, error) {
	var user models.User
	if err := s.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}

	var purchases []models.PurchaseResponse
	if err := s.DB.Table("purchases").
		Select("purchases.id, purchases.merch_id, merch.type, merch.price, purchases.amount").
		Joins("JOIN merch ON purchases.merch_id = merch.id").
		Where("purchases.user_id = ?", userID).
		Scan(&purchases).Error; err != nil {
		return nil, err
	}

	var sentOperations []models.OperationResponse
	if err := s.DB.Table("operations").
		Select("id, from_user AS from, to_user AS to, amount").
		Where("from_user = ?", userID).
		Scan(&sentOperations).Error; err != nil {
		return nil, err
	}

	var receivedOperations []models.OperationResponse
	if err := s.DB.Table("operations").
		Select("id, from_user AS from, to_user AS to, amount").
		Where("to_user = ?", userID).
		Scan(&receivedOperations).Error; err != nil {
		return nil, err
	}

	return &models.InfoResponse{
		Balance:            user.Balance,
		Purchases:          purchases,
		SentOperations:     sentOperations,
		ReceivedOperations: receivedOperations,
	}, nil
}
