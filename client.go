package plainrouter

import (
	"net/http"
	"strings"

	"github.com/plainrouter/sdk-go/openapi"
)

// Client is the curated entry point for the PlainRouter API.
//
// Events, Operations, and Sandbox expose the generated service groups without
// placing every generated model and accessor in the root package index.
type Client struct {
	Events     *openapi.EventAPIService
	Operations *openapi.OperationsAPIService
	Sandbox    *openapi.SandboxAPIService

	openapi *openapi.APIClient
}

// Option configures a Client created by [New].
type Option func(*openapi.Configuration)

// ContextAccessToken is the generated client's bearer-token context key.
//
// Deprecated: create a curated client with [New], which applies its token to
// every request automatically.
var ContextAccessToken = openapi.ContextAccessToken

// New creates a PlainRouter client. Pass an empty token for the zero-auth
// synthetic sandbox operations.
func New(token string, options ...Option) *Client {
	configuration := openapi.NewConfiguration()
	if token != "" {
		configuration.AddDefaultHeader("Authorization", "Bearer "+token)
	}
	for _, option := range options {
		option(configuration)
	}

	generated := openapi.NewAPIClient(configuration)
	return &Client{
		Events:     generated.EventAPI,
		Operations: generated.OperationsAPI,
		Sandbox:    generated.SandboxAPI,
		openapi:    generated,
	}
}

// NewConfiguration creates low-level generated client configuration.
//
// Deprecated: use [New] with [WithBaseURL], [WithHTTPClient], or
// [WithUserAgent].
func NewConfiguration() *openapi.Configuration {
	return openapi.NewConfiguration()
}

// NewAPIClient creates the complete generated client.
//
// Deprecated: use [New] for the curated entry point or import the
// github.com/plainrouter/sdk-go/openapi package directly.
func NewAPIClient(configuration *openapi.Configuration) *openapi.APIClient {
	return openapi.NewAPIClient(configuration)
}

// OpenAPI returns the complete generated client for advanced use cases.
func (client *Client) OpenAPI() *openapi.APIClient {
	return client.openapi
}

// WithBaseURL overrides the default PlainRouter API base URL. It is intended
// for tests, proxies, and self-hosted compatible endpoints.
func WithBaseURL(baseURL string) Option {
	return func(configuration *openapi.Configuration) {
		configuration.Servers = openapi.ServerConfigurations{{
			URL:         strings.TrimRight(baseURL, "/"),
			Description: "Custom",
		}}
	}
}

// WithHTTPClient supplies the HTTP client used for requests. A nil client is
// ignored and leaves the standard HTTP client in place.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(configuration *openapi.Configuration) {
		if httpClient != nil {
			configuration.HTTPClient = httpClient
		}
	}
}

// WithUserAgent overrides the generated client's default User-Agent header.
func WithUserAgent(userAgent string) Option {
	return func(configuration *openapi.Configuration) {
		configuration.UserAgent = userAgent
	}
}
