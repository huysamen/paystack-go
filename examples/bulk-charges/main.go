package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	fmt.Println("=== Bulk Charges API - Basic Examples ===\n")

	// 1. Initiate Bulk Charge
	fmt.Println("1. Initiating bulk charge...")
	charges := bulkcharges.InitiateBulkChargeRequest{
		{
			Authorization: "AUTH_example_123",
			Amount:        2500, // ₦25.00 in kobo
			Reference:     "bulk-charge-ref-001",
		},
		{
			Authorization: "AUTH_example_456",
			Amount:        5000, // ₦50.00 in kobo
			Reference:     "bulk-charge-ref-002",
		},
		{
			Authorization: "AUTH_example_789",
			Amount:        7500, // ₦75.00 in kobo
			Reference:     "bulk-charge-ref-003",
		},
	}

	batch, err := client.BulkCharges.Initiate(ctx, charges)
	if err != nil {
		log.Printf("Error initiating bulk charge: %v", err)
		// Continue with other examples using a placeholder batch code
		fmt.Println("Using placeholder batch code for demonstration")
	} else {
		fmt.Printf("Bulk charge initiated successfully!\n")
		fmt.Printf("Batch Code: %s\n", batch.Data.BatchCode)
		fmt.Printf("Reference: %s\n", batch.Data.Reference)
		fmt.Printf("Status: %s\n", batch.Data.Status)
		fmt.Printf("Total Charges: %d\n", batch.Data.TotalCharges)
		fmt.Printf("Pending Charges: %d\n", batch.Data.PendingCharges)
	}
	fmt.Println()

	// 2. List Bulk Charge Batches
	fmt.Println("2. Listing bulk charge batches...")
	perPage := 10
	page := 1
	listReq := &bulkcharges.ListBulkChargeBatchesRequest{
		PerPage: &perPage,
		Page:    &page,
	}

	batches, err := client.BulkCharges.List(ctx, listReq)
	if err != nil {
		log.Printf("Error listing bulk charge batches: %v", err)
	} else {
		fmt.Printf("Found %d bulk charge batches\n", len(batches.Data))
		for i, batch := range batches.Data {
			if i < 3 { // Show first 3 batches
				fmt.Printf("  - %s: %s (%d charges)\n",
					batch.BatchCode, batch.Status, batch.TotalCharges)
			}
		}
	}
	fmt.Println()

	// 3. Fetch Specific Bulk Charge Batch
	fmt.Println("3. Fetching bulk charge batch details...")
	if batch != nil && batch.Data.BatchCode != "" {
		fetchedBatch, err := client.BulkCharges.Fetch(ctx, batch.Data.BatchCode)
		if err != nil {
			log.Printf("Error fetching batch: %v", err)
		} else {
			fmt.Printf("Batch Details:\n")
			fmt.Printf("  Code: %s\n", fetchedBatch.Data.BatchCode)
			fmt.Printf("  Status: %s\n", fetchedBatch.Data.Status)
			fmt.Printf("  Total Charges: %d\n", fetchedBatch.Data.TotalCharges)
			fmt.Printf("  Pending Charges: %d\n", fetchedBatch.Data.PendingCharges)
		}
	} else {
		fmt.Println("No valid batch code available for fetching")
	}
	fmt.Println()

	// 4. Fetch Charges in a Batch
	fmt.Println("4. Fetching charges in batch...")
	if batch != nil && batch.Data.BatchCode != "" {
		chargesReq := &bulkcharges.FetchChargesInBatchRequest{
			PerPage: &perPage,
			Page:    &page,
		}

		charges, err := client.BulkCharges.FetchChargesInBatch(ctx, batch.Data.BatchCode, chargesReq)
		if err != nil {
			log.Printf("Error fetching charges in batch: %v", err)
		} else {
			fmt.Printf("Found %d charges in batch\n", len(charges.Data))
			for i, charge := range charges.Data {
				if i < 3 { // Show first 3 charges
					fmt.Printf("  - %s: ₦%.2f (%s)\n",
						charge.Reference,
						float64(charge.Amount)/100,
						charge.Status)
				}
			}
		}
	} else {
		fmt.Println("No valid batch code available for fetching charges")
	}
	fmt.Println()

	// Note: Pause and Resume operations require active batches
	fmt.Println("Note: Pause and Resume operations require active bulk charge batches.")
	fmt.Println("These would be used to control processing of large bulk charge batches.")

	fmt.Println("\n=== Bulk Charges Basic Examples Complete ===")
}
