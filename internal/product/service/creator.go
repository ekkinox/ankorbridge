package service

import (
	"context"
	"github.com/ankorstore/yokai/trace"
	"github.com/ekkinox/ankorbridge/internal/product"
	"go.opentelemetry.io/otel/attribute"
)

type Creator struct {
	repository *product.Repository
}

func NewCreator(repository *product.Repository) *Creator {
	return &Creator{
		repository: repository,
	}
}

type CreateProductParams struct {
	Name string
}

func (c *Creator) Create(ctx context.Context, params CreateProductParams) (product.Product, error) {
	ctx, span := trace.CtxTracer(ctx).Start(ctx, "Creating a new product")
	defer span.End()

	model, err := c.repository.Create(ctx, product.CreateProductParams{
		Name: params.Name,
	})
	if err != nil {
		return product.Product{}, err
	}
	span.SetAttributes(attribute.Int("product_id", model.ID))

	return model, nil
}
