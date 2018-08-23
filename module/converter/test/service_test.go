package converter

import (
	"strconv"
	"testing"

	messages "github.com/currency-converter/module/converter/messages"

	"github.com/stretchr/testify/assert"
)

func TestAddExchange(t *testing.T) {
	Init()
	result, err := service.AddExchange("USD", "GBP")
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestAddDailyData(t *testing.T) {
	Init()
	request := &messages.AddDailyRateRequest{
		From:         "USD",
		To:           "GBP",
		ExchangeDate: "2018-01-09",
		Rate:         "0.171717"}

	f, err := strconv.ParseFloat(request.Rate, 32)
	result, err := service.AddDailyExchange(request.From, request.To, request.ExchangeDate, float64(f))
	assert.Equal(t, nil, err)
	assert.NotEqual(t, "", result.String())
}

func TestExchangeRateLast7(t *testing.T) {
	Init()
	request := &messages.ExchangeRequest7Day{From: "USD", To: "GBP", Date: "2018-08-01"}

	result, err := service.ExchangeRateLast7(request.From, request.To, request.Date)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, result.Data)
}

func TestTrackedRates(t *testing.T) {
	Init()
	request := &messages.TrackedRequest{Date: "2018-08-01"}

	result, err := service.TracketRates(request.From, request.To, request.Date)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, result.Data)
}
