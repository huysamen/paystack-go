package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/huysamen/paystack-go"
	directdebit "github.com/huysamen/paystack-go/api/direct-debit"
)

func main() {
	// Get secret key from environment variable
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable is required")
	}

	// Initialize client
	client := paystack.DefaultClient(secretKey)
	ctx := context.Background()

	fmt.Println("=== Advanced Direct Debit API Examples ===")
	fmt.Println()

	// Example 1: Comprehensive mandate listing with filtering
	err := comprehensiveMandateListing(ctx, client)
	if err != nil {
		log.Printf("Comprehensive mandate listing failed: %v", err)
	}

	// Example 2: Batch activation with error handling
	err = batchActivationWithErrorHandling(ctx, client)
	if err != nil {
		log.Printf("Batch activation example failed: %v", err)
	}

	// Example 3: Status-based mandate management
	err = statusBasedMandateManagement(ctx, client)
	if err != nil {
		log.Printf("Status-based management example failed: %v", err)
	}

	fmt.Println("\n=== Advanced Direct Debit Examples Complete ===")
}

// comprehensiveMandateListing demonstrates advanced filtering and pagination
func comprehensiveMandateListing(ctx context.Context, client *paystack.Client) error {
	fmt.Println("1. Comprehensive Mandate Listing with Advanced Filtering")
	fmt.Println("--------------------------------------------------------")

	// List all mandates with pagination
	allBuilder := directdebit.NewListMandateAuthorizationsBuilder().
		PerPage(50)

	allMandates, err := client.DirectDebit.ListMandateAuthorizations(ctx, allBuilder)
	if err != nil {
		return fmt.Errorf("failed to list all mandates: %w", err)
	}

	fmt.Printf("Total mandates found: %d\n", len(allMandates.Data))

	// Filter by each status
	statuses := []directdebit.MandateAuthorizationStatus{
		directdebit.MandateAuthorizationStatusActive,
		directdebit.MandateAuthorizationStatusPending,
		directdebit.MandateAuthorizationStatusRevoked,
	}

	statusCounts := make(map[directdebit.MandateAuthorizationStatus]int)

	for _, status := range statuses {
		builder := directdebit.NewListMandateAuthorizationsBuilder().
			Status(status).
			PerPage(100)

		mandates, err := client.DirectDebit.ListMandateAuthorizations(ctx, builder)
		if err != nil {
			log.Printf("Warning: Failed to list %s mandates: %v", status, err)
			continue
		}

		count := len(mandates.Data)
		statusCounts[status] = count
		fmt.Printf("- %s mandates: %d\n", status, count)

		// Show details for first few of each status
		if count > 0 {
			limit := min(count, 3)
			fmt.Printf("  Sample %s mandates:\n", status)
			for i := 0; i < limit; i++ {
				mandate := mandates.Data[i]
				fmt.Printf("    • %s (%s) - %s - %s\n",
					mandate.Customer.Email,
					mandate.BankName,
					mandate.AccountNumber,
					mandate.CreatedAt)
			}
		}
	}

	fmt.Println()
	return nil
}

// batchActivationWithErrorHandling demonstrates bulk operations with proper error handling
func batchActivationWithErrorHandling(ctx context.Context, client *paystack.Client) error {
	fmt.Println("2. Batch Activation with Error Handling")
	fmt.Println("---------------------------------------")

	// First, get pending mandates
	pendingBuilder := directdebit.NewListMandateAuthorizationsBuilder().
		Status(directdebit.MandateAuthorizationStatusPending).
		PerPage(20)

	pendingMandates, err := client.DirectDebit.ListMandateAuthorizations(ctx, pendingBuilder)
	if err != nil {
		return fmt.Errorf("failed to list pending mandates: %w", err)
	}

	if len(pendingMandates.Data) == 0 {
		fmt.Println("No pending mandates found for activation")
		fmt.Println()
		return nil
	}

	fmt.Printf("Found %d pending mandates\n", len(pendingMandates.Data))

	// Process in batches of 5
	batchSize := 5
	customerIDs := make([]uint64, 0)

	for i, mandate := range pendingMandates.Data {
		customerIDs = append(customerIDs, mandate.Customer.ID)

		// Process batch when we reach batch size or end of list
		if len(customerIDs) == batchSize || i == len(pendingMandates.Data)-1 {
			fmt.Printf("Processing batch of %d customers...\n", len(customerIDs))

			// Build activation request
			activationBuilder := directdebit.NewTriggerActivationChargeBuilder().
				CustomerIDs(customerIDs)

			// Execute with timeout context
			batchCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			activationResp, err := client.DirectDebit.TriggerActivationCharge(batchCtx, activationBuilder)
			cancel()

			if err != nil {
				log.Printf("Batch activation failed for %d customers: %v", len(customerIDs), err)
				// Continue with next batch instead of failing completely
			} else {
				fmt.Printf("✓ Activation triggered successfully: %s\n", activationResp.Message)
			}

			// Reset for next batch
			customerIDs = make([]uint64, 0)

			// Small delay between batches to avoid rate limiting
			time.Sleep(1 * time.Second)
		}
	}

	fmt.Println()
	return nil
}

// statusBasedMandateManagement demonstrates advanced mandate management workflows
func statusBasedMandateManagement(ctx context.Context, client *paystack.Client) error {
	fmt.Println("3. Status-Based Mandate Management")
	fmt.Println("----------------------------------")

	// Get mandates by status and analyze patterns
	statuses := []directdebit.MandateAuthorizationStatus{
		directdebit.MandateAuthorizationStatusActive,
		directdebit.MandateAuthorizationStatusPending,
		directdebit.MandateAuthorizationStatusRevoked,
	}

	mandatesByStatus := make(map[directdebit.MandateAuthorizationStatus][]directdebit.MandateAuthorization)

	for _, status := range statuses {
		builder := directdebit.NewListMandateAuthorizationsBuilder().
			Status(status).
			PerPage(50)

		mandates, err := client.DirectDebit.ListMandateAuthorizations(ctx, builder)
		if err != nil {
			log.Printf("Warning: Failed to get %s mandates: %v", status, err)
			continue
		}

		mandatesByStatus[status] = mandates.Data
	}

	// Analyze and report on mandate patterns
	fmt.Println("Mandate Status Analysis:")

	for status, mandates := range mandatesByStatus {
		fmt.Printf("\n%s Mandates (%d total):\n", status, len(mandates))

		if len(mandates) == 0 {
			fmt.Println("  No mandates in this status")
			continue
		}

		// Group by bank
		bankCounts := make(map[string]int)
		for _, mandate := range mandates {
			bankCounts[mandate.BankName]++
		}

		fmt.Println("  Top banks:")
		count := 0
		for bank, bankCount := range bankCounts {
			if count >= 5 { // Show top 5 banks
				break
			}
			fmt.Printf("    • %s: %d mandates\n", bank, bankCount)
			count++
		}

		// Show recommended actions based on status
		switch status {
		case directdebit.MandateAuthorizationStatusPending:
			if len(mandates) > 0 {
				fmt.Println("  💡 Recommended Action: Consider triggering activation charges")
				fmt.Printf("     Sample customer emails: ")
				for i := 0; i < min(len(mandates), 3); i++ {
					if i > 0 {
						fmt.Print(", ")
					}
					fmt.Print(mandates[i].Customer.Email)
				}
				fmt.Println()
			}
		case directdebit.MandateAuthorizationStatusActive:
			fmt.Printf("  ✅ Status: Healthy - %d active mandates ready for direct debit\n", len(mandates))
		case directdebit.MandateAuthorizationStatusRevoked:
			if len(mandates) > 0 {
				fmt.Printf("  ⚠️  Status: %d revoked mandates may need customer re-engagement\n", len(mandates))
			}
		}
	}

	fmt.Println()
	return nil
}

// Helper function for minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
