package ogone

import (
	"log"
	"net/http"

	"github.com/ewgRa/paymentservices/ogone/aliasdirectlink"
	"golang.org/x/net/context"
	"github.com/ewgRa/paymentservices/service/metric"
)

// Service for payments to ogone gateway
type Service struct {
}

// Start service
func (s *Service) Start(ctx context.Context) {
	m := &metric.Metric{}
	ep := &aliasdirectlink.Endpoint{M: m}

	ks := aliasdirectlink.NewKitHandler(ctx, ep)
	http.Handle("/alias_direct_link.json", ks)

	metricEp := &metric.Endpoint{M: m}
	metricKs := metric.NewKitHandler(ctx, metricEp)
	http.Handle("/alias_direct_link.metric.json", metricKs)

	// FIXME XXX: configuration
	log.Fatal(http.ListenAndServe(":8080", nil))
}
