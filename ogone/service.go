package ogone

import (
	"log"
	"net/http"

	"github.com/ewgRa/paymentservices/ogone/aliasdirectlink"
	"golang.org/x/net/context"
	"github.com/ewgRa/paymentservices/service/metric"
	"github.com/ewgRa/paymentservices/service"
	"fmt"
)

// Service for payments to ogone gateway
type Service struct {
}

// Start service
func (s *Service) Start(ctx *context.Context, conf *service.Config) {
	m := &metric.Metric{}
	ep := &aliasdirectlink.Endpoint{M: m}

	ks := aliasdirectlink.NewKitHandler(ctx, ep)
	http.Handle(conf.GetPath(), ks)

	metricEp := &metric.Endpoint{M: m}
	metricKs := metric.NewKitHandler(ctx, metricEp)
	http.Handle(conf.GetMetricPath(), metricKs)

	log.Fatal(http.ListenAndServe(conf.GetHost() + ":" + fmt.Sprint(conf.GetPort()), nil))
}
