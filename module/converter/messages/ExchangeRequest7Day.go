package messages

//ExchangeRequest7Day :request model for add exchange
type ExchangeRequest7Day struct {
	From string `json:"from" binding:"required,len=3"`
	To   string `json:"to" binding:"required,len=3"`
	Date string `json:"date" binding:"required"`
}
