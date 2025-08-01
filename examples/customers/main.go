package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/customers"
)

func main() {
	// Initialize client
	client := paystack.DefaultClient("sk_test_your_secret_key")

	// Create a customer using builder pattern
	createBuilder := customers.NewCreateCustomerRequest("customer@example.com").
		FirstName("John").
		LastName("Doe").
		Phone("+2348123456789").
		Metadata(map[string]any{
			"user_id": "12345",
			"type":    "premium",
		})

	fmt.Println("Creating customer...")
	createResp, err := client.Customers.Create(context.Background(), createBuilder)
	if err != nil {
		log.Fatalf("Failed to create customer: %v", err)
	}

	fmt.Printf("Customer created: %s (%s)\n",
		createResp.Data.CustomerCode,
		createResp.Data.Email)

	// List customers
	fmt.Println("\nListing customers...")
	builder := customers.NewCustomerListRequest().
		PerPage(10).
		Page(1)

	listResp, err := client.Customers.List(context.Background(), builder)
	if err != nil {
		log.Fatalf("Failed to list customers: %v", err)
	}

	fmt.Printf("Found %d customers\n", len(listResp.Data.Data))
	for _, customer := range listResp.Data.Data {
		fmt.Printf("- %s: %s\n", customer.CustomerCode, customer.Email)
	}

	// Fetch a specific customer
	fmt.Println("\nFetching customer details...")
	fetchResp, err := client.Customers.Fetch(context.Background(), createResp.Data.CustomerCode)
	if err != nil {
		log.Fatalf("Failed to fetch customer: %v", err)
	}

	fmt.Printf("Customer details: %s (%s)\n",
		fetchResp.Data.CustomerCode,
		fetchResp.Data.Email)
	fmt.Printf("- Authorizations: %d\n", len(fetchResp.Data.Authorizations))
	fmt.Printf("- Transactions: %d\n", len(fetchResp.Data.Transactions))

	// Update customer using builder pattern
	fmt.Println("\nUpdating customer...")
	updateBuilder := customers.NewUpdateCustomerRequest().
		FirstName("Jane").
		Phone("+2348987654321").
		Metadata(map[string]any{
			"user_id": "12345",
			"type":    "premium",
			"updated": true,
		})

	updateResp, err := client.Customers.Update(context.Background(), createResp.Data.CustomerCode, updateBuilder)
	if err != nil {
		log.Fatalf("Failed to update customer: %v", err)
	}

	fmt.Printf("Customer updated: %s\n", updateResp.Data.FirstName)

	// Set risk action (whitelist customer) using builder pattern
	fmt.Println("\nWhitelisting customer...")
	riskBuilder := customers.NewSetRiskActionRequest(createResp.Data.CustomerCode).
		RiskAction(customers.RiskActionAllow)

	riskResp, err := client.Customers.SetRiskAction(context.Background(), riskBuilder)
	if err != nil {
		log.Fatalf("Failed to set risk action: %v", err)
	}

	fmt.Printf("Customer risk action set to: %s\n", riskResp.Data.RiskAction)

	// Initialize authorization for direct debit using builder pattern
	fmt.Println("\nInitializing authorization...")
	authBuilder := customers.NewInitializeAuthorizationRequest(createResp.Data.Email, "direct_debit")

	authResp, err := client.Customers.InitializeAuthorization(context.Background(), authBuilder)
	if err != nil {
		log.Fatalf("Failed to initialize authorization: %v", err)
	}

	fmt.Printf("Authorization initialized: %s\n", authResp.Data.Reference)
	fmt.Printf("Redirect URL: %s\n", authResp.Data.RedirectURL)

	fmt.Println("\nCustomer API example completed successfully!")
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
