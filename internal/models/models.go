package models

// Пользователь
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
	Balance  int    `gorm:"not null;default:1000"`
}

// Мерч
type Merch struct {
	ID    uint   `gorm:"primaryKey"`
	Type  string `gorm:"not null"`
	Price int    `gorm:"not null"`
}

// Покупки
type Purchase struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint `gorm:"not null"`
	MerchID uint `gorm:"not null"`
	Amount  int  `gorm:"not null"`

	User  User  `gorm:"foreignKey:UserID"`
	Merch Merch `gorm:"foreignKey:MerchID"`
}

// Операции (переводы монет)
type Operation struct {
	ID       uint `gorm:"primaryKey"`
	FromUser uint `gorm:"not null"`
	ToUser   uint `gorm:"not null"`
	Amount   int  `gorm:"not null"`

	From User `gorm:"foreignKey:FromUser"`
	To   User `gorm:"foreignKey:ToUser"`
}
