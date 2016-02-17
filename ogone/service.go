package ogone

import (
	"golang.org/x/net/context"
	"net/http"
	"log"
	alias_direct_link "./alias_direct_link"
)

type Service struct {
}

func (s *Service) Start(ctx context.Context) {
	ks := alias_direct_link.NewKitHandler(ctx)
	http.Handle("/alias_direct_link.json", ks)

	// FIXME XXX: configuration
	log.Fatal(http.ListenAndServe(":8080", nil))
}
