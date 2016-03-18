package aliasdirectlink

import (
	"github.com/ewgRa/ogone"
	"github.com/ewgRa/paymentservices/service/metric"
)

// Endpoint for Alias Direct Link
type Endpoint struct {
	M *metric.Metric
}

// Response for Alias Direct Link request
func (ep *Endpoint) Response(r *Request) (*Response, error) {
	if ep.M != nil {
		ep.M.IncRequestsCount()
	}

	if err := r.isValid(); err != nil {
		return nil, err
	}

	dlr := ogone.NewDirectLinkRequest()

	dlr.
		SetAlias(r.Alias).
		SetAmount(r.Amount).
		SetReserveOperation().
		SetCurrency("EUR").
		SetOrderID(r.OrderID)

	dlg := ogone.NewDirectLinkGateway()

	// FIXME XXX: configuration here
	dlr.
		SetPspID("ewgraogone").
		SetUserID("ewgragolang").
		SetPassword("123123aa").
		Sign("qwdqwoidj29812d9")

	dlResp, _ := dlg.SandboxSend(dlr)

	if !dlResp.IsAuthorised() {
		return &Response{"NOK", ""}, nil
	}

	return &Response{"OK", ""}, nil
}
