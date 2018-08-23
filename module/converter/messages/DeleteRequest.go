package messages

//DeleteRequest : request message to delete list of exchanges
type DeleteRequest struct {
	Exchanges []*ExchangeRequest `json:"exchanges" binding:"required"`
}
