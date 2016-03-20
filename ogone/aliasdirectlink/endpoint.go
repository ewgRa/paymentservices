package aliasdirectlink

import (
	"github.com/ewgRa/ogone"
	ogoneservice "github.com/ewgRa/paymentservices/ogone"
	"github.com/ewgRa/paymentservices/service/metric"
)

// Endpoint for Alias Direct Link
type Endpoint struct {
	M *metric.Metric
	C *ogoneservice.Config
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

	dlr.
		SetPspID(ep.C.GetPspID()).
		SetUserID(ep.C.GetUserID()).
		SetPassword(ep.C.GetPassword()).
		Sign(ep.C.GetSign())

	var dlResp *ogone.DirectLinkResponse

	if ep.C.IsSandbox() {
		dlResp, _ = dlg.SandboxSend(dlr)
	} else {
		dlResp, _ = dlg.Send(dlr)
	}

	if !dlResp.IsAuthorised() {
		return &Response{"NOK", ""}, nil
	}

	return &Response{"OK", ""}, nil
}
