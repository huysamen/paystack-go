package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/refunds"
)

func main() {
	client := paystack.DefaultClient("sk_test_your_secret_key_here")
	ctx := context.Background()

	// Example 1: Full refund with customer communication
	fmt.Println("=== Creating Full Refund with Customer Communication ===")
	fullRefundReq := &refunds.RefundCreateRequest{
		Transaction:  "T685312322670591", // Replace with actual transaction reference
		CustomerNote: &[]string{"Your refund has been processed due to product unavailability. You should see the amount back in your account within 3-5 business days."}[0],
		MerchantNote: &[]string{"Product out of stock - approved for full refund by customer service team"}[0],
	}

	fullRefund, err := client.Refunds.Create(ctx, fullRefundReq)
	if err != nil {
		log.Printf("Full refund error: %v\n", err)
	} else {
		fmt.Printf("Full refund created: ₦%.2f for transaction %d\n",
			float64(fullRefund.Data.Amount)/100,
			fullRefund.Data.Transaction.ID)
	}

	// Example 2: Partial refund for shipping costs
	fmt.Println("\n=== Creating Partial Refund (Shipping Cost) ===")
	partialRefundReq := &refunds.RefundCreateRequest{
		Transaction:  "T123456789012345", // Replace with actual transaction reference
		Amount:       &[]int{1500}[0],    // ₦15.00 shipping refund
		Currency:     &[]string{"NGN"}[0],
		CustomerNote: &[]string{"Shipping fee refund - your item was delayed beyond our delivery promise"}[0],
		MerchantNote: &[]string{"Goodwill gesture for delivery delay - partial refund for shipping"}[0],
	}

	partialRefund, err := client.Refunds.Create(ctx, partialRefundReq)
	if err != nil {
		log.Printf("Partial refund error: %v\n", err)
	} else {
		fmt.Printf("Partial refund created: ₦%.2f\n", float64(partialRefund.Data.Amount)/100)
	}

	// Example 3: List refunds with advanced filtering
	fmt.Println("\n=== Advanced Refund Filtering and Analytics ===")

	// Filter by transaction
	transactionFilter := &refunds.RefundListRequest{
		Transaction: &[]string{"T685312322670591"}[0],
		PerPage:     &[]int{50}[0],
	}

	transactionRefunds, err := client.Refunds.List(ctx, transactionFilter)
	if err != nil {
		log.Printf("Transaction filter error: %v\n", err)
	} else {
		fmt.Printf("Refunds for specific transaction: %d\n", len(transactionRefunds.Data))
	}

	// Filter by date range (last 30 days)
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	now := time.Now()
	dateRangeFilter := &refunds.RefundListRequest{
		From:     &thirtyDaysAgo,
		To:       &now,
		Currency: &[]string{"NGN"}[0],
		PerPage:  &[]int{100}[0],
	}

	recentRefunds, err := client.Refunds.List(ctx, dateRangeFilter)
	if err != nil {
		log.Printf("Date range filter error: %v\n", err)
	} else {
		fmt.Printf("Refunds in last 30 days: %d\n", len(recentRefunds.Data))

		// Calculate analytics
		if len(recentRefunds.Data) > 0 {
			totalRefunded := 0
			processedCount := 0
			pendingCount := 0
			failedCount := 0

			for _, refund := range recentRefunds.Data {
				totalRefunded += refund.Amount

				switch refund.Status {
				case refunds.RefundStatusProcessed:
					processedCount++
				case refunds.RefundStatusPending:
					pendingCount++
				case refunds.RefundStatusFailed:
					failedCount++
				}
			}

			fmt.Printf("\nRefund Analytics (Last 30 Days):\n")
			fmt.Printf("Total Amount Refunded: ₦%.2f\n", float64(totalRefunded)/100)
			fmt.Printf("Processed: %d\n", processedCount)
			fmt.Printf("Pending: %d\n", pendingCount)
			fmt.Printf("Failed: %d\n", failedCount)
		}
	}

	// Example 4: Detailed refund analysis with channel breakdown
	fmt.Println("\n=== Channel Analysis ===")
	allRefunds, err := client.Refunds.List(ctx, &refunds.RefundListRequest{
		PerPage: &[]int{100}[0],
	})

	if err != nil {
		log.Printf("All refunds error: %v\n", err)
	} else {
		channelStats := make(map[refunds.RefundChannel]struct {
			Count  int
			Amount int
		})

		for _, refund := range allRefunds.Data {
			stats := channelStats[refund.Channel]
			stats.Count++
			stats.Amount += refund.Amount
			channelStats[refund.Channel] = stats
		}

		fmt.Printf("Refunds by Payment Channel:\n")
		for channel, stats := range channelStats {
			fmt.Printf("- %s: %d refunds, ₦%.2f total\n",
				channel.String(),
				stats.Count,
				float64(stats.Amount)/100)
		}
	}

	// Example 5: Refund status monitoring
	fmt.Println("\n=== Refund Status Monitoring ===")
	if len(allRefunds.Data) > 0 {
		for i, refund := range allRefunds.Data {
			if i >= 3 { // Show only first 3 for demo
				break
			}

			detailedRefund, err := client.Refunds.Fetch(ctx, fmt.Sprintf("%d", refund.ID))
			if err != nil {
				log.Printf("Error fetching refund %d: %v\n", refund.ID, err)
				continue
			}

			fmt.Printf("\nRefund %d Status:\n", refund.ID)
			fmt.Printf("  Amount: ₦%.2f (Deducted: ₦%.2f)\n",
				float64(detailedRefund.Data.Amount)/100,
				float64(detailedRefund.Data.DeductedAmount)/100)
			fmt.Printf("  Status: %s\n", detailedRefund.Data.Status)
			fmt.Printf("  Channel: %s\n", detailedRefund.Data.Channel)
			fmt.Printf("  Fully Deducted: %t\n", detailedRefund.Data.FullyDeducted)

			if detailedRefund.Data.RefundedAt != nil {
				fmt.Printf("  Processed: %s\n", detailedRefund.Data.RefundedAt.Time.Format("2006-01-02 15:04:05"))
			}

			if detailedRefund.Data.ExpectedAt != nil {
				fmt.Printf("  Expected: %s\n", detailedRefund.Data.ExpectedAt.Time.Format("2006-01-02 15:04:05"))
			}

			// Time-based analysis
			if detailedRefund.Data.CreatedAt != nil && detailedRefund.Data.RefundedAt != nil {
				processingTime := detailedRefund.Data.RefundedAt.Time.Sub(detailedRefund.Data.CreatedAt.Time)
				fmt.Printf("  Processing Time: %v\n", processingTime.Round(time.Minute))
			}
		}
	}

	// Example 6: Refund dispute correlation
	fmt.Println("\n=== Dispute-Related Refunds ===")
	disputeRefunds := 0
	for _, refund := range allRefunds.Data {
		if refund.Dispute != nil {
			disputeRefunds++
			fmt.Printf("Refund %d linked to dispute %d (₦%.2f)\n",
				refund.ID,
				*refund.Dispute,
				float64(refund.Amount)/100)
		}
	}

	if disputeRefunds == 0 {
		fmt.Printf("No dispute-related refunds found\n")
	} else {
		fmt.Printf("Total dispute-related refunds: %d\n", disputeRefunds)
	}

	fmt.Println("\n=== Refund Management Complete ===")
}
