package checkout_test

import (
	"testing"

	"github.com/tigerlily-e-bakery-server/internal/db"
	"github.com/tigerlily-e-bakery-server/internal/db/models"
)


func TestCreate(t *testing.T) {
	orderItems := models.Order{
		OrderId: "orderId",
		SkuId: "skuid",
		CustomerId: "customerId",
		DiscountCode: "discountcode",
	}
	db := db.NewDB()
	db.Create(orderItems)

}