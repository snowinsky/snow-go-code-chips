package api

import "fmt"
import "github.com/go-resty/resty/v2"

func RestyDemo() {
	client := resty.New()

	resp, err := client.
		R().
		EnableTrace().
		Get("https://httpbin.org/get")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
}
