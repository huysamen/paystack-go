package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/huysamen/paystack-go"
)

func main() {
	// Get API key from environment
	apiKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if apiKey == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable is required")
	}

	// Create client
	client := paystack.DefaultClient(apiKey)
	ctx := context.Background()

	fmt.Println("=== Transfer Control API - Basic Examples ===\n")

	// 1. Check Balance
	fmt.Println("1. Checking balance...")
	balance, err := client.TransferControl.CheckBalance(ctx)
	if err != nil {
		log.Printf("Error checking balance: %v", err)
	} else {
		fmt.Printf("Balance retrieved successfully!\n")
		for _, bal := range balance.Data {
			fmt.Printf("Currency: %s, Balance: %d\n", bal.Currency, bal.Balance)
		}
	}
	fmt.Println()

	// 2. Fetch Balance Ledger
	fmt.Println("2. Fetching balance ledger...")
	ledger, err := client.TransferControl.FetchBalanceLedger(ctx)
	if err != nil {
		log.Printf("Error fetching balance ledger: %v", err)
	} else {
		fmt.Printf("Balance ledger retrieved successfully!\n")
		fmt.Printf("Found %d ledger entries\n", len(ledger.Data))
		if len(ledger.Data) > 0 {
			first := ledger.Data[0]
			fmt.Printf("Latest entry: Balance=%d, Currency=%s, Reason=%s\n",
				first.Balance, first.Currency, first.Reason)
		}
	}
	fmt.Println()

	// 3. Enable OTP (Safe operation)
	fmt.Println("3. Enabling OTP requirement...")
	enableResp, err := client.TransferControl.EnableOTP(ctx)
	if err != nil {
		log.Printf("Error enabling OTP: %v", err)
	} else {
		fmt.Printf("OTP enabled successfully: %s\n", enableResp.Message)
	}
	fmt.Println()

	// Note: Resend OTP, Disable OTP, and Finalize Disable OTP require existing transfers or OTP sessions
	fmt.Println("Note: Resend OTP, Disable OTP, and Finalize Disable OTP operations")
	fmt.Println("require existing transfer codes or active OTP sessions to demonstrate.")

	fmt.Println("\n=== Transfer Control Basic Examples Complete ===")
}
