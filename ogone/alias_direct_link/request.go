package ogone_alias_direct_link

type Request struct {
	orderId  int    `json:"orderId"`
	amount   int    `json:","`
//	currency string `json:"currency"`
	alias    string `json:"alias"`
}

func (r *Request) isValid() error {
	if "" == r.alias {
		return nil
		//return errors.New("Alias required")
	}

	return nil
}

func NewRequest() interface {} {
	return &Request{}
}