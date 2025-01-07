package product

import (
	"fmt"
	"github.com/ekkinox/ankorbridge/internal/product/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateProductHandler struct {
	creator *service.Creator
}

func NewCreateProductHandler(creator *service.Creator) *CreateProductHandler {
	return &CreateProductHandler{
		creator: creator,
	}
}

type CreateProductParams struct {
	Name string `json:"name"`
}

func (h *CreateProductHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		params := new(CreateProductParams)
		if err := c.Bind(params); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid request body parameters %v", err))
		}

		if len(params.Name) < 2 {
			return echo.NewHTTPError(http.StatusBadRequest, "product name is too short")
		}

		if len(params.Name) > 100 {
			return echo.NewHTTPError(http.StatusBadRequest, "product name is too long")
		}

		p, err := h.creator.Create(ctx, service.CreateProductParams{
			Name: params.Name,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to create product: %v", err))
		}

		return c.JSON(http.StatusCreated, p)
	}
}
