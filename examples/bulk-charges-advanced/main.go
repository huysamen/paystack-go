package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/huysamen/paystack-go"
	bulkcharges "github.com/huysamen/paystack-go/api/bulk-charges"
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

	fmt.Println("=== Bulk Charges API - Advanced Examples ===")
	fmt.Println()

	// 1. Large Scale Bulk Charge with Validation
	fmt.Println("1. Large Scale Bulk Charge Processing")
	fmt.Println("===================================")

	// Simulate a large batch of charges (e.g., salary payments)
	charges := generateSampleCharges(50) // Generate 50 sample charges

	fmt.Printf("Preparing bulk charge with %d items...\n", len(charges))

	// Validate before sending
	totalAmount := int64(0)
	for _, charge := range charges {
		totalAmount += charge.Amount
	}

	fmt.Printf("Total amount to be charged: ₦%.2f\n", float64(totalAmount)/100)
	fmt.Printf("Average per charge: ₦%.2f\n", float64(totalAmount)/float64(len(charges))/100)

	batch, err := client.BulkCharges.Initiate(ctx, charges)
	if err != nil {
		log.Printf("Error initiating bulk charge: %v", err)
	} else {
		fmt.Printf("✓ Bulk charge initiated successfully!\n")
		fmt.Printf("  Batch Code: %s\n", batch.Data.BatchCode)
		fmt.Printf("  Status: %s\n", batch.Data.Status)
		fmt.Printf("  Total Charges: %d\n", batch.Data.TotalCharges)
		fmt.Printf("  Pending Charges: %d\n", batch.Data.PendingCharges)
	}
	fmt.Println()

	// 2. Batch Management and Monitoring
	fmt.Println("2. Batch Management and Monitoring")
	fmt.Println("=================================")

	if batch != nil {
		batchCode := batch.Data.BatchCode

		// Monitor batch progress
		fmt.Printf("Monitoring batch progress: %s\n", batchCode)

		for i := 0; i < 3; i++ {
			time.Sleep(2 * time.Second)

			fetchedBatch, err := client.BulkCharges.Fetch(ctx, batchCode)
			if err != nil {
				log.Printf("Error fetching batch: %v", err)
				break
			}

			fmt.Printf("  Check %d - Status: %s, Pending: %d/%d\n",
				i+1,
				fetchedBatch.Data.Status,
				fetchedBatch.Data.PendingCharges,
				fetchedBatch.Data.TotalCharges)

			// If all charges are processed, break
			if fetchedBatch.Data.PendingCharges == 0 {
				fmt.Printf("  ✓ All charges processed!\n")
				break
			}
		}

		// Demonstrate pause and resume functionality
		fmt.Printf("\nDemonstrating pause/resume functionality:\n")

		// Pause batch
		pauseResp, err := client.BulkCharges.Pause(ctx, batchCode)
		if err != nil {
			log.Printf("Error pausing batch: %v", err)
		} else {
			fmt.Printf("✓ Batch paused: %s\n", pauseResp.Message)
		}

		time.Sleep(1 * time.Second)

		// Resume batch
		resumeResp, err := client.BulkCharges.Resume(ctx, batchCode)
		if err != nil {
			log.Printf("Error resuming batch: %v", err)
		} else {
			fmt.Printf("✓ Batch resumed: %s\n", resumeResp.Message)
		}
	}
	fmt.Println()

	// 3. Detailed Charge Analysis
	fmt.Println("3. Detailed Charge Analysis")
	fmt.Println("===========================")

	if batch != nil {
		batchCode := batch.Data.BatchCode

		// Fetch all charges in the batch
		chargesReq := &bulkcharges.FetchChargesInBatchRequest{
			PerPage: func() *int { p := 100; return &p }(),
			Page:    func() *int { p := 1; return &p }(),
		}

		charges, err := client.BulkCharges.FetchChargesInBatch(ctx, batchCode, chargesReq)
		if err != nil {
			log.Printf("Error fetching charges: %v", err)
		} else {
			// Analyze charge statuses
			statusCounts := make(map[string]int)
			totalSuccessAmount := int64(0)

			for _, charge := range charges.Data {
				statusCounts[charge.Status]++
				if charge.Status == "success" {
					totalSuccessAmount += charge.Amount
				}
			}

			fmt.Printf("Charge Analysis:\n")
			fmt.Printf("  Total Charges: %d\n", len(charges.Data))
			for status, count := range statusCounts {
				fmt.Printf("  %s: %d charges\n", status, count)
			}
			fmt.Printf("  Successful Amount: ₦%.2f\n", float64(totalSuccessAmount)/100)

			// Show sample charges by status
			fmt.Printf("\nSample Charges by Status:\n")
			showSamplesByStatus(charges.Data, "success", 2)
			showSamplesByStatus(charges.Data, "failed", 2)
			showSamplesByStatus(charges.Data, "pending", 2)
		}
	}
	fmt.Println()

	// 4. Historical Batch Analysis
	fmt.Println("4. Historical Batch Analysis")
	fmt.Println("============================")

	// Get recent batches with date filtering
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30).Format("2006-01-02")
	today := time.Now().Format("2006-01-02")

	listReq := &bulkcharges.ListBulkChargeBatchesRequest{
		PerPage: func() *int { p := 20; return &p }(),
		Page:    func() *int { p := 1; return &p }(),
		From:    &thirtyDaysAgo,
		To:      &today,
	}

	batches, err := client.BulkCharges.List(ctx, listReq)
	if err != nil {
		log.Printf("Error listing batches: %v", err)
	} else {
		fmt.Printf("Historical Batches (Last 30 days): %d\n", len(batches.Data))

		// Analyze batch statuses
		batchStatusCounts := make(map[string]int)
		totalBatchCharges := 0

		for _, batch := range batches.Data {
			batchStatusCounts[batch.Status]++
			totalBatchCharges += batch.TotalCharges
		}

		fmt.Printf("Batch Status Summary:\n")
		for status, count := range batchStatusCounts {
			fmt.Printf("  %s: %d batches\n", status, count)
		}
		fmt.Printf("Total Charges Across All Batches: %d\n", totalBatchCharges)

		// Show recent batches
		fmt.Printf("\nRecent Batches:\n")
		for i, batch := range batches.Data {
			if i < 5 { // Show latest 5
				fmt.Printf("  - %s: %s (%d charges) - %s\n",
					batch.BatchCode,
					batch.Status,
					batch.TotalCharges,
					batch.CreatedAt)
			}
		}
	}
	fmt.Println()

	// 5. Best Practices and Integration Patterns
	fmt.Println("5. Best Practices and Integration Patterns")
	fmt.Println("=========================================")
	fmt.Println("Bulk Charges Best Practices:")
	fmt.Println("1. Batch Size: Keep batches under 200 charges for optimal processing")
	fmt.Println("2. Validation: Always validate authorization codes before bulk processing")
	fmt.Println("3. Monitoring: Implement monitoring for batch progress and status")
	fmt.Println("4. Error Handling: Process failed charges separately for retry logic")
	fmt.Println("5. Rate Limiting: Space out large bulk operations to avoid API limits")
	fmt.Println("6. Reconciliation: Use charge status analysis for financial reconciliation")
	fmt.Println()

	fmt.Println("Integration Patterns:")
	fmt.Println("- Payroll Processing: Monthly salary payments to employees")
	fmt.Println("- Subscription Billing: Batch process recurring subscription charges")
	fmt.Println("- Vendor Payments: Bulk payments to multiple suppliers")
	fmt.Println("- Refund Processing: Batch process customer refunds")

	fmt.Println("\n=== Bulk Charges Advanced Examples Complete ===")
}

// generateSampleCharges creates sample charges for demonstration
func generateSampleCharges(count int) bulkcharges.InitiateBulkChargeRequest {
	charges := make(bulkcharges.InitiateBulkChargeRequest, count)

	baseAmounts := []int64{2500, 5000, 7500, 10000, 15000} // Various amounts

	for i := 0; i < count; i++ {
		charges[i] = bulkcharges.BulkChargeItem{
			Authorization: fmt.Sprintf("AUTH_example_%03d", i+1),
			Amount:        baseAmounts[i%len(baseAmounts)],
			Reference:     fmt.Sprintf("bulk-charge-ref-%03d", i+1),
		}
	}

	return charges
}

// showSamplesByStatus displays sample charges filtered by status
func showSamplesByStatus(charges []bulkcharges.BulkChargeCharge, status string, limit int) {
	count := 0
	for _, charge := range charges {
		if charge.Status == status && count < limit {
			fmt.Printf("    %s (%s): ₦%.2f - %s\n",
				status,
				charge.Reference,
				float64(charge.Amount)/100,
				charge.Customer.Email)
			count++
		}
	}
	if count == 0 {
		fmt.Printf("    No %s charges found\n", status)
	}
}
