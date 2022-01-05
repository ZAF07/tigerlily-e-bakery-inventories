package models

import "time"

type Order struct {
	ID int32 `gorm:"column:id"`
	OrderID string `gorm:"column:order_id"`
	SkuID string	`gorm:"column:sku_id"`
	CustomerID string	`gorm:"column:customer_id"`
	DiscountCode string	`gorm:"column:discount_code"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	// DeletedAt time.Time `gorm:"column:deleted_at"`
}