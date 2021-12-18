package models

type Order struct {
	Id int
	ProductName string
	SkuId int
	CustomerId int
	Price int64
	Discount int
}