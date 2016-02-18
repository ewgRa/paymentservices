package ogone

import (
	"log"
	"net/http"

	"./aliasdirectlink"
	"golang.org/x/net/context"
)

// Service for payments to ogone gateway
type Service struct {
}

// Start service
func (s *Service) Start(ctx context.Context) {
	ks := aliasdirectlink.NewKitHandler(ctx)
	http.Handle("/alias_direct_link.json", ks)

	// FIXME XXX: configuration
	log.Fatal(http.ListenAndServe(":8080", nil))
}
