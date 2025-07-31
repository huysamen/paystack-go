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

	// Initialize direct debit
	fmt.Println("Initializing direct debit...")
	directDebitReq := &customers.DirectDebitInitializeRequest{
		Account: customers.Account{
			Number:   "0123456789",
			BankCode: "058", // GTBank
		},
		Address: customers.Address{
			Street: "123 Main Street",
			City:   "Lagos",
			State:  "Lagos",
		},
	}

	directDebitResp, err := client.Customers.InitializeDirectDebit(context.Background(), customerID, directDebitReq)
	if err != nil {
		log.Fatalf("Failed to initialize direct debit: %v", err)
	}

	fmt.Printf("Direct debit initialized: %s\n", directDebitResp.Data.Reference)
	fmt.Printf("Access code: %s\n", directDebitResp.Data.AccessCode)

	// Validate customer identity
	fmt.Println("\nValidating customer identity...")
	validateReq := &customers.CustomerValidateRequest{
		FirstName:     "John",
		LastName:      "Doe",
		Type:          "bank_account",
		Value:         "0123456789",
		Country:       "NG",
		BVN:           "20012345677",
		BankCode:      "058",
		AccountNumber: "0123456789",
	}

	validateResp, err := client.Customers.Validate(context.Background(), customerCode, validateReq)
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

	// Trigger activation charge for inactive mandate
	if len(mandateResp.Data) > 0 {
		authID := mandateResp.Data[0].AuthorizationID
		fmt.Printf("\nTriggering activation charge for authorization %d...\n", authID)

		chargeReq := &customers.DirectDebitActivationChargeRequest{
			AuthorizationID: authID,
		}

		chargeResp, err := client.Customers.DirectDebitActivationCharge(context.Background(), customerID, chargeReq)
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

	// Finally, deactivate an authorization
	fmt.Println("\nDeactivating authorization...")
	deactivateReq := &customers.DeactivateAuthorizationRequest{
		AuthorizationCode: verifyResp.Data.AuthorizationCode,
	}

	deactivateResp, err := client.Customers.DeactivateAuthorization(context.Background(), deactivateReq)
	if err != nil {
		log.Fatalf("Failed to deactivate authorization: %v", err)
	}

	fmt.Printf("Deactivation: %s\n", deactivateResp.Data.Message)

	fmt.Println("\nAdvanced customer operations example completed!")
}
