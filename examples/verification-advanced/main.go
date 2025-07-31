package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/verification"
)

func main() {
	// Create a client with your secret key
	client := paystack.DefaultClient("sk_test_your_secret_key_here")

	// Example 1: Account verification workflow
	fmt.Println("=== Account Verification Workflow ===")

	accounts := []struct {
		number string
		bank   string
		name   string
	}{
		{"0022728151", "063", "Diamond Bank Account"},
		{"0123456789", "058", "GTBank Account"},
		{"9876543210", "044", "Access Bank Account"},
	}

	verifiedAccounts := make([]string, 0)

	for _, account := range accounts {
		fmt.Printf("\nVerifying %s (%s)...\n", account.name, account.number)

		// Step 1: Resolve the account to get the account name
		resolveReq := &verification.AccountResolveRequest{
			AccountNumber: account.number,
			BankCode:      account.bank,
		}

		resolveResp, err := client.Verification.ResolveAccount(context.Background(), resolveReq)
		if err != nil {
			log.Printf("❌ Failed to resolve account: %v", err)
			continue
		}

		fmt.Printf("✅ Account resolved: %s\n", resolveResp.Data.AccountName)
		verifiedAccounts = append(verifiedAccounts, resolveResp.Data.AccountName)

		// Add a small delay to avoid rate limiting
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Printf("\n📊 Successfully verified %d out of %d accounts\n", len(verifiedAccounts), len(accounts))

	// Example 2: Card BIN intelligence analysis
	fmt.Println("\n=== Card BIN Intelligence Analysis ===")

	cardBINs := []struct {
		bin  string
		name string
	}{
		{"539983", "MasterCard"},
		{"408408", "Visa Card"},
		{"506099", "Verve Card"},
		{"627416", "Verve Card"},
		{"411111", "Test Visa"},
	}

	cardAnalysis := make(map[string]int)
	bankAnalysis := make(map[string]int)

	for _, card := range cardBINs {
		fmt.Printf("\nAnalyzing %s (%s)...\n", card.name, card.bin)

		binResp, err := client.Verification.ResolveCardBIN(context.Background(), card.bin)
		if err != nil {
			log.Printf("❌ Failed to resolve BIN: %v", err)
			continue
		}

		fmt.Printf("✅ Brand: %s | Type: %s | Bank: %s | Country: %s\n",
			binResp.Data.Brand,
			binResp.Data.CardType,
			binResp.Data.Bank,
			binResp.Data.CountryName)

		// Track statistics
		cardAnalysis[binResp.Data.Brand]++
		bankAnalysis[binResp.Data.Bank]++

		// Add a small delay to avoid rate limiting
		time.Sleep(500 * time.Millisecond)
	}

	// Display analysis results
	fmt.Println("\n📈 Card Brand Analysis:")
	for brand, count := range cardAnalysis {
		fmt.Printf("  %s: %d cards\n", brand, count)
	}

	fmt.Println("\n🏦 Bank Distribution:")
	for bank, count := range bankAnalysis {
		fmt.Printf("  %s: %d cards\n", bank, count)
	}

	// Example 3: Payment method verification workflow
	fmt.Println("\n=== Payment Method Verification Workflow ===")

	type CustomerPaymentMethod struct {
		CustomerID    string
		AccountNumber string
		BankCode      string
		CardBIN       string
	}

	customers := []CustomerPaymentMethod{
		{
			CustomerID:    "CUST001",
			AccountNumber: "0022728151",
			BankCode:      "063",
			CardBIN:       "539983",
		},
		{
			CustomerID:    "CUST002",
			AccountNumber: "0123456789",
			BankCode:      "058",
			CardBIN:       "408408",
		},
	}

	fmt.Printf("Verifying payment methods for %d customers...\n", len(customers))

	for _, customer := range customers {
		fmt.Printf("\n👤 Customer %s:\n", customer.CustomerID)

		// Verify bank account
		if customer.AccountNumber != "" && customer.BankCode != "" {
			resolveReq := &verification.AccountResolveRequest{
				AccountNumber: customer.AccountNumber,
				BankCode:      customer.BankCode,
			}

			resolveResp, err := client.Verification.ResolveAccount(context.Background(), resolveReq)
			if err != nil {
				fmt.Printf("  🏦 Bank Account: ❌ Verification failed\n")
			} else {
				fmt.Printf("  🏦 Bank Account: ✅ %s\n", resolveResp.Data.AccountName)
			}
		}

		// Verify card
		if customer.CardBIN != "" {
			binResp, err := client.Verification.ResolveCardBIN(context.Background(), customer.CardBIN)
			if err != nil {
				fmt.Printf("  💳 Card: ❌ BIN verification failed\n")
			} else {
				fmt.Printf("  💳 Card: ✅ %s %s from %s\n",
					binResp.Data.Brand,
					binResp.Data.CardType,
					binResp.Data.Bank)
			}
		}

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("\n✨ Verification workflow completed!")
}
