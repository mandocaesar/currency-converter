package messages

//AddDailyRateRequest : request message to add daily exchange
type AddDailyRateRequest struct {
	From         string  `json:"from" binding:"required,len=3"`
	To           string  `json:"to" binding:"required,len=3"`
	ExchangeDate string  `json:"date" binding:"required"`
	Rate         float32 `json:"rate" binding:"required"`
}
