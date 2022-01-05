// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"time"
)

// Cart is an object representing the database table.
type Cart struct {
	ID           int32     `gorm:"column:id"`
	CartID       string    `gorm:"column:cart_id"`
	SkuID        string    `gorm:"column:sku_id"`
	CustomerID   string    `gorm:"column:customer_id"`
	DiscountCode string    `gorm:"column:discount_code"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
	DeletedAt    time.Time `gorm:"column:deleted_at"`
}
