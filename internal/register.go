package internal

import (
	"github.com/ekkinox/ankorbridge/internal/product"
	"go.uber.org/fx"
)

// Register is used to register the application dependencies.
func Register() fx.Option {
	return fx.Options(
		fx.Provide(
			product.NewProductRepository,
			product.NewProductService,
		),
	)
}
