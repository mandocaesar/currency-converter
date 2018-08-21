package converter

//Controller for contact management
type Controller struct {
	converterService *Service
}

//NewController : function to instantiate new controller
func NewController(service *Service) (*Controller, error) {
	return &Controller{converterService: service}, nil
}
