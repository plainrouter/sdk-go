package plainrouter_test

import (
	"context"
	"fmt"
	"log"
	"os"

	plainrouter "github.com/plainrouter/sdk-go"
)

func Example() {
	client := plainrouter.New(os.Getenv("PLAINROUTER_TOKEN"))

	report, response, err := client.Operations.GetEmqReport(context.Background()).Execute()
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", report)
}
