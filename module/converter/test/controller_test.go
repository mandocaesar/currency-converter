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

func TestAPIAddDailyExchangeData(t *testing.T) {
	Init()
	payload := bytes.NewBuffer([]byte(`{"Date":"2018-07-12","From":"USD","To":"GBP","Rate":"0.75709"}`))
	response := shared.DispatchRequest(router, "POST", "/api/v1/daily", payload)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestAPIExchangeRateLast7(t *testing.T) {
	Init()
	payload := bytes.NewBuffer([]byte(`{"From":"USD","To":"GBP","Date":"2018-08-01"}`))
	response := shared.DispatchRequest(router, "POST", "/api/v1/last7", payload)
	assert.Equal(t, http.StatusOK, response.Code)
}
