package converter

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/currency-converter/shared"
)

func TestAPIAddExchange(t *testing.T) {
	Init()
	payload := bytes.NewBuffer([]byte(`{"from":"USD","to":"GBP"}`))
	response := shared.DispatchRequest(router, "POST", "/api/v1/exchange", payload)
	assert.Equal(t, http.StatusOK, response.Code)
}
