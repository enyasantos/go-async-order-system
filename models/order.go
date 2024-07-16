package models

import (
	"math"

	"github.com/google/uuid"
)

type Order struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	OrderCode    int64     `json:"order_code"`
	CustomerCode string    `json:"customer_code"`
	Items        []Item    `gorm:"foreignkey:OrderID"`
}

func (order Order) TotalPrice() float64 {
	total := 0.0
	for _, item := range order.Items {
		total += float64(item.Quantity) * item.Price
	}
	return math.Round(total*100) / 100
}
