package router

import (
	"net/http"

	"github.com/enyasantos/go-async-order-system/handler"
)

func initializeRoutes(mx *http.ServeMux) {
	handler.InitializeHandler()

	mx.HandleFunc("GET /orders/{orderCode}", handler.ShowOrderPriceHandler)
	mx.HandleFunc("GET /orders", handler.IndexOrdersHandler)
}
