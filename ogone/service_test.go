package ogone


import (
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/net/context"
	"github.com/ewgRa/paymentservices/ogone/aliasdirectlink"
	"strings"
	"io/ioutil"
	"fmt"
)

func TestAliasDirectLinkEndpoint(t *testing.T) {
	ks := aliasdirectlink.NewKitHandler(context.Background())

	server := httptest.NewServer(ks)

	defer server.Close()

	resp, _ := http.Post(server.URL, "application/json", strings.NewReader("{\"orderId\":10,\"alias\":\"testalias\"}"))

	buf, _ := ioutil.ReadAll(resp.Body)

	if string(buf) != "{\"v\":\"OK\"}\n" {
		fmt.Println(string(buf))
		t.Fatalf("Wrong alias direct link response")
	}
}
