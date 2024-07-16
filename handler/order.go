package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/enyasantos/go-async-order-system/models"
	"gorm.io/gorm"
)

type OrderGet struct{}

func (handler *OrderGet) ServerHTTP(w http.ResponseWriter, r *http.Request) {

}

func ShowOrderPriceHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	orderIdStr := parts[len(parts)-1]

	var order *models.Order
	orderCode, err := strconv.ParseInt(orderIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid orderId", http.StatusBadRequest)
		return
	}

	if err := db.Preload("Items").Where("order_code = ?", orderCode).First(&order).Error; err != nil {
		log.Printf("Failed to query order: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(w, http.StatusNotFound, "Order not found")
			return
		}
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	totalPrice := order.TotalPrice()
	data := struct {
		OrderCode  int64   `json:"order_code"`
		TotalPrice float64 `json:"total_price"`
	}{
		OrderCode:  orderCode,
		TotalPrice: totalPrice,
	}

	responseJSON(w, http.StatusOK, data)
}

func IndexOrdersHandler(w http.ResponseWriter, r *http.Request) {
	var order []*models.Order

	if err := db.Preload("Items").Find(&order).Error; err != nil {
		log.Printf("Failed to query orders: %v", err)
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJSON(w, http.StatusOK, order)
}
