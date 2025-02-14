package tests

import (
	"avito-shop/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Merch{}, &models.Purchase{}, &models.Operation{})

	db.Exec("DELETE FROM operations")
	db.Exec("DELETE FROM purchases")
	db.Exec("DELETE FROM merch")
	db.Exec("DELETE FROM users")

	return db
}
