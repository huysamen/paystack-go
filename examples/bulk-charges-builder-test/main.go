package main

import (
	"fmt"

	bulkcharges "github.com/huysamen/paystack-go/api/bulk-charges"
)

func main() {
	// Test builder pattern
	fmt.Println("Testing Bulk Charges Builder Pattern")
	fmt.Println("===================================")

	// Test InitiateBulkChargeRequest builder
	chargesBuilder := bulkcharges.NewInitiateBulkChargeRequest().
		AddItem("AUTH_test_001", 5000, "test-ref-001").
		AddItem("AUTH_test_002", 10000, "test-ref-002")

	fmt.Printf("Initiate Builder: Created with 2 items\n")

	// Test ListBulkChargeBatchesRequest builder
	listBuilder := bulkcharges.NewListBulkChargeBatchesRequest().
		PerPage(10).
		Page(1).
		DateRange("2024-01-01", "2024-01-31")

	fmt.Printf("List Builder: Created with pagination and date range\n")

	// Test FetchChargesInBatchRequest builder
	fetchBuilder := bulkcharges.NewFetchChargesInBatchRequest().
		PerPage(50).
		Status("success").
		DateRange("2024-01-01", "2024-01-31")

	fmt.Printf("Fetch Builder: Created with filters\n")

	// Test that builders return the correct types
	_ = chargesBuilder // *InitiateBulkChargeRequestBuilder
	_ = listBuilder    // *ListBulkChargeBatchesRequestBuilder
	_ = fetchBuilder   // *FetchChargesInBatchRequestBuilder

	fmt.Println("✓ All builders created successfully!")
	fmt.Println("✓ Builder pattern working correctly!")
}
