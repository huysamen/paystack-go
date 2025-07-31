package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/huysamen/paystack-go"
	payment_requests "github.com/huysamen/paystack-go/api/payment-requests"
)

func main() {
	// Initialize client
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable is required")
	}

	client := paystack.DefaultClient(secretKey)
	ctx := context.Background()

	// Example 1: Create a simple payment request
	fmt.Println("=== Creating Simple Payment Request ===")
	amount := 50000 // ₦500.00 in kobo
	createReq := &payment_requests.CreatePaymentRequestRequest{
		Customer:    "CUS_customer_code_here", // Replace with actual customer code
		Amount:      &amount,
		Description: "Payment for premium service subscription",
		DueDate:     time.Now().AddDate(0, 0, 7).Format("2006-01-02"), // Due in 7 days
		Currency:    "NGN",
	}

	paymentRequest, err := client.PaymentRequests.Create(ctx, createReq)
	if err != nil {
		log.Printf("Failed to create simple payment request (expected if customer doesn't exist): %v", err)

		// Create with line items instead (doesn't require existing customer for demo)
		fmt.Println("\n=== Creating Payment Request with Line Items ===")
		createReq2 := &payment_requests.CreatePaymentRequestRequest{
			Customer:    "demo@example.com", // Can be email for demo
			Description: "Invoice for professional services",
			DueDate:     time.Now().AddDate(0, 0, 14).Format("2006-01-02"), // Due in 14 days
			Currency:    "NGN",
			LineItems: []payment_requests.LineItem{
				{
					Name:     "Website Development",
					Amount:   150000, // ₦1,500.00
					Quantity: 1,
				},
				{
					Name:     "SEO Optimization",
					Amount:   75000, // ₦750.00
					Quantity: 1,
				},
				{
					Name:     "Monthly Maintenance",
					Amount:   25000, // ₦250.00
					Quantity: 3,
				},
			},
			Tax: []payment_requests.Tax{
				{
					Name:   "VAT (7.5%)",
					Amount: 22500, // ₦225.00
				},
			},
		}

		paymentRequest, err = client.PaymentRequests.Create(ctx, createReq2)
		if err != nil {
			log.Fatalf("Failed to create payment request with line items: %v", err)
		}
	}

	fmt.Printf("Payment Request Created:\n")
	fmt.Printf("  ID: %d\n", paymentRequest.ID)
	fmt.Printf("  Request Code: %s\n", paymentRequest.RequestCode)
	fmt.Printf("  Amount: ₦%.2f\n", float64(paymentRequest.Amount)/100)
	fmt.Printf("  Currency: %s\n", paymentRequest.Currency)
	fmt.Printf("  Status: %s\n", paymentRequest.Status)
	fmt.Printf("  Due Date: %s\n", paymentRequest.DueDate)

	requestCode := paymentRequest.RequestCode

	// Example 2: Fetch the created payment request
	fmt.Println("\n=== Fetching Payment Request ===")
	fetchedRequest, err := client.PaymentRequests.Fetch(ctx, requestCode)
	if err != nil {
		log.Fatalf("Failed to fetch payment request: %v", err)
	}

	fmt.Printf("Fetched Request:\n")
	fmt.Printf("  Description: %s\n", fetchedRequest.Description)
	fmt.Printf("  Line Items: %d\n", len(fetchedRequest.LineItems))
	for i, item := range fetchedRequest.LineItems {
		fmt.Printf("    %d. %s: ₦%.2f x %d\n", i+1, item.Name, float64(item.Amount)/100, item.Quantity)
	}
	fmt.Printf("  Tax Items: %d\n", len(fetchedRequest.Tax))
	for i, tax := range fetchedRequest.Tax {
		fmt.Printf("    %d. %s: ₦%.2f\n", i+1, tax.Name, float64(tax.Amount)/100)
	}

	// Example 3: Verify payment request
	fmt.Println("\n=== Verifying Payment Request ===")
	verifiedRequest, err := client.PaymentRequests.Verify(ctx, requestCode)
	if err != nil {
		log.Fatalf("Failed to verify payment request: %v", err)
	}

	fmt.Printf("Verified Request:\n")
	fmt.Printf("  Status: %s\n", verifiedRequest.Status)
	fmt.Printf("  Paid: %t\n", verifiedRequest.Paid)
	fmt.Printf("  Has Invoice: %t\n", verifiedRequest.HasInvoice)
	if verifiedRequest.InvoiceNumber != nil {
		fmt.Printf("  Invoice Number: %d\n", *verifiedRequest.InvoiceNumber)
	}

	// Example 4: Update payment request
	fmt.Println("\n=== Updating Payment Request ===")
	updateReq := &payment_requests.UpdatePaymentRequestRequest{
		Description: "Updated: Invoice for professional services (Rush Order)",
		DueDate:     time.Now().AddDate(0, 0, 3).Format("2006-01-02"), // Due in 3 days (urgent)
		LineItems: []payment_requests.LineItem{
			{
				Name:     "Website Development (Express)",
				Amount:   180000, // ₦1,800.00 (rush fee)
				Quantity: 1,
			},
			{
				Name:     "SEO Optimization",
				Amount:   75000, // ₦750.00
				Quantity: 1,
			},
			{
				Name:     "Monthly Maintenance",
				Amount:   25000, // ₦250.00
				Quantity: 3,
			},
		},
		Tax: []payment_requests.Tax{
			{
				Name:   "VAT (7.5%)",
				Amount: 24750, // ₦247.50 (updated based on new total)
			},
		},
	}

	updatedRequest, err := client.PaymentRequests.Update(ctx, requestCode, updateReq)
	if err != nil {
		log.Fatalf("Failed to update payment request: %v", err)
	}

	fmt.Printf("Updated Request:\n")
	fmt.Printf("  Description: %s\n", updatedRequest.Description)
	fmt.Printf("  New Amount: ₦%.2f\n", float64(updatedRequest.Amount)/100)
	fmt.Printf("  New Due Date: %s\n", updatedRequest.DueDate)

	// Example 5: List payment requests
	fmt.Println("\n=== Listing Payment Requests ===")
	listReq := &payment_requests.ListPaymentRequestsRequest{
		PerPage: 5,
		Page:    1,
		Status:  "pending", // Only show pending requests
	}

	requestsResp, err := client.PaymentRequests.List(ctx, listReq)
	if err != nil {
		log.Fatalf("Failed to list payment requests: %v", err)
	}

	fmt.Printf("Found %d payment requests:\n", len(requestsResp.Data))
	for i, req := range requestsResp.Data {
		if i >= 3 { // Show only first 3 for brevity
			fmt.Printf("  ... and %d more\n", len(requestsResp.Data)-3)
			break
		}
		fmt.Printf("  %d. %s - ₦%.2f (%s)\n",
			i+1, req.Description, float64(req.Amount)/100, req.Status)
	}

	// Example 6: Get payment request totals
	fmt.Println("\n=== Payment Request Totals ===")
	totals, err := client.PaymentRequests.GetTotals(ctx)
	if err != nil {
		log.Fatalf("Failed to get payment request totals: %v", err)
	}

	fmt.Printf("Payment Request Analytics:\n")
	fmt.Printf("  Pending Requests:\n")
	for _, pending := range totals.Pending {
		fmt.Printf("    %s: ₦%.2f\n", pending.Currency, float64(pending.Amount)/100)
	}
	fmt.Printf("  Successful Requests:\n")
	for _, successful := range totals.Successful {
		fmt.Printf("    %s: ₦%.2f\n", successful.Currency, float64(successful.Amount)/100)
	}
	fmt.Printf("  Total Requests:\n")
	for _, total := range totals.Total {
		fmt.Printf("    %s: ₦%.2f\n", total.Currency, float64(total.Amount)/100)
	}

	// Example 7: Send notification (optional)
	fmt.Println("\n=== Sending Notification ===")
	notifyResp, err := client.PaymentRequests.SendNotification(ctx, requestCode)
	if err != nil {
		log.Printf("Failed to send notification (expected for test): %v", err)
	} else {
		fmt.Printf("Notification Status: %s\n", notifyResp.Message)
	}

	// Example 8: Finalize payment request (if it's a draft)
	fmt.Println("\n=== Finalizing Payment Request (Demo) ===")
	sendNotification := true
	finalizeReq := &payment_requests.FinalizePaymentRequestRequest{
		SendNotification: &sendNotification,
	}

	finalizedRequest, err := client.PaymentRequests.Finalize(ctx, requestCode, finalizeReq)
	if err != nil {
		log.Printf("Failed to finalize (expected if not a draft): %v", err)
	} else {
		fmt.Printf("Finalized Request: %s\n", finalizedRequest.Status)
	}

	// Example 9: Archive payment request (careful - this hides it from lists)
	fmt.Println("\n=== Archive Payment Request (Demo) ===")
	fmt.Printf("Note: Archiving will hide the payment request from future list operations\n")
	fmt.Printf("Skipping archive operation to keep example request visible\n")
	// archiveResp, err := client.PaymentRequests.Archive(ctx, requestCode)
	// if err != nil {
	//     log.Printf("Failed to archive: %v", err)
	// } else {
	//     fmt.Printf("Archive Status: %s\n", archiveResp.Message)
	// }

	// Print summary
	fmt.Println("\n=== Summary ===")
	fmt.Printf("✅ Created payment request: %s\n", paymentRequest.RequestCode)
	fmt.Printf("✅ Amount: ₦%.2f (%s)\n", float64(updatedRequest.Amount)/100, updatedRequest.Currency)
	fmt.Printf("✅ Status: %s\n", updatedRequest.Status)
	fmt.Printf("✅ Due Date: %s\n", updatedRequest.DueDate)
	fmt.Printf("✅ Line Items: %d\n", len(updatedRequest.LineItems))
	fmt.Printf("✅ Tax Items: %d\n", len(updatedRequest.Tax))

	// Print the full request data as JSON for reference
	fmt.Println("\n=== Full Payment Request Data (JSON) ===")
	requestJSON, _ := json.MarshalIndent(updatedRequest, "", "  ")
	fmt.Println(string(requestJSON))
}
