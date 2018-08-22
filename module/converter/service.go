package converter

import (
	"errors"
	"time"

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

	if !tx.First(&exchange, "Source = ? AND Target = ?", from, to).RecordNotFound() {
		return uuid.Nil, errors.New("Exchange already registered")
	}

	if err := tx.Create(exchange).Error; err != nil {
		tx.Rollback()

		return uuid.Nil, err
	}
	tx.Commit()
	return exchange.ID, nil

}

//AddDailyExchange : function to add daily exchange rate
func (s *Service) AddDailyExchange(from string, to string, exchangeDate string, rate float32) (uuid.UUID, error) {
	if from == "" || to == "" || exchangeDate == "" || rate <= 0 {
		return uuid.Nil, errors.New("one of parameter is nil")
	}

	dt, err := time.Parse("2006-01-02", exchangeDate)
	if err != nil {
		return uuid.Nil, err
	}

	tx := s.db.Begin()
	var exchange model.Exchange
	if tx.First(&exchange, "Source = ? AND Target = ?", from, to).RecordNotFound() {
		return uuid.Nil, errors.New("Exchange not listed, please register it first ")
	}

	var check model.DailyRate
	if !tx.First(&check, "exchange_id = ? AND exchange_date = ?", exchange.ID, dt).RecordNotFound() {
		return uuid.Nil, errors.New("Daily Exchange Rate Exist")
	}

	data := &model.DailyRate{ExchangeID: exchange.ID, ExchangeDate: &dt, Rate: rate}
	if err := tx.Create(data).Error; err != nil {
		tx.Rollback()
		return uuid.Nil, err
	}
	tx.Commit()

	return data.ID, nil
}
