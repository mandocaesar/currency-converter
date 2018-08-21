package messages

//AddExchangeRequest :request model for add exchange
type AddExchangeRequest struct {
	From string `json:"from" binding:"required,len=3"`
	To   string `json:"to" binding:"required,len=3"`
}
