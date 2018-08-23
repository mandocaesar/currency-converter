package messages

//ExchangeRequest :request model for add exchange
type ExchangeRequest struct {
	From string `json:"from" binding:"required,len=3"`
	To   string `json:"to" binding:"required,len=3"`
}
