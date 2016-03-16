package metric

import (
	"net/http"

	"github.com/ewgRa/paymentservices/service"
	kithttp "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

// NewKitHandler create new kit handler
func NewKitHandler(ctx context.Context, ep *Endpoint) *kithttp.Server {
	return kithttp.NewServer(
		ctx,
		func(ctx context.Context, request interface{}) (interface{}, error) {
			return ep.Response(request.(*Request))
		},
		func(r *http.Request) (request interface{}, err error) {
			if r.ContentLength == 0 {
				return NewRequest(), nil
			}

			return service.KitJSONDecodeFunc(r, NewRequest())
		},
		service.KitJSONEncodeFunc,
	)
}
