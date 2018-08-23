package messages

//TrackedRequest : request message to get tracked list of exchanges
type TrackedRequest struct {
	Date      string             `json:"date" binding:"required"`
	Exchanges []*ExchangeRequest `json:"exchanges" binding:"required"`
}
