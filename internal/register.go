package internal

import (
	"github.com/ekkinox/ankorbridge/internal/product"
	"github.com/ekkinox/ankorbridge/internal/product/service"
	"go.uber.org/fx"
)

// Register is used to register the application dependencies.
func Register() fx.Option {
	return fx.Options(
		fx.Provide(
			product.NewRepository,
			service.NewFetcher,
			service.NewCreator,
			service.NewUpdater,
			service.NewDeleter,
		),
	)
}
