//go:build live

package plainrouter

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/plainrouter/sdk-go/openapi"
)

// TestLiveSandboxContract calls the live zero-auth sandbox, which persists
// nothing and contacts no advertising provider. Run it with -tags live.
func TestLiveSandboxContract(t *testing.T) {
	options := []Option{}
	if baseURL := os.Getenv("PLAINROUTER_BASE_URL"); baseURL != "" {
		options = append(options, WithBaseURL(baseURL))
	}
	ctx := context.Background()

	sandbox, _, err := New("", options...).Sandbox.GetSandbox(ctx).Execute()
	if err != nil {
		t.Fatal(err)
	}
	if sandbox.GetPersistsData() || sandbox.GetProviderDelivery() {
		t.Fatal("the live sandbox no longer declares itself isolated")
	}

	encoded, err := json.Marshal(sandbox.GetTry().Body)
	if err != nil {
		t.Fatal(err)
	}
	var body openapi.ValidateSandboxEventRequest
	if err := json.Unmarshal(encoded, &body); err != nil {
		t.Fatal(err)
	}

	assertDiscarded := func(result *openapi.ValidateSandboxEvent200Response, err error) {
		t.Helper()
		if err != nil {
			t.Fatal(err)
		}
		if !result.GetSandbox() || !result.GetAccepted() || result.GetPersisted() || result.GetProviderDelivery() {
			t.Fatalf("synthetic event was not validated and discarded: %+v", result)
		}
	}

	result, _, err := New("", options...).Sandbox.ValidateSandboxEvent(ctx).ValidateSandboxEventRequest(body).Execute()
	assertDiscarded(result, err)

	issuedKey := sandbox.GetSelfServeKey().IssuedKey
	result, _, err = New(issuedKey.GetApiKey(), options...).Sandbox.ValidateSandboxEventWithKey(ctx).ValidateSandboxEventRequest(body).Execute()
	assertDiscarded(result, err)
}

// TestLiveAuthenticatedRead lists events with PLAINROUTER_SMOKE_SECRET and
// skips when the secret is absent.
func TestLiveAuthenticatedRead(t *testing.T) {
	secret := os.Getenv("PLAINROUTER_SMOKE_SECRET")
	if secret == "" {
		t.Skip("PLAINROUTER_SMOKE_SECRET is not set; only the zero-auth sandbox was exercised")
	}
	options := []Option{}
	if baseURL := os.Getenv("PLAINROUTER_BASE_URL"); baseURL != "" {
		options = append(options, WithBaseURL(baseURL))
	}

	if _, _, err := New(secret, options...).Operations.ListEvents(context.Background()).PerPage(5).Execute(); err != nil {
		t.Fatal(err)
	}
}
