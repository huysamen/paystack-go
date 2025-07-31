package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/settlements"
)

func main() {
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable is required")
	}

	// Create client
	client := paystack.DefaultClient(secretKey)

	// Example 1: List all settlements
	fmt.Println("=== Listing All Settlements ===")
	perPage := 10
	page := 1
	listParams := &settlements.SettlementListRequest{
		PerPage: &perPage,
		Page:    &page,
	}

	response, err := client.Settlements.List(context.Background(), listParams)
	if err != nil {
		log.Printf("Error listing settlements: %v", err)
	} else {
		fmt.Printf("Found %d settlements\n", len(response.Data))
		for _, settlement := range response.Data {
			settledBy := "N/A"
			if settlement.SettledBy != nil {
				settledBy = *settlement.SettledBy
			}
			fmt.Printf("Settlement ID: %d, Amount: %d, Status: %s, Settled By: %s\n",
				settlement.ID, settlement.EffectiveAmount, settlement.Status, settledBy)
		}
	}

	// Example 2: List settlements with status filter
	fmt.Println("\n=== Listing Successful Settlements ===")
	successStatus := settlements.SettlementStatusSuccess
	successPerPage := 5
	successPage := 1
	successParams := &settlements.SettlementListRequest{
		Status:  &successStatus,
		PerPage: &successPerPage,
		Page:    &successPage,
	}

	successResponse, err := client.Settlements.List(context.Background(), successParams)
	if err != nil {
		log.Printf("Error listing successful settlements: %v", err)
	} else {
		fmt.Printf("Found %d successful settlements\n", len(successResponse.Data))
		for _, settlement := range successResponse.Data {
			fmt.Printf("Settlement ID: %d, Amount: %d, Fees: %d\n",
				settlement.ID, settlement.EffectiveAmount, settlement.TotalFees)
		}
	}

	// Example 3: List settlements for a specific subaccount
	fmt.Println("\n=== Listing Settlements for Subaccount ===")
	subaccountCode := "ACCT_xxxxxxxxxx" // Replace with actual subaccount code
	subaccountPerPage := 5
	subaccountPage := 1
	subaccountParams := &settlements.SettlementListRequest{
		Subaccount: &subaccountCode,
		PerPage:    &subaccountPerPage,
		Page:       &subaccountPage,
	}

	subaccountResponse, err := client.Settlements.List(context.Background(), subaccountParams)
	if err != nil {
		log.Printf("Error listing subaccount settlements: %v", err)
	} else {
		fmt.Printf("Found %d settlements for subaccount\n", len(subaccountResponse.Data))
		for _, settlement := range subaccountResponse.Data {
			fmt.Printf("Settlement ID: %d, Amount: %d, Status: %s\n",
				settlement.ID, settlement.EffectiveAmount, settlement.Status)
		}
	}

	// Example 4: List settlement transactions
	if len(response.Data) > 0 {
		settlementID := fmt.Sprintf("%d", response.Data[0].ID)
		fmt.Printf("\n=== Listing Transactions for Settlement %s ===\n", settlementID)

		txPerPage := 10
		txPage := 1
		txParams := &settlements.SettlementTransactionListRequest{
			PerPage: &txPerPage,
			Page:    &txPage,
		}

		txResponse, err := client.Settlements.ListTransactions(context.Background(), settlementID, txParams)
		if err != nil {
			log.Printf("Error listing settlement transactions: %v", err)
		} else {
			fmt.Printf("Found %d transactions in settlement\n", len(txResponse.Data))
			for _, tx := range txResponse.Data {
				fmt.Printf("Transaction ID: %d, Amount: %d, Status: %s, Channel: %s\n",
					tx.ID, tx.Amount, tx.Status, tx.Channel)
			}
		}
	}

	// Example 5: List settlements within date range
	fmt.Println("\n=== Listing Settlements from Last Month ===")
	lastMonth := time.Now().AddDate(0, -1, 0)
	now := time.Now()
	datePerPage := 10
	datePage := 1
	dateParams := &settlements.SettlementListRequest{
		From:    &lastMonth,
		To:      &now,
		PerPage: &datePerPage,
		Page:    &datePage,
	}

	dateResponse, err := client.Settlements.List(context.Background(), dateParams)
	if err != nil {
		log.Printf("Error listing settlements by date: %v", err)
	} else {
		fmt.Printf("Found %d settlements from last month\n", len(dateResponse.Data))
		for _, settlement := range dateResponse.Data {
			fmt.Printf("Settlement Date: %s, Amount: %d, Status: %s\n",
				settlement.SettlementDate.Format("2006-01-02"), settlement.EffectiveAmount, settlement.Status)
		}
	}
}
