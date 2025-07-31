package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huysamen/paystack-go"
	transaction_splits "github.com/huysamen/paystack-go/api/transaction-splits"
	"github.com/huysamen/paystack-go/types"
)

func main() {
	// Initialize client
	client := paystack.DefaultClient("sk_test_your_secret_key")

	fmt.Println("=== Transaction Splits API Examples ===")

	// Example 1: Create a new transaction split
	fmt.Println("\n=== Creating Transaction Split ===")
	createReq := &transaction_splits.TransactionSplitCreateRequest{
		Name:     "Revenue Split",
		Type:     transaction_splits.TransactionSplitTypePercentage,
		Currency: types.CurrencyNGN,
		Subaccounts: []transaction_splits.TransactionSplitSubaccount{
			{
				Subaccount: "ACCT_xxxxxxxxxx", // Replace with actual subaccount code
				Share:      50,                // 50% share
			},
			{
				Subaccount: "ACCT_yyyyyyyyyy", // Replace with actual subaccount code
				Share:      30,                // 30% share
			},
		},
		// Optional: Specify who bears the charges
		BearerType: &[]transaction_splits.TransactionSplitBearerType{
			transaction_splits.TransactionSplitBearerTypeAllProportional,
		}[0],
	}

	createResp, err := client.TransactionSplits.Create(context.Background(), createReq)
	if err != nil {
		log.Printf("Error creating transaction split: %v", err)
	} else {
		fmt.Printf("Transaction split created: %s\n", createResp.Data.SplitCode)
		fmt.Printf("Name: %s\n", createResp.Data.Name)
		fmt.Printf("Type: %s\n", createResp.Data.Type)
		fmt.Printf("Bearer Type: %s\n", createResp.Data.BearerType)
		fmt.Printf("Total Subaccounts: %d\n", createResp.Data.TotalSubaccounts)
	}

	// Example 2: List transaction splits
	fmt.Println("\n=== Listing Transaction Splits ===")
	listReq := &transaction_splits.TransactionSplitListRequest{
		PerPage: &[]int{10}[0],
		Page:    &[]int{1}[0],
		Active:  &[]bool{true}[0], // Only active splits
	}

	listResp, err := client.TransactionSplits.List(context.Background(), listReq)
	if err != nil {
		log.Printf("Error listing transaction splits: %v", err)
	} else {
		fmt.Printf("Found %d transaction splits\n", len(listResp.Data))
		for i, split := range listResp.Data {
			if i >= 3 { // Show only first 3 for demo
				break
			}
			fmt.Printf("Split %d: %s (%s) - %s\n",
				i+1, split.Name, split.SplitCode, split.Type)
		}
	}

	// Example 3: Fetch a specific transaction split
	if createResp != nil && createResp.Data.SplitCode != "" {
		fmt.Println("\n=== Fetching Transaction Split ===")
		fetchResp, err := client.TransactionSplits.Fetch(context.Background(), createResp.Data.SplitCode)
		if err != nil {
			log.Printf("Error fetching transaction split: %v", err)
		} else {
			fmt.Printf("Fetched split: %s\n", fetchResp.Data.Name)
			fmt.Printf("Active: %v\n", fetchResp.Data.Active)
			fmt.Printf("Subaccounts:\n")
			for _, subaccount := range fetchResp.Data.Subaccounts {
				fmt.Printf("  - %s: %d%%\n", subaccount.Subaccount, subaccount.Share)
			}
		}

		// Example 4: Update the transaction split
		fmt.Println("\n=== Updating Transaction Split ===")
		updateReq := &transaction_splits.TransactionSplitUpdateRequest{
			Name: &[]string{"Updated Revenue Split"}[0],
			BearerType: &[]transaction_splits.TransactionSplitBearerType{
				transaction_splits.TransactionSplitBearerTypeAccount,
			}[0],
		}

		updateResp, err := client.TransactionSplits.Update(context.Background(), createResp.Data.SplitCode, updateReq)
		if err != nil {
			log.Printf("Error updating transaction split: %v", err)
		} else {
			fmt.Printf("Split updated: %s\n", updateResp.Data.Name)
			fmt.Printf("New bearer type: %s\n", updateResp.Data.BearerType)
		}

		// Example 5: Add a new subaccount to the split
		fmt.Println("\n=== Adding Subaccount to Split ===")
		addSubaccountReq := &transaction_splits.TransactionSplitSubaccountAddRequest{
			Subaccount: "ACCT_zzzzzzzzzz", // Replace with actual subaccount code
			Share:      20,                // 20% share
		}

		addResp, err := client.TransactionSplits.AddSubaccount(context.Background(), createResp.Data.SplitCode, addSubaccountReq)
		if err != nil {
			log.Printf("Error adding subaccount to split: %v", err)
		} else {
			fmt.Printf("Subaccount added to split: %s\n", addResp.Data.SplitCode)
			fmt.Printf("Total subaccounts now: %d\n", addResp.Data.TotalSubaccounts)
		}

		// Example 6: Remove a subaccount from the split
		fmt.Println("\n=== Removing Subaccount from Split ===")
		removeSubaccountReq := &transaction_splits.TransactionSplitSubaccountRemoveRequest{
			Subaccount: "ACCT_zzzzzzzzzz", // Remove the subaccount we just added
		}

		removeResp, err := client.TransactionSplits.RemoveSubaccount(context.Background(), createResp.Data.SplitCode, removeSubaccountReq)
		if err != nil {
			log.Printf("Error removing subaccount from split: %v", err)
		} else {
			fmt.Printf("Subaccount removed from split: %s\n", removeResp.Message)
		}
	}

	fmt.Println("\n=== Transaction Splits Examples Complete ===")
}
