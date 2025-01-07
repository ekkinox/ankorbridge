package service

import (
	"context"
	"github.com/ankorstore/yokai/trace"
	"github.com/ekkinox/ankorbridge/internal/product"
	"go.opentelemetry.io/otel/attribute"
)

type Updater struct {
	repository *product.Repository
}

func NewUpdater(repository *product.Repository) *Updater {
	return &Updater{
		repository: repository,
	}
}

type UpdateProductParams struct {
	Name string
}

func (u *Updater) Update(ctx context.Context, p product.Product, params UpdateProductParams) (product.Product, error) {
	ctx, span := trace.CtxTracer(ctx).Start(ctx, "Updating a product")
	span.SetAttributes(attribute.Int("product_id", p.ID))
	defer span.End()

	model, err := u.repository.Update(ctx, product.UpdateProductParams{
		ID:   p.ID,
		Name: params.Name,
	})
	if err != nil {
		return product.Product{}, err
	}

	return model, nil
}
