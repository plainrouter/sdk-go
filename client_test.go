package plainrouter

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (function roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return function(request)
}

func TestDefaultServerIsPlainRouter(t *testing.T) {
	config := NewConfiguration()
	serverURL, err := config.ServerURL(0, nil)
	if err != nil {
		t.Fatal(err)
	}
	if serverURL != "https://plainrouter.com/api/v1" {
		t.Fatalf("unexpected default server: %s", serverURL)
	}
}

func TestBearerTokenAndReportRequest(t *testing.T) {
	transport := roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if request.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", request.Method)
		}
		if request.URL.Path != "/api/v1/reports/emq" {
			t.Errorf("unexpected path: %s", request.URL.Path)
		}
		if got := request.Header.Get("Authorization"); got != "Bearer test-token" {
			t.Errorf("unexpected authorization header: %s", got)
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{"snapshots":[]}`)),
			Request:    request,
		}, nil
	})

	config := NewConfiguration()
	config.HTTPClient = &http.Client{Transport: transport}
	client := NewAPIClient(config)
	ctx := context.WithValue(context.Background(), ContextAccessToken, "test-token")

	result, httpResponse, err := client.OperationsAPI.GetEmqReport(ctx).Execute()
	if err != nil {
		t.Fatal(err)
	}
	defer httpResponse.Body.Close()
	if result == nil {
		t.Fatal("expected a decoded response")
	}
}
