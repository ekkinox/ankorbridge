package product

import (
	"context"
	"database/sql"

	"github.com/huandu/go-sqlbuilder"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Find(ctx context.Context, id int) (Product, error) {
	var product Product

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(
		"p.id",
		"p.name",
	)
	sb.From("products p")
	sb.Where(sb.Equal("p.id", id))
	sb.Limit(1)

	query, args := sb.Build()

	row := r.db.QueryRowContext(ctx, query, args...)

	err := row.Scan(
		&product.ID,
		&product.Name,
	)

	return product, err
}
