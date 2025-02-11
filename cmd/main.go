package main

import (
	"avito-shop/internal/database"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	godotenv.Load()
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
	}

	fmt.Println("База данных успешно инициализирована!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
