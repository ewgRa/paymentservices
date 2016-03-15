package ogone

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ewgRa/paymentservices/ogone/aliasdirectlink"
	"github.com/ewgRa/paymentservices/service/metric"
	"golang.org/x/net/context"
)

func TestAliasDirectLinkEndpoint(t *testing.T) {
	metric := &metric.Metric{}
	ep := &aliasdirectlink.Endpoint{M: metric}

	ks := aliasdirectlink.NewKitHandler(context.Background(), ep)

	server := httptest.NewServer(ks)

	defer server.Close()

	resp, _ := http.Post(server.URL, "application/json", strings.NewReader("{\"orderId\":10,\"alias\":\"testalias\"}"))

	buf, _ := ioutil.ReadAll(resp.Body)

	if string(buf) != "{\"v\":\"OK\"}\n" {
		fmt.Println(string(buf))
		t.Fatalf("Wrong alias direct link response")
	}

	if metric.GetRequestsCount() != 1 {
		t.Fatalf("Wrong metric for requests count")
	}
}
