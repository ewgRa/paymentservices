package main

import (
	"./ogone"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	http.Handle("/alias_accept", ogone.NewAliasAcceptHandler(ctx))
	http.Handle("/alias_direct_link", ogone.NewAliasDirectLinkHandler(ctx))

	// FIXME XXX: port configuration
	log.Fatal(http.ListenAndServe(":8080", nil))
}
