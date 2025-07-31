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

	// Create a customer
	createReq := &customers.CustomerCreateRequest{
		Email:     "customer@example.com",
		FirstName: stringPtr("John"),
		LastName:  stringPtr("Doe"),
		Phone:     stringPtr("+2348123456789"),
		Metadata: map[string]any{
			"user_id": "12345",
			"type":    "premium",
		},
	}

	fmt.Println("Creating customer...")
	createResp, err := client.Customers.Create(context.Background(), createReq)
	if err != nil {
		log.Fatalf("Failed to create customer: %v", err)
	}

	fmt.Printf("Customer created: %s (%s)\n",
		createResp.Data.CustomerCode,
		createResp.Data.Email)

	// List customers
	fmt.Println("\nListing customers...")
	listReq := &customers.CustomerListRequest{
		PerPage: intPtr(10),
		Page:    intPtr(1),
	}

	listResp, err := client.Customers.List(context.Background(), listReq)
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

	// Update customer
	fmt.Println("\nUpdating customer...")
	updateReq := &customers.CustomerUpdateRequest{
		FirstName: stringPtr("Jane"),
		Phone:     stringPtr("+2348987654321"),
		Metadata: map[string]any{
			"user_id": "12345",
			"type":    "premium",
			"updated": true,
		},
	}

	updateResp, err := client.Customers.Update(context.Background(), createResp.Data.CustomerCode, updateReq)
	if err != nil {
		log.Fatalf("Failed to update customer: %v", err)
	}

	fmt.Printf("Customer updated: %s\n", *updateResp.Data.FirstName)

	// Set risk action (whitelist customer)
	fmt.Println("\nWhitelisting customer...")
	riskReq := &customers.CustomerRiskActionRequest{
		Customer:   createResp.Data.CustomerCode,
		RiskAction: (*customers.RiskAction)(&[]customers.RiskAction{customers.RiskActionAllow}[0]),
	}

	riskResp, err := client.Customers.SetRiskAction(context.Background(), riskReq)
	if err != nil {
		log.Fatalf("Failed to set risk action: %v", err)
	}

	fmt.Printf("Customer risk action set to: %s\n", riskResp.Data.RiskAction)

	// Initialize authorization for direct debit
	fmt.Println("\nInitializing authorization...")
	authReq := &customers.AuthorizationInitializeRequest{
		Email:   createResp.Data.Email,
		Channel: "direct_debit",
	}

	authResp, err := client.Customers.InitializeAuthorization(context.Background(), authReq)
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
