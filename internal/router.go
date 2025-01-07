package internal

import (
	"github.com/ankorstore/yokai/fxhttpserver"
	"github.com/ekkinox/ankorbridge/internal/api/handler/product"
	"go.uber.org/fx"
)

// Router is used to register the application HTTP middlewares and handlers.
func Router() fx.Option {
	return fx.Options(
		fxhttpserver.AsHandler("GET", "/products", product.NewListProductsHandler),
		fxhttpserver.AsHandler("POST", "/products", product.NewCreateProductHandler),
		fxhttpserver.AsHandler("GET", "/products/:id", product.NewGetProductHandler),
		fxhttpserver.AsHandler("PATCH", "/products/:id", product.NewUpdateProductHandler),
		fxhttpserver.AsHandler("DELETE", "/products/:id", product.NewDeleteProductHandler),
	)
}
