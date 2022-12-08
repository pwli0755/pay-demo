// Package biz ...
package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Product struct {
	ID        int64
	Title     string
	Price     uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductRepo interface {
	// db
	ListProduct(ctx context.Context) ([]*Product, error)
}

type ProductUsecase struct {
	repo ProductRepo
}

func NewProductUsecase(repo ProductRepo, logger log.Logger) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

func (uc *ProductUsecase) List(ctx context.Context) (ps []*Product, err error) {
	ps, err = uc.repo.ListProduct(ctx)
	if err != nil {
		return
	}
	return
}
