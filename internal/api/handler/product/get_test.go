package product_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ankorstore/yokai/fxsql"
	"github.com/ankorstore/yokai/log/logtest"
	"github.com/ankorstore/yokai/trace/tracetest"
	"github.com/ekkinox/ankorbridge/db/seeds"
	"github.com/ekkinox/ankorbridge/internal"
	"github.com/ekkinox/ankorbridge/internal/product"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func TestGetProductHandler(t *testing.T) {
	var httpServer *echo.Echo
	var logBuffer logtest.TestLogBuffer
	var traceExporter tracetest.TestTraceExporter

	internal.RunTest(
		t,
		fxsql.RunFxSQLSeeds(
			seeds.ProductsSeedName,
		),
		fx.Populate(&httpServer, &logBuffer, &traceExporter),
	)

	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	rec := httptest.NewRecorder()
	httpServer.ServeHTTP(rec, req)

	// response assertion
	assert.Equal(t, http.StatusOK, rec.Code)

	var p product.Product

	err := json.Unmarshal(rec.Body.Bytes(), &p)
	assert.NoError(t, err)

	assert.Equal(t, 1, p.ID)
	assert.Equal(t, "foo", p.Name)

	// logs assertion
	logtest.AssertHasLogRecord(t, logBuffer, map[string]interface{}{
		"level":   "info",
		"message": "Fetching single product",
	})

	// trace assertion
	tracetest.AssertHasTraceSpan(t, traceExporter, "Fetching single product")
}
