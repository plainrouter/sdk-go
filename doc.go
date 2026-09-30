// Package plainrouter provides the official Go client for the Plainrouter
// Signals Conversion API.
//
// # Getting started
//
// Create one Client and reuse it. New targets the production Plainrouter API
// by default and applies the bearer token to every request:
//
//	client := plainrouter.New(os.Getenv("PLAINROUTER_TOKEN"))
//	report, response, err := client.Operations.GetEmqReport(context.Background()).Execute()
//
// Keep tokens out of source code and load them from an environment variable or
// secret manager. Most methods return a decoded result, the underlying
// HTTP response, and an error. Inspect the response when handling API errors or
// reading response metadata.
//
// # API groups
//
// Client exposes three service groups:
//
//   - Client.Events submits conversion events and reads event delivery
//     state.
//   - Client.Operations provides reporting, reconciliation, replay,
//     deletion, and destination test operations.
//   - Client.Sandbox discovers and validates isolated synthetic sandbox
//     events without granting access to production data.
//
// The complete generated contract is available from the
// [github.com/plainrouter/sdk-go/openapi] subpackage and through Client.OpenAPI.
//
// See the [Plainrouter documentation] for authentication, consent-aware event
// shapes, sandbox usage, and operational guidance.
//
// [Plainrouter documentation]: https://plainrouter.com/docs
package plainrouter
