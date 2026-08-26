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

func TestNewConfiguresServiceGroupsAndRawClient(t *testing.T) {
	client := New("")
	if client.Events == nil || client.Operations == nil || client.Sandbox == nil {
		t.Fatal("expected all service groups")
	}
	if client.OpenAPI() == nil {
		t.Fatal("expected generated client")
	}
}

func TestDeprecatedRawClientEntryPointsRemainUsable(t *testing.T) {
	configuration := NewConfiguration()
	client := NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), ContextAccessToken, "test-token")
	if client == nil || ctx.Value(ContextAccessToken) != "test-token" {
		t.Fatal("expected raw client compatibility entry points")
	}
}

func TestDefaultServerIsPlainRouter(t *testing.T) {
	configuration := NewConfiguration()
	serverURL, err := configuration.ServerURL(0, nil)
	if err != nil {
		t.Fatal(err)
	}
	if serverURL != "https://plainrouter.com/api/v1" {
		t.Fatalf("unexpected default server: %s", serverURL)
	}
}

func TestBearerTokenAndOptionsApplyToRequests(t *testing.T) {
	transport := roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if request.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", request.Method)
		}
		if request.URL.String() != "https://example.test/reports/emq" {
			t.Errorf("unexpected URL: %s", request.URL)
		}
		if got := request.Header.Get("Authorization"); got != "Bearer test-token" {
			t.Errorf("unexpected authorization header: %s", got)
		}
		if got := request.Header.Get("User-Agent"); got != "plainrouter-test/1.0" {
			t.Errorf("unexpected user agent: %s", got)
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{"snapshots":[]}`)),
			Request:    request,
		}, nil
	})

	client := New(
		"test-token",
		WithBaseURL("https://example.test/"),
		WithHTTPClient(&http.Client{Transport: transport}),
		WithUserAgent("plainrouter-test/1.0"),
	)

	result, response, err := client.Operations.GetEmqReport(context.Background()).Execute()
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if result == nil {
		t.Fatal("expected a decoded response")
	}
}

func TestEmptyTokenDoesNotAddAuthorization(t *testing.T) {
	transport := roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if got := request.Header.Get("Authorization"); got != "" {
			t.Errorf("unexpected authorization header: %s", got)
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader("")),
			Request:    request,
		}, nil
	})

	client := New(
		"",
		WithBaseURL("https://example.test"),
		WithHTTPClient(&http.Client{Transport: transport}),
	)
	_, response, err := client.Sandbox.GetSandbox(context.Background()).Execute()
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
}
