package service

import (
	pb "pay/api/pay/v1"
	"pay/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPayService)

type PayService struct {
	pb.UnimplementedPayServiceServer

	log *log.Helper

	product *biz.ProductUsecase
}
