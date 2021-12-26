package checkout

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tigerlily-e-bakery-server/internal/pkg/logger"
	rpc "github.com/tigerlily-e-bakery-server/internal/pkg/protos"
)

type Service struct {
	db *gorm.DB
	logs logger.Logger
}

func NewCheckoutService(DB *gorm.DB) *Service {
	return&Service{
		db: DB,
		logs: *logger.NewLogger(),
	}
}

func (srv Service) Checkout(ctx context.Context, req *rpc.CheckoutReq) (resp *rpc.CheckoutResp, err error) {
	srv.logs.InfoLogger.Printf(" [SERVICE] Checkout service ran %+v", req)
	fmt.Printf("request : %+v", &req)
	resp = &rpc.CheckoutResp{
		Success: true,
	}
	return 
} 
