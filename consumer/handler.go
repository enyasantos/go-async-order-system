package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/enyasantos/go-async-order-system/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

type itemData struct {
	Product  string  `json:"produto"`
	Quantity int     `json:"quantidade"`
	Price    float64 `json:"preco"`
}

type orderData struct {
	OrderCode    int64      `json:"codigoPedido"`
	CustomerCode int        `json:"codigoCliente"`
	Items        []itemData `json:"itens"`
}

func consumeMessage(d amqp.Delivery) {
	var orderData orderData

	err := json.Unmarshal(d.Body, &orderData)
	if err != nil {
		log.Printf("Failed to unmarshal message body: %v", err)
	}

	var items []models.Item
	for _, item := range orderData.Items {
		items = append(items, models.Item{
			Product:  item.Product,
			Quantity: item.Quantity,
			Price:    item.Price,
		})
	}

	fmt.Printf("items: %v", items)

	order := models.Order{
		OrderCode:    orderData.OrderCode,
		CustomerCode: strconv.Itoa(orderData.CustomerCode),
		Items:        items,
	}

	if err := db.Create(&order).Error; err != nil {
		log.Printf("Failed to save order to database: %v", err)
	}
}
