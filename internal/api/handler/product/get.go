package product

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ekkinox/ankorbridge/internal/product/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GetProductHandler struct {
	service *service.Fetcher
}

func NewGetProductHandler(service *service.Fetcher) *GetProductHandler {
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
				return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("product with ID %d does not exist", productID))
			}
		}

		return c.JSON(http.StatusOK, p)
	}
}
