package messages

//TrackedResponse : Response struct for tracked request
type TrackedResponse struct {
	From string `json:"from" binding:"required,len=3"`
	To   string `json:"to" binding:"required,len=3"`
	Rate string `json:"rate" binding:"required"`
	Avg  string `json:"7-day-avg" binding:"required"`
}
