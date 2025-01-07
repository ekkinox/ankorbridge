package product

import (
	"fmt"
	"github.com/ekkinox/ankorbridge/internal/product/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type DeleteProductHandler struct {
	deleter *service.Deleter
}

func NewDeleteProductHandler(deleter *service.Deleter) *DeleteProductHandler {
	return &DeleteProductHandler{
		deleter: deleter,
	}
}

func (h *DeleteProductHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		productID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		err = h.deleter.Delete(ctx, productID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to delete product: %v", err))
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
