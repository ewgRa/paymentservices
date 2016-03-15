package metric

import "fmt"

// Endpoint for Alias Direct Link
type Endpoint struct {
	M *Metric
}

// Response for Alias Direct Link request
func (ep *Endpoint) Response(r *Request) (*Response, error) {
	return &Response{"requests count: " + fmt.Sprint(ep.M.GetRequestsCount()), ""}, nil
}