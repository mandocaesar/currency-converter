package data

import (
	converterModel "github.com/currency-converter/module/converter/model"

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
	)

	if seed {
		glog.Info("Start Database Seed")

		tx := d.connection.Begin()
		if err := tx.Exec("Truncate Table time_slots").Error; err != nil {
			tx.Rollback()
			glog.Infof("seed error:%s", err.Error())
		}

		//PUT SEED HERE

		tx.Commit()
		glog.Info("Finish Database Seed")
	}

	glog.Info("Database migration finished")
	return true, nil
}
