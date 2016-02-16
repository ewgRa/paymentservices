package ogone_alias_direct_link

type Endpoint struct {
}

func (ep *Endpoint) Response(r interface{}) (interface{}, error) {
	if err := r.(*Request).isValid(); err != nil {
		return nil, err
	}

	// FIXME XXX: implementation
	return &Response{"OK", ""}, nil
}
