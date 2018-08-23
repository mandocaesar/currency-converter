package data

import (
	"time"

	converterModel "github.com/currency-converter/module/converter/model"
	"github.com/satori/go.uuid"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"

	"github.com/currency-converter/shared/config"
)

//DbMigration : instance to hold dbmigration instance
type DbMigration struct {
	connection *gorm.DB
}

//NewDbMigration : intantiate new DBMigration instance
func NewDbMigration(cfg *config.Configuration) (*DbMigration, error) {
	dbFactory, err := NewDbFactory(cfg)

	if err != nil {
		glog.Errorf("%s", err)
		return nil, err
	}

	conn, err := dbFactory.DBConnection()

	if err != nil {
		glog.Errorf("%s", err)
		return nil, err
	}

	return &DbMigration{connection: conn}, nil
}

//Migrate : function to invoke gorm's automigrate
func (d *DbMigration) Migrate(seed bool) (bool, error) {
	glog.Info("Start Database Migration")

	d.connection.AutoMigrate(
		&converterModel.Exchange{},
		&converterModel.DailyRate{},
	)

	if seed {
		glog.Info("Start Database Seed")

		tx := d.connection.Begin()

		usdGbp := &converterModel.Exchange{Source: "USD", Target: "GBP"}
		usdIdr := &converterModel.Exchange{Source: "USD", Target: "IDR"}
		jpyIdr := &converterModel.Exchange{Source: "JPY", Target: "IDR"}

		tx.Create(usdGbp)
		tx.Create(usdIdr)
		tx.Create(jpyIdr)

		dailyRate := &converterModel.DailyRate{}

		dt, _ := time.Parse("2006-01-02", "2018-08-01")
		dailyRate.ID = uuid.NewV4()
		dailyRate.ExchangeID = usdGbp.ID
		dailyRate.ExchangeDate = &dt
		dailyRate.Rate = 0.711
		tx.Create(dailyRate)

		dt, _ = time.Parse("2006-01-02", "2018-08-02")
		dailyRate.ID = uuid.NewV4()
		dailyRate.ExchangeID = usdGbp.ID
		dailyRate.ExchangeDate = &dt
		dailyRate.Rate = 0.111
		tx.Create(dailyRate)

		dt, _ = time.Parse("2006-01-02", "2018-08-03")
		dailyRate.ID = uuid.NewV4()
		dailyRate.ExchangeID = usdGbp.ID
		dailyRate.ExchangeDate = &dt
		dailyRate.Rate = 0.411
		tx.Create(dailyRate)

		dt, _ = time.Parse("2006-01-02", "2018-08-04")
		dailyRate.ID = uuid.NewV4()
		dailyRate.ExchangeID = usdGbp.ID
		dailyRate.ExchangeDate = &dt
		dailyRate.Rate = 0.511
		tx.Create(dailyRate)

		dt, _ = time.Parse("2006-01-02", "2018-08-05")
		dailyRate.ID = uuid.NewV4()
		dailyRate.ExchangeID = usdGbp.ID
		dailyRate.ExchangeDate = &dt
		dailyRate.Rate = 0.3311
		tx.Create(dailyRate)

		dt, _ = time.Parse("2006-01-02", "2018-08-06")
		dailyRate.ID = uuid.NewV4()
		dailyRate.ExchangeID = usdGbp.ID
		dailyRate.ExchangeDate = &dt
		dailyRate.Rate = 0.3566
		tx.Create(dailyRate)

		dt, _ = time.Parse("2006-01-02", "2018-08-07")
		dailyRate.ID = uuid.NewV4()
		dailyRate.ExchangeID = usdGbp.ID
		dailyRate.ExchangeDate = &dt
		dailyRate.Rate = 0.787
		tx.Create(dailyRate)

		tx.Commit()
		glog.Info("Finish Database Seed")
	}

	glog.Info("Database migration finished")
	return true, nil
}
