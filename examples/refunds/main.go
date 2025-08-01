package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/refunds"
)

func main() {
	client := paystack.DefaultClient("sk_test_your_secret_key_here")
	ctx := context.Background()

	// Create a refund for a transaction
	refundReq := &refunds.RefundCreateRequest{
		Transaction: "T685312322670591", // Replace with actual transaction reference
		// Amount is optional - omitting it will refund the full transaction amount
		// Amount: &[]int{5000}[0], // Partial refund of ₦50.00 in kobo
		// Currency: &[]string{"NGN"}[0], // Optional, defaults to transaction currency
		CustomerNote: &[]string{"Product returned due to defect"}[0],
		MerchantNote: &[]string{"Quality control issue, full refund approved"}[0],
	}

	refund, err := client.Refunds.Create(ctx, refundReq)
	if err != nil {
		log.Fatal("Error creating refund:", err)
	}

	fmt.Printf("Refund initiated successfully!\n")
	fmt.Printf("Transaction ID: %d\n", refund.Data.Transaction.ID)
	fmt.Printf("Refund Amount: ₦%.2f\n", float64(refund.Data.Amount)/100)
	fmt.Printf("Currency: %s\n", refund.Data.Currency)
	fmt.Printf("Status: Queued for processing\n")

	// List recent refunds
	listReq := &refunds.RefundListRequest{
		PerPage: &[]int{10}[0],
		Page:    &[]int{1}[0],
		// Currency: &[]string{"NGN"}[0], // Filter by currency
	}

	refundsList, err := client.Refunds.List(ctx, listReq)
	if err != nil {
		log.Fatal("Error listing refunds:", err)
	}

	fmt.Printf("\nRecent refunds (%d found):\n", len(refundsList.Data))
	for _, r := range refundsList.Data {
		status := "Processed"
		if r.Status == refunds.RefundStatusPending {
			status = "Pending"
		} else if r.Status == refunds.RefundStatusFailed {
			status = "Failed"
		}

		fmt.Printf("- ID %d: ₦%.2f (%s) - %s\n",
			r.ID,
			float64(r.Amount)/100,
			r.Currency,
			status)
	}

	// Fetch detailed refund information
	if len(refundsList.Data) > 0 {
		refundID := fmt.Sprintf("%d", refundsList.Data[0].ID)

		detailedRefund, err := client.Refunds.Fetch(ctx, refundID)
		if err != nil {
			log.Printf("Error fetching refund details: %v\n", err)
		} else {
			fmt.Printf("\nRefund Details (ID: %d):\n", detailedRefund.Data.ID)
			fmt.Printf("Transaction: %d\n", detailedRefund.Data.Transaction)
			fmt.Printf("Amount: ₦%.2f\n", float64(detailedRefund.Data.Amount)/100)
			fmt.Printf("Deducted: ₦%.2f\n", float64(detailedRefund.Data.DeductedAmount)/100)
			fmt.Printf("Channel: %s\n", detailedRefund.Data.Channel)
			fmt.Printf("Status: %s\n", detailedRefund.Data.Status)
			fmt.Printf("Refunded By: %s\n", detailedRefund.Data.RefundedBy)

			if detailedRefund.Data.RefundedAt != nil {
				fmt.Printf("Refunded At: %s\n", detailedRefund.Data.RefundedAt.Time.Format("2006-01-02 15:04:05"))
			}

			if detailedRefund.Data.CustomerNote != nil {
				fmt.Printf("Customer Note: %s\n", *detailedRefund.Data.CustomerNote)
			}

			if detailedRefund.Data.MerchantNote != nil {
				fmt.Printf("Merchant Note: %s\n", *detailedRefund.Data.MerchantNote)
			}
		}
	}
}
