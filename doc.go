// Package plainrouter provides the official Go client for the PlainRouter
// Signals Conversion API.
//
// # Getting started
//
// Create one APIClient and reuse it. NewConfiguration targets the production
// PlainRouter API by default. Account-scoped operations require a bearer token
// in the request context:
//
//	config := plainrouter.NewConfiguration()
//	client := plainrouter.NewAPIClient(config)
//	ctx := context.WithValue(
//		context.Background(),
//		plainrouter.ContextAccessToken,
//		os.Getenv("PLAINROUTER_TOKEN"),
//	)
//	report, response, err := client.OperationsAPI.GetEmqReport(ctx).Execute()
//
// Keep tokens out of source code and load them from an environment variable or
// secret manager. Most methods return a decoded result, the underlying
// HTTP response, and an error. Inspect the response when handling API errors or
// reading response metadata.
//
// # API groups
//
// The client exposes three service groups:
//
//   - APIClient.EventAPI submits conversion events and reads event delivery
//     state.
//   - APIClient.OperationsAPI provides reporting, reconciliation, replay,
//     deletion, and destination test operations.
//   - APIClient.SandboxAPI discovers and validates isolated synthetic sandbox
//     events without granting access to production data.
//
// See the [PlainRouter documentation] for authentication, consent-aware event
// shapes, sandbox usage, and operational guidance.
//
// [PlainRouter documentation]: https://plainrouter.com/docs
package plainrouter
