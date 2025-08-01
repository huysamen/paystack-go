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

	fmt.Println("🧾 Payment Requests Advanced Example")
	fmt.Println("===================================")

	// Scenario 1: Freelancer Invoice System
	fmt.Println("\n💼 Scenario 1: Freelancer Invoice System")
	fmt.Println("Creating detailed invoice for web development project...")

	hasInvoice := true
	sendNotification := true
	freelancerInvoiceReq := &payment_requests.CreatePaymentRequestRequest{
		Customer:         "client@techstartup.com",
		Description:      "Web Development Project - E-commerce Platform",
		DueDate:          time.Now().AddDate(0, 0, 30).Format("2006-01-02"), // 30 days payment terms
		Currency:         "NGN",
		HasInvoice:       &hasInvoice,
		SendNotification: &sendNotification,
		LineItems: []payment_requests.LineItem{
			{
				Name:     "Frontend Development (React.js)",
				Amount:   500000, // ₦5,000.00
				Quantity: 1,
			},
			{
				Name:     "Backend API Development (Node.js)",
				Amount:   400000, // ₦4,000.00
				Quantity: 1,
			},
			{
				Name:     "Database Design & Setup",
				Amount:   200000, // ₦2,000.00
				Quantity: 1,
			},
			{
				Name:     "Payment Integration (Paystack)",
				Amount:   150000, // ₦1,500.00
				Quantity: 1,
			},
			{
				Name:     "Testing & Quality Assurance",
				Amount:   100000, // ₦1,000.00
				Quantity: 1,
			},
		},
		Tax: []payment_requests.Tax{
			{
				Name:   "VAT (7.5%)",
				Amount: 103750, // 7.5% of subtotal
			},
		},
	}

	freelancerInvoice, err := client.PaymentRequests.Create(ctx, freelancerInvoiceReq)
	if err != nil {
		log.Fatalf("Failed to create freelancer invoice: %v", err)
	}

	fmt.Printf("✅ Freelancer invoice created:\n")
	fmt.Printf("   Request Code: %s\n", freelancerInvoice.RequestCode)
	fmt.Printf("   Total Amount: ₦%.2f\n", float64(freelancerInvoice.Amount)/100)
	fmt.Printf("   Invoice Number: %d\n", *freelancerInvoice.InvoiceNumber)

	// Scenario 2: SaaS Subscription Management
	fmt.Println("\n💻 Scenario 2: SaaS Subscription Management")
	fmt.Println("Creating subscription invoice with usage-based billing...")

	saasInvoiceReq := &payment_requests.CreatePaymentRequestRequest{
		Customer:    "admin@growthcompany.com",
		Description: "SaaS Platform Subscription - March 2024",
		DueDate:     time.Now().AddDate(0, 0, 7).Format("2006-01-02"), // 7 days payment terms
		Currency:    "USD",
		HasInvoice:  &hasInvoice,
		LineItems: []payment_requests.LineItem{
			{
				Name:     "Pro Plan Monthly Subscription",
				Amount:   9900, // $99.00
				Quantity: 1,
			},
			{
				Name:     "Additional Users (5 users)",
				Amount:   2500, // $25.00
				Quantity: 5,
			},
			{
				Name:     "API Calls Overage (per 1k calls)",
				Amount:   500, // $5.00
				Quantity: 15,  // 15k extra API calls
			},
			{
				Name:     "Premium Support",
				Amount:   4900, // $49.00
				Quantity: 1,
			},
		},
		Tax: []payment_requests.Tax{
			{
				Name:   "Digital Services Tax (6%)",
				Amount: 1074, // 6% of subtotal
			},
		},
	}

	saasInvoice, err := client.PaymentRequests.Create(ctx, saasInvoiceReq)
	if err != nil {
		log.Fatalf("Failed to create SaaS invoice: %v", err)
	}

	fmt.Printf("✅ SaaS subscription invoice created:\n")
	fmt.Printf("   Request Code: %s\n", saasInvoice.RequestCode)
	fmt.Printf("   Total Amount: $%.2f\n", float64(saasInvoice.Amount)/100)

	// Scenario 3: Event Management & Ticketing
	fmt.Println("\n🎫 Scenario 3: Event Management & Ticketing")
	fmt.Println("Creating group booking invoice for corporate event...")

	eventInvoiceReq := &payment_requests.CreatePaymentRequestRequest{
		Customer:    "events@corporatefirm.com",
		Description: "Corporate Training Workshop - Leadership Development",
		DueDate:     time.Now().AddDate(0, 0, 14).Format("2006-01-02"), // 14 days payment terms
		Currency:    "NGN",
		HasInvoice:  &hasInvoice,
		LineItems: []payment_requests.LineItem{
			{
				Name:     "Workshop Registration (Executive Level)",
				Amount:   75000, // ₦750.00 per person
				Quantity: 15,    // 15 executives
			},
			{
				Name:     "Workshop Materials & Workbooks",
				Amount:   5000, // ₦50.00 per person
				Quantity: 15,
			},
			{
				Name:     "Catering - Executive Lunch",
				Amount:   8000, // ₦80.00 per person
				Quantity: 15,
			},
			{
				Name:     "Venue Upgrade - Premium Conference Room",
				Amount:   100000, // ₦1,000.00 flat fee
				Quantity: 1,
			},
			{
				Name:     "Professional Photography",
				Amount:   50000, // ₦500.00 flat fee
				Quantity: 1,
			},
		},
		Tax: []payment_requests.Tax{
			{
				Name:   "VAT (7.5%)",
				Amount: 118125, // 7.5% of subtotal
			},
		},
	}

	eventInvoice, err := client.PaymentRequests.Create(ctx, eventInvoiceReq)
	if err != nil {
		log.Fatalf("Failed to create event invoice: %v", err)
	}

	fmt.Printf("✅ Event booking invoice created:\n")
	fmt.Printf("   Request Code: %s\n", eventInvoice.RequestCode)
	fmt.Printf("   Total Amount: ₦%.2f\n", float64(eventInvoice.Amount)/100)

	// Scenario 4: E-commerce Bulk Order
	fmt.Println("\n🛒 Scenario 4: E-commerce Bulk Order")
	fmt.Println("Creating wholesale order invoice with volume discounts...")

	wholesaleInvoiceReq := &payment_requests.CreatePaymentRequestRequest{
		Customer:    "purchasing@retailchain.com",
		Description: "Wholesale Order - Q1 2024 Inventory Restock",
		DueDate:     time.Now().AddDate(0, 0, 45).Format("2006-01-02"), // 45 days payment terms
		Currency:    "NGN",
		HasInvoice:  &hasInvoice,
		LineItems: []payment_requests.LineItem{
			{
				Name:     "Wireless Headphones (Bulk Rate)",
				Amount:   15000, // ₦150.00 each (discounted from ₦200)
				Quantity: 100,
			},
			{
				Name:     "Smartphone Cases (Assorted)",
				Amount:   2500, // ₦25.00 each
				Quantity: 200,
			},
			{
				Name:     "Charging Cables (USB-C)",
				Amount:   1200, // ₦12.00 each
				Quantity: 500,
			},
			{
				Name:     "Screen Protectors (Tempered Glass)",
				Amount:   800, // ₦8.00 each
				Quantity: 300,
			},
			{
				Name:     "Bulk Order Packaging",
				Amount:   25000, // ₦250.00 flat fee
				Quantity: 1,
			},
		},
		Tax: []payment_requests.Tax{
			{
				Name:   "VAT (7.5%)",
				Amount: 188625, // 7.5% of subtotal
			},
		},
	}

	wholesaleInvoice, err := client.PaymentRequests.Create(ctx, wholesaleInvoiceReq)
	if err != nil {
		log.Fatalf("Failed to create wholesale invoice: %v", err)
	}

	fmt.Printf("✅ Wholesale order invoice created:\n")
	fmt.Printf("   Request Code: %s\n", wholesaleInvoice.RequestCode)
	fmt.Printf("   Total Amount: ₦%.2f\n", float64(wholesaleInvoice.Amount)/100)

	// Scenario 5: Advanced Analytics & Reporting
	fmt.Println("\n📊 Scenario 5: Advanced Analytics & Reporting")
	fmt.Println("Analyzing payment request performance...")

	// Get comprehensive totals
	totals, err := client.PaymentRequests.GetTotals(ctx)
	if err != nil {
		log.Fatalf("Failed to get totals: %v", err)
	}

	fmt.Printf("📈 Payment Request Analytics:\n")

	// Calculate totals by currency
	pendingTotal := make(map[string]float64)
	successfulTotal := make(map[string]float64)
	grandTotal := make(map[string]float64)

	for _, pending := range totals.Pending {
		pendingTotal[pending.Currency] = float64(pending.Amount) / 100
	}
	for _, successful := range totals.Successful {
		successfulTotal[successful.Currency] = float64(successful.Amount) / 100
	}
	for _, total := range totals.Total {
		grandTotal[total.Currency] = float64(total.Amount) / 100
	}

	fmt.Printf("   💰 Financial Summary:\n")
	for currency := range grandTotal {
		pending := pendingTotal[currency]
		successful := successfulTotal[currency]
		total := grandTotal[currency]

		fmt.Printf("     %s:\n", currency)
		fmt.Printf("       - Pending: %.2f\n", pending)
		fmt.Printf("       - Successful: %.2f\n", successful)
		fmt.Printf("       - Total: %.2f\n", total)

		if total > 0 {
			successRate := (successful / total) * 100
			fmt.Printf("       - Success Rate: %.1f%%\n", successRate)
		}
	}

	// Scenario 6: Bulk Operations Management
	fmt.Println("\n⚙️ Scenario 6: Bulk Operations Management")
	fmt.Println("Demonstrating bulk payment request operations...")

	// List all recent payment requests
	listReq := &payment_requests.ListPaymentRequestsRequest{
		PerPage: 20,
		Page:    1,
	}

	allRequests, err := client.PaymentRequests.List(ctx, listReq)
	if err != nil {
		log.Fatalf("Failed to list payment requests: %v", err)
	}

	fmt.Printf("📋 Found %d payment requests:\n", len(allRequests.Data))

	// Categorize requests by status
	statusCount := make(map[string]int)
	currencyTotals := make(map[string]float64)

	for _, req := range allRequests.Data {
		statusCount[req.Status]++
		currencyTotals[req.Currency] += float64(req.Amount) / 100
	}

	fmt.Printf("   📊 Status Distribution:\n")
	for status, count := range statusCount {
		fmt.Printf("     - %s: %d requests\n", status, count)
	}

	fmt.Printf("   💱 Currency Breakdown:\n")
	for currency, total := range currencyTotals {
		fmt.Printf("     - %s: %.2f\n", currency, total)
	}

	// Scenario 7: Payment Request Lifecycle Management
	fmt.Println("\n🔄 Scenario 7: Payment Request Lifecycle Management")
	fmt.Println("Demonstrating complete payment request lifecycle...")

	// Update the freelancer invoice with project changes
	fmt.Printf("Updating freelancer invoice with scope changes...\n")
	updateReq := &payment_requests.UpdatePaymentRequestRequest{
		Description: "Web Development Project - E-commerce Platform (Revised Scope)",
		LineItems: []payment_requests.LineItem{
			{
				Name:     "Frontend Development (React.js + PWA)",
				Amount:   600000, // ₦6,000.00 (increased scope)
				Quantity: 1,
			},
			{
				Name:     "Backend API Development (Node.js + GraphQL)",
				Amount:   500000, // ₦5,000.00 (added GraphQL)
				Quantity: 1,
			},
			{
				Name:     "Database Design & Setup",
				Amount:   200000, // ₦2,000.00
				Quantity: 1,
			},
			{
				Name:     "Payment Integration (Paystack + Flutterwave)",
				Amount:   200000, // ₦2,000.00 (added Flutterwave)
				Quantity: 1,
			},
			{
				Name:     "Testing & Quality Assurance",
				Amount:   150000, // ₦1,500.00 (increased scope)
				Quantity: 1,
			},
			{
				Name:     "Mobile App (React Native)",
				Amount:   400000, // ₦4,000.00 (new requirement)
				Quantity: 1,
			},
		},
		Tax: []payment_requests.Tax{
			{
				Name:   "VAT (7.5%)",
				Amount: 153750, // 7.5% of new subtotal
			},
		},
	}

	updatedFreelancerInvoice, err := client.PaymentRequests.Update(ctx, freelancerInvoice.RequestCode, updateReq)
	if err != nil {
		log.Fatalf("Failed to update freelancer invoice: %v", err)
	}

	fmt.Printf("✅ Updated freelancer invoice:\n")
	fmt.Printf("   Original Amount: ₦%.2f\n", float64(freelancerInvoice.Amount)/100)
	fmt.Printf("   Updated Amount: ₦%.2f\n", float64(updatedFreelancerInvoice.Amount)/100)
	fmt.Printf("   Increase: ₦%.2f\n", float64(updatedFreelancerInvoice.Amount-freelancerInvoice.Amount)/100)

	// Send notification for the updated invoice
	fmt.Printf("Sending notification for updated invoice...\n")
	notifyResp, err := client.PaymentRequests.SendNotification(ctx, freelancerInvoice.RequestCode)
	if err != nil {
		log.Printf("Failed to send notification (expected for demo): %v", err)
	} else {
		fmt.Printf("✅ Notification sent: %s\n", notifyResp.Message)
	}

	// Final summary and configuration export
	fmt.Println("\n📋 Management Summary")
	fmt.Println("====================")

	createdInvoices := []struct {
		name        string
		code        string
		amount      float64
		currency    string
		description string
	}{
		{"Freelancer Invoice", freelancerInvoice.RequestCode, float64(updatedFreelancerInvoice.Amount) / 100, updatedFreelancerInvoice.Currency, "Web development project"},
		{"SaaS Invoice", saasInvoice.RequestCode, float64(saasInvoice.Amount) / 100, saasInvoice.Currency, "Monthly subscription"},
		{"Event Invoice", eventInvoice.RequestCode, float64(eventInvoice.Amount) / 100, eventInvoice.Currency, "Corporate training"},
		{"Wholesale Invoice", wholesaleInvoice.RequestCode, float64(wholesaleInvoice.Amount) / 100, wholesaleInvoice.Currency, "Inventory restock"},
	}

	totalValue := 0.0
	for i, invoice := range createdInvoices {
		fmt.Printf("%d. %s\n", i+1, invoice.name)
		fmt.Printf("   Code: %s\n", invoice.code)
		fmt.Printf("   Amount: %.2f %s\n", invoice.amount, invoice.currency)
		fmt.Printf("   Purpose: %s\n", invoice.description)
		fmt.Printf("   Status: ✅ Created and ready for payment\n")
		fmt.Println()

		// Convert to NGN for total calculation (simplified)
		if invoice.currency == "NGN" {
			totalValue += invoice.amount
		} else if invoice.currency == "USD" {
			totalValue += invoice.amount * 800 // Rough conversion for demo
		}
	}

	// Export configuration for business intelligence
	fmt.Println("💾 Business Intelligence Export")
	fmt.Println("==============================")

	config := map[string]any{
		"created_at": time.Now().Format(time.RFC3339),
		"invoices": map[string]any{
			"freelancer": map[string]any{
				"code":     freelancerInvoice.RequestCode,
				"amount":   updatedFreelancerInvoice.Amount,
				"currency": updatedFreelancerInvoice.Currency,
				"type":     "professional_services",
			},
			"saas": map[string]any{
				"code":     saasInvoice.RequestCode,
				"amount":   saasInvoice.Amount,
				"currency": saasInvoice.Currency,
				"type":     "subscription",
			},
			"event": map[string]any{
				"code":     eventInvoice.RequestCode,
				"amount":   eventInvoice.Amount,
				"currency": eventInvoice.Currency,
				"type":     "event_booking",
			},
			"wholesale": map[string]any{
				"code":     wholesaleInvoice.RequestCode,
				"amount":   wholesaleInvoice.Amount,
				"currency": wholesaleInvoice.Currency,
				"type":     "bulk_order",
			},
		},
		"analytics": map[string]any{
			"total_requests":      len(createdInvoices),
			"estimated_value_ngn": totalValue,
			"status_distribution": statusCount,
			"currency_breakdown":  currencyTotals,
		},
	}

	configJSON, _ := json.MarshalIndent(config, "", "  ")
	fmt.Println(string(configJSON))

	fmt.Printf("\n🎉 Advanced Payment Requests Example Complete!\n")
	fmt.Printf("Created %d professional invoices across different business scenarios.\n", len(createdInvoices))
	fmt.Printf("Total estimated value: ₦%.2f\n", totalValue)
	fmt.Printf("All payment requests are ready for customer payment and can be tracked through the Paystack dashboard.\n")
}
