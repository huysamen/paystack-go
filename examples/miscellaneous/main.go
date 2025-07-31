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

	// Example 1: List all banks
	fmt.Println("=== Listing All Banks ===")
	perPage := 20
	banksReq := &miscellaneous.BankListRequest{
		PerPage: &perPage,
	}

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

	// Example 2: List banks by country
	fmt.Println("\n=== Listing Nigerian Banks ===")
	country := "nigeria"
	nigerianBanksReq := &miscellaneous.BankListRequest{
		Country: &country,
		PerPage: &perPage,
	}

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

	// Example 4: List states for address verification
	fmt.Println("\n=== Listing States (US) ===")
	statesReq := &miscellaneous.StateListRequest{
		Country: "US",
	}

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

	// Example 5: Filter banks for verification
	fmt.Println("\n=== Listing Banks Supporting Verification ===")
	enabledForVerification := true
	verificationBanksReq := &miscellaneous.BankListRequest{
		Country:                &country,
		EnabledForVerification: &enabledForVerification,
		PerPage:                &perPage,
	}

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

	// Example 6: Filter banks for direct payment
	fmt.Println("\n=== Listing Banks Supporting Direct Payment ===")
	payWithBank := true
	directPayBanksReq := &miscellaneous.BankListRequest{
		Country:     &country,
		PayWithBank: &payWithBank,
		PerPage:     &perPage,
	}

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
