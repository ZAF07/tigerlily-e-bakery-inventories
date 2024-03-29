// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	// "strconv"

	// "github.com/friendsofgo/errors"
		"github.com/volatiletech/null/v8"
	// "github.com/volatiletech/sqlboiler/v4/boil"
	// "github.com/volatiletech/sqlboiler/v4/queries"
	// "github.com/volatiletech/sqlboiler/v4/queries/qm"
	// "github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	// "github.com/volatiletech/sqlboiler/v4/types"
	// "github.com/volatiletech/strmangle"
)

// Skus is an object representing the database table.
type Skus struct {
	ID          int32             `gorm:"column:id"`
	SkuID       string            `gorm:"column:sku_id"`
	Name        string            `gorm:"column:name"`
	Price       float64 					`gorm:"column:price"`
	Type        string            `gorm:"column:type"`
	Description string            `gorm:"column:description"`
	ImageURL    string            `gorm:"column:image_url"`
	Quantity    int32          `gorm:"column:quantity" json:"quan"`
	CreatedAt   null.Time         `gorm:"column:created_at" `
	UpdatedAt   null.Time         `gorm:"column:updated_at"`
	DeletedAt   null.Time         `gorm:"column:deleted_at"`
}
