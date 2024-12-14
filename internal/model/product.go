// internal/model/product.go
package model

import (
	"time"
)

type Product struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Status      int8      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SeckillActivity struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	ProductID    uint      `json:"product_id"`
	SeckillPrice float64   `json:"seckill_price"`
	SeckillStock int       `json:"seckill_stock"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Status       int8      `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type SeckillOrder struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	UserID     uint      `json:"user_id"`
	ActivityID uint      `json:"activity_id"`
	ProductID  uint      `json:"product_id"`
	OrderNo    string    `json:"order_no"`
	Price      float64   `json:"price"`
	Status     int8      `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
