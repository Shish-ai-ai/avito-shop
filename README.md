## Инструкция по запуску
### 1. Клонирование репозитория
```sh
$ git clone https://github.com/your-repo/avito-shop.git
$ cd avito-shop
```

### 2. Запуск с помощью Docker Compose
```sh
$ docker-compose up --build
```

После успешного запуска сервис будет доступен по адресу `http://localhost:8080`.

## API эндпоинты
### 1. Авторизация
**POST** `http://localhost:8080/api/auth`

#### Тело запроса:
```json
{
    "name": "test",
    "password": "password"
}
```
#### Ответ:
```json
{
    "token": "<JWT_TOKEN>"
}
```
Этот токен нужно передавать в заголовке `Authorization` для всех дальнейших запросов.

---

### 2. Покупка товара
**GET** `http://localhost:8080/api/buy/cup`

#### Заголовки:
```http
Authorization: Bearer <JWT_TOKEN>
```
#### Ответ:
```json
{
    "message": "Successfully purchased cup",
    "new_balance": 100
}
```

---

### 3. Получение информации о пользователе
**GET** `http://localhost:8080/api/info`

#### Заголовки:
```http
Authorization: Bearer <JWT_TOKEN>
```
#### Ответ:
```json
{
    "user_id": 1,
    "name": "test",
    "balance": 900,
    "purchased_items": ["cup"]
}
```

---

### 4. Передача монет другому пользователю
**POST** `http://localhost:8080/api/sendCoin`

#### Заголовки:
```http
Authorization: Bearer <JWT_TOKEN>
```
#### Тело запроса:
```json
{
    "to_user": 2,
    "amount": 100
}
```
#### Ответ:
```json
{
    "message": "Successfully sent 100 coins to user 2",
    "new_balance": 800
}
```
Важно, чтобы пользователь с user_id = 2 существовал.

