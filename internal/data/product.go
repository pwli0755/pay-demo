package data

import (
	"context"
	"pay/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type productRepo struct {
	data *Data
	log  *log.Helper
}

// NewProductRepo .
func NewProductRepo(data *Data, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (p *productRepo) ListProduct(ctx context.Context) ([]*biz.Product, error) {
	ps, err := p.data.db.Product.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Product, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Product{
			ID:        p.ID,
			Title:     p.Title,
			Price:     p.Price,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return rv, nil
}
