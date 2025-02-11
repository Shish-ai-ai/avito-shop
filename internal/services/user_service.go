package services

import (
	"avito-shop/internal/models"
	"avito-shop/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

func (service *AuthService) Authenticate(name, password string) (string, error) {
	var user models.User

	err := service.DB.Where("name = ?", name).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {

			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				log.Println("Error hashing password:", err)
				return "", err
			}

			user = models.User{
				Name:     name,
				Password: string(hashedPassword),
			}
			if err := service.DB.Create(&user).Error; err != nil {
				log.Println("Error creating user:", err)
				return "", err
			}

			token, err := utils.GenerateJWT(user.ID, user.Name)
			if err != nil {
				log.Println("Error generating JWT:", err)
				return "", err
			}
			return token, nil
		} else {
			log.Println("Error finding user:", err)
			return "", err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Println("Invalid password")
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Name)
	if err != nil {
		log.Println("Error generating JWT:", err)
		return "", err
	}

	return token, nil
}
