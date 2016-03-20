package aliasdirectlink

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ewgRa/paymentservices/ogone"
	"github.com/ewgRa/paymentservices/service"
	"github.com/ewgRa/paymentservices/service/metric"
	"golang.org/x/net/context"
)

// Run the alias direct link service
func Run() {
	conf := service.NewConfig("", 8080, "/alias_direct_link.json", "/alias_direct_link.metric.json")

	// FIXME XXX: must be configured
	endpointConf := ogone.NewConfig(
		"ewgraogone",
		"ewgragolang",
		"123123aa",
		"qwdqwoidj29812d9",
		true,
	)

	m := &metric.Metric{}
	ep := &Endpoint{M: m, C: endpointConf}

	ks := NewKitHandler(context.Background(), ep)
	http.Handle(conf.GetPath(), ks)

	metricEp := &metric.Endpoint{M: m}
	metricKs := metric.NewKitHandler(context.Background(), metricEp)
	http.Handle(conf.GetMetricPath(), metricKs)

	log.Fatal(http.ListenAndServe(conf.GetHost()+":"+fmt.Sprint(conf.GetPort()), nil))
}
