package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/customers"
	"github.com/huysamen/paystack-go/api/plans"
	"github.com/huysamen/paystack-go/api/subscriptions"
	"github.com/huysamen/paystack-go/types"
)

func main() {
	// Initialize client
	client := paystack.DefaultClient("sk_test_your_secret_key")

	// Complete workflow: Create customer -> Create plan -> Create subscription
	fmt.Println("=== Complete Subscription Workflow ===")

	// 1. Create a customer
	fmt.Println("\n1. Creating customer...")
	customerReq := &customers.CustomerCreateRequest{
		Email:     "subscriber@example.com",
		FirstName: stringPtr("Premium"),
		LastName:  stringPtr("Subscriber"),
		Phone:     stringPtr("+2348123456789"),
		Metadata: map[string]any{
			"subscription_tier": "premium",
			"signup_source":     "website",
		},
	}

	customerResp, err := client.Customers.Create(context.Background(), customerReq)
	if err != nil {
		log.Fatalf("Failed to create customer: %v", err)
	}

	fmt.Printf("Customer created: %s (%s)\n",
		customerResp.Data.CustomerCode,
		customerResp.Data.Email)

	// 2. Create a subscription plan
	fmt.Println("\n2. Creating subscription plan...")
	planReq := &plans.PlanCreateRequest{
		Name:         "Premium Monthly Subscription",
		Amount:       5000000, // ₦50,000 in kobo
		Interval:     types.IntervalMonthly,
		Currency:     types.CurrencyNGN,
		Description:  "Premium monthly subscription with all features",
		SendInvoices: boolPtr(true),
		SendSMS:      boolPtr(true),
	}

	planResp, err := client.Plans.Create(context.Background(), planReq)
	if err != nil {
		log.Fatalf("Failed to create plan: %v", err)
	}

	fmt.Printf("Plan created: %s (₦%.2f %s)\n",
		planResp.Data.PlanCode,
		float64(planResp.Data.Amount)/100,
		planResp.Data.Interval)

	// 3. Create subscription
	fmt.Println("\n3. Creating subscription...")
	subscriptionReq := &subscriptions.SubscriptionCreateRequest{
		Customer:  customerResp.Data.CustomerCode,
		Plan:      planResp.Data.PlanCode,
		StartDate: timePtr(time.Now().AddDate(0, 0, 7)), // Start in 1 week
	}

	subscriptionResp, err := client.Subscriptions.Create(context.Background(), subscriptionReq)
	if err != nil {
		log.Fatalf("Failed to create subscription: %v", err)
	}

	fmt.Printf("Subscription created: %s\n", subscriptionResp.Data.SubscriptionCode)
	fmt.Printf("Status: %s\n", subscriptionResp.Data.Status)
	fmt.Printf("Amount: ₦%.2f\n", float64(subscriptionResp.Data.Amount)/100)

	// 4. Fetch detailed subscription information
	fmt.Println("\n4. Fetching subscription details...")
	detailResp, err := client.Subscriptions.Fetch(context.Background(), subscriptionResp.Data.SubscriptionCode)
	if err != nil {
		log.Fatalf("Failed to fetch subscription details: %v", err)
	}

	fmt.Printf("=== Subscription Details ===\n")
	fmt.Printf("Code: %s\n", detailResp.Data.SubscriptionCode)
	fmt.Printf("Status: %s\n", detailResp.Data.Status)
	fmt.Printf("Domain: %s\n", detailResp.Data.Domain)

	if detailResp.Data.NextPaymentDate != nil {
		fmt.Printf("Next Payment: %s\n", detailResp.Data.NextPaymentDate.Format("2006-01-02 15:04:05"))
	}

	fmt.Printf("\n=== Customer Info ===\n")
	fmt.Printf("Name: %s %s\n",
		stringValue(detailResp.Data.Customer.FirstName),
		stringValue(detailResp.Data.Customer.LastName))
	fmt.Printf("Email: %s\n", detailResp.Data.Customer.Email)
	fmt.Printf("Customer Code: %s\n", detailResp.Data.Customer.CustomerCode)

	fmt.Printf("\n=== Plan Info ===\n")
	fmt.Printf("Name: %s\n", detailResp.Data.Plan.Name)
	fmt.Printf("Amount: ₦%.2f\n", float64(detailResp.Data.Plan.Amount)/100)
	fmt.Printf("Interval: %s\n", detailResp.Data.Plan.Interval)
	fmt.Printf("Currency: %s\n", detailResp.Data.Plan.Currency)

	fmt.Printf("\n=== Authorization Info ===\n")
	fmt.Printf("Code: %s\n", detailResp.Data.Authorization.AuthorizationCode)
	fmt.Printf("Bank: %s\n", detailResp.Data.Authorization.Bank)
	fmt.Printf("Card: **** **** **** %s\n", detailResp.Data.Authorization.Last4)
	fmt.Printf("Brand: %s\n", detailResp.Data.Authorization.Brand)
	fmt.Printf("Channel: %s\n", detailResp.Data.Authorization.Channel)

	fmt.Printf("\n=== Invoices ===\n")
	if len(detailResp.Data.Invoices) == 0 {
		fmt.Println("No invoices yet")
	} else {
		for i, invoice := range detailResp.Data.Invoices {
			fmt.Printf("Invoice %d:\n", i+1)
			fmt.Printf("  Code: %s\n", invoice.InvoiceCode)
			fmt.Printf("  Amount: ₦%.2f\n", float64(invoice.Amount)/100)
			fmt.Printf("  Status: %s\n", invoice.Status)
			fmt.Printf("  Period: %s to %s\n",
				invoice.PeriodStart.Format("2006-01-02"),
				invoice.PeriodEnd.Format("2006-01-02"))
		}
	}

	// 5. List all subscriptions for reporting
	fmt.Println("\n5. Listing all subscriptions...")
	allSubsReq := &subscriptions.SubscriptionListRequest{
		PerPage: intPtr(20),
		Page:    intPtr(1),
	}

	allSubsResp, err := client.Subscriptions.List(context.Background(), allSubsReq)
	if err != nil {
		log.Fatalf("Failed to list subscriptions: %v", err)
	}

	fmt.Printf("Total subscriptions: %d\n", len(allSubsResp.Data.Data))

	// Group by status
	statusCount := make(map[string]int)
	totalRevenue := 0

	for _, sub := range allSubsResp.Data.Data {
		statusCount[sub.Status]++
		totalRevenue += sub.Amount
	}

	fmt.Printf("\n=== Subscription Analytics ===\n")
	for status, count := range statusCount {
		fmt.Printf("%s: %d subscriptions\n", status, count)
	}
	fmt.Printf("Total Monthly Revenue: ₦%.2f\n", float64(totalRevenue)/100)

	// 6. Generate management link for customer
	fmt.Println("\n6. Generating customer management link...")
	manageLinkResp, err := client.Subscriptions.GenerateUpdateLink(context.Background(), subscriptionResp.Data.SubscriptionCode)
	if err != nil {
		log.Fatalf("Failed to generate management link: %v", err)
	}

	fmt.Printf("Customer can manage subscription at: %s\n", manageLinkResp.Data.Link)

	// 7. Send management link via email
	fmt.Println("\n7. Sending management link to customer...")
	emailResp, err := client.Subscriptions.SendUpdateLink(context.Background(), subscriptionResp.Data.SubscriptionCode)
	if err != nil {
		log.Fatalf("Failed to send management link: %v", err)
	}

	fmt.Printf("Email sent: %s\n", emailResp.Data.Message)

	fmt.Println("\n=== Subscription Workflow Complete ===")
	fmt.Printf("Customer %s is now subscribed to %s\n",
		customerResp.Data.Email,
		planResp.Data.Name)
	fmt.Printf("Subscription Code: %s\n", subscriptionResp.Data.SubscriptionCode)
	fmt.Printf("Management Link sent to customer's email\n")

	fmt.Println("\nAdvanced subscriptions example completed successfully!")
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func stringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
