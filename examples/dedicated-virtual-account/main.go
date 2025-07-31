package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huysamen/paystack-go"
	dedicatedvirtualaccount "github.com/huysamen/paystack-go/api/dedicated-virtual-account"
)

func main() {
	// Initialize client with test secret key
	client := paystack.DefaultClient("sk_test_your_secret_key")
	ctx := context.Background()

	fmt.Println("=== Paystack Dedicated Virtual Account API Examples ===")
	fmt.Println()

	// Example 1: Fetch Bank Providers
	fmt.Println("1. Fetching available bank providers...")
	providers, err := client.DedicatedVirtualAccount.FetchBankProviders(ctx)
	if err != nil {
		log.Printf("Error fetching bank providers: %v", err)
		return
	}
	fmt.Printf("Found %d bank providers\n", len(providers.Data))
	var preferredBank string
	for _, provider := range providers.Data {
		fmt.Printf("- %s (%s) - ID: %d\n", provider.BankName, provider.ProviderSlug, provider.BankID)
		if preferredBank == "" {
			preferredBank = provider.ProviderSlug // Use first provider for examples
		}
	}

	// Example 2: Create Dedicated Virtual Account for existing customer
	fmt.Println("\n2. Creating dedicated virtual account for existing customer...")
	createReq := &dedicatedvirtualaccount.CreateDedicatedVirtualAccountRequest{
		Customer:      "CUS_xnxdt6s1zg5f4nx", // Replace with actual customer code
		PreferredBank: preferredBank,
		FirstName:     "John",
		LastName:      "Doe",
		Phone:         "+2348100000000",
	}

	account, err := client.DedicatedVirtualAccount.Create(ctx, createReq)
	if err != nil {
		log.Printf("Error creating dedicated virtual account: %v", err)
		// Continue with other examples
	} else {
		fmt.Printf("Created account: %s (%s)\n", account.AccountNumber, account.AccountName)
		fmt.Printf("Bank: %s, Currency: %s, Active: %t\n", account.Bank.Name, account.Currency, account.Active)
	}

	// Example 3: Assign Dedicated Virtual Account (create customer and assign)
	fmt.Println("\n3. Assigning dedicated virtual account (create customer and assign)...")
	assignReq := &dedicatedvirtualaccount.AssignDedicatedVirtualAccountRequest{
		Email:         "jane.doe@example.com",
		FirstName:     "Jane",
		LastName:      "Doe",
		Phone:         "+2348100000001",
		PreferredBank: preferredBank,
		Country:       "NG",
		MiddleName:    "Ann",
	}

	assignResp, err := client.DedicatedVirtualAccount.Assign(ctx, assignReq)
	if err != nil {
		log.Printf("Error assigning dedicated virtual account: %v", err)
	} else {
		fmt.Printf("Assignment response: %s\n", assignResp.Message)
	}

	// Example 4: List Dedicated Virtual Accounts
	fmt.Println("\n4. Listing dedicated virtual accounts...")
	active := true
	listReq := &dedicatedvirtualaccount.ListDedicatedVirtualAccountsRequest{
		Active:   &active,
		Currency: "NGN",
	}

	accounts, err := client.DedicatedVirtualAccount.List(ctx, listReq)
	if err != nil {
		log.Printf("Error listing dedicated virtual accounts: %v", err)
		return
	}
	fmt.Printf("Found %d active dedicated virtual accounts\n", len(accounts.Data))

	var accountID string
	for _, acc := range accounts.Data {
		fmt.Printf("- %s: %s (%s) - %s\n", acc.AccountNumber, acc.AccountName, acc.Bank.Name, acc.Currency)
		if acc.Customer != nil {
			fmt.Printf("  Customer: %s (%s)\n", acc.Customer.Email, acc.Customer.CustomerCode)
		}
		if accountID == "" && acc.ID > 0 {
			accountID = fmt.Sprintf("%d", acc.ID) // Use first account for other examples
		}
	}

	if accountID != "" {
		// Example 5: Fetch Dedicated Virtual Account
		fmt.Println("\n5. Fetching dedicated virtual account details...")
		fetchedAccount, err := client.DedicatedVirtualAccount.Fetch(ctx, accountID)
		if err != nil {
			log.Printf("Error fetching dedicated virtual account: %v", err)
		} else {
			fmt.Printf("Account details: %s (%s)\n", fetchedAccount.AccountNumber, fetchedAccount.AccountName)
			fmt.Printf("Bank: %s, Assigned: %t, Active: %t\n",
				fetchedAccount.Bank.Name, fetchedAccount.Assigned, fetchedAccount.Active)
		}

		// Example 6: Split Transaction
		fmt.Println("\n6. Adding split to dedicated virtual account...")
		splitReq := &dedicatedvirtualaccount.SplitDedicatedAccountTransactionRequest{
			Customer:      "CUS_xnxdt6s1zg5f4nx", // Replace with actual customer code
			SplitCode:     "SPL_98WF13Zu8w5",     // Replace with actual split code
			PreferredBank: preferredBank,
		}

		splitAccount, err := client.DedicatedVirtualAccount.SplitTransaction(ctx, splitReq)
		if err != nil {
			log.Printf("Error adding split to account: %v", err)
		} else {
			fmt.Printf("Split added to account: %s\n", splitAccount.AccountNumber)
		}

		// Example 7: Remove Split
		fmt.Println("\n7. Removing split from dedicated virtual account...")
		removeSplitReq := &dedicatedvirtualaccount.RemoveSplitFromDedicatedAccountRequest{
			AccountNumber: fetchedAccount.AccountNumber,
		}

		removeSplitResp, err := client.DedicatedVirtualAccount.RemoveSplit(ctx, removeSplitReq)
		if err != nil {
			log.Printf("Error removing split: %v", err)
		} else {
			fmt.Printf("Split removed: %s\n", removeSplitResp.Message)
		}
	}

	// Example 8: Requery Dedicated Account
	if len(accounts.Data) > 0 {
		fmt.Println("\n8. Requerying dedicated virtual account...")
		requeryReq := &dedicatedvirtualaccount.RequeryDedicatedAccountRequest{
			AccountNumber: accounts.Data[0].AccountNumber,
			ProviderSlug:  accounts.Data[0].Bank.Slug,
			Date:          "2024-01-15", // Replace with actual date
		}

		requeryResp, err := client.DedicatedVirtualAccount.Requery(ctx, requeryReq)
		if err != nil {
			log.Printf("Error requerying account: %v", err)
		} else {
			fmt.Printf("Requery response: %s\n", requeryResp.Message)
		}
	}

	// Example 9: List by provider (filter by bank)
	fmt.Println("\n9. Listing accounts by provider...")
	providerListReq := &dedicatedvirtualaccount.ListDedicatedVirtualAccountsRequest{
		ProviderSlug: "wema-bank", // Replace with actual provider slug
		Currency:     "NGN",
	}

	providerAccounts, err := client.DedicatedVirtualAccount.List(ctx, providerListReq)
	if err != nil {
		log.Printf("Error listing accounts by provider: %v", err)
	} else {
		fmt.Printf("Found %d accounts with Wema Bank provider\n", len(providerAccounts.Data))
	}

	// Example 10: Deactivate Account (commented out to prevent accidental deactivation)
	/*
		if accountID != "" {
			fmt.Println("\n10. Deactivating dedicated virtual account...")
			deactivatedAccount, err := client.DedicatedVirtualAccount.Deactivate(ctx, accountID)
			if err != nil {
				log.Printf("Error deactivating account: %v", err)
			} else {
				fmt.Printf("Account deactivated: %s (Assigned: %t)\n",
					deactivatedAccount.AccountNumber, deactivatedAccount.Assigned)
			}
		}
	*/
	fmt.Println("\n10. Account deactivation example skipped (uncomment to test)")

	fmt.Println("\n=== Dedicated Virtual Account API Examples Completed ===")
}
