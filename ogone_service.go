package main

import (
	"github.com/ewgRa/paymentservices/ogone"
	"golang.org/x/net/context"
)

func main() {
	s := &ogone.Service{}
	s.Start(context.Background())
}
