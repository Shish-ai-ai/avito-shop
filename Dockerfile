FROM golang:1.22

WORKDIR /app

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o avito-shop ./cmd/main.go

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./avito-shop"]