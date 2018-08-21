package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/currency-converter/shared"
	"github.com/currency-converter/shared/config"
	"github.com/currency-converter/shared/data"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

var (
	runMigration  bool
	runSeeder     bool
	db            *gorm.DB
	configuration config.Configuration
	router        *gin.Engine
)

func init() {
	flag.BoolVar(&runMigration, "migrate", false, "run db migration")
	flag.BoolVar(&runSeeder, "seed", false, "run db seeder after db migration")
	flag.Parse()

	glog.V(2).Info("Migration status : %s", runMigration)
	glog.V(2).Info("Initilizing server...")

	//Setup Configuration
	cfg, err := config.New("")
	if err != nil {
		glog.Fatalf("Failed to load configuration: %s", err)
		panic(fmt.Errorf("Fatal error on load configuration : %s ", err))
	}
	configuration = *cfg

	//Setup Router
	routerInstance := shared.NewRouter(configuration)
	router = routerInstance.SetupRouter()

	if runMigration == true {
		dbMigration, error := data.NewDbMigration(&configuration)
		if error != nil {
			glog.Fatalf("Failed instantiate dbmigration : %s", error)
		}

		var count = 0

		if migrate(dbMigration, runSeeder) != nil {
			if count == 10 {
				return
			}
			count++
			migrate(dbMigration, runSeeder)
		}
	}
}

func migrate(db *data.DbMigration, seed bool) error {
	success, error := db.Migrate(seed)

	if error != nil {
		glog.Fatalf("Failed migrate: %s", error)
	}

	glog.V(2).Infof("database migration : %s", success)
	return error
}

func main() {
	glog.V(2).Infof("Server run on mode: %s", configuration.Server.Mode)
	gin.SetMode(configuration.Server.Mode)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(fmt.Errorf("Fatal error failed to start server : %s", err))
	}
}
