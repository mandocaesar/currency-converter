package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddExchange(t *testing.T) {
	Init()
	result, err := service.AddExchange("GBP", "USD")
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
