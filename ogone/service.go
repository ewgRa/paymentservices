package ogone

import (
	"golang.org/x/net/context"
	"net/http"
	"log"
	alias_direct_link "./alias_direct_link"
	"../service/kit"
)

type Service struct {
}

func (s *Service) Start(ctx context.Context) {
	ks := kit_service.NewHandler(
		ctx,
		&alias_direct_link.Endpoint{},
		kit_service.NewJsonTransport(func() interface{} {return alias_direct_link.NewRequest()}),
	)

	http.Handle("/alias_direct_link.json", ks)

	// FIXME XXX: configuration
	log.Fatal(http.ListenAndServe(":8080", nil))
}
