package models

type Order struct {
	OrderId string
	ProductName string
	SkuId int
	CustomerId int
	Price int64
	Discount int
}