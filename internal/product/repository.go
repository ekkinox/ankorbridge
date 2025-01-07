package product

import (
	"context"
	"database/sql"

	"github.com/huandu/go-sqlbuilder"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Find(ctx context.Context, id int) (Product, error) {
	var product Product

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("p.id", "p.name")
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

func (r *Repository) FindAll(ctx context.Context) ([]Product, error) {
	var products []Product

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("p.id", "p.name")
	sb.From("products p")

	query, args := sb.Build()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

type CreateProductParams struct {
	Name string
}

func (r *Repository) Create(ctx context.Context, params CreateProductParams) (Product, error) {
	var product Product

	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("products")
	ib.Cols("name")
	ib.Values(params.Name)

	query, args := ib.Build()

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return product, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return product, err
	}

	product.ID = int(id)
	product.Name = params.Name

	return product, nil
}

type UpdateProductParams struct {
	ID   int
	Name string
}

func (r *Repository) Update(ctx context.Context, params UpdateProductParams) (Product, error) {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("products")
	ub.Set(ub.Assign("name", params.Name))
	ub.Where(ub.Equal("id", params.ID))

	query, args := ub.Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return Product{}, err
	}

	product := Product{
		ID:   params.ID,
		Name: params.Name,
	}

	return product, nil
}

type DeleteProductParams struct {
	ID int
}

func (r *Repository) Delete(ctx context.Context, params DeleteProductParams) (int, error) {
	db := sqlbuilder.NewDeleteBuilder()
	db.DeleteFrom("products")
	db.Where(db.Equal("id", params.ID))

	query, args := db.Build()

	res, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	return int(rowsAffected), err
}
