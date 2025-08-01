package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/transactions"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

func main() {
	fmt.Println("=== Paystack Go Library - Advanced Error Handling Examples ===")
	fmt.Println()

	// Example 1: Authentication Error
	fmt.Println("1. Testing authentication error with invalid API key")
	fmt.Println("===============================================")

	client := paystack.DefaultClient("sk_test_invalid_key")
	ctx := context.Background()

	req := &transactions.TransactionInitializeRequest{
		Amount:   50000, // ₦500.00 in kobo
		Email:    "customer@example.com",
		Currency: types.CurrencyNGN,
	}

	_, err := client.Transactions.Initialize(ctx, req)
	if err != nil {
		if paystackErr, ok := err.(*net.PaystackError); ok {
			fmt.Printf("✅ Caught Paystack error: %s\n", paystackErr.Error())
			fmt.Printf("   Status Code: %d\n", paystackErr.StatusCode)
			fmt.Printf("   Error Code: %s\n", paystackErr.Code)
			fmt.Printf("   Error Type: %s\n", paystackErr.Type)

			// Check error characteristics
			if paystackErr.IsAuthenticationError() {
				fmt.Printf("   ✅ Correctly identified as authentication error\n")
				if nextStep := paystackErr.GetNextStep(); nextStep != "" {
					fmt.Printf("   💡 Next step: %s\n", nextStep)
				}
			}
			if paystackErr.IsClientError() {
				fmt.Printf("   ✅ Correctly identified as client error (4xx)\n")
			}
		} else {
			fmt.Printf("❌ Expected PaystackError, got: %v\n", err)
		}
	} else {
		fmt.Printf("❌ Expected error, but got success\n")
	}

	fmt.Println()

	// Example 2: Validation Error
	fmt.Println("2. Testing validation error with missing email")
	fmt.Println("===========================================")

	// Use a valid test key format but with invalid request
	clientValid := paystack.DefaultClient("sk_test_" + string(make([]byte, 32))) // Fake but properly formatted key

	invalidReq := &transactions.TransactionInitializeRequest{
		Amount:   0,  // Invalid amount
		Email:    "", // Missing email
		Currency: types.CurrencyNGN,
	}

	_, err = clientValid.Transactions.Initialize(ctx, invalidReq)
	if err != nil {
		if paystackErr, ok := err.(*net.PaystackError); ok {
			fmt.Printf("✅ Caught Paystack error: %s\n", paystackErr.Error())

			// Check error characteristics
			if paystackErr.IsValidationError() {
				fmt.Printf("   ✅ Correctly identified as validation error\n")
			}
			if paystackErr.IsClientError() {
				fmt.Printf("   ✅ Correctly identified as client error\n")
			}
			if !paystackErr.IsServerError() {
				fmt.Printf("   ✅ Correctly identified as NOT a server error\n")
			}
		}
	}

	fmt.Println()

	// Example 3: Demonstrating error constructor functions
	fmt.Println("3. Testing error constructor functions")
	fmt.Println("===================================")

	// Create different types of errors
	validationErr := net.NewValidationError("Email is required and must be valid")
	authErr := net.NewAuthenticationError("API key is missing or invalid")
	notFoundErr := net.NewNotFoundError("Transaction with reference 'xyz123' not found")

	fmt.Printf("Validation Error: %s\n", validationErr.Error())
	fmt.Printf("  - Is Validation Error: %v\n", validationErr.IsValidationError())
	fmt.Printf("  - Is Client Error: %v\n", validationErr.IsClientError())
	fmt.Printf("  - Status Code: %d\n", validationErr.StatusCode)

	fmt.Printf("\nAuthentication Error: %s\n", authErr.Error())
	fmt.Printf("  - Is Authentication Error: %v\n", authErr.IsAuthenticationError())
	fmt.Printf("  - Error Code: %s\n", authErr.Code)

	fmt.Printf("\nNot Found Error: %s\n", notFoundErr.Error())
	fmt.Printf("  - Is Not Found Error: %v\n", notFoundErr.IsNotFoundError())
	fmt.Printf("  - Status Code: %d\n", notFoundErr.StatusCode)

	fmt.Println()

	// Example 4: Error constants
	fmt.Println("4. Available error constants")
	fmt.Println("==========================")
	fmt.Printf("Common Error Codes:\n")
	fmt.Printf("  - Invalid Key: %s\n", net.ErrorCodeInvalidKey)
	fmt.Printf("  - Validation Error: %s\n", net.ErrorCodeValidationError)
	fmt.Printf("  - Insufficient Funds: %s\n", net.ErrorCodeInsufficientFunds)
	fmt.Printf("  - Duplicate Reference: %s\n", net.ErrorCodeDuplicateReference)
	fmt.Printf("  - Transaction Not Found: %s\n", net.ErrorCodeTransactionNotFound)

	fmt.Printf("\nCommon Error Types:\n")
	fmt.Printf("  - Validation: %s\n", net.ErrorTypeValidation)
	fmt.Printf("  - Authentication: %s\n", net.ErrorTypeAuthentication)
	fmt.Printf("  - Not Found: %s\n", net.ErrorTypeNotFound)
	fmt.Printf("  - Server Error: %s\n", net.ErrorTypeServerError)

	fmt.Println()
	fmt.Println("=== Error Handling Examples Complete ===")
}

// demonstrateErrorHandlingPatterns shows common error handling patterns
func demonstrateErrorHandlingPatterns() {
	// This function shows how to handle different types of errors
	// in a production application

	client := paystack.DefaultClient("your_secret_key_here")
	ctx := context.Background()

	req := &transactions.TransactionInitializeRequest{
		Amount:   100000,
		Email:    "customer@example.com",
		Currency: types.CurrencyNGN,
	}

	resp, err := client.Transactions.Initialize(ctx, req)
	if err != nil {
		if paystackErr, ok := err.(*net.PaystackError); ok {
			// Handle different error types appropriately
			switch {
			case paystackErr.IsAuthenticationError():
				log.Printf("Authentication failed: %s", paystackErr.Message)
				// Handle auth errors - maybe refresh API key or notify admin

			case paystackErr.IsValidationError():
				log.Printf("Validation failed: %s", paystackErr.Message)
				if nextStep := paystackErr.GetNextStep(); nextStep != "" {
					log.Printf("Suggestion: %s", nextStep)
				}
				// Handle validation errors - show user-friendly messages

			case paystackErr.IsRateLimitError():
				log.Printf("Rate limited: %s", paystackErr.Message)
				// Handle rate limiting - implement backoff/retry

			case paystackErr.IsNotFoundError():
				log.Printf("Resource not found: %s", paystackErr.Message)
				// Handle not found - check if resource ID is correct

			case paystackErr.IsServerError():
				log.Printf("Server error: %s", paystackErr.Message)
				// Handle server errors - retry with exponential backoff
				// and potentially notify Paystack support

			default:
				log.Printf("Unexpected error: %s", paystackErr.Error())
				// Handle other errors
			}
		} else {
			// Handle non-Paystack errors (network issues, JSON parsing, etc.)
			log.Printf("Non-Paystack error: %v", err)
		}
		return
	}

	// Success - process the response
	fmt.Printf("Transaction initialized: %s\n", resp.Data.AuthorizationURL)
}
