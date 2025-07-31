package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/huysamen/paystack-go"
	payment_pages "github.com/huysamen/paystack-go/api/payment-pages"
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

	fmt.Println("🚀 Payment Pages Advanced Example")
	fmt.Println("=================================")

	// Scenario 1: E-commerce product showcase page
	fmt.Println("\n📦 Scenario 1: E-commerce Product Showcase")
	fmt.Println("Creating a flexible payment page for multiple products...")

	collectPhone := true
	productPageReq := &payment_pages.CreatePaymentPageRequest{
		Name:         "Tech Store - Premium Electronics",
		Description:  "Browse and purchase premium electronics with flexible pricing",
		Type:         "product",
		Currency:     "NGN",
		CollectPhone: &collectPhone,
		CustomFields: []payment_pages.CustomField{
			{
				DisplayName:  "Delivery Address",
				VariableName: "delivery_address",
				Required:     true,
			},
			{
				DisplayName:  "Phone Number",
				VariableName: "phone_number",
				Required:     true,
			},
			{
				DisplayName:  "Special Instructions",
				VariableName: "special_instructions",
				Required:     false,
			},
		},
		Metadata: &types.Metadata{
			"store_id":     "tech-store-001",
			"category":     "electronics",
			"allow_pickup": "true",
			"contact":      "+234123456789",
		},
		RedirectURL:       "https://techstore.example.com/thank-you",
		SuccessMessage:    "Thank you! Your order has been received. We'll contact you within 24 hours for delivery arrangements.",
		NotificationEmail: "orders@techstore.example.com",
	}

	productPage, err := client.PaymentPages.Create(ctx, productPageReq)
	if err != nil {
		log.Fatalf("Failed to create product page: %v", err)
	}

	fmt.Printf("✅ Product showcase page created:\n")
	fmt.Printf("   Name: %s\n", productPage.Name)
	fmt.Printf("   URL: https://paystack.com/pay/%s\n", productPage.Slug)
	fmt.Printf("   Type: %s\n", productPage.Type)

	// Scenario 2: Subscription-based service
	fmt.Println("\n💳 Scenario 2: Subscription Service")
	fmt.Println("Creating a subscription payment page...")

	// Note: This would require a valid plan ID from the Plans API
	subscriptionPageReq := &payment_pages.CreatePaymentPageRequest{
		Name:        "Premium SaaS Subscription",
		Description: "Monthly subscription to our premium SaaS platform with advanced features",
		Type:        "subscription",
		Currency:    "USD",
		// Plan:        "PLN_subscription_plan_id", // Would need actual plan ID
		Metadata: &types.Metadata{
			"service_type":  "saas",
			"billing_cycle": "monthly",
			"trial_period":  "14_days",
			"support_level": "premium",
		},
		RedirectURL:    "https://saas.example.com/welcome",
		SuccessMessage: "Welcome to Premium! Your subscription is now active. Check your email for next steps.",
	}

	// For demo purposes, create as payment type since we don't have a plan
	subscriptionPageReq.Type = "payment"
	subscriptionPageReq.Amount = func() *int { amount := 2000; return &amount }() // $20.00
	fixedAmount := true
	subscriptionPageReq.FixedAmount = &fixedAmount

	subscriptionPage, err := client.PaymentPages.Create(ctx, subscriptionPageReq)
	if err != nil {
		log.Fatalf("Failed to create subscription page: %v", err)
	}

	fmt.Printf("✅ Subscription page created:\n")
	fmt.Printf("   Name: %s\n", subscriptionPage.Name)
	fmt.Printf("   URL: https://paystack.com/pay/%s\n", subscriptionPage.Slug)

	// Scenario 3: Donation/Fundraising page
	fmt.Println("\n💝 Scenario 3: Donation Campaign")
	fmt.Println("Creating a flexible donation page...")

	donationPageReq := &payment_pages.CreatePaymentPageRequest{
		Name:        "Help Build Schools in Nigeria",
		Description: "Your donation helps provide quality education to children in underserved communities. Every contribution makes a difference!",
		Type:        "payment",
		Currency:    "NGN",
		// No fixed amount - donors can choose
		CustomFields: []payment_pages.CustomField{
			{
				DisplayName:  "Donor Name (for recognition)",
				VariableName: "donor_name",
				Required:     false,
			},
			{
				DisplayName:  "Message of Support",
				VariableName: "support_message",
				Required:     false,
			},
			{
				DisplayName:  "Anonymous Donation",
				VariableName: "anonymous",
				Required:     false,
			},
		},
		Metadata: &types.Metadata{
			"campaign_id":   "school-building-2024",
			"cause":         "education",
			"target_amount": "5000000", // 50,000 NGN target
			"location":      "Lagos State, Nigeria",
		},
		SuccessMessage: "Thank you for your generous donation! Together, we're building a brighter future for Nigerian children.",
	}

	donationPage, err := client.PaymentPages.Create(ctx, donationPageReq)
	if err != nil {
		log.Fatalf("Failed to create donation page: %v", err)
	}

	fmt.Printf("✅ Donation page created:\n")
	fmt.Printf("   Name: %s\n", donationPage.Name)
	fmt.Printf("   URL: https://paystack.com/pay/%s\n", donationPage.Slug)

	// Scenario 4: Event ticket sales
	fmt.Println("\n🎫 Scenario 4: Event Ticketing")
	fmt.Println("Creating an event ticket sales page...")

	eventPageReq := &payment_pages.CreatePaymentPageRequest{
		Name:        "Tech Conference Lagos 2024",
		Description: "Join Nigeria's premier technology conference featuring industry leaders, workshops, and networking opportunities",
		Amount:      func() *int { amount := 15000; return &amount }(), // 150.00 NGN
		Currency:    "NGN",
		Type:        "payment",
		FixedAmount: &fixedAmount,
		Slug:        "tech-conf-lagos-2024",
		CustomFields: []payment_pages.CustomField{
			{
				DisplayName:  "Full Name",
				VariableName: "attendee_name",
				Required:     true,
			},
			{
				DisplayName:  "Company/Organization",
				VariableName: "company",
				Required:     false,
			},
			{
				DisplayName:  "Dietary Restrictions",
				VariableName: "dietary_restrictions",
				Required:     false,
			},
			{
				DisplayName:  "T-Shirt Size",
				VariableName: "tshirt_size",
				Required:     true,
			},
		},
		Metadata: &types.Metadata{
			"event_id":       "tech-conf-2024",
			"venue":          "Lagos Continental Hotel",
			"date":           "2024-11-15",
			"capacity":       "500",
			"includes_lunch": "true",
			"includes_kit":   "true",
		},
		RedirectURL:    "https://techconf.example.com/ticket-confirmation",
		SuccessMessage: "Your ticket is confirmed! Check your email for the e-ticket and event details.",
	}

	eventPage, err := client.PaymentPages.Create(ctx, eventPageReq)
	if err != nil {
		log.Fatalf("Failed to create event page: %v", err)
	}

	fmt.Printf("✅ Event ticketing page created:\n")
	fmt.Printf("   Name: %s\n", eventPage.Name)
	fmt.Printf("   URL: https://paystack.com/pay/%s\n", eventPage.Slug)

	// Scenario 5: Advanced page management
	fmt.Println("\n⚙️  Scenario 5: Advanced Page Management")
	fmt.Println("Demonstrating bulk operations and analytics...")

	// List all pages with pagination
	allPages := []payment_pages.PaymentPage{}
	page := 1
	for {
		listReq := &payment_pages.ListPaymentPagesRequest{
			PerPage: 10,
			Page:    page,
		}

		pagesResp, err := client.PaymentPages.List(ctx, listReq)
		if err != nil {
			log.Printf("Error listing pages: %v", err)
			break
		}

		allPages = append(allPages, pagesResp.Data...)

		if len(pagesResp.Data) < 10 {
			break // Last page
		}
		page++

		if page > 5 { // Safety limit for demo
			break
		}
	}

	fmt.Printf("📊 Found %d total payment pages\n", len(allPages))

	// Analyze pages by type
	typeCount := make(map[string]int)
	activeCount := 0
	totalRevenue := 0

	for _, p := range allPages {
		typeCount[p.Type]++
		if p.Active {
			activeCount++
		}
		if p.Amount != nil {
			totalRevenue += *p.Amount
		}
	}

	fmt.Printf("   📈 Analytics:\n")
	fmt.Printf("   - Active pages: %d/%d\n", activeCount, len(allPages))
	for pageType, count := range typeCount {
		fmt.Printf("   - %s pages: %d\n", pageType, count)
	}
	fmt.Printf("   - Potential revenue (fixed amount pages): ₦%.2f\n", float64(totalRevenue)/100)

	// Scenario 6: Batch slug availability check
	fmt.Println("\n🔍 Scenario 6: Batch Slug Validation")
	fmt.Println("Checking availability of multiple slugs...")

	proposedSlugs := []string{
		"premium-course-2024",
		"black-friday-deals",
		"startup-bootcamp",
		"charity-run-2024",
		"tech-meetup-monthly",
	}

	availableSlugs := []string{}
	for _, slug := range proposedSlugs {
		resp, err := client.PaymentPages.CheckSlugAvailability(ctx, slug)
		if err != nil {
			fmt.Printf("   ❌ %s: Error checking (%v)\n", slug, err)
		} else {
			if resp.Status {
				fmt.Printf("   ✅ %s: Available\n", slug)
				availableSlugs = append(availableSlugs, slug)
			} else {
				fmt.Printf("   ❌ %s: Not available\n", slug)
			}
		}
		time.Sleep(100 * time.Millisecond) // Rate limiting courtesy
	}

	fmt.Printf("   Available slugs: %d/%d\n", len(availableSlugs), len(proposedSlugs))

	// Scenario 7: Page lifecycle management
	fmt.Println("\n🔄 Scenario 7: Page Lifecycle Management")
	fmt.Println("Demonstrating page updates and management...")

	// Update the event page with early bird pricing
	fmt.Printf("Updating event page with early bird discount...\n")
	updateEventReq := &payment_pages.UpdatePaymentPageRequest{
		Name:        "Tech Conference Lagos 2024 - Early Bird Special",
		Description: "🐦 Early Bird Special: Save 33%! Join Nigeria's premier technology conference featuring industry leaders, workshops, and networking opportunities",
		Amount:      func() *int { amount := 10000; return &amount }(), // 100.00 NGN (discounted)
	}

	updatedEventPage, err := client.PaymentPages.Update(ctx, eventPage.Slug, updateEventReq)
	if err != nil {
		log.Printf("Failed to update event page: %v", err)
	} else {
		fmt.Printf("✅ Updated event page with early bird pricing\n")
		fmt.Printf("   New price: ₦%.2f (was ₦%.2f)\n",
			float64(*updatedEventPage.Amount)/100,
			float64(*eventPage.Amount)/100)
	}

	// Final summary with management insights
	fmt.Println("\n📋 Management Summary")
	fmt.Println("====================")

	createdPages := []struct {
		name     string
		slug     string
		purpose  string
		pageType string
	}{
		{productPage.Name, productPage.Slug, "E-commerce product sales", productPage.Type},
		{subscriptionPage.Name, subscriptionPage.Slug, "Subscription billing", subscriptionPage.Type},
		{donationPage.Name, donationPage.Slug, "Fundraising campaign", donationPage.Type},
		{eventPage.Name, eventPage.Slug, "Event ticket sales", eventPage.Type},
	}

	for i, page := range createdPages {
		fmt.Printf("%d. %s\n", i+1, page.name)
		fmt.Printf("   Purpose: %s\n", page.purpose)
		fmt.Printf("   Type: %s\n", page.pageType)
		fmt.Printf("   URL: https://paystack.com/pay/%s\n", page.slug)
		fmt.Printf("   Status: ✅ Ready to accept payments\n")
		fmt.Println()
	}

	// Export configuration for reference
	fmt.Println("💾 Configuration Export")
	fmt.Println("======================")

	config := map[string]interface{}{
		"created_at": time.Now().Format(time.RFC3339),
		"pages": map[string]interface{}{
			"product_showcase": map[string]interface{}{
				"slug":    productPage.Slug,
				"type":    "product",
				"purpose": "Multi-product e-commerce",
			},
			"subscription_service": map[string]interface{}{
				"slug":    subscriptionPage.Slug,
				"type":    "payment",
				"purpose": "SaaS subscription",
			},
			"donation_campaign": map[string]interface{}{
				"slug":    donationPage.Slug,
				"type":    "payment",
				"purpose": "Charitable donations",
			},
			"event_ticketing": map[string]interface{}{
				"slug":    eventPage.Slug,
				"type":    "payment",
				"purpose": "Event ticket sales",
			},
		},
		"analytics":    typeCount,
		"total_pages":  len(allPages),
		"active_pages": activeCount,
	}

	configJSON, _ := json.MarshalIndent(config, "", "  ")
	fmt.Println(string(configJSON))

	fmt.Printf("\n🎉 Advanced Payment Pages Example Complete!\n")
	fmt.Printf("Created %d payment pages for different use cases.\n", len(createdPages))
	fmt.Printf("All pages are ready to accept payments and can be customized further.\n")
}
