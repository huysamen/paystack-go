package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

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

	// Example 1: Create a payment page
	fmt.Println("=== Creating Payment Page ===")
	fixedAmount := true
	collectPhone := false
	createReq := &payment_pages.CreatePaymentPageRequest{
		Name:         "Premium Course Access",
		Description:  "One-time payment for premium course access",
		Amount:       func() *int { amount := 25000; return &amount }(), // 250.00 NGN
		Currency:     "NGN",
		Type:         "payment",
		FixedAmount:  &fixedAmount,
		CollectPhone: &collectPhone,
		Metadata: &types.Metadata{
			"course_id":  "premium-001",
			"instructor": "John Doe",
			"duration":   "6 months",
		},
		SuccessMessage: "Thank you for your payment! You now have access to the premium course.",
	}

	page, err := client.PaymentPages.Create(ctx, createReq)
	if err != nil {
		log.Fatalf("Failed to create payment page: %v", err)
	}

	fmt.Printf("Payment Page Created:\n")
	fmt.Printf("  ID: %d\n", page.ID)
	fmt.Printf("  Name: %s\n", page.Name)
	fmt.Printf("  Slug: %s\n", page.Slug)
	fmt.Printf("  Amount: %v NGN\n", float64(*page.Amount)/100)
	fmt.Printf("  URL: https://paystack.com/pay/%s\n", page.Slug)

	pageID := page.ID
	pageSlug := page.Slug

	// Example 2: Check slug availability
	fmt.Println("\n=== Checking Slug Availability ===")
	testSlug := "premium-course-2024"
	slugResp, err := client.PaymentPages.CheckSlugAvailability(ctx, testSlug)
	if err != nil {
		log.Printf("Error checking slug availability: %v", err)
	} else {
		fmt.Printf("Slug '%s' status: %s\n", testSlug, slugResp.Message)
	}

	// Example 3: Fetch the created payment page
	fmt.Println("\n=== Fetching Payment Page ===")
	fetchedPage, err := client.PaymentPages.Fetch(ctx, pageSlug)
	if err != nil {
		log.Fatalf("Failed to fetch payment page: %v", err)
	}

	fmt.Printf("Fetched Page:\n")
	fmt.Printf("  Name: %s\n", fetchedPage.Name)
	fmt.Printf("  Description: %s\n", fetchedPage.Description)
	fmt.Printf("  Active: %t\n", fetchedPage.Active)
	fmt.Printf("  Type: %s\n", fetchedPage.Type)

	// Example 4: Update the payment page
	fmt.Println("\n=== Updating Payment Page ===")
	updateReq := &payment_pages.UpdatePaymentPageRequest{
		Name:        "Premium Course Access - Limited Time",
		Description: "Special offer: One-time payment for premium course access with bonus materials",
		Amount:      func() *int { amount := 20000; return &amount }(), // 200.00 NGN (discounted)
	}

	updatedPage, err := client.PaymentPages.Update(ctx, pageSlug, updateReq)
	if err != nil {
		log.Fatalf("Failed to update payment page: %v", err)
	}

	fmt.Printf("Updated Page:\n")
	fmt.Printf("  Name: %s\n", updatedPage.Name)
	fmt.Printf("  Description: %s\n", updatedPage.Description)
	fmt.Printf("  Amount: %v NGN\n", float64(*updatedPage.Amount)/100)

	// Example 5: List payment pages
	fmt.Println("\n=== Listing Payment Pages ===")
	listReq := &payment_pages.ListPaymentPagesRequest{
		PerPage: 5,
		Page:    1,
	}

	pagesResp, err := client.PaymentPages.List(ctx, listReq)
	if err != nil {
		log.Fatalf("Failed to list payment pages: %v", err)
	}

	fmt.Printf("Found %d payment pages:\n", len(pagesResp.Data))
	for i, p := range pagesResp.Data {
		if i >= 3 { // Show only first 3 for brevity
			fmt.Printf("  ... and %d more\n", len(pagesResp.Data)-3)
			break
		}
		amountStr := "Variable"
		if p.Amount != nil {
			amountStr = fmt.Sprintf("%.2f %s", float64(*p.Amount)/100, p.Currency)
		}
		fmt.Printf("  %d. %s (%s) - %s - Active: %t\n",
			i+1, p.Name, p.Slug, amountStr, p.Active)
	}

	// Example 6: Add products to page (requires existing products)
	fmt.Println("\n=== Adding Products to Page (Demo) ===")
	// This is a demo of how to add products - requires actual product IDs
	// addProductsReq := &payment_pages.AddProductsToPageRequest{
	//     Product: []int{123, 456}, // Replace with actual product IDs
	// }
	//
	// productPage, err := client.PaymentPages.AddProducts(ctx, pageID, addProductsReq)
	// if err != nil {
	//     log.Printf("Note: Failed to add products (expected if no products exist): %v", err)
	// } else {
	//     fmt.Printf("Added products to page: %s\n", productPage.Name)
	// }

	fmt.Printf("Note: To add products, first create products using the Products API, then use their IDs\n")
	fmt.Printf("Example: client.PaymentPages.AddProducts(ctx, %d, &payment_pages.AddProductsToPageRequest{Product: []int{productID1, productID2}})\n", pageID)

	// Print final summary
	fmt.Println("\n=== Summary ===")
	fmt.Printf("✅ Created payment page: %s\n", page.Name)
	fmt.Printf("✅ Payment URL: https://paystack.com/pay/%s\n", page.Slug)
	fmt.Printf("✅ Updated with discount pricing\n")
	fmt.Printf("✅ Ready to accept payments!\n")

	// Print the full page data as JSON for reference
	fmt.Println("\n=== Full Page Data (JSON) ===")
	pageJSON, _ := json.MarshalIndent(updatedPage, "", "  ")
	fmt.Println(string(pageJSON))
}
