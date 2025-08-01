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
	"github.com/huysamen/paystack-go/types"
)

func main() {
	// Initialize client
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable is required")
	}

	client := paystack.DefaultClient(secretKey)
	ctx := context.Background()

	fmt.Println("🧾 Payment Requests Advanced Example - Builder Pattern Edition")
	fmt.Println("===========================================================")

	// Scenario 1: Freelancer Invoice System using Builder Pattern
	fmt.Println("\n💼 Scenario 1: Freelancer Invoice System")
	fmt.Println("Creating detailed invoice for web development project...")

	freelancerInvoice, err := client.PaymentRequests.Create(ctx,
		payment_requests.NewCreatePaymentRequestRequest().
			Customer("client@techstartup.com").
			Description("Web Development Project - E-commerce Platform").
			DueDate(time.Now().AddDate(0, 0, 30).Format("2006-01-02")). // 30 days payment terms
			Currency("NGN").
			HasInvoice(true).
			SendNotification(true).
			AddLineItem(payment_requests.LineItem{
				Name:     "Frontend Development (React.js)",
				Amount:   500000, // ₦5,000.00
				Quantity: 1,
			}).
			AddLineItem(payment_requests.LineItem{
				Name:     "Backend API Development (Node.js)",
				Amount:   400000, // ₦4,000.00
				Quantity: 1,
			}).
			AddLineItem(payment_requests.LineItem{
				Name:     "Database Design & Setup",
				Amount:   200000, // ₦2,000.00
				Quantity: 1,
			}).
			AddLineItem(payment_requests.LineItem{
				Name:     "Payment Gateway Integration",
				Amount:   150000, // ₦1,500.00
				Quantity: 1,
			}).
			AddLineItem(payment_requests.LineItem{
				Name:     "Testing & QA",
				Amount:   100000, // ₦1,000.00
				Quantity: 1,
			}).
			AddTax(payment_requests.Tax{
				Name:   "VAT (7.5%)",
				Amount: 103125, // ₦1,031.25
			}).
			Metadata(types.Metadata{
				"project_type": "web_development",
				"client_tier":  "enterprise",
				"rush_order":   "false",
			}),
	)
	if err != nil {
		log.Fatalf("Failed to create freelancer invoice: %v", err)
	}

	fmt.Printf("✅ Freelancer invoice created:\n")
	fmt.Printf("  Request Code: %s\n", freelancerInvoice.RequestCode)
	fmt.Printf("  Total: ₦%.2f (including VAT)\n", float64(freelancerInvoice.Amount)/100)

	// Scenario 2: SaaS Subscription Invoice using Builder Pattern
	fmt.Println("\n💻 Scenario 2: SaaS Subscription Invoice")
	fmt.Println("Creating recurring invoice for software subscription...")

	saasInvoice, err := client.PaymentRequests.Create(ctx,
		payment_requests.NewCreatePaymentRequestRequest().
			Customer("admin@smallbusiness.ng").
			Description("Project Management Software - Annual Subscription").
			DueDate(time.Now().AddDate(0, 0, 15).Format("2006-01-02")). // 15 days payment terms
			Currency("NGN").
			HasInvoice(true).
			InvoiceNumber(202400001).
			SendNotification(true).
			AddLineItem(payment_requests.LineItem{
				Name:     "Professional Plan (50 users)",
				Amount:   240000, // ₦2,400.00
				Quantity: 12, // 12 months
			}).
			AddLineItem(payment_requests.LineItem{
				Name:     "Premium Support",
				Amount:   60000, // ₦600.00
				Quantity: 12, // 12 months
			}).
			AddLineItem(payment_requests.LineItem{
				Name:     "Setup & Migration",
				Amount:   50000, // ₦500.00 (one-time)
				Quantity: 1,
			}).
			AddTax(payment_requests.Tax{
				Name:   "VAT (7.5%)",
				Amount: 262500, // ₦2,625.00
			}).
			Metadata(types.Metadata{
				"subscription_type": "annual",
				"user_count":        "50",
				"billing_cycle":     "yearly",
			}),
	)
	if err != nil {
		log.Fatalf("Failed to create SaaS invoice: %v", err)
	}

	fmt.Printf("✅ SaaS subscription invoice created:\n")
	fmt.Printf("  Request Code: %s\n", saasInvoice.RequestCode)
	fmt.Printf("  Total: ₦%.2f (annual subscription)\n", float64(saasInvoice.Amount)/100)

	// Scenario 3: Event Booking Invoice using Builder Pattern
	fmt.Println("\n🎭 Scenario 3: Event Booking Invoice")
	fmt.Println("Creating invoice for conference registration...")

	eventInvoice, err := client.PaymentRequests.Create(ctx,
		payment_requests.NewCreatePaymentRequestRequest().
			Customer("registration@techconf.ng").
			Description("Tech Conference 2024 - Registration & Services").
			DueDate(time.Now().AddDate(0, 0, 10).Format("2006-01-02")). // 10 days before event
			Currency("NGN").
			HasInvoice(true).
			SendNotification(true).
			AddLineItem(payment_requests.LineItem{
				Name:     "Early Bird Registration",
				Amount:   75000, // ₦750.00
				Quantity: 5, // 5 attendees
			}).
			AddLineItem(payment_requests.LineItem{
				Name:     "Workshop Access (AI/ML Track)",
				Amount:   25000, // ₦250.00
				Quantity: 3, // 3 attendees
			}).
			AddLineItem(payment_requests.LineItem{
				Name:     "Premium Networking Dinner",
				Amount:   15000, // ₦150.00
				Quantity: 5, // 5 attendees
			}).
			AddLineItem(payment_requests.LineItem{
				Name:     "Conference Swag Package",
				Amount:   8000, // ₦80.00
				Quantity: 5, // 5 attendees
			}).
			AddTax(payment_requests.Tax{
				Name:   "Service Charge (5%)",
				Amount: 26500, // ₦265.00
			}).
			Metadata(types.Metadata{
				"event_name":     "Tech Conference 2024",
				"attendee_count": "5",
				"registration_type": "corporate",
			}),
	)
	if err != nil {
		log.Fatalf("Failed to create event invoice: %v", err)
	}

	fmt.Printf("✅ Event booking invoice created:\n")
	fmt.Printf("  Request Code: %s\n", eventInvoice.RequestCode)
	fmt.Printf("  Total: ₦%.2f (for 5 attendees)\n", float64(eventInvoice.Amount)/100)

	// Scenario 4: Wholesale Order Invoice using Builder Pattern
	fmt.Println("\n📦 Scenario 4: Wholesale Order Invoice")
	fmt.Println("Creating invoice for bulk product order...")

	wholesaleInvoice, err := client.PaymentRequests.Create(ctx,
		payment_requests.NewCreatePaymentRequestRequest().
			Customer("orders@retailstore.ng").
			Description("Wholesale Electronics Order - Q1 2024").
			DueDate(time.Now().AddDate(0, 0, 45).Format("2006-01-02")). // 45 days payment terms
			Currency("NGN").
			HasInvoice(true).
			InvoiceNumber(202400002).
			SendNotification(true).
			AddLineItem(payment_requests.LineItem{
				Name:     "Wireless Headphones (Premium)",
				Amount:   25000, // ₦250.00 each
				Quantity: 50,
			}).
			AddLineItem(payment_requests.LineItem{
				Name:     "Bluetooth Speakers (Portable)",
				Amount:   35000, // ₦350.00 each
				Quantity: 30,
			}).
			AddLineItem(payment_requests.LineItem{
				Name:     "Phone Cases (Assorted)",
				Amount:   3000, // ₦30.00 each
				Quantity: 200,
			}).
			AddLineItem(payment_requests.LineItem{
				Name:     "Charging Cables (USB-C)",
				Amount:   2500, // ₦25.00 each
				Quantity: 100,
			}).
			AddLineItem(payment_requests.LineItem{
				Name:     "Shipping & Handling",
				Amount:   75000, // ₦750.00
				Quantity: 1,
			}).
			AddTax(payment_requests.Tax{
				Name:   "VAT (7.5%)",
				Amount: 183750, // ₦1,837.50
			}).
			AddTax(payment_requests.Tax{
				Name:   "Import Duty (2%)",
				Amount: 49000, // ₦490.00
			}).
			Metadata(types.Metadata{
				"order_type":       "wholesale",
				"customer_tier":    "gold",
				"bulk_discount":    "15%",
				"payment_terms":    "45_days",
			}),
	)
	if err != nil {
		log.Fatalf("Failed to create wholesale invoice: %v", err)
	}

	fmt.Printf("✅ Wholesale order invoice created:\n")
	fmt.Printf("  Request Code: %s\n", wholesaleInvoice.RequestCode)
	fmt.Printf("  Total: ₦%.2f (bulk order)\n", float64(wholesaleInvoice.Amount)/100)

	// Collect all request codes for further operations
	requestCodes := []string{
		freelancerInvoice.RequestCode,
		saasInvoice.RequestCode,
		eventInvoice.RequestCode,
		wholesaleInvoice.RequestCode,
	}

	// Advanced Listing Examples using Builder Pattern
	fmt.Println("\n📊 Advanced Payment Request Management")
	fmt.Println("=====================================")

	// List all recent payment requests using builder pattern
	allRequests, err := client.PaymentRequests.List(ctx,
		payment_requests.NewListPaymentRequestsRequest().
			PerPage(20).
			Page(1).
			Status("pending"),
	)
	if err != nil {
		log.Fatalf("Failed to list payment requests: %v", err)
	}

	fmt.Printf("\n📋 All Recent Payment Requests (%d found):\n", len(allRequests.Data))
	for i, req := range allRequests.Data {
		if i >= 10 { // Show only first 10
			fmt.Printf("  ... and %d more\n", len(allRequests.Data)-10)
			break
		}
		fmt.Printf("  %d. %s - ₦%.2f (%s)\n",
			i+1, req.Description, float64(req.Amount)/100, req.Status)
	}

	// List high-value requests using builder pattern
	highValueRequests, err := client.PaymentRequests.List(ctx,
		payment_requests.NewListPaymentRequestsRequest().
			PerPage(10).
			Currency("NGN").
			DateRange(
				time.Now().AddDate(0, 0, -7).Format("2006-01-02"), // Last 7 days
				time.Now().Format("2006-01-02"),
			),
	)
	if err != nil {
		log.Printf("Failed to list high-value requests: %v", err)
	} else {
		fmt.Printf("\n💰 Recent High-Value Requests (%d found):\n", len(highValueRequests.Data))
		for _, req := range highValueRequests.Data {
			if req.Amount >= 1000000 { // Only show requests >= ₦10,000
				fmt.Printf("  • %s: ₦%.2f\n", req.Description, float64(req.Amount)/100)
			}
		}
	}

	// Update one of the invoices using builder pattern
	fmt.Println("\n✏️ Updating Freelancer Invoice")
	updatedFreelancerInvoice, err := client.PaymentRequests.Update(ctx, freelancerInvoice.RequestCode,
		payment_requests.NewUpdatePaymentRequestRequest().
			Description("Web Development Project - E-commerce Platform (URGENT)").
			DueDate(time.Now().AddDate(0, 0, 15).Format("2006-01-02")). // Shortened to 15 days
			AddLineItem(payment_requests.LineItem{
				Name:     "Rush Delivery Bonus",
				Amount:   100000, // ₦1,000.00
				Quantity: 1,
			}).
			SendNotification(true),
	)
	if err != nil {
		log.Printf("Failed to update freelancer invoice: %v", err)
	} else {
		fmt.Printf("✅ Updated freelancer invoice: ₦%.2f\n", float64(updatedFreelancerInvoice.Amount)/100)
	}

	// Get payment request totals
	fmt.Println("\n📈 Payment Request Analytics")
	totals, err := client.PaymentRequests.GetTotals(ctx)
	if err != nil {
		log.Printf("Failed to get totals: %v", err)
	} else {
		fmt.Printf("Analytics Summary:\n")
		for _, total := range totals.Total {
			fmt.Printf("  Total %s: ₦%.2f\n", total.Currency, float64(total.Amount)/100)
		}
		for _, pending := range totals.Pending {
			fmt.Printf("  Pending %s: ₦%.2f\n", pending.Currency, float64(pending.Amount)/100)
		}
		for _, successful := range totals.Successful {
			fmt.Printf("  Successful %s: ₦%.2f\n", successful.Currency, float64(successful.Amount)/100)
		}
	}

	// Verify a few payment requests
	fmt.Println("\n🔍 Verification Examples")
	for i, code := range requestCodes[:2] { // Verify first 2
		verified, err := client.PaymentRequests.Verify(ctx, code)
		if err != nil {
			log.Printf("Failed to verify %s: %v", code, err)
			continue
		}
		fmt.Printf("  %d. %s: %s (Paid: %t)\n",
			i+1, verified.Description, verified.Status, verified.Paid)
	}

	// Send notifications using existing method (no builder needed)
	fmt.Println("\n📧 Sending Notifications")
	for i, code := range requestCodes[:2] { // Send to first 2
		_, err := client.PaymentRequests.SendNotification(ctx, code)
		if err != nil {
			fmt.Printf("  %d. Failed to send notification for %s: %v\n", i+1, code, err)
		} else {
			fmt.Printf("  %d. ✅ Notification sent for %s\n", i+1, code)
		}
	}

	// Finalize demo (using builder pattern)
	fmt.Println("\n🎯 Finalization Demo")
	_, err = client.PaymentRequests.Finalize(ctx, requestCodes[0],
		payment_requests.NewFinalizePaymentRequestRequest().
			SendNotification(true),
	)
	if err != nil {
		fmt.Printf("Finalization demo (expected to fail if not draft): %v\n", err)
	} else {
		fmt.Printf("✅ Request finalized successfully\n")
	}

	// Summary
	fmt.Println("\n📋 Advanced Example Summary")
	fmt.Println("===========================")
	fmt.Printf("✅ Created 4 different types of payment requests:\n")
	fmt.Printf("   • Freelancer Project Invoice: %s\n", freelancerInvoice.RequestCode)
	fmt.Printf("   • SaaS Subscription Invoice: %s\n", saasInvoice.RequestCode)
	fmt.Printf("   • Event Booking Invoice: %s\n", eventInvoice.RequestCode)
	fmt.Printf("   • Wholesale Order Invoice: %s\n", wholesaleInvoice.RequestCode)
	
	fmt.Printf("\n💡 Builder Pattern Benefits Demonstrated:\n")
	fmt.Printf("   • Fluent method chaining for intuitive API usage\n")
	fmt.Printf("   • Type-safe request construction\n")
	fmt.Printf("   • Easy addition of line items and taxes\n")
	fmt.Printf("   • Flexible metadata and configuration options\n")
	fmt.Printf("   • Advanced filtering in list operations\n")
	
	fmt.Printf("\n🚀 All operations used modern builder patterns!\n")

	// Export one invoice as JSON for reference
	fmt.Println("\n📄 Sample Invoice JSON (Freelancer):")
	invoiceJSON, _ := json.MarshalIndent(freelancerInvoice, "", "  ")
	fmt.Println(string(invoiceJSON))
}
