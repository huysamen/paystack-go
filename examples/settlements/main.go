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
	listReq := settlements.NewSettlementListRequest().
		PerPage(10).
		Page(1)

	response, err := client.Settlements.List(context.Background(), listReq)
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
	successReq := settlements.NewSettlementListRequest().
		Status(settlements.SettlementStatusSuccess).
		PerPage(5).
		Page(1)

	successResponse, err := client.Settlements.List(context.Background(), successReq)
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
	subaccountReq := settlements.NewSettlementListRequest().
		Subaccount(subaccountCode).
		PerPage(5).
		Page(1)

	subaccountResponse, err := client.Settlements.List(context.Background(), subaccountReq)
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

		txReq := settlements.NewSettlementTransactionListRequest().
			PerPage(10).
			Page(1)

		txResponse, err := client.Settlements.ListTransactions(context.Background(), settlementID, txReq)
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
	dateReq := settlements.NewSettlementListRequest().
		DateRange(lastMonth, now).
		PerPage(10).
		Page(1)

	dateResponse, err := client.Settlements.List(context.Background(), dateReq)
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
