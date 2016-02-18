package aliasdirectlink

// Endpoint for Alias Direct Link
type Endpoint struct {
}

// Response for Alias Direct Link request
func (ep *Endpoint) Response(r *Request) (*Response, error) {
	if err := r.isValid(); err != nil {
		return nil, err
	}

	// FIXME XXX: implementation
	return &Response{"OK", ""}, nil
}
