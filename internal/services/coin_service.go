package services

import (
	"avito-shop/internal/models"
	"errors"
	"gorm.io/gorm"
)

type CoinService struct {
	DB *gorm.DB
}

func NewCoinService(db *gorm.DB) *CoinService {
	return &CoinService{DB: db}
}

func (s *CoinService) SendCoins(fromUserID, toUserID uint, amount int) error {
	if fromUserID == toUserID {
		return errors.New("нельзя перевести монеты самому себе")
	}

	tx := s.DB.Begin()

	var sender models.User
	if err := tx.First(&sender, fromUserID).Error; err != nil {
		tx.Rollback()
		return errors.New("отправитель не найден")
	}

	if sender.Balance < amount {
		tx.Rollback()
		return errors.New("недостаточно средств")
	}

	var receiver models.User
	if err := tx.First(&receiver, toUserID).Error; err != nil {
		tx.Rollback()
		return errors.New("получатель не найден")
	}

	if err := tx.Model(&sender).Update("balance", sender.Balance-amount).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&receiver).Update("balance", receiver.Balance+amount).Error; err != nil {
		tx.Rollback()
		return err
	}

	operation := models.Operation{
		FromUser: fromUserID,
		ToUser:   toUserID,
		Amount:   amount,
	}
	if err := tx.Create(&operation).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
