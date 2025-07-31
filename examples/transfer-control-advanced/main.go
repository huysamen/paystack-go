package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/huysamen/paystack-go"
	transfercontrol "github.com/huysamen/paystack-go/api/transfer-control"
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

	fmt.Println("=== Transfer Control API - Advanced Examples ===\n")

	// 1. Comprehensive Balance Monitoring
	fmt.Println("1. Comprehensive Balance Monitoring")
	fmt.Println("=================================")

	// Check current balance
	balance, err := client.TransferControl.CheckBalance(ctx)
	if err != nil {
		log.Printf("Error checking balance: %v", err)
	} else {
		fmt.Printf("Current Account Balances:\n")
		for _, bal := range balance.Data {
			// Convert kobo to major currency units
			majorAmount := float64(bal.Balance) / 100
			fmt.Printf("  %s: %.2f (Raw: %d kobo)\n", bal.Currency, majorAmount, bal.Balance)
		}
	}
	fmt.Println()

	// 2. Detailed Balance Ledger Analysis
	fmt.Println("2. Detailed Balance Ledger Analysis")
	fmt.Println("=================================")

	ledger, err := client.TransferControl.FetchBalanceLedger(ctx)
	if err != nil {
		log.Printf("Error fetching balance ledger: %v", err)
	} else {
		fmt.Printf("Balance Ledger Analysis:\n")
		fmt.Printf("Total entries: %d\n", len(ledger.Data))

		if len(ledger.Data) > 0 {
			// Analyze recent transactions
			totalInflows := int64(0)
			totalOutflows := int64(0)

			for i, entry := range ledger.Data {
				if i < 5 { // Show latest 5 entries
					fmt.Printf("\nEntry %d:\n", i+1)
					fmt.Printf("  Balance: %d %s\n", entry.Balance, entry.Currency)
					fmt.Printf("  Change: %d\n", entry.Difference)
					fmt.Printf("  Reason: %s\n", entry.Reason)
					fmt.Printf("  Model: %s\n", entry.ModelResponsible)
					fmt.Printf("  Date: %s\n", entry.CreatedAt)
				}

				// Aggregate flows
				if entry.Difference > 0 {
					totalInflows += entry.Difference
				} else {
					totalOutflows += entry.Difference
				}
			}

			fmt.Printf("\nSummary:\n")
			fmt.Printf("  Total Inflows: %d\n", totalInflows)
			fmt.Printf("  Total Outflows: %d\n", totalOutflows)
			fmt.Printf("  Net Change: %d\n", totalInflows+totalOutflows)
		}
	}
	fmt.Println()

	// 3. OTP Management Workflow
	fmt.Println("3. OTP Management Workflow")
	fmt.Println("========================")

	// Enable OTP (safe operation)
	fmt.Println("Step 1: Enabling OTP requirement...")
	enableResp, err := client.TransferControl.EnableOTP(ctx)
	if err != nil {
		log.Printf("Error enabling OTP: %v", err)
	} else {
		fmt.Printf("✓ OTP enabled: %s\n", enableResp.Message)
	}

	// Demonstrate OTP operations (with placeholder data since we need real transfer codes)
	fmt.Println("\nStep 2: OTP Operations (Demo with placeholder data)")
	fmt.Println("Note: These operations require actual transfer codes from real transfers")

	// Example ResendOTP call (will fail without valid transfer code)
	resendReq := &transfercontrol.ResendOTPRequest{
		TransferCode: "TRF_example_code_123",
		Reason:       "resend_otp",
	}

	fmt.Printf("Attempting to resend OTP for transfer: %s\n", resendReq.TransferCode)
	_, err = client.TransferControl.ResendOTP(ctx, resendReq)
	if err != nil {
		fmt.Printf("Expected error (invalid transfer code): %v\n", err)
	}

	// Example FinalizeDisableOTP call (will fail without valid OTP)
	finalizeReq := &transfercontrol.FinalizeDisableOTPRequest{
		OTP: "123456",
	}

	fmt.Printf("Attempting to finalize OTP disable with OTP: %s\n", finalizeReq.OTP)
	_, err = client.TransferControl.FinalizeDisableOTP(ctx, finalizeReq)
	if err != nil {
		fmt.Printf("Expected error (invalid OTP): %v\n", err)
	}

	fmt.Println()

	// 4. Balance Monitoring with Retry Logic
	fmt.Println("4. Balance Monitoring with Retry Logic")
	fmt.Println("====================================")

	maxRetries := 3
	for attempt := 1; attempt <= maxRetries; attempt++ {
		fmt.Printf("Attempt %d/%d: Checking balance...\n", attempt, maxRetries)

		balance, err := client.TransferControl.CheckBalance(ctx)
		if err != nil {
			fmt.Printf("Error on attempt %d: %v\n", attempt, err)
			if attempt < maxRetries {
				fmt.Printf("Retrying in 2 seconds...\n")
				time.Sleep(2 * time.Second)
				continue
			}
		} else {
			fmt.Printf("✓ Balance retrieved successfully on attempt %d\n", attempt)
			for _, bal := range balance.Data {
				fmt.Printf("  %s: %d\n", bal.Currency, bal.Balance)
			}
			break
		}
	}
	fmt.Println()

	// 5. Integration with Transfer Operations
	fmt.Println("5. Integration Patterns")
	fmt.Println("=====================")
	fmt.Println("Best Practices for Transfer Control Integration:")
	fmt.Println("1. Check balance before initiating large transfers")
	fmt.Println("2. Monitor balance ledger for transaction reconciliation")
	fmt.Println("3. Enable OTP for security-sensitive accounts")
	fmt.Println("4. Use resend OTP for customer support scenarios")
	fmt.Println("5. Implement retry logic for network resilience")
	fmt.Println()

	// Example integration pattern
	fmt.Println("Example: Pre-transfer Balance Check")
	balance, err = client.TransferControl.CheckBalance(ctx)
	if err == nil && len(balance.Data) > 0 {
		ngnBalance := balance.Data[0].Balance // Assuming NGN is first
		transferAmount := int64(100000)       // 1000 NGN in kobo

		if ngnBalance >= transferAmount {
			fmt.Printf("✓ Sufficient balance for transfer of %d kobo\n", transferAmount)
			fmt.Printf("  Current balance: %d kobo\n", ngnBalance)
			fmt.Printf("  Remaining after transfer: %d kobo\n", ngnBalance-transferAmount)
		} else {
			fmt.Printf("⚠ Insufficient balance for transfer of %d kobo\n", transferAmount)
			fmt.Printf("  Current balance: %d kobo\n", ngnBalance)
			fmt.Printf("  Shortfall: %d kobo\n", transferAmount-ngnBalance)
		}
	}

	fmt.Println("\n=== Transfer Control Advanced Examples Complete ===")
}
