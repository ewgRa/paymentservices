package main

import (
	"./ogone"
	"golang.org/x/net/context"
)

func main() {
	s := &ogone.Service{}
	s.Start(context.Background())
}
