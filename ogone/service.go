package ogone

import (
	"log"
	"net/http"

	"fmt"

	"github.com/ewgRa/paymentservices/ogone/aliasdirectlink"
	"github.com/ewgRa/paymentservices/service"
	"github.com/ewgRa/paymentservices/service/metric"
	"golang.org/x/net/context"
)

// Service for payments to ogone gateway
type Service struct {
}

// Start service
func (s *Service) Start(ctx context.Context, conf *service.Config) {
	m := &metric.Metric{}
	ep := &aliasdirectlink.Endpoint{M: m}

	ks := aliasdirectlink.NewKitHandler(ctx, ep)
	http.Handle(conf.GetPath(), ks)

	metricEp := &metric.Endpoint{M: m}
	metricKs := metric.NewKitHandler(ctx, metricEp)
	http.Handle(conf.GetMetricPath(), metricKs)

	log.Fatal(http.ListenAndServe(conf.GetHost()+":"+fmt.Sprint(conf.GetPort()), nil))
}
