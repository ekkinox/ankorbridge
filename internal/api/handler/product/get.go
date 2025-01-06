package product

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ekkinox/ankorbridge/internal/product"
	"github.com/labstack/echo/v4"
)

type GetProductHandler struct {
	service *product.ProductService
}

func NewGetProductHandler(service *product.ProductService) *GetProductHandler {
	return &GetProductHandler{
		service: service,
	}
}

func (h *GetProductHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		productID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid product ID")
		}

		p, err := h.service.Find(ctx, productID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("product ID %d does not exist", productID))
			}
		}

		return c.JSON(http.StatusOK, p)
	}
}
