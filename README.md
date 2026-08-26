# PlainRouter Go SDK

The official Go SDK for the [PlainRouter](https://plainrouter.com) Signals Conversion API. It is generated from PlainRouter's signed OpenAPI contract and published as a standard Go module.

## Install

```bash
go get github.com/plainrouter/sdk-go@v0.5.0
```

## Authenticate

Pass your PlainRouter bearer token through the request context:

```go
package main

import (
	"context"
	"log"

	plainrouter "github.com/plainrouter/sdk-go"
)

func main() {
	config := plainrouter.NewConfiguration()
	client := plainrouter.NewAPIClient(config)
	ctx := context.WithValue(context.Background(), plainrouter.ContextAccessToken, "YOUR_TOKEN")

	report, _, err := client.OperationsAPI.GetEmqReport(ctx).Execute()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", report)
}
```

Keep tokens out of source control; load them from your environment or secret manager.

## Contract and generation

- API contract: signed OpenAPI `0.5.0`
- Generator: OpenAPI Generator `7.25.0`, checksum-pinned by `scripts/generate.sh`
- Module path: `github.com/plainrouter/sdk-go`
- Documentation: [plainrouter.com/docs](https://plainrouter.com/docs)

Run `scripts/generate.sh` to regenerate the client and `scripts/check-generated.sh` to verify that committed output matches the signed contract.

## License

Apache-2.0
