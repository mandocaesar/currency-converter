package converter

import (
	"errors"

	model "github.com/currency-converter/module/converter/model"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
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

//AddExchange : function to add exchange to list
func (s *Service) AddExchange(from string, to string) (uuid.UUID, error) {
	if (from == "" || to == "") || (len(from) != 3 || len(to) != 3) {
		return uuid.Nil, errors.New("Source or Target is nil")
	}

	tx := s.db.Begin()
	exchange := &model.Exchange{Source: from, Target: to}
	if err := tx.Create(exchange).Error; err != nil {
		tx.Rollback()

		return uuid.Nil, err
	}
	tx.Commit()
	return exchange.ID, nil

}

//AddDailyRequest : function to add daily exchange rate
func (s *Service) AddDailyRequest(from string, to string, exchangeDate string, rate float32) (uuid.UUID, error) {
	return uuid.Nil, nil
}
