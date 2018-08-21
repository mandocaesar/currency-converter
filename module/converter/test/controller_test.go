package converter

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/golang/glog"
	"github.com/stretchr/testify/assert"

	"github.com/currency-converter/shared"
	"github.com/currency-converter/shared/config"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func Init() {
	cfg, err := config.New("../../../shared/config/")
	glog.Error(err)
	configuration := *cfg

	routerInstance := shared.NewRouter(configuration)
	router = routerInstance.SetupRouter()
}

func TestAddExchange(t *testing.T) {
	Init()
	payload := bytes.NewBuffer([]byte(`{"from":"USD","to":"GBP"}`))
	response := shared.DispatchRequest(router, "POST", "/api/v1/exchange", payload)
	assert.Equal(t, http.StatusOK, response.Code)
}
