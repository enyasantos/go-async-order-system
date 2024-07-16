package models

import "github.com/google/uuid"

type Item struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	OrderID  uuid.UUID `gorm:"type:uuid;index"`
	Product  string    `json:"product"`
	Quantity int       `json:"quantity"`
	Price    float64   `json:"price"`
}
