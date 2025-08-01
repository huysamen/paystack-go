package main

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/transactions"
	"github.com/huysamen/paystack-go/types"
)

func main() {
	ctx := context.Background()

	// Create client with invalid key to demonstrate error handling
	client := paystack.DefaultClient("sk_test_invalid_key")

	// Create a transaction request
	req := &transactions.TransactionInitializeRequest{
		Amount:   10000, // 100 NGN in kobo
		Email:    "test@example.com",
		Currency: types.CurrencyNGN,
	}

	// Make API call
	resp, err := client.Transactions.Initialize(ctx, req)
	if err != nil {
		// This is a real error (network, parsing, etc.)
		fmt.Printf("❌ System error: %v\n", err)
		return
	}

	// Check if the API call was successful
	if resp.IsError() {
		fmt.Printf("✅ API Error: %s\n", resp.GetErrorMessage())
		fmt.Printf("   Code: %s\n", resp.GetErrorCode())
		fmt.Printf("   Type: %s\n", resp.GetErrorType())

		if resp.IsAuthenticationError() {
			fmt.Printf("   🔑 This is an authentication error\n")
		}
		if resp.IsValidationError() {
			fmt.Printf("   📝 This is a validation error\n")
		}
		if resp.IsNotFoundError() {
			fmt.Printf("   🔍 This is a not found error\n")
		}
		if resp.IsRateLimitError() {
			fmt.Printf("   ⏰ This is a rate limit error\n")
		}

		if nextStep := resp.GetNextStep(); nextStep != "" {
			fmt.Printf("   💡 Next step: %s\n", nextStep)
		}
	} else {
		fmt.Printf("✅ Success! Transaction reference: %s\n", resp.Data.Reference)
	}
}
