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
	ogonelib "github.com/ewgRa/ogone"
	"math/rand"
	"strconv"
	"time"
	"encoding/json"
)

// FIXME XXX: test non valid request, like without orderId and so on
func TestAliasDirectLinkEndpoint(t *testing.T) {
	metric := &metric.Metric{}
	ep := &aliasdirectlink.Endpoint{M: metric}

	ks := aliasdirectlink.NewKitHandler(context.Background(), ep)

	server := httptest.NewServer(ks)

	defer server.Close()

	orderID := "GOLANGTEST" + time.Now().Format("20060102150405_") + strconv.Itoa(rand.Intn(100000))

	aResp := ogonelib.NewTestAliasResponse(orderID, t)

	jsonRequest, _ := json.Marshal(map[string]string{"orderId": orderID, "alias": aResp.Alias(), "amount": "105"})

	resp, _ := http.Post(server.URL, "application/json", strings.NewReader(string(jsonRequest)))

	buf, _ := ioutil.ReadAll(resp.Body)

	if string(buf) != "{\"v\":\"OK\"}\n" {
		fmt.Println(string(buf))
		t.Fatalf("Wrong alias direct link response")
	}

	if metric.GetRequestsCount() != 1 {
		t.Fatalf("Wrong metric for requests count")
	}
}
