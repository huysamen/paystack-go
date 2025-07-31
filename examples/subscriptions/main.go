package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/subscriptions"
)

func main() {
	// Initialize client
	client := paystack.DefaultClient("sk_test_your_secret_key")

	// Create a subscription
	fmt.Println("Creating subscription...")
	createReq := &subscriptions.SubscriptionCreateRequest{
		Customer:      "customer@example.com",               // or customer code like "CUS_xnxdt6s1zg1f4nx"
		Plan:          "PLN_gx2wn530m0i3w3m",                // plan code
		Authorization: stringPtr("AUTH_6tmt288t0o"),         // optional - specific authorization
		StartDate:     timePtr(time.Now().AddDate(0, 0, 1)), // start tomorrow
	}

	createResp, err := client.Subscriptions.Create(context.Background(), createReq)
	if err != nil {
		log.Fatalf("Failed to create subscription: %v", err)
	}

	fmt.Printf("Subscription created: %s\n", createResp.Data.SubscriptionCode)
	fmt.Printf("Status: %s\n", createResp.Data.Status)
	fmt.Printf("Amount: ₦%.2f\n", float64(createResp.Data.Amount)/100)
	fmt.Printf("Customer: %s (%s)\n",
		createResp.Data.Customer.Email,
		createResp.Data.Customer.CustomerCode)
	fmt.Printf("Plan: %s (%s)\n",
		createResp.Data.Plan.Name,
		createResp.Data.Plan.PlanCode)

	if createResp.Data.NextPaymentDate != nil {
		fmt.Printf("Next payment: %s\n", createResp.Data.NextPaymentDate.Format("2006-01-02"))
	}

	// List subscriptions
	fmt.Println("\nListing subscriptions...")
	listReq := &subscriptions.SubscriptionListRequest{
		PerPage:  intPtr(10),
		Page:     intPtr(1),
		Customer: intPtr(createResp.Data.Customer.ID), // filter by customer
	}

	listResp, err := client.Subscriptions.List(context.Background(), listReq)
	if err != nil {
		log.Fatalf("Failed to list subscriptions: %v", err)
	}

	fmt.Printf("Found %d subscriptions\n", len(listResp.Data.Data))
	for _, subscription := range listResp.Data.Data {
		fmt.Printf("- %s: %s (₦%.2f - %s)\n",
			subscription.SubscriptionCode,
			subscription.Status,
			float64(subscription.Amount)/100,
			subscription.Plan.Name)
	}

	// Fetch subscription details with invoices
	fmt.Println("\nFetching subscription details...")
	fetchResp, err := client.Subscriptions.Fetch(context.Background(), createResp.Data.SubscriptionCode)
	if err != nil {
		log.Fatalf("Failed to fetch subscription: %v", err)
	}

	fmt.Printf("Subscription: %s\n", fetchResp.Data.SubscriptionCode)
	fmt.Printf("Status: %s\n", fetchResp.Data.Status)
	fmt.Printf("Customer: %s\n", fetchResp.Data.Customer.Email)
	fmt.Printf("Plan: %s (₦%.2f %s)\n",
		fetchResp.Data.Plan.Name,
		float64(fetchResp.Data.Plan.Amount)/100,
		fetchResp.Data.Plan.Interval)
	fmt.Printf("Authorization: %s (%s **** %s)\n",
		fetchResp.Data.Authorization.AuthorizationCode,
		fetchResp.Data.Authorization.Bank,
		fetchResp.Data.Authorization.Last4)
	fmt.Printf("Invoices: %d\n", len(fetchResp.Data.Invoices))

	// Generate update link for customer to update their payment method
	fmt.Println("\nGenerating update link...")
	updateLinkResp, err := client.Subscriptions.GenerateUpdateLink(context.Background(), createResp.Data.SubscriptionCode)
	if err != nil {
		log.Fatalf("Failed to generate update link: %v", err)
	}

	fmt.Printf("Update link: %s\n", updateLinkResp.Data.Link)

	// Send update link via email to customer
	fmt.Println("\nSending update link via email...")
	sendLinkResp, err := client.Subscriptions.SendUpdateLink(context.Background(), createResp.Data.SubscriptionCode)
	if err != nil {
		log.Fatalf("Failed to send update link: %v", err)
	}

	fmt.Printf("Email result: %s\n", sendLinkResp.Data.Message)

	// Example of disabling a subscription (requires email token)
	fmt.Println("\nNote: To enable/disable subscriptions, you need the email token")
	fmt.Println("Email token is typically obtained from webhook notifications or customer emails")

	// Example disable request (commented out as it requires valid token)
	/*
		disableReq := &subscriptions.SubscriptionDisableRequest{
			Code:  createResp.Data.SubscriptionCode,
			Token: "d7gofp6yppn3qz7", // from email or webhook
		}

		disableResp, err := client.Subscriptions.Disable(context.Background(), disableReq)
		if err != nil {
			log.Fatalf("Failed to disable subscription: %v", err)
		}

		fmt.Printf("Disable result: %s\n", disableResp.Data.Message)

		// Re-enable subscription
		enableReq := &subscriptions.SubscriptionEnableRequest{
			Code:  createResp.Data.SubscriptionCode,
			Token: "d7gofp6yppn3qz7", // same token
		}

		enableResp, err := client.Subscriptions.Enable(context.Background(), enableReq)
		if err != nil {
			log.Fatalf("Failed to enable subscription: %v", err)
		}

		fmt.Printf("Enable result: %s\n", enableResp.Data.Message)
	*/

	fmt.Println("\nSubscriptions API example completed successfully!")
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func timePtr(t time.Time) *time.Time {
	return &t
}
