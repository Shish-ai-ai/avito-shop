package tests

import (
	"avito-shop/internal/models"
	"avito-shop/internal/services"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestAuthService_Authenticate(t *testing.T) {
	db := setupTestDB()
	authService := services.NewAuthService(db)

	password := "password"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Ошибка хэширования пароля: %v", err)
	}

	db.Create(&models.User{Name: "testuser", Password: string(hashedPassword), Balance: 1000})

	_, err = authService.Authenticate("testuser", password)
	if err != nil {
		t.Errorf("Ошибка аутентификации: %v", err)
	}
}
