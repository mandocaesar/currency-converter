package shared

import (
	"net/http"
	"time"

	"github.com/currency-converter/module/converter"

	"github.com/gin-contrib/cors"
	"github.com/jinzhu/gorm"

	"github.com/golang/glog"

	"github.com/currency-converter/shared/config"
	"github.com/currency-converter/shared/data"
	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
)

//Router : Instance struct for router model
type Router struct {
	database *gorm.DB
	config   *config.Configuration

	converterController converter.Controller
	converterService    converter.Service
}

//NewRouter : Instantiate new Router
func NewRouter(configuration config.Configuration) *Router {
	cfg, _ := config.New("../../shared/config/")

	dbInstance, err := data.NewDbFactory(cfg)
	dbase, err := dbInstance.DBConnection()

	if err != nil {
		glog.Fatalf("Fatal Error on create database instance : %s", err.Error())
	}

	_converterService, errD := converter.NewService(dbase)
	if errD != nil {
		glog.Fatalf("Fatal Error on create converter Service : %s", errD.Error())
	}

	_converterController, errD := converter.NewController(_converterService)
	if errD != nil {
		glog.Fatalf("Fatal Error on create converter Controller : %s", errD.Error())
	}

	return &Router{
		converterController: *_converterController,
		converterService:    *_converterService,
	}
}

//SetupRouter : function that return registered end point
func (r *Router) SetupRouter() *gin.Engine {
	router := gin.New()

	//middleware setup
	duration := time.Duration(5) * time.Second
	router.Use(ginglog.Logger(duration), gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//diagnostic endpoint
	diagnostic := router.Group("api/v1")
	{
		diagnostic.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Name":       "BE-TEST-Server",
				"message":    "OK",
				"serverTime": time.Now().UTC(),
				"version":    "0.1",
			})
		})
	}

	exchange := router.Group("api/v1")
	{
		exchange.POST("/exchange", r.converterController.AddExchange)
		exchange.POST("/daily", r.converterController.AddDailyRate)
		exchange.POST("/last7", r.converterController.TrendLast7)
		exchange.POST("/tracked", r.converterController.Tracked)
	}

	return router
}
