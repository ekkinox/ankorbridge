package service

import (
	"context"
	"github.com/ankorstore/yokai/trace"
	"github.com/ekkinox/ankorbridge/internal/product"
	"go.opentelemetry.io/otel/attribute"
)

type Deleter struct {
	repository *product.Repository
}

func NewDeleter(repository *product.Repository) *Deleter {
	return &Deleter{
		repository: repository,
	}
}

func (u *Deleter) Delete(ctx context.Context, productID int) error {
	ctx, span := trace.CtxTracer(ctx).Start(ctx, "Deleting a product")
	span.SetAttributes(attribute.Int("product_id", productID))
	defer span.End()

	_, err := u.repository.Delete(ctx, product.DeleteProductParams{
		ID: productID,
	})
	if err != nil {
		return err
	}

	return nil
}
