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

	var purchases []struct {
		ID      uint   `json:"id"`
		MerchID uint   `json:"merch_id"`
		Type    string `json:"type"`
		Price   int    `json:"price"`
		Amount  int    `json:"amount"`
	}
	if err := s.DB.Table("purchases").
		Select("purchases.id, purchases.merch_id, merch.type, merch.price, purchases.amount").
		Joins("JOIN merch ON purchases.merch_id = merch.id").
		Where("purchases.user_id = ?", userID).
		Scan(&purchases).Error; err != nil {
		return nil, err
	}

	var sentOperations []struct {
		ID       uint `json:"id"`
		FromUser uint `json:"from_user"`
		ToUser   uint `json:"to_user"`
		Amount   int  `json:"amount"`
	}
	if err := s.DB.Table("operations").
		Select("id, from_user, to_user, amount").
		Where("from_user = ?", userID).
		Scan(&sentOperations).Error; err != nil {
		return nil, err
	}

	var receivedOperations []struct {
		ID       uint `json:"id"`
		FromUser uint `json:"from_user"`
		ToUser   uint `json:"to_user"`
		Amount   int  `json:"amount"`
	}
	if err := s.DB.Table("operations").
		Select("id, from_user, to_user, amount").
		Where("to_user = ?", userID).
		Scan(&receivedOperations).Error; err != nil {
		return nil, err
	}

	result := map[string]interface{}{
		"balance":             user.Balance,
		"purchases":           purchases,
		"sent_operations":     sentOperations,
		"received_operations": receivedOperations,
	}

	return result, nil
}
