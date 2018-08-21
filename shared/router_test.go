package shared

import (
	"net/http"
	"testing"

	"github.com/currency-converter/shared/config"
	"github.com/stretchr/testify/assert"
)

//TestRouterDiagnosticEndPoint : Test Diagnostic Endpoint
func TestRouterDiagnosticEndPoint(t *testing.T) {
	cfg, err := config.New("../shared/config/")

	assert.Empty(t, err)

	configuration := *cfg
	routerInstance := NewRouter(configuration)
	router := routerInstance.SetupRouter()

	response := DispatchRequest(router, "GET", "/api/v1/ping", nil)

	assert.Equal(t, http.StatusOK, response.Code)
}
