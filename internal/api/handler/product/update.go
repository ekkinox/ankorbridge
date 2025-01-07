package product

import (
	"fmt"
	"github.com/ekkinox/ankorbridge/internal/product/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UpdateProductHandler struct {
	fetcher *service.Fetcher
	updater *service.Updater
}

func NewUpdateProductHandler(fetcher *service.Fetcher, updater *service.Updater) *UpdateProductHandler {
	return &UpdateProductHandler{
		fetcher: fetcher,
		updater: updater,
	}
}

type UpdateProductParams struct {
	Name string `json:"name"`
}

func (p *UpdateProductParams) Validate() error {
	if len(p.Name) < 2 {
		return fmt.Errorf("product name is too short")
	}

	if len(p.Name) > 100 {
		return fmt.Errorf("product name is too long")
	}

	return nil
}

func (h *UpdateProductHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		params := new(UpdateProductParams)
		if err := c.Bind(params); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid request body parameters %v", err))
		}

		if err := params.Validate(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		productID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid product ID")
		}

		p, err := h.fetcher.Find(ctx, productID)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("product with ID %d does not exist", productID))
		}

		p, err = h.updater.Update(ctx, p, service.UpdateProductParams{
			Name: params.Name,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to update product: %v", err))
		}

		return c.JSON(http.StatusOK, p)
	}
}
