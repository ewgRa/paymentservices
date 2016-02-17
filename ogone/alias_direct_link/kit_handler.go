package ogone_alias_direct_link

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"../../service"
	"net/http"
)

func NewKitHandler(ctx context.Context) *kithttp.Server {
	ep := &Endpoint{}

	return kithttp.NewServer(
		ctx,
		func(ctx context.Context, request interface{}) (interface{}, error) {
			return ep.Response(request.(*Request))
		},
		func(r *http.Request) (request interface{}, err error) {
			return service.KitJsonDecodeFunc(r, NewRequest())
		},
		service.KitJsonEncodeFunc,
	)
}
