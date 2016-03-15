package main

import (
	"github.com/ewgRa/paymentservices/ogone"
	"golang.org/x/net/context"
	"github.com/ewgRa/paymentservices/service"
)

func main() {
	s := &ogone.Service{}
	conf := service.NewConfig("", 8080, "/alias_direct_link.json", "/alias_direct_link.metric.json")
	s.Start(context.Background(), conf)
}
