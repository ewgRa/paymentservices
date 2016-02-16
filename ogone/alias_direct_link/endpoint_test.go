package ogone_alias_direct_link

import (
	"testing"
	"golang.org/x/net/context"
	"net/http"
	"fmt"
	"net/http/httptest"
)

func TestAliasDirectLinkEndpoint(t *testing.T) {
	ep := &Endpoint{}

	r := &http.Request{}
	response := &httptest.ResponseRecorder{}

	h := ep.Handler(context.Background())

	h.ServeHTTP(response, r)

	fmt.Println(response)
}
