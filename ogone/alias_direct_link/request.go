package ogone_alias_direct_link

type Request struct {
	OrderId  int    `json:"orderId"`
	Amount   int    `json:","`
//	Currency string `json:"currency"`
	Alias    string `json:"alias"`
}

func (r *Request) isValid() error {
	if "" == r.Alias {
		return nil
		//return errors.New("Alias required")
	}

	return nil
}

func NewRequest() *Request {
	return &Request{}
}
