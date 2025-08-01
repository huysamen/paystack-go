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

	// Example customer ID (replace with actual customer ID)
	customerID := "123456"
	customerCode := "CUS_xnxdt6s1zg1f4nx"

	// Initialize direct debit using builder pattern
	fmt.Println("Initializing direct debit...")
	directDebitBuilder := customers.NewInitializeDirectDebitRequest(
		"0123456789",      // account number
		"058",             // GTBank code
		"123 Main Street", // street
		"Lagos",           // city
		"Lagos",           // state
	)

	directDebitResp, err := client.Customers.InitializeDirectDebit(context.Background(), customerID, directDebitBuilder)
	if err != nil {
		log.Fatalf("Failed to initialize direct debit: %v", err)
	}

	fmt.Printf("Direct debit initialized: %s\n", directDebitResp.Data.Reference)
	fmt.Printf("Access code: %s\n", directDebitResp.Data.AccessCode)

	// Validate customer identity using builder pattern
	fmt.Println("\nValidating customer identity...")
	validateBuilder := customers.NewValidateCustomerRequest(
		"John",         // firstName
		"Doe",          // lastName
		"bank_account", // type
		"0123456789",   // value
		"NG",           // country
		"20012345677",  // bvn
	).BankCode("058").AccountNumber("0123456789")

	validateResp, err := client.Customers.Validate(context.Background(), customerCode, validateBuilder)
	if err != nil {
		log.Fatalf("Failed to validate customer: %v", err)
	}

	fmt.Printf("Validation initiated: %s\n", validateResp.Data.Message)

	// Fetch mandate authorizations
	fmt.Println("\nFetching mandate authorizations...")
	mandateResp, err := client.Customers.FetchMandateAuthorizations(context.Background(), customerID)
	if err != nil {
		log.Fatalf("Failed to fetch mandate authorizations: %v", err)
	}

	fmt.Printf("Found %d mandate authorizations\n", len(mandateResp.Data))
	for _, mandate := range mandateResp.Data {
		fmt.Printf("- Authorization %s: %s (%s)\n",
			mandate.AuthorizationCode,
			mandate.Status,
			mandate.AccountNumber)
	}

	// Trigger activation charge for inactive mandate using builder pattern
	if len(mandateResp.Data) > 0 {
		authID := mandateResp.Data[0].AuthorizationID
		fmt.Printf("\nTriggering activation charge for authorization %d...\n", authID)

		chargeBuilder := customers.NewDirectDebitActivationChargeRequest(authID)

		chargeResp, err := client.Customers.DirectDebitActivationCharge(context.Background(), customerID, chargeBuilder)
		if err != nil {
			log.Fatalf("Failed to trigger activation charge: %v", err)
		}

		fmt.Printf("Activation charge: %s\n", chargeResp.Data.Message)
	}

	// Verify authorization (using reference from earlier initialization)
	fmt.Println("\nVerifying authorization...")
	reference := directDebitResp.Data.Reference
	verifyResp, err := client.Customers.VerifyAuthorization(context.Background(), reference)
	if err != nil {
		log.Fatalf("Failed to verify authorization: %v", err)
	}

	fmt.Printf("Authorization verified: %s\n", verifyResp.Data.AuthorizationCode)
	fmt.Printf("Bank: %s\n", verifyResp.Data.Bank)
	fmt.Printf("Active: %v\n", verifyResp.Data.Active)

	// Finally, deactivate an authorization using builder pattern
	fmt.Println("\nDeactivating authorization...")
	deactivateBuilder := customers.NewDeactivateAuthorizationRequest(verifyResp.Data.AuthorizationCode)

	deactivateResp, err := client.Customers.DeactivateAuthorization(context.Background(), deactivateBuilder)
	if err != nil {
		log.Fatalf("Failed to deactivate authorization: %v", err)
	}

	fmt.Printf("Deactivation: %s\n", deactivateResp.Data.Message)

	fmt.Println("\nAdvanced customer operations example completed!")
}
