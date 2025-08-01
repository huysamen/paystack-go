// go:build example
//go:build example
// +build example

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/disputes"
)

func main() {
	// Create client with custom configuration
	config := paystack.NewConfig("sk_test_your_secret_key_here").
		WithTimeout(30 * time.Second)

	client := paystack.NewClient(config)
	ctx := context.Background()

	fmt.Println("=== Advanced Disputes API Demo ===\n")

	// Advanced dispute management scenarios
	examples := []struct {
		name        string
		description string
		runFn       func() error
	}{
		{
			name:        "Dispute Analytics",
			description: "Analyze dispute patterns and statistics",
			runFn:       func() error { return demonstrateDisputeAnalytics(client, ctx) },
		},
		{
			name:        "Evidence Management",
			description: "Comprehensive evidence handling workflow",
			runFn:       func() error { return demonstrateEvidenceManagement(client, ctx) },
		},
		{
			name:        "Bulk Dispute Operations",
			description: "Handle multiple disputes efficiently",
			runFn:       func() error { return demonstrateBulkOperations(client, ctx) },
		},
		{
			name:        "Resolution Strategies",
			description: "Different approaches to dispute resolution",
			runFn:       func() error { return demonstrateResolutionStrategies(client, ctx) },
		},
		{
			name:        "Error Handling",
			description: "Robust error handling patterns",
			runFn:       func() error { return demonstrateErrorHandling(client, ctx) },
		},
	}

	// Run each example
	for i, example := range examples {
		fmt.Printf("%d. %s\n", i+1, example.name)
		fmt.Printf("   %s\n", example.description)

		if err := example.runFn(); err != nil {
			fmt.Printf("   ❌ Error: %v\n", err)
		} else {
			fmt.Printf("   ✅ Completed successfully\n")
		}
		fmt.Println()
		time.Sleep(1 * time.Second) // Brief pause between examples
	}

	fmt.Println("=== Advanced Disputes Demo Complete ===")
}

// demonstrateDisputeAnalytics shows how to analyze dispute data
func demonstrateDisputeAnalytics(client *paystack.Client, ctx context.Context) error {
	// Get dispute statistics by status
	statuses := []disputes.DisputeStatus{
		disputes.DisputeStatusPending,
		disputes.DisputeStatusResolved,
		disputes.DisputeStatusAwaitingMerchantFeedback,
		disputes.DisputeStatusAwaitingBankFeedback,
	}

	stats := make(map[disputes.DisputeStatus]int)
	totalAmount := 0

	for _, status := range statuses {
		req := &disputes.DisputeListRequest{
			Status:  &status,
			PerPage: intPtr(100),
		}

		result, err := client.Disputes.List(ctx, req)
		if err != nil {
			return fmt.Errorf("failed to get %s disputes: %w", status, err)
		}

		stats[status] = len(result.Data)

		// Calculate total disputed amount for this status
		for _, dispute := range result.Data {
			if dispute.RefundAmount != nil {
				totalAmount += *dispute.RefundAmount
			}
		}
	}

	fmt.Printf("   📊 Dispute Statistics:\n")
	for status, count := range stats {
		fmt.Printf("     - %s: %d disputes\n", status, count)
	}
	fmt.Printf("     - Total disputed amount: ₦%.2f\n", float64(totalAmount)/100)

	return nil
}

// demonstrateEvidenceManagement shows comprehensive evidence handling
func demonstrateEvidenceManagement(client *paystack.Client, ctx context.Context) error {
	// Get first available dispute
	disputeList, err := client.Disputes.List(ctx, &disputes.DisputeListRequest{
		PerPage: intPtr(1),
	})
	if err != nil {
		return fmt.Errorf("failed to get disputes: %w", err)
	}

	if len(disputeList.Data) == 0 {
		fmt.Printf("   📝 No disputes available for evidence demonstration\n")
		return nil
	}

	dispute := disputeList.Data[0]
	disputeID := fmt.Sprintf("%d", dispute.ID)

	// Evidence types for different dispute categories
	evidenceTemplates := map[string]*disputes.DisputeEvidenceRequest{
		"fraud": {
			CustomerEmail:   "verified-customer@example.com",
			CustomerName:    "John Doe",
			CustomerPhone:   "+2348123456789",
			ServiceDetails:  "Transaction verified with customer directly via phone call. Customer confirmed purchase and provided transaction details.",
			DeliveryAddress: stringPtr("123 Verified Street, Lagos, Nigeria"),
		},
		"general": {
			CustomerEmail:  "satisfied-customer@example.com",
			CustomerName:   "Jane Smith",
			CustomerPhone:  "+2349087654321",
			ServiceDetails: "Service delivered as agreed. Customer satisfied with outcome. Payment processed correctly.",
		},
		"authorization": {
			CustomerEmail:  "authorized-customer@example.com",
			CustomerName:   "Bob Johnson",
			CustomerPhone:  "+2347012345678",
			ServiceDetails: "Customer authorization confirmed. Transaction processed with valid payment method and customer consent.",
		},
	}

	// Use appropriate evidence template
	evidenceReq := evidenceTemplates[string(dispute.Category)]
	if evidenceReq == nil {
		evidenceReq = evidenceTemplates["general"] // fallback
	}

	fmt.Printf("   📋 Adding evidence for %s dispute (Category: %s)\n", dispute.Status, dispute.Category)

	evidence, err := client.Disputes.AddEvidence(ctx, disputeID, evidenceReq)
	if err != nil {
		return fmt.Errorf("failed to add evidence: %w", err)
	}

	fmt.Printf("   ✅ Evidence added: ID %d\n", evidence.Data.ID)

	// Get upload URL for supporting documents
	uploadReq := &disputes.DisputeUploadURLRequest{
		UploadFileName: "supporting-document.pdf",
	}

	uploadURL, err := client.Disputes.GetUploadURL(ctx, disputeID, uploadReq)
	if err != nil {
		return fmt.Errorf("failed to get upload URL: %w", err)
	}

	fmt.Printf("   📎 Upload URL generated (expires in %d seconds)\n", uploadURL.Data.ExpiresIn)

	return nil
}

// demonstrateBulkOperations shows efficient handling of multiple disputes
func demonstrateBulkOperations(client *paystack.Client, ctx context.Context) error {
	// Get multiple disputes for bulk operations
	disputesList, err := client.Disputes.List(ctx, &disputes.DisputeListRequest{
		PerPage: intPtr(5),
	})
	if err != nil {
		return fmt.Errorf("failed to get disputes: %w", err)
	}

	if len(disputesList.Data) == 0 {
		fmt.Printf("   📦 No disputes available for bulk operations\n")
		return nil
	}

	fmt.Printf("   📦 Processing %d disputes in bulk...\n", len(disputesList.Data))

	// Bulk evidence addition
	successCount := 0
	errorCount := 0

	for i, dispute := range disputesList.Data {
		disputeID := fmt.Sprintf("%d", dispute.ID)

		evidenceReq := &disputes.DisputeEvidenceRequest{
			CustomerEmail:  fmt.Sprintf("bulk-customer-%d@example.com", i+1),
			CustomerName:   fmt.Sprintf("Customer %d", i+1),
			CustomerPhone:  fmt.Sprintf("+234801234567%d", i),
			ServiceDetails: fmt.Sprintf("Bulk processing - Service provided for dispute %d", dispute.ID),
		}

		_, err := client.Disputes.AddEvidence(ctx, disputeID, evidenceReq)
		if err != nil {
			errorCount++
			fmt.Printf("     ❌ Failed to add evidence for dispute %d: %v\n", dispute.ID, err)
		} else {
			successCount++
			fmt.Printf("     ✅ Evidence added for dispute %d\n", dispute.ID)
		}

		// Rate limiting - avoid overwhelming the API
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Printf("   📊 Bulk operation results: %d success, %d errors\n", successCount, errorCount)

	return nil
}

// demonstrateResolutionStrategies shows different resolution approaches
func demonstrateResolutionStrategies(client *paystack.Client, ctx context.Context) error {
	// Get disputes by different statuses
	pendingDisputeList, err := client.Disputes.List(ctx, &disputes.DisputeListRequest{
		Status:  &[]disputes.DisputeStatus{disputes.DisputeStatusPending}[0],
		PerPage: intPtr(2),
	})
	if err != nil {
		return fmt.Errorf("failed to get pending disputes: %w", err)
	}

	strategies := []struct {
		name       string
		resolution disputes.DisputeResolution
		message    string
		condition  func(dispute disputes.Dispute) bool
	}{
		{
			name:       "Accept with Refund",
			resolution: disputes.DisputeResolutionMerchantAccepted,
			message:    "Customer contacted and issue resolved with partial refund",
			condition: func(d disputes.Dispute) bool {
				return d.RefundAmount != nil && *d.RefundAmount < 10000 // Less than ₦100
			},
		},
		{
			name:       "Decline with Evidence",
			resolution: disputes.DisputeResolutionDeclined,
			message:    "Comprehensive evidence provided showing valid transaction",
			condition: func(d disputes.Dispute) bool {
				return string(d.Category) == "fraud"
			},
		},
	}

	fmt.Printf("   🎯 Demonstrating resolution strategies:\n")

	for _, strategy := range strategies {
		fmt.Printf("     Strategy: %s\n", strategy.name)

		// Find suitable dispute for this strategy
		var targetDispute *disputes.Dispute
		for i := range pendingDisputeList.Data {
			if strategy.condition(pendingDisputeList.Data[i]) {
				targetDispute = &pendingDisputeList.Data[i]
				break
			}
		}

		if targetDispute == nil {
			fmt.Printf("       No suitable dispute found for this strategy\n")
			continue
		}

		// Note: In production, be very careful with resolution operations
		fmt.Printf("       Would resolve dispute %d using strategy: %s\n", targetDispute.ID, strategy.name)
		fmt.Printf("       Resolution: %s\n", strategy.resolution)
		fmt.Printf("       Message: %s\n", strategy.message)

		// Uncomment to actually resolve (use with caution!)
		/*
			resolveReq := &disputes.DisputeResolveRequest{
				Resolution:       strategy.resolution,
				Message:          strategy.message,
				RefundAmount:     5000, // ₦50.00
				UploadedFileName: "evidence.pdf",
			}

			result, err := client.Disputes.Resolve(ctx, disputeID, resolveReq)
			if err != nil {
				fmt.Printf("       ❌ Resolution failed: %v\n", err)
			} else {
				fmt.Printf("       ✅ Dispute resolved: %s\n", result.Data.Status)
			}
		*/
	}

	return nil
}

// demonstrateErrorHandling shows robust error handling patterns
func demonstrateErrorHandling(client *paystack.Client, ctx context.Context) error {
	fmt.Printf("   🛡️ Testing error handling scenarios:\n")

	// Test various error conditions
	errorTests := []struct {
		name string
		test func() error
	}{
		{
			name: "Invalid Dispute ID",
			test: func() error {
				_, err := client.Disputes.Fetch(ctx, "invalid-id")
				return err
			},
		},
		{
			name: "Empty Evidence Request",
			test: func() error {
				_, err := client.Disputes.AddEvidence(ctx, "123", &disputes.DisputeEvidenceRequest{})
				return err
			},
		},
		{
			name: "Invalid Email in Evidence",
			test: func() error {
				req := &disputes.DisputeEvidenceRequest{
					CustomerEmail:  "invalid-email",
					CustomerName:   "Test Customer",
					CustomerPhone:  "+2348123456789",
					ServiceDetails: "Test service",
				}
				_, err := client.Disputes.AddEvidence(ctx, "123", req)
				return err
			},
		},
		{
			name: "Invalid Date Range",
			test: func() error {
				future := time.Now().AddDate(0, 0, 1)
				past := time.Now().AddDate(0, 0, -1)
				req := &disputes.DisputeListRequest{
					From: &future, // Future date as "from"
					To:   &past,   // Past date as "to"
				}
				_, err := client.Disputes.List(ctx, req)
				return err
			},
		},
	}

	for _, test := range errorTests {
		err := test.test()
		if err != nil {
			fmt.Printf("     ✅ %s: Correctly caught error - %v\n", test.name, err)
		} else {
			fmt.Printf("     ❌ %s: Should have failed but didn't\n", test.name)
		}
	}

	return nil
}

// Helper functions
func intPtr(i int) *int {
	return &i
}

func stringPtr(s string) *string {
	return &s
}
