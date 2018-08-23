package messages

import (
	"time"
)

//Exchange7DaysResponse :request model to show 7 days exchange
type Exchange7DaysResponse struct {
	From     string         `json:"from" binding:"required,len=3"`
	To       string         `json:"to" binding:"required,len=3"`
	Average  float64        `json:"average"`
	Variance float64        `json:"variance"`
	Data     []ExchangeData `json:"data"`
}

//ExchangeData : model to show date and rate
type ExchangeData struct {
	Date *time.Time `json:"Date"`
	Rate float64    `json:"Rate"`
}
