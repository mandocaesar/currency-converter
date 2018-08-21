package converter

import (
	"github.com/currency-converter/module/converter"
	"github.com/currency-converter/shared"
	"github.com/currency-converter/shared/config"
	"github.com/currency-converter/shared/data"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var (
	router  *gin.Engine
	service *converter.Service
)

func Init() {
	cfg, err := config.New("../../../shared/config/")
	glog.Error(err)
	configuration := *cfg

	dbInstance, _ := data.NewDbFactory(cfg)
	conn, _ := dbInstance.DBConnection()
	service, _ = converter.NewService(conn)

	routerInstance := shared.NewRouter(configuration)
	router = routerInstance.SetupRouter()
}
