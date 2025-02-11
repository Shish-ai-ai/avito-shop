package handlers

import (
	"avito-shop/internal/models"
	"avito-shop/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterUserHandler - регистрация пользователя
func RegisterUserHandler(repo *repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if err := repo.CreateUser(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	}
}

// GetUserInfoHandler - получение информации о пользователе
func GetUserInfoHandler(repo *repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		user, err := repo.GetUserByID(userID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// TransferCoinsHandler - передача монет другому пользователю
func TransferCoinsHandler(repo *repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var transfer models.TransferRequest
		if err := c.ShouldBindJSON(&transfer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if err := repo.TransferCoins(transfer.FromUserID, transfer.ToUserID, transfer.Amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
	}
}

// PurchaseMerchHandler - покупка мерча
func PurchaseMerchHandler(repo *repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var purchase models.PurchaseRequest
		if err := c.ShouldBindJSON(&purchase); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		if err := repo.PurchaseMerch(purchase.UserID, purchase.MerchID, purchase.Amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Purchase successful"})
	}
}

// GetMerchListHandler - получение списка доступного мерча
func GetMerchListHandler(repo *repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		merchList, err := repo.GetMerchList()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve merch list"})
			return
		}
		c.JSON(http.StatusOK, merchList)
	}
}
