// go:build example
//go:build example
// +build example

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/integration"
)

func main() {
	// Create client with custom configuration
	config := paystack.NewConfig("sk_test_your_secret_key_here").
		WithTimeout(30 * time.Second)

	client := paystack.NewClient(config)
	ctx := context.Background()

	fmt.Println("=== Integration Management Demo ===\n")

	// 1. Fetch current timeout settings
	fmt.Println("1. Fetching current payment session timeout...")
	originalTimeout, err := client.Integration.FetchTimeout(ctx)
	if err != nil {
		log.Fatal("Failed to fetch timeout:", err)
	}

	fmt.Printf("   Current timeout: %d seconds\n\n", originalTimeout.Data.PaymentSessionTimeout)

	// 2. Configure timeout for different scenarios
	scenarios := []struct {
		name        string
		timeout     int
		description string
	}{
		{
			name:        "Quick Checkout",
			timeout:     15,
			description: "Short timeout for simple payments (15 seconds)",
		},
		{
			name:        "Standard Checkout",
			timeout:     30,
			description: "Standard timeout for regular transactions (30 seconds)",
		},
		{
			name:        "Extended Checkout",
			timeout:     60,
			description: "Extended timeout for complex forms (60 seconds)",
		},
		{
			name:        "Unlimited Session",
			timeout:     0,
			description: "Disable session timeout (0 = unlimited)",
		},
	}

	for i, scenario := range scenarios {
		fmt.Printf("%d. Configuring %s...\n", i+2, scenario.name)
		fmt.Printf("   %s\n", scenario.description)

		updateReq := &integration.UpdateTimeoutRequest{
			Timeout: scenario.timeout,
		}

		updatedTimeout, err := client.Integration.UpdateTimeout(ctx, updateReq)
		if err != nil {
			log.Printf("   ❌ Failed to update timeout: %v\n", err)
			continue
		}

		timeoutText := fmt.Sprintf("%d seconds", updatedTimeout.Data.PaymentSessionTimeout)
		if updatedTimeout.Data.PaymentSessionTimeout == 0 {
			timeoutText = "unlimited"
		}

		fmt.Printf("   ✅ Successfully updated to: %s\n", timeoutText)

		// Verify the change
		verifyTimeout, err := client.Integration.FetchTimeout(ctx)
		if err != nil {
			log.Printf("   ⚠️  Could not verify timeout change: %v\n", err)
		} else if verifyTimeout.Data.PaymentSessionTimeout == scenario.timeout {
			fmt.Printf("   ✅ Verified: Timeout is now %s\n", timeoutText)
		} else {
			fmt.Printf("   ⚠️  Verification failed: Expected %d, got %d\n",
				scenario.timeout, verifyTimeout.Data.PaymentSessionTimeout)
		}

		fmt.Println()

		// Add a small delay between updates
		time.Sleep(1 * time.Second)
	}

	// 3. Restore original timeout
	fmt.Println("5. Restoring original timeout settings...")
	restoreReq := &integration.UpdateTimeoutRequest{
		Timeout: originalTimeout.Data.PaymentSessionTimeout,
	}

	restoredTimeout, err := client.Integration.UpdateTimeout(ctx, restoreReq)
	if err != nil {
		log.Fatal("Failed to restore original timeout:", err)
	}

	fmt.Printf("   ✅ Restored to original timeout: %d seconds\n\n", restoredTimeout.Data.PaymentSessionTimeout)

	// 4. Best practices and recommendations
	fmt.Println("=== Best Practices for Payment Session Timeouts ===")
	fmt.Println()
	fmt.Println("• 15-30 seconds: Ideal for simple payment forms")
	fmt.Println("• 30-60 seconds: Good for forms with multiple fields")
	fmt.Println("• 60+ seconds: Use for complex checkout processes")
	fmt.Println("• 0 (unlimited): Use with caution, may impact security")
	fmt.Println()
	fmt.Println("Consider your customer experience:")
	fmt.Println("• Shorter timeouts improve security but may frustrate slow users")
	fmt.Println("• Longer timeouts are user-friendly but may expose sessions")
	fmt.Println("• Monitor your conversion rates when adjusting timeouts")
	fmt.Println()

	// 5. Example error handling
	fmt.Println("=== Error Handling Example ===")

	// Try to set an invalid timeout (negative value)
	invalidReq := &integration.UpdateTimeoutRequest{
		Timeout: -1,
	}

	_, err = client.Integration.UpdateTimeout(ctx, invalidReq)
	if err != nil {
		fmt.Printf("✅ Correctly handled invalid timeout: %v\n", err)
	} else {
		fmt.Println("⚠️  Expected error for negative timeout, but update succeeded")
	}

	fmt.Println("\n=== Integration Management Demo Complete ===")
}
