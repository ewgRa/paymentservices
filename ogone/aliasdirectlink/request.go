package aliasdirectlink

// Request for Alias direct link
type Request struct {
	OrderID int `json:"orderId"`
	Amount  int `json:","`
	//	Currency string `json:"currency"`
	Alias string `json:"alias"`
}

// isValid validate request
func (r *Request) isValid() error {
	if "" == r.Alias {
		return nil
		//return errors.New("Alias required")
	}

	return nil
}

// NewRequest create new request
func NewRequest() *Request {
	return &Request{}
}
