package converter

import (
	"errors"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//Service : authentication service
type Service struct {
	db *gorm.DB
}

//NewService : instantiate new contact service
func NewService(Db *gorm.DB) (*Service, error) {
	if Db == nil {
		glog.Error("failed to intantiate Booking Service , Db instance is null")
		return nil, errors.New("failed to intantiate Booking Service , Db instance is null")
	}
	return &Service{db: Db}, nil
}
