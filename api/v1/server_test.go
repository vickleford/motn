package v1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testclient struct {
	hut    *http.ServeMux
	server *httptest.Server
	client *http.Client
	t      *testing.T
}

func newtestclient(handlerUnderTest *http.ServeMux, t *testing.T) *testclient {
	tc := new(testclient)
	tc.hut = handlerUnderTest
	tc.server = httptest.NewServer(handlerUnderTest)
	tc.client = tc.server.Client()
	tc.t = t
	return tc
}

func (tc *testclient) assertEndpointWired(path, substr string) {
	res, err := tc.client.Get(fmt.Sprintf("%s%s", tc.server.URL, path))
	if err != nil {
		tc.t.Errorf("Error making request. %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		tc.t.Errorf("Error reading response body: %s", err.Error())
	}

	if strings.Contains(string(body), substr) != true {
		tc.t.Errorf("Wrong response body. Got: '%s'", body)
	}
}

func TestServeV1(t *testing.T) {
	v1handler := ServeV1()
	tester := newtestclient(v1handler, t)

	tester.assertEndpointWired("/v1/healthz", "status")
	tester.assertEndpointWired("/v1/motn", "message")
}
