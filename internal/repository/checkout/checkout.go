package checkout

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tigerlily-e-bakery-server/internal/models"
	"github.com/tigerlily-e-bakery-server/internal/pkg/logger"
	rpc "github.com/tigerlily-e-bakery-server/internal/pkg/protos"
)

type CheckoutRepo struct {
	db *gorm.DB
	logs logger.Logger
}

func NewCheckoutRepo(DB *gorm.DB) *CheckoutRepo {
	return&CheckoutRepo{
		db: DB,
		logs: *logger.NewLogger(),
	}
}

func (repo CheckoutRepo) CreateNewOrder(checkoutItems []*rpc.Checkout) (success bool, err error) {

	// RUN A TRANSACTION FOR CREATION, IF FAIL, WILL FALLBACK
	repo.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range checkoutItems {
			orderItem := &models.Order{
				DiscountCode: item.DiscountCode,
				OrderID: item.OrderId,
				CustomerID: item.CustomerId,
				SkuID: item.SkuId,
			}
	
			if err := tx.Debug().Omit("DeletedAt").Create(&orderItem).Error; err != nil {
				repo.logs.WarnLogger.Printf("[REPO] Error batch creating order items : %+v", err)
				success = false
				return err
			}
		}
		success = true
		return nil
	})
	
	return
}