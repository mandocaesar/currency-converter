package converter

import (
	"net/http"
	"strconv"

	"github.com/currency-converter/module/converter/messages"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//Controller for contact management
type Controller struct {
	converterService *Service
}

//NewController : function to instantiate new controller
func NewController(service *Service) (*Controller, error) {
	return &Controller{converterService: service}, nil
}

//AddExchange : function to add exchange to list
func (c *Controller) AddExchange(ctx *gin.Context) {
	var req messages.AddExchangeRequest
	err := ctx.ShouldBindWith(&req, binding.JSON)
	if err == nil {
		id, err := c.converterService.AddExchange(req.From, req.To)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{"result": "success", "id": id})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"result": "success", "error": err.Error(), "id": 0})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": "failed", "error": err.Error(), "id": 0})
	}

	return
}

//AddDailyRate : function to add daily rate to registered exchange
func (c *Controller) AddDailyRate(ctx *gin.Context) {
	var req messages.AddDailyRateRequest
	err := ctx.ShouldBindWith(&req, binding.JSON)
	if err == nil {
		rate, err := strconv.ParseFloat(req.Rate, 32)

		id, err := c.converterService.AddDailyExchange(req.From, req.To, req.ExchangeDate, float32(rate))
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{"result": "success", "id": id})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"result": "success", "error": err.Error(), "id": 0})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": "failed", "error": err.Error(), "id": 0})
	}

	return
}
