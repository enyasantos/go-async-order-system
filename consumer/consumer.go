package consumer

import (
	"github.com/enyasantos/go-async-order-system/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeConsumer() {
	db = config.GetPostgres()
}
