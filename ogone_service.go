package main

import (
	"github.com/ewgRa/paymentservices/ogone"
	"github.com/ewgRa/paymentservices/service"
	"golang.org/x/net/context"
)

func main() {
	s := &ogone.Service{}
	conf := service.NewConfig("", 8080, "/alias_direct_link.json", "/alias_direct_link.metric.json")
	s.Start(context.Background(), conf)
}
