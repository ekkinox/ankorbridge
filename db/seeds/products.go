package seeds

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
)

const ProductsSeedName = "products"

type ProductsSeed struct{}

func NewProductsSeed() *ProductsSeed {
	return &ProductsSeed{}
}

func (s *ProductsSeed) Name() string {
	return ProductsSeedName
}

func (s *ProductsSeed) Run(ctx context.Context, db *sql.DB) error {
	ib := sqlbuilder.NewInsertBuilder()

	ib.InsertInto("products").
		Cols("id", "name").
		Values(1, "foo").
		Values(2, "bar")

	query, args := ib.Build()

	_, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to seed products: %w", err)
	}

	return nil
}
