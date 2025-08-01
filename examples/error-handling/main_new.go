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

	fmt.Println("Paystack Go SDK - Error Handling Examples")
	fmt.Println("========================================")
	fmt.Println()

	// Example 1: Authentication Error
	fmt.Println("1. Testing authentication error with invalid key")
	fmt.Println("==============================================")

	// Use an invalid API key to trigger authentication error
	client := paystack.DefaultClient("sk_test_invalid_key")

	req := &transactions.TransactionInitializeRequest{
		Amount:   10000, // 100 NGN in kobo
		Email:    "test@example.com",
		Currency: types.CurrencyNGN,
	}

	resp, err := client.Transactions.Initialize(ctx, req)
	if err != nil {
		// This is a real error (network, parsing, etc.)
		fmt.Printf("❌ System error: %v\n", err)
		return
	}

	// Check if the API call was successful
	if resp.IsError() {
		fmt.Printf("✅ Caught Paystack API error: %s\n", resp.GetErrorMessage())
		fmt.Printf("   Error Code: %s\n", resp.GetErrorCode())
		fmt.Printf("   Error Type: %s\n", resp.GetErrorType())

		if nextStep := resp.GetNextStep(); nextStep != "" {
			fmt.Printf("   Next Step: %s\n", nextStep)
		}

		// Check error characteristics
		if resp.IsAuthenticationError() {
			fmt.Printf("   ✅ Correctly identified as authentication error\n")
		}
		if resp.IsValidationError() {
			fmt.Printf("   ✅ Correctly identified as validation error\n")
		}
		if resp.IsNotFoundError() {
			fmt.Printf("   ✅ Correctly identified as not found error\n")
		}
		if resp.IsRateLimitError() {
			fmt.Printf("   ✅ Correctly identified as rate limit error\n")
		}
	} else {
		fmt.Printf("✅ API call was successful!\n")
		fmt.Printf("   Transaction Reference: %s\n", resp.Data.Reference)
	}

	fmt.Println()

	// Example 2: Validation Error
	fmt.Println("2. Testing validation error with invalid data")
	fmt.Println("===========================================")

	// Use a properly formatted test key (won't work but format is correct)
	clientValid := paystack.DefaultClient("sk_test_b74b3d3d4b755d336c9b8d8e4e3a2a1c4b5d6e7f")

	invalidReq := &transactions.TransactionInitializeRequest{
		Amount:   0,  // Invalid amount
		Email:    "", // Missing email - will cause validation error
		Currency: types.CurrencyNGN,
	}

	resp2, err := clientValid.Transactions.Initialize(ctx, invalidReq)
	if err != nil {
		fmt.Printf("❌ System error: %v\n", err)
		return
	}

	if resp2.IsError() {
		fmt.Printf("✅ Caught Paystack API error: %s\n", resp2.GetErrorMessage())
		fmt.Printf("   Error Code: %s\n", resp2.GetErrorCode())
		fmt.Printf("   Error Type: %s\n", resp2.GetErrorType())

		if resp2.IsValidationError() {
			fmt.Printf("   ✅ Correctly identified as validation error\n")
		}
	} else {
		fmt.Printf("✅ API call was successful (unexpected!)\n")
	}

	fmt.Println()

	// Example 3: Successful Response
	fmt.Println("3. Example of successful response handling")
	fmt.Println("=======================================")

	// This would normally work with a real API key and valid data
	fmt.Println("📝 In a real scenario with valid API key and data:")
	fmt.Println("   resp, err := client.Transactions.Initialize(ctx, validReq)")
	fmt.Println("   if err != nil {")
	fmt.Println("       // Handle system errors (network, parsing, etc.)")
	fmt.Println("       log.Fatal(err)")
	fmt.Println("   }")
	fmt.Println("")
	fmt.Println("   if resp.IsError() {")
	fmt.Println("       // Handle API errors")
	fmt.Println("       fmt.Printf(\"API Error: %s\\n\", resp.GetErrorMessage())")
	fmt.Println("       return")
	fmt.Println("   }")
	fmt.Println("")
	fmt.Println("   // Use successful response")
	fmt.Println("   fmt.Printf(\"Success! Reference: %s\\n\", resp.Data.Reference)")

	fmt.Println()
	fmt.Println("Key Benefits of New Error Handling:")
	fmt.Println("====================================")
	fmt.Println("✅ API errors are not Go errors - they're part of the Response")
	fmt.Println("✅ Only real system errors (network, parsing) return Go errors")
	fmt.Println("✅ Easy to check error types with helper methods")
	fmt.Println("✅ All error information available in Response struct")
	fmt.Println("✅ Cleaner, more intuitive error handling")
}
