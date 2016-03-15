package aliasdirectlink

import "github.com/ewgRa/paymentservices/service/metric"

// Endpoint for Alias Direct Link
type Endpoint struct {
	M *metric.Metric
}

// Response for Alias Direct Link request
func (ep *Endpoint) Response(r *Request) (*Response, error) {
	if err := r.isValid(); err != nil {
		return nil, err
	}

	if ep.M != nil {
		ep.M.IncRequestsCount()
	}

	// FIXME XXX: implementation
	return &Response{"OK", ""}, nil
}
