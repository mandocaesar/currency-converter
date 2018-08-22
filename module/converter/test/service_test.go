package converter

import (
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
		Rate:         0.171717}

	result, err := service.AddDailyExchange(request.From, request.To, request.ExchangeDate, request.Rate)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, "", result.String())
}
