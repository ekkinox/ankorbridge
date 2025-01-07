package service

import (
	"context"
	"github.com/ankorstore/yokai/log"
	"github.com/ankorstore/yokai/trace"
	"github.com/ekkinox/ankorbridge/internal/product"
)

type Fetcher struct {
	repository *product.Repository
}

func NewFetcher(repository *product.Repository) *Fetcher {
	return &Fetcher{
		repository: repository,
	}
}

func (s *Fetcher) Find(ctx context.Context, id int) (product.Product, error) {
	ctx, span := trace.CtxTracer(ctx).Start(ctx, "Fetching a single product")
	defer span.End()

	log.CtxLogger(ctx).Info().Msg("Fetching a single product")

	return s.repository.Find(ctx, id)
}

func (s *Fetcher) FindAll(ctx context.Context) ([]product.Product, error) {
	ctx, span := trace.CtxTracer(ctx).Start(ctx, "Fetching a list of products")
	defer span.End()

	log.CtxLogger(ctx).Info().Msg("Fetching a list of products")

	return s.repository.FindAll(ctx)
}
