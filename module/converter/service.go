package converter

import (
	"errors"
	"sort"
	"time"

	"github.com/currency-converter/module/converter/messages"
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
func (s *Service) AddDailyExchange(from string, to string, exchangeDate string, rate float64) (uuid.UUID, error) {
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

//ExchangeRateLast7 : function to get rate exchange trend for last 7 day
func (s *Service) ExchangeRateLast7(from string, to string, exchangeDate string) (*messages.Exchange7DaysResponse, error) {
	var data []messages.ExchangeData
	var rates []float64
	var total float64

	if (from == "" || to == "") || (len(from) != 3 || len(to) != 3) {
		return nil, errors.New("Source or Target is nil")
	}

	var dailyrates []model.DailyRate
	var exchange model.Exchange

	tx := s.db.Begin()
	dt, err := time.Parse("2006-01-02", exchangeDate)
	if err != nil {
		return nil, err
	}

	if err := tx.Find(&exchange, "Source = ? AND Target = ?", from, to).Error; err != nil {
		return nil, err
	}

	end := dt.AddDate(0, 0, 6)

	if err := tx.Order("exchange_date desc").Find(&dailyrates, "exchange_id = ? AND exchange_date BETWEEN ? AND ?", exchange.ID, dt.Format("2006-01-02"), end.Format("2006-01-02")).Error; err != nil {
		return nil, err
	}

	if len(dailyrates) == 0 {
		return nil, errors.New("No Daily exchange rate available")
	}

	for index := 0; index < len(dailyrates); index++ {
		info := messages.ExchangeData{Date: dailyrates[index].ExchangeDate, Rate: dailyrates[index].Rate}
		data = append(data, info)
		rates = append(rates, dailyrates[index].Rate)
		total += dailyrates[index].Rate
	}

	sort.Float64s(rates)

	variance := rates[len(rates)-1] - rates[0]
	avg := total / float64(len(dailyrates))

	r := &messages.Exchange7DaysResponse{From: from, To: to, Average: avg, Variance: variance, Data: data}

	return r, nil
}
