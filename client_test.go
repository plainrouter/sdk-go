package plainrouter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"sync/atomic"

	"github.com/plainrouter/sdk-go/openapi"
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

// TestContractWireValidation uses a local HTTP fixture; existing transport mocks
// cannot prove that invalid bodies send zero requests or that optional pages are
// absent from a real request URL.
func TestContractWireValidation(t *testing.T) {
	var requests atomic.Int32
	var captured atomic.Value
	var capturedBody atomic.Value
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		captured.Store(r.URL.RawQuery)
		body, _ := io.ReadAll(r.Body)
		capturedBody.Store(body)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()
	client := New("", WithBaseURL(server.URL))
	api := client.OpenAPI()
	ctx := context.Background()
	lists := []struct {
		name string
		call func(*int64) error
	}{
		{"actions", func(page *int64) error {
			r := api.ActionReadApiAPI.ActionsApiIndex(ctx, 1)
			if page != nil {
				r = r.Page(*page)
			}
			_, _, err := r.Execute()
			return err
		}},
		{"action-rules", func(page *int64) error {
			r := api.ActionRuleApiAPI.ActionsRulesIndex(ctx, 1)
			if page != nil {
				r = r.Page(*page)
			}
			_, _, err := r.Execute()
			return err
		}},
		{"tests", func(page *int64) error {
			r := api.AdTestApiAPI.ActionsTestsApiIndex(ctx, 1)
			if page != nil {
				r = r.Page(*page)
			}
			_, _, err := r.Execute()
			return err
		}},
		{"creatives", func(page *int64) error {
			r := api.CreativeAPI.LaunchCreativesIndex(ctx, 1)
			if page != nil {
				r = r.Page(*page)
			}
			_, _, err := r.Execute()
			return err
		}},
		{"plans", func(page *int64) error {
			r := api.DeploymentPlanAPI.LaunchPlansIndex(ctx, 1)
			if page != nil {
				r = r.Page(*page)
			}
			_, _, err := r.Execute()
			return err
		}},
		{"policies", func(page *int64) error {
			r := api.LaunchPolicyAPI.LaunchPoliciesIndex(ctx, 1)
			if page != nil {
				r = r.Page(*page)
			}
			_, _, err := r.Execute()
			return err
		}},
	}
	for _, list := range lists {
		for _, page := range []int64{0, 1, 10000000000000000} {
			t.Run(fmt.Sprintf("%s/page=%d", list.name, page), func(t *testing.T) {
				before := requests.Load()
				var value *int64
				if page != 0 {
					value = &page
				}
				if err := list.call(value); err != nil {
					t.Fatal(err)
				}
				if requests.Load() != before+1 {
					t.Fatal("expected exactly one request")
				}
				query, err := url.ParseQuery(captured.Load().(string))
				if err != nil {
					t.Fatal(err)
				}
				if page == 0 {
					if _, present := query["page"]; present {
						t.Fatalf("unset page sent: %q", captured.Load())
					}
				} else if got := query["page"]; !reflect.DeepEqual(got, []string{strconv.FormatInt(page, 10)}) {
					t.Fatalf("wrong page: %v", got)
				}
			})
		}
	}

	// Exercise all root entry points, including the deprecated raw constructor.
	configuration := NewConfiguration()
	configuration.Servers = openapi.ServerConfigurations{{URL: server.URL}}
	raw := NewAPIClient(configuration)
	clients := []*openapi.APIClient{api, raw}
	for i, candidate := range clients {
		for _, value := range []interface{}{true, false, 1, "1"} {
			t.Run(fmt.Sprintf("client=%d/account_id=%v(%T)", i, value, value), func(t *testing.T) {
				input := openapi.ActionProposalInput{AdditionalProperties: map[string]interface{}{"account_id": value}}
				before := requests.Load()
				_, _, err := candidate.ActionProposalApiAPI.ActionsApiPropose(ctx, 1).ActionProposalInput(input).Execute()
				if _, boolean := value.(bool); boolean {
					if err == nil || !strings.Contains(err.Error(), "account_id") {
						t.Fatalf("expected named integer error, got %v", err)
					}
					if requests.Load() != before {
						t.Fatal("boolean sent a request")
					}
				} else {
					if err != nil {
						t.Fatal(err)
					}
					if requests.Load() != before+1 {
						t.Fatal("integer or numeric string refused")
					}
					var body map[string]json.RawMessage
					if err := json.Unmarshal(capturedBody.Load().([]byte), &body); err != nil {
						t.Fatal(err)
					}
					expected, err := json.Marshal(value)
					if err != nil {
						t.Fatal(err)
					}
					if string(body["account_id"]) != string(expected) {
						t.Fatalf("integer value coerced: got %s, want %s", body["account_id"], expected)
					}
				}
			})
		}
	}
	for _, field := range []string{"daily_budget", "bid_amount", "lifetime_budget"} {
		for _, value := range []interface{}{true, false, 200, "200"} {
			t.Run(fmt.Sprintf("ad_set_spec.%s=%v(%T)", field, value, value), func(t *testing.T) {
				input := openapi.UpdateDeploymentPlanRequest{AdditionalProperties: map[string]interface{}{
					"ad_set_spec": map[string]interface{}{field: value},
				}}
				before := requests.Load()
				_, _, err := api.DeploymentPlanAPI.LaunchPlansUpdate(ctx, 1, "plan").UpdateDeploymentPlanRequest(input).Execute()
				if _, boolean := value.(bool); boolean {
					if err == nil || !strings.Contains(err.Error(), "ad_set_spec."+field) {
						t.Fatalf("expected named integer error, got %v", err)
					}
					if requests.Load() != before {
						t.Fatal("nested boolean sent a request")
					}
				} else {
					if err != nil {
						t.Fatal(err)
					}
					if requests.Load() != before+1 {
						t.Fatal("nested integer or numeric string refused")
					}
				}
			})
		}
	}

	// A parent override must not bypass integers inside closed array members.
	t.Run("evidence.0.action_index/boolean", func(t *testing.T) {
		input := openapi.ActionProposalInput{AdditionalProperties: map[string]interface{}{
			"evidence": []interface{}{map[string]interface{}{"action_index": true}},
		}}
		before := requests.Load()
		_, _, err := api.ActionProposalApiAPI.ActionsApiPropose(ctx, 1).ActionProposalInput(input).Execute()
		if err == nil || !strings.Contains(err.Error(), "evidence.0.action_index") {
			t.Fatalf("expected named integer error, got %v", err)
		}
		if requests.Load() != before {
			t.Fatal("array-member boolean sent a request")
		}
	})

	specBytes, err := os.ReadFile("spec/openapi.json")
	if err != nil {
		t.Fatal(err)
	}
	var spec map[string]interface{}
	if err := json.Unmarshal(specBytes, &spec); err != nil {
		t.Fatal(err)
	}
	schemas := spec["components"].(map[string]interface{})["schemas"].(map[string]interface{})
	proposal := schemas["ActionProposalInput"].(map[string]interface{})
	variants := proposal["properties"].(map[string]interface{})["actions"].(map[string]interface{})["items"].(map[string]interface{})["anyOf"].([]interface{})
	union := reflect.TypeOf(openapi.ActionProposalInputActionsInner{})
	if len(variants) != 14 || union.NumField() != len(variants) {
		t.Fatalf("expected 14 action variants, spec=%d model=%d", len(variants), union.NumField())
	}
	for i, variant := range variants {
		properties := variant.(map[string]interface{})["properties"].(map[string]interface{})
		actionType := properties["type"].(map[string]interface{})["enum"].([]interface{})[0].(string)
		t.Run(actionType+"/closed-params", func(t *testing.T) {
			paramsSchema := properties["params"].(map[string]interface{})
			if paramsSchema["additionalProperties"] != false {
				t.Fatal("params must be closed in signed spec")
			}
			modelName := "ActionProposalInputActionsInnerAnyOf"
			if i != 0 {
				modelName += strconv.Itoa(i)
			}
			model, found := union.FieldByName(modelName)
			if !found {
				t.Fatalf("missing generated variant %s", modelName)
			}
			paramsField, ok := model.Type.Elem().FieldByName("Params")
			if !ok || paramsField.Type.Kind() != reflect.Struct {
				t.Fatalf("params can carry arbitrary keys: %v", paramsField.Type)
			}
			if _, extra := paramsField.Type.FieldByName("AdditionalProperties"); extra {
				t.Fatal("closed params has AdditionalProperties")
			}
			// Fill required scalar fields, then prove the generated decoder refuses
			// an extra key (including the empty pause/resume variants).
			params := map[string]interface{}{}
			paramsProperties, _ := paramsSchema["properties"].(map[string]interface{})
			for key, child := range paramsProperties {
				kind := child.(map[string]interface{})["type"]
				if kinds, ok := kind.([]interface{}); ok {
					kind = kinds[0]
				}
				if kind == "integer" || kind == "number" {
					params[key] = 1
				} else if kind == "array" {
					params[key] = []string{}
				} else {
					params[key] = "1"
				}
			}
			params["unexpected"] = 1
			encoded, err := json.Marshal(params)
			if err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal(encoded, reflect.New(paramsField.Type).Interface()); err == nil || !strings.Contains(err.Error(), "unexpected") {
				t.Fatalf("generated closed decoder: %v", err)
			}
			input := openapi.ActionProposalInput{AdditionalProperties: map[string]interface{}{
				"actions": []interface{}{map[string]interface{}{"type": actionType, "params": params}},
			}}
			before := requests.Load()
			_, _, err = api.ActionProposalApiAPI.ActionsApiPropose(ctx, 1).ActionProposalInput(input).Execute()
			if err == nil || !strings.Contains(err.Error(), "actions.0.params.unexpected") {
				t.Fatalf("expected named closed field error, got %v", err)
			}
			if requests.Load() != before {
				t.Fatal("extra action param sent a request")
			}
		})
	}
}
