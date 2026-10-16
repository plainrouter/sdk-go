# Plainrouter Go SDK

The official Go SDK for the [Plainrouter](https://plainrouter.com) Signals Conversion API. The root package provides a compact, idiomatic entry point; the complete generated contract is available from [`github.com/plainrouter/sdk-go/openapi`](https://pkg.go.dev/github.com/plainrouter/sdk-go/openapi).

Plainrouter is the paid ads platform for developers and agents. Its hosted [Meta Ads MCP server](https://plainrouter.com/solutions/meta-ads-mcp) lets Claude, ChatGPT, Codex, Cursor and other MCP clients read a Meta ad account and propose changes that pass policy checks. MCP setup, Agent Skills and the other SDKs live in [plainrouter/sdk](https://github.com/plainrouter/sdk).

## Install

```bash
go get github.com/plainrouter/sdk-go@v0.7.0
```

## Authenticate

Create one client with your Plainrouter bearer token:

```go
package main

import (
	"context"
	"log"

	plainrouter "github.com/plainrouter/sdk-go"
)

func main() {
	client := plainrouter.New("YOUR_TOKEN")

	report, _, err := client.Operations.GetEmqReport(context.Background()).Execute()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", report)
}
```

Keep tokens out of source control; load them from your environment or secret manager.

Pass an empty token to use zero-auth synthetic sandbox operations. Applications that need generated wire models, request builders, or low-level configuration can use `client.OpenAPI()` or import the `openapi` subpackage directly.

## Migrating from v0.5

The generated contract moved from the root package to `github.com/plainrouter/sdk-go/openapi`. Existing `NewConfiguration`, `NewAPIClient`, and `ContextAccessToken` calls remain as deprecated migration helpers. Code that names generated request or response types should import the `openapi` subpackage and change the qualifier from `plainrouter` to `openapi`.

## Contract and generation

- API contract: signed OpenAPI `0.5.0` (core main `2d15dc970d92ca22e00d369ab544c113c8cdde77`)
- Generator: OpenAPI Generator `7.25.0`, checksum-pinned by `scripts/generate.sh`
- Module path: `github.com/plainrouter/sdk-go`
- Documentation: [plainrouter.com/docs](https://plainrouter.com/docs)

Run `scripts/generate.sh` to regenerate the `openapi` subpackage and `scripts/check-generated.sh` to verify that committed output matches the signed contract. The root package is repository-owned and intentionally small. Its constructors reject JSON booleans in integer fields and unknown keys in closed request objects before sending, including values supplied through permissive maps. These guards are generated from the signed spec; numbers and numeric strings are preserved.

The six Actions and Launch list operations accept `int64` pages through `10^16`. Leave page unset to omit it and use the server default of page 1. The server also accepts integer strings; the Go setters take integers. See [CHANGELOG.md](CHANGELOG.md) for the complete contract delta.

## License

Apache-2.0
