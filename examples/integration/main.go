// go:build example
//go:build example
// +build example

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/integration"
)

func main() {
	// Create client
	client := paystack.DefaultClient("sk_test_your_secret_key_here")
	ctx := context.Background()

	// Fetch current payment session timeout
	timeout, err := client.Integration.FetchTimeout(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current payment session timeout: %d seconds\n", timeout.Data.PaymentSessionTimeout)

	// Update payment session timeout to 45 seconds using builder pattern
	updateReq := integration.NewUpdateTimeoutRequest(45)

	updatedTimeout, err := client.Integration.UpdateTimeout(ctx, updateReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated payment session timeout: %d seconds\n", updatedTimeout.Data.PaymentSessionTimeout)

	// Set timeout to 0 to disable session timeouts
	disableReq := integration.NewUpdateTimeoutRequest(0)

	disabledTimeout, err := client.Integration.UpdateTimeout(ctx, disableReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Session timeouts disabled (timeout: %d)\n", disabledTimeout.Data.PaymentSessionTimeout)

	// Restore original timeout
	restoreReq := integration.NewUpdateTimeoutRequest(timeout.Data.PaymentSessionTimeout)

	restoredTimeout, err := client.Integration.UpdateTimeout(ctx, restoreReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Restored timeout to original value: %d seconds\n", restoredTimeout.Data.PaymentSessionTimeout)
}
