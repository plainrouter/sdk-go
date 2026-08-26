package plainrouter_test

import (
	"context"
	"fmt"
	"log"
	"os"

	plainrouter "github.com/plainrouter/sdk-go"
)

func Example() {
	config := plainrouter.NewConfiguration()
	client := plainrouter.NewAPIClient(config)
	ctx := context.WithValue(
		context.Background(),
		plainrouter.ContextAccessToken,
		os.Getenv("PLAINROUTER_TOKEN"),
	)

	report, response, err := client.OperationsAPI.GetEmqReport(ctx).Execute()
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", report)
}
