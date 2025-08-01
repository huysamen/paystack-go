// go:build example
//go:build example
// +build example

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/disputes"
)

func main() {
	// Create client
	client := paystack.DefaultClient("sk_test_your_secret_key_here")
	ctx := context.Background()

	fmt.Println("=== Disputes API Demo ===\n")

	// 1. List all disputes
	fmt.Println("1. Listing all disputes...")
	listReq := &disputes.DisputeListRequest{
		PerPage: intPtr(10),
		Page:    intPtr(1),
	}

	disputesList, err := client.Disputes.List(ctx, listReq)
	if err != nil {
		log.Printf("Error listing disputes: %v", err)
	} else {
		fmt.Printf("Found %d disputes\n", len(disputesList.Data))
		for i, dispute := range disputesList.Data {
			fmt.Printf("  %d. ID: %d, Status: %s, Amount: ", i+1, dispute.ID, dispute.Status)
			if dispute.RefundAmount != nil {
				fmt.Printf("₦%.2f", float64(*dispute.RefundAmount)/100)
			} else {
				fmt.Print("N/A")
			}
			fmt.Printf(" (%s)\n", dispute.CreatedAt.Time.Format("2006-01-02"))
		}
	}
	fmt.Println()

	// 2. List disputes by status
	fmt.Println("2. Listing pending disputes...")
	pendingReq := &disputes.DisputeListRequest{
		Status:  &[]disputes.DisputeStatus{disputes.DisputeStatusPending}[0],
		PerPage: intPtr(5),
	}

	pendingDisputes, err := client.Disputes.List(ctx, pendingReq)
	if err != nil {
		log.Printf("Error listing pending disputes: %v", err)
	} else {
		fmt.Printf("Found %d pending disputes\n", len(pendingDisputes.Data))
	}
	fmt.Println()

	// 3. List disputes with date filter
	fmt.Println("3. Listing recent disputes (last 30 days)...")
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	now := time.Now()

	recentReq := &disputes.DisputeListRequest{
		From:    &thirtyDaysAgo,
		To:      &now,
		PerPage: intPtr(5),
	}

	recentDisputes, err := client.Disputes.List(ctx, recentReq)
	if err != nil {
		log.Printf("Error listing recent disputes: %v", err)
	} else {
		fmt.Printf("Found %d disputes in the last 30 days\n", len(recentDisputes.Data))
	}
	fmt.Println()

	// If we have disputes, demonstrate other operations
	if len(disputesList.Data) > 0 {
		dispute := disputesList.Data[0]
		disputeID := fmt.Sprintf("%d", dispute.ID)

		// 4. Fetch specific dispute
		fmt.Printf("4. Fetching dispute details for ID: %s...\n", disputeID)
		fetchedDispute, err := client.Disputes.Fetch(ctx, disputeID)
		if err != nil {
			log.Printf("Error fetching dispute: %v", err)
		} else {
			fmt.Printf("Dispute Status: %s\n", fetchedDispute.Data.Status)
			fmt.Printf("Category: %s\n", fetchedDispute.Data.Category)
			if fetchedDispute.Data.Transaction != nil {
				fmt.Printf("Transaction Reference: %s\n", fetchedDispute.Data.Transaction.Reference)
			}
		}
		fmt.Println()

		// 5. List transaction disputes (if we have a transaction)
		if dispute.Transaction != nil {
			transactionID := fmt.Sprintf("%d", dispute.Transaction.ID)
			fmt.Printf("5. Listing disputes for transaction ID: %s...\n", transactionID)

			txDisputes, err := client.Disputes.ListTransactionDisputes(ctx, transactionID)
			if err != nil {
				log.Printf("Error listing transaction disputes: %v", err)
			} else {
				fmt.Printf("Found %d history records\n", len(txDisputes.Data.History))
				fmt.Printf("Found %d messages\n", len(txDisputes.Data.Messages))
			}
			fmt.Println()
		}

		// 6. Add evidence (example with validation)
		fmt.Printf("6. Adding evidence to dispute ID: %s...\n", disputeID)
		evidenceReq := &disputes.DisputeEvidenceRequest{
			CustomerEmail:   "customer@example.com",
			CustomerName:    "John Doe",
			CustomerPhone:   "+2348123456789",
			ServiceDetails:  "Product delivered successfully on time with receipt",
			DeliveryAddress: stringPtr("123 Main Street, Lagos, Nigeria"),
		}

		evidence, err := client.Disputes.AddEvidence(ctx, disputeID, evidenceReq)
		if err != nil {
			log.Printf("Error adding evidence: %v", err)
		} else {
			fmt.Printf("Evidence added successfully: ID %d\n", evidence.Data.ID)
		}
		fmt.Println()

		// 7. Get upload URL for file evidence
		fmt.Println("7. Getting upload URL for evidence file...")
		uploadReq := &disputes.DisputeUploadURLRequest{
			UploadFileName: "receipt.pdf",
		}

		uploadURL, err := client.Disputes.GetUploadURL(ctx, disputeID, uploadReq)
		if err != nil {
			log.Printf("Error getting upload URL: %v", err)
		} else {
			fmt.Printf("Upload URL generated: %s\n", uploadURL.Data.SignedURL[:50]+"...")
			fmt.Printf("Expires in: %d seconds\n", uploadURL.Data.ExpiresIn)
		}
		fmt.Println()

		// 8. Update dispute (example)
		fmt.Printf("8. Updating dispute ID: %s...\n", disputeID)
		updateReq := &disputes.DisputeUpdateRequest{
			RefundAmount: intPtr(5000), // ₦50.00 refund
		}

		updatedDispute, err := client.Disputes.Update(ctx, disputeID, updateReq)
		if err != nil {
			log.Printf("Error updating dispute: %v", err)
		} else {
			fmt.Printf("Dispute updated successfully: %d records\n", len(updatedDispute.Data))
		}
		fmt.Println()

		// 9. Resolve dispute (example - be careful with this in production)
		fmt.Printf("9. Resolving dispute ID: %s...\n", disputeID)
		resolveReq := &disputes.DisputeResolveRequest{
			Resolution:       disputes.DisputeResolutionMerchantAccepted,
			Message:          "Customer contacted directly and issue resolved amicably",
			RefundAmount:     5000,          // ₦50.00
			UploadedFileName: "receipt.pdf", // from upload URL step
		}

		resolvedDispute, err := client.Disputes.Resolve(ctx, disputeID, resolveReq)
		if err != nil {
			log.Printf("Error resolving dispute: %v", err)
		} else {
			fmt.Printf("Dispute resolved: %s\n", resolvedDispute.Data.Status)
			if resolvedDispute.Data.Resolution != nil {
				fmt.Printf("Resolution: %s\n", *resolvedDispute.Data.Resolution)
			}
		}
		fmt.Println()
	}

	// 10. Export disputes
	fmt.Println("10. Exporting disputes...")
	exportReq := &disputes.DisputeExportRequest{
		From:    &thirtyDaysAgo,
		To:      &now,
		PerPage: intPtr(100),
	}

	exportResult, err := client.Disputes.Export(ctx, exportReq)
	if err != nil {
		log.Printf("Error exporting disputes: %v", err)
	} else {
		fmt.Printf("Export generated: %s\n", exportResult.Data.Path)
		if exportResult.Data.ExpiresAt != nil {
			fmt.Printf("Export expires at: %s\n", exportResult.Data.ExpiresAt.Time.Format(time.RFC3339))
		}
	}

	fmt.Println("\n=== Disputes Demo Complete ===")
}

// Helper functions
func intPtr(i int) *int {
	return &i
}

func stringPtr(s string) *string {
	return &s
}
