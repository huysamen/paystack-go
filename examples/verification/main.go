package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/verification"
)

func main() {
	// Create a client with your secret key
	client := paystack.DefaultClient("sk_test_your_secret_key_here")

	// Example 1: Resolve an account number
	fmt.Println("=== Account Resolution ===")
	resolveReq := &verification.AccountResolveRequest{
		AccountNumber: "0022728151",
		BankCode:      "063", // Diamond Bank
	}

	resolveResp, err := client.Verification.ResolveAccount(context.Background(), resolveReq)
	if err != nil {
		log.Printf("Error resolving account: %v", err)
	} else {
		fmt.Printf("Account Name: %s\n", resolveResp.Data.AccountName)
		fmt.Printf("Account Number: %s\n", resolveResp.Data.AccountNumber)
	}

	// Example 2: Validate an account with KYC information
	fmt.Println("\n=== Account Validation ===")
	validateReq := &verification.AccountValidateRequest{
		AccountName:    "John Doe",
		AccountNumber:  "0123456789",
		AccountType:    "personal",
		BankCode:       "058",
		CountryCode:    "NG",
		DocumentType:   "bvn",
		DocumentNumber: &[]string{"12345678901"}[0],
	}

	validateResp, err := client.Verification.ValidateAccount(context.Background(), validateReq)
	if err != nil {
		log.Printf("Error validating account: %v", err)
	} else {
		fmt.Printf("Verification Status: %v\n", validateResp.Data.Verified)
		fmt.Printf("Message: %s\n", validateResp.Data.VerificationMessage)
	}

	// Example 3: Resolve card BIN information
	fmt.Println("\n=== Card BIN Resolution ===")
	binResp, err := client.Verification.ResolveCardBIN(context.Background(), "539983")
	if err != nil {
		log.Printf("Error resolving card BIN: %v", err)
	} else {
		fmt.Printf("BIN: %s\n", binResp.Data.BIN)
		fmt.Printf("Brand: %s\n", binResp.Data.Brand)
		fmt.Printf("Card Type: %s\n", binResp.Data.CardType)
		fmt.Printf("Bank: %s\n", binResp.Data.Bank)
		fmt.Printf("Country: %s (%s)\n", binResp.Data.CountryName, binResp.Data.CountryCode)
	}
}
