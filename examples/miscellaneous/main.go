package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/miscellaneous"
)

func main() {
	// Create a client with your secret key
	client := paystack.DefaultClient("sk_test_your_secret_key_here")

	// Example 1: List all banks using builder pattern
	fmt.Println("=== Listing All Banks ===")
	banksReq := miscellaneous.NewBankListRequest().
		PerPage(20)

	banksResp, err := client.Miscellaneous.ListBanks(context.Background(), banksReq)
	if err != nil {
		log.Printf("Error listing banks: %v", err)
	} else {
		fmt.Printf("Found %d banks\n", len(banksResp.Data))
		for i, bank := range banksResp.Data {
			if i >= 5 { // Show only first 5 for demo
				break
			}
			fmt.Printf("Bank: %s, Code: %s, Country: %s, Currency: %s\n",
				bank.Name, bank.Code, bank.Country, bank.Currency)
		}
	}

	// Example 2: List banks by country using builder pattern
	fmt.Println("\n=== Listing Nigerian Banks ===")
	nigerianBanksReq := miscellaneous.NewBankListRequest().
		Country("nigeria").
		PerPage(20)

	nigerianBanksResp, err := client.Miscellaneous.ListBanks(context.Background(), nigerianBanksReq)
	if err != nil {
		log.Printf("Error listing Nigerian banks: %v", err)
	} else {
		fmt.Printf("Found %d Nigerian banks\n", len(nigerianBanksResp.Data))
		for i, bank := range nigerianBanksResp.Data {
			if i >= 5 { // Show only first 5 for demo
				break
			}
			fmt.Printf("Bank: %s, Code: %s\n", bank.Name, bank.Code)
		}
	}

	// Example 3: List countries
	fmt.Println("\n=== Listing Countries ===")
	countriesResp, err := client.Miscellaneous.ListCountries(context.Background())
	if err != nil {
		log.Printf("Error listing countries: %v", err)
	} else {
		fmt.Printf("Found %d countries supported by Paystack\n", len(countriesResp.Data))
		for i, country := range countriesResp.Data {
			if i >= 5 { // Show only first 5 for demo
				break
			}
			fmt.Printf("Country: %s (%s), Currency: %s\n",
				country.Name, country.ISOCode, country.DefaultCurrencyCode)
		}
	}

	// Example 4: List states for address verification using builder pattern
	fmt.Println("\n=== Listing States (US) ===")
	statesReq := miscellaneous.NewStateListRequest("US")

	statesResp, err := client.Miscellaneous.ListStates(context.Background(), statesReq)
	if err != nil {
		log.Printf("Error listing US states: %v", err)
	} else {
		fmt.Printf("Found %d US states\n", len(statesResp.Data))
		for i, state := range statesResp.Data {
			if i >= 5 { // Show only first 5 for demo
				break
			}
			fmt.Printf("State: %s (%s)\n", state.Name, state.Abbreviation)
		}
	}

	// Example 5: Filter banks for verification using builder pattern
	fmt.Println("\n=== Listing Banks Supporting Verification ===")
	verificationBanksReq := miscellaneous.NewBankListRequest().
		Country("nigeria").
		EnabledForVerification(true).
		PerPage(20)

	verificationBanksResp, err := client.Miscellaneous.ListBanks(context.Background(), verificationBanksReq)
	if err != nil {
		log.Printf("Error listing verification banks: %v", err)
	} else {
		fmt.Printf("Found %d banks supporting verification\n", len(verificationBanksResp.Data))
		for i, bank := range verificationBanksResp.Data {
			if i >= 3 { // Show only first 3 for demo
				break
			}
			fmt.Printf("Bank: %s, Code: %s\n", bank.Name, bank.Code)
		}
	}

	// Example 6: Filter banks for direct payment using builder pattern
	fmt.Println("\n=== Listing Banks Supporting Direct Payment ===")
	directPayBanksReq := miscellaneous.NewBankListRequest().
		Country("nigeria").
		PayWithBank(true).
		PerPage(20)

	directPayBanksResp, err := client.Miscellaneous.ListBanks(context.Background(), directPayBanksReq)
	if err != nil {
		log.Printf("Error listing direct payment banks: %v", err)
	} else {
		fmt.Printf("Found %d banks supporting direct payment\n", len(directPayBanksResp.Data))
		for i, bank := range directPayBanksResp.Data {
			if i >= 3 { // Show only first 3 for demo
				break
			}
			fmt.Printf("Bank: %s, Code: %s, Direct Payment: %t\n",
				bank.Name, bank.Code, bank.PayWithBank)
		}
	}
}
