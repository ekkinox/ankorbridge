package product

import (
	"context"

	"github.com/ankorstore/yokai/log"
	"github.com/ankorstore/yokai/trace"
)

type ProductService struct {
	repository *ProductRepository
}

func NewProductService(repository *ProductRepository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) Find(ctx context.Context, id int) (Product, error) {
	ctx, span := trace.CtxTracer(ctx).Start(ctx, "ProductService::Find")
	defer span.End()

	log.CtxLogger(ctx).Info().Msg("ProductService::Find")

	return s.repository.Find(ctx, id)
}
