package models

// Ответ на запрос информации о пользователе
type InfoResponse struct {
	Balance            int                 `json:"balance"`
	Purchases          []PurchaseResponse  `json:"purchases"`
	SentOperations     []OperationResponse `json:"sent_operations"`
	ReceivedOperations []OperationResponse `json:"received_operations"`
}

// Упрощенный ответ о покупке (без вложенного User)
type PurchaseResponse struct {
	ID      uint   `json:"id"`
	MerchID uint   `json:"merch_id"`
	Type    string `json:"type"`
	Price   int    `json:"price"`
	Amount  int    `json:"amount"`
}

// Упрощенный ответ о переводе монет (без вложенных пользователей)
type OperationResponse struct {
	ID     uint `json:"id"`
	From   uint `json:"from_user"`
	To     uint `json:"to_user"`
	Amount int  `json:"amount"`
}
