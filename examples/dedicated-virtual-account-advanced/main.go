package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/huysamen/paystack-go"
	dedicatedvirtualaccount "github.com/huysamen/paystack-go/api/dedicated-virtual-account"
)

func main() {
	// Get secret key from environment variable
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable is required")
	}

	// Initialize client
	client := paystack.DefaultClient(secretKey)
	ctx := context.Background()

	fmt.Println("=== Advanced Dedicated Virtual Account API Examples ===")
	fmt.Println()

	// Example 1: Create account with error handling and validation
	err := createAccountWithValidation(ctx, client)
	if err != nil {
		log.Printf("Create account example failed: %v", err)
	}

	// Example 2: Manage account lifecycle
	err = manageAccountLifecycle(ctx, client)
	if err != nil {
		log.Printf("Account lifecycle example failed: %v", err)
	}

	// Example 3: Handle splits
	err = manageSplitTransactions(ctx, client)
	if err != nil {
		log.Printf("Split transaction example failed: %v", err)
	}

	// Example 4: Monitor accounts
	err = monitorAccounts(ctx, client)
	if err != nil {
		log.Printf("Account monitoring example failed: %v", err)
	}

	fmt.Println("\n=== Advanced Examples Completed ===")
}

func createAccountWithValidation(ctx context.Context, client *paystack.Client) error {
	fmt.Println("1. Creating dedicated virtual account with validation...")

	// First check available providers
	providers, err := client.DedicatedVirtualAccount.FetchBankProviders(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch bank providers: %w", err)
	}

	if len(providers.Data) == 0 {
		return fmt.Errorf("no bank providers available")
	}

	// Select provider
	selectedProvider := providers.Data[0]
	fmt.Printf("Using provider: %s (%s)\n", selectedProvider.BankName, selectedProvider.ProviderSlug)

	// Create account request with full validation
	createReq := &dedicatedvirtualaccount.CreateDedicatedVirtualAccountRequest{
		Customer:      "CUS_example123",
		PreferredBank: selectedProvider.ProviderSlug,
		FirstName:     "John",
		LastName:      "Doe",
		Phone:         "+234801000000",
	}

	account, err := client.DedicatedVirtualAccount.Create(ctx, createReq)
	if err != nil {
		return fmt.Errorf("failed to create account: %w", err)
	}

	fmt.Printf("✓ Account created: %s (%s)\n", account.AccountNumber, account.Bank.Name)
	return nil
}

func manageAccountLifecycle(ctx context.Context, client *paystack.Client) error {
	fmt.Println("\n2. Managing account lifecycle...")

	// List active accounts
	active := true
	listReq := &dedicatedvirtualaccount.ListDedicatedVirtualAccountsRequest{
		Active:   &active,
		Currency: "NGN",
	}

	accounts, err := client.DedicatedVirtualAccount.List(ctx, listReq)
	if err != nil {
		return fmt.Errorf("failed to list accounts: %w", err)
	}

	fmt.Printf("Found %d active accounts\n", len(accounts.Data))

	if len(accounts.Data) == 0 {
		fmt.Println("No accounts found for lifecycle management")
		return nil
	}

	// Get details for first account
	accountID := fmt.Sprintf("%d", accounts.Data[0].ID)
	account, err := client.DedicatedVirtualAccount.Fetch(ctx, accountID)
	if err != nil {
		return fmt.Errorf("failed to fetch account details: %w", err)
	}

	fmt.Printf("✓ Account details: %s - Active: %t, Assigned: %t\n",
		account.AccountNumber, account.Active, account.Assigned)

	// Requery account status
	if account.Bank.Slug != "" {
		requeryReq := &dedicatedvirtualaccount.RequeryDedicatedAccountRequest{
			AccountNumber: account.AccountNumber,
			ProviderSlug:  account.Bank.Slug,
			Date:          time.Now().Format("2006-01-02"),
		}

		requeryResp, err := client.DedicatedVirtualAccount.Requery(ctx, requeryReq)
		if err != nil {
			return fmt.Errorf("failed to requery account: %w", err)
		}

		fmt.Printf("✓ Requery completed: %s\n", requeryResp.Message)
	}

	return nil
}

func manageSplitTransactions(ctx context.Context, client *paystack.Client) error {
	fmt.Println("\n3. Managing split transactions...")

	// List accounts to work with
	listReq := &dedicatedvirtualaccount.ListDedicatedVirtualAccountsRequest{
		Currency: "NGN",
	}

	accounts, err := client.DedicatedVirtualAccount.List(ctx, listReq)
	if err != nil {
		return fmt.Errorf("failed to list accounts: %w", err)
	}

	if len(accounts.Data) == 0 {
		fmt.Println("No accounts available for split management")
		return nil
	}

	// Example: Add split to account (requires valid split code)
	splitReq := &dedicatedvirtualaccount.SplitDedicatedAccountTransactionRequest{
		Customer:      "CUS_example123",
		SplitCode:     "SPL_example123", // Replace with actual split code
		PreferredBank: "wema-bank",      // Replace with actual provider
	}

	splitAccount, err := client.DedicatedVirtualAccount.SplitTransaction(ctx, splitReq)
	if err != nil {
		// Expected to fail with test data, log but continue
		log.Printf("Split addition failed (expected with test data): %v", err)
	} else {
		fmt.Printf("✓ Split added to account: %s\n", splitAccount.AccountNumber)

		// Remove split from account
		removeSplitReq := &dedicatedvirtualaccount.RemoveSplitFromDedicatedAccountRequest{
			AccountNumber: splitAccount.AccountNumber,
		}

		removeSplitResp, err := client.DedicatedVirtualAccount.RemoveSplit(ctx, removeSplitReq)
		if err != nil {
			return fmt.Errorf("failed to remove split: %w", err)
		}

		fmt.Printf("✓ Split removed: %s\n", removeSplitResp.Message)
	}

	return nil
}

func monitorAccounts(ctx context.Context, client *paystack.Client) error {
	fmt.Println("\n4. Monitoring accounts...")

	// Get all providers
	providers, err := client.DedicatedVirtualAccount.FetchBankProviders(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch providers: %w", err)
	}

	fmt.Printf("Available providers (%d):\n", len(providers.Data))
	for _, provider := range providers.Data {
		fmt.Printf("  - %s (%s)\n", provider.BankName, provider.ProviderSlug)
	}

	// Monitor accounts by provider
	for _, provider := range providers.Data[:min(3, len(providers.Data))] { // Check first 3 providers
		listReq := &dedicatedvirtualaccount.ListDedicatedVirtualAccountsRequest{
			ProviderSlug: provider.ProviderSlug,
			Currency:     "NGN",
		}

		accounts, err := client.DedicatedVirtualAccount.List(ctx, listReq)
		if err != nil {
			log.Printf("Failed to list accounts for provider %s: %v", provider.BankName, err)
			continue
		}

		fmt.Printf("Provider %s: %d accounts\n", provider.BankName, len(accounts.Data))

		// Show active vs inactive counts
		activeCount := 0
		assignedCount := 0
		for _, account := range accounts.Data {
			if account.Active {
				activeCount++
			}
			if account.Assigned {
				assignedCount++
			}
		}

		fmt.Printf("  Active: %d, Assigned: %d\n", activeCount, assignedCount)
	}

	fmt.Println("✓ Account monitoring completed")
	return nil
}

// Helper function for Go versions without built-in min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
