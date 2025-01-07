package product

import (
	"github.com/ekkinox/ankorbridge/internal/product/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ListProductsHandler struct {
	service *service.Fetcher
}

func NewListProductsHandler(service *service.Fetcher) *ListProductsHandler {
	return &ListProductsHandler{
		service: service,
	}
}

func (h *ListProductsHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		products, err := h.service.FindAll(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to list products")
		}

		return c.JSON(http.StatusOK, products)
	}
}
