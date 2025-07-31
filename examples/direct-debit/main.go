package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huysamen/paystack-go"
	directdebit "github.com/huysamen/paystack-go/api/direct-debit"
)

func main() {
	// Initialize client with test secret key
	client := paystack.DefaultClient("sk_test_your_secret_key")
	ctx := context.Background()

	fmt.Println("=== Paystack Direct Debit API Examples ===")
	fmt.Println()

	// Example 1: List Mandate Authorizations
	fmt.Println("1. Listing mandate authorizations...")
	listReq := &directdebit.ListMandateAuthorizationsRequest{
		Status:  directdebit.MandateAuthorizationStatusActive,
		PerPage: 10,
	}

	mandates, err := client.DirectDebit.ListMandateAuthorizations(ctx, listReq)
	if err != nil {
		log.Printf("Error listing mandate authorizations: %v", err)
		return
	}
	fmt.Printf("Found %d mandate authorizations\n", len(mandates.Data))
	for _, mandate := range mandates.Data {
		fmt.Printf("- ID: %d, Status: %s, Bank: %s (%s)\n",
			mandate.ID, mandate.Status, mandate.BankName, mandate.BankCode)
		fmt.Printf("  Customer: %s (%s)\n", mandate.Customer.Email, mandate.Customer.CustomerCode)
		fmt.Printf("  Account: %s, Auth Code: %s\n", mandate.AccountNumber, mandate.AuthorizationCode)
	}

	// Example 2: List all mandate authorizations (no filters)
	fmt.Println("\n2. Listing all mandate authorizations...")
	allMandates, err := client.DirectDebit.ListMandateAuthorizations(ctx, nil)
	if err != nil {
		log.Printf("Error listing all mandate authorizations: %v", err)
		return
	}
	fmt.Printf("Found %d total mandate authorizations\n", len(allMandates.Data))

	// Example 3: List pending mandate authorizations
	fmt.Println("\n3. Listing pending mandate authorizations...")
	pendingReq := &directdebit.ListMandateAuthorizationsRequest{
		Status:  directdebit.MandateAuthorizationStatusPending,
		PerPage: 5,
	}

	pendingMandates, err := client.DirectDebit.ListMandateAuthorizations(ctx, pendingReq)
	if err != nil {
		log.Printf("Error listing pending mandate authorizations: %v", err)
		return
	}
	fmt.Printf("Found %d pending mandate authorizations\n", len(pendingMandates.Data))

	// Example 4: Trigger Activation Charge (if there are pending mandates)
	if len(pendingMandates.Data) > 0 {
		fmt.Println("\n4. Triggering activation charge for pending mandates...")

		// Collect customer IDs from pending mandates
		var customerIDs []int
		for _, mandate := range pendingMandates.Data {
			customerIDs = append(customerIDs, mandate.Customer.ID)
			// Limit to first 3 customers for example
			if len(customerIDs) >= 3 {
				break
			}
		}

		activationReq := &directdebit.TriggerActivationChargeRequest{
			CustomerIDs: customerIDs,
		}

		activationResp, err := client.DirectDebit.TriggerActivationCharge(ctx, activationReq)
		if err != nil {
			log.Printf("Error triggering activation charge: %v", err)
			// This might fail in test mode, continue with example
		} else {
			fmt.Printf("Activation charge triggered: %s\n", activationResp.Message)
		}
	} else {
		fmt.Println("\n4. No pending mandates found to trigger activation charge")
	}

	// Example 5: List revoked mandate authorizations
	fmt.Println("\n5. Listing revoked mandate authorizations...")
	revokedReq := &directdebit.ListMandateAuthorizationsRequest{
		Status:  directdebit.MandateAuthorizationStatusRevoked,
		PerPage: 10,
	}

	revokedMandates, err := client.DirectDebit.ListMandateAuthorizations(ctx, revokedReq)
	if err != nil {
		log.Printf("Error listing revoked mandate authorizations: %v", err)
		return
	}
	fmt.Printf("Found %d revoked mandate authorizations\n", len(revokedMandates.Data))

	// Example 6: Demonstrate pagination with cursor
	if allMandates.Meta.Next != "" {
		fmt.Println("\n6. Demonstrating pagination with cursor...")
		paginationReq := &directdebit.ListMandateAuthorizationsRequest{
			Cursor:  allMandates.Meta.Next,
			PerPage: 5,
		}

		nextPage, err := client.DirectDebit.ListMandateAuthorizations(ctx, paginationReq)
		if err != nil {
			log.Printf("Error fetching next page: %v", err)
		} else {
			fmt.Printf("Next page contains %d mandate authorizations\n", len(nextPage.Data))
		}
	} else {
		fmt.Println("\n6. No additional pages available for pagination")
	}

	fmt.Println("\n=== Direct Debit API Examples Completed ===")
}
