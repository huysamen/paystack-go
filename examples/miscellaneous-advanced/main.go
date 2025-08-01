package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/miscellaneous"
)

func main() {
	// Create client
	client := paystack.DefaultClient("sk_test_your_secret_key_here")

	// Advanced Example 1: Bank Analysis by Country
	fmt.Println("=== Bank Analysis by Country ===")
	analyzeBanksByCountry(client)

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Advanced Example 2: Payment Method Discovery
	fmt.Println("=== Payment Method Discovery ===")
	discoverPaymentMethods(client)

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Advanced Example 3: Multi-Country Setup Assistant
	fmt.Println("=== Multi-Country Setup Assistant ===")
	setupMultiCountrySupport(client)

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Advanced Example 4: Geographic Information System
	fmt.Println("=== Geographic Information System ===")
	buildGeographicInformation(client)
}

func analyzeBanksByCountry(client *paystack.Client) {
	countries := []string{"nigeria", "ghana", "kenya", "south africa"}

	banksByCountry := make(map[string][]miscellaneous.Bank)
	totalBanks := 0

	for _, country := range countries {
		fmt.Printf("\n📍 Analyzing banks in %s...\n", strings.Title(country))

		banksReq := miscellaneous.NewBankListRequest().
			Country(country)

		banksResp, err := client.Miscellaneous.ListBanks(context.Background(), banksReq)
		if err != nil {
			log.Printf("❌ Error fetching banks for %s: %v", country, err)
			continue
		}

		banksByCountry[country] = banksResp.Data
		totalBanks += len(banksResp.Data)

		// Analyze bank capabilities
		var directPayBanks int
		for _, bank := range banksResp.Data {
			if bank.PayWithBank {
				directPayBanks++
			}
		}

		fmt.Printf("  📊 Total banks: %d\n", len(banksResp.Data))
		fmt.Printf("  💳 Direct payment capable: %d (%.1f%%)\n",
			directPayBanks,
			float64(directPayBanks)/float64(len(banksResp.Data))*100)

		// Show top 3 banks
		fmt.Printf("  🏆 Top banks:\n")
		for i, bank := range banksResp.Data {
			if i >= 3 {
				break
			}
			fmt.Printf("    - %s (Code: %s)\n", bank.Name, bank.Code)
		}
	}

	// Summary analysis
	fmt.Printf("\n📈 SUMMARY ANALYSIS:\n")
	fmt.Printf("Total countries analyzed: %d\n", len(countries))
	fmt.Printf("Total banks discovered: %d\n", totalBanks)

	fmt.Printf("\nBanks by country:\n")
	for country, banks := range banksByCountry {
		fmt.Printf("  %s: %d banks\n", strings.Title(country), len(banks))
	}
}

func discoverPaymentMethods(client *paystack.Client) {
	country := "nigeria"

	// Discover different types of payment-capable banks
	paymentTypes := []struct {
		name   string
		filter func() *miscellaneous.BankListRequestBuilder
	}{
		{
			name: "Direct Payment Banks",
			filter: func() *miscellaneous.BankListRequestBuilder {
				return miscellaneous.NewBankListRequest().
					Country(country).
					PayWithBank(true).
					PerPage(50)
			},
		},
		{
			name: "Bank Transfer Capable",
			filter: func() *miscellaneous.BankListRequestBuilder {
				return miscellaneous.NewBankListRequest().
					Country(country).
					PayWithBankTransfer(true).
					PerPage(50)
			},
		},
		{
			name: "Verification Enabled Banks",
			filter: func() *miscellaneous.BankListRequestBuilder {
				return miscellaneous.NewBankListRequest().
					Country(country).
					EnabledForVerification(true).
					PerPage(50)
			},
		},
	}

	paymentMethods := make(map[string][]miscellaneous.Bank)

	for _, paymentType := range paymentTypes {
		fmt.Printf("\n🔍 Discovering %s...\n", paymentType.name)

		req := paymentType.filter()
		resp, err := client.Miscellaneous.ListBanks(context.Background(), req)
		if err != nil {
			log.Printf("❌ Error discovering %s: %v", paymentType.name, err)
			continue
		}

		paymentMethods[paymentType.name] = resp.Data
		fmt.Printf("✅ Found %d banks supporting %s\n", len(resp.Data), strings.ToLower(paymentType.name))

		// Show first few banks
		for i, bank := range resp.Data {
			if i >= 3 {
				break
			}
			fmt.Printf("  • %s (%s)\n", bank.Name, bank.Code)
		}
	}

	// Cross-analysis
	fmt.Printf("\n🔄 CROSS-ANALYSIS:\n")
	if directBanks, ok := paymentMethods["Direct Payment Banks"]; ok {
		if transferBanks, ok := paymentMethods["Bank Transfer Capable"]; ok {
			// Find banks that support both
			commonBanks := findCommonBanks(directBanks, transferBanks)
			fmt.Printf("Banks supporting both direct payment and transfer: %d\n", len(commonBanks))
		}
	}
}

func setupMultiCountrySupport(client *paystack.Client) {
	// Get all supported countries
	fmt.Println("🌍 Fetching supported countries...")

	countriesResp, err := client.Miscellaneous.ListCountries(context.Background())
	if err != nil {
		log.Printf("❌ Error fetching countries: %v", err)
		return
	}

	fmt.Printf("✅ Paystack supports %d countries\n", len(countriesResp.Data))

	// Analyze each country's setup
	countryAnalysis := make(map[string]map[string]any)

	for i, country := range countriesResp.Data {
		if i >= 4 { // Limit to first 4 for demo
			break
		}

		fmt.Printf("\n🏁 Setting up for %s (%s)...\n", country.Name, country.ISOCode)

		analysis := make(map[string]any)
		analysis["currency"] = country.DefaultCurrencyCode
		analysis["iso_code"] = country.ISOCode

		// Get banks for this country if available
		countrySlug := strings.ToLower(country.Name)
		banksReq := miscellaneous.NewBankListRequest().
			Country(countrySlug).
			PerPage(20)

		banksResp, err := client.Miscellaneous.ListBanks(context.Background(), banksReq)
		if err != nil {
			fmt.Printf("  ⚠️  Could not fetch banks (might not be available): %v\n", err)
			analysis["banks_available"] = false
		} else {
			analysis["banks_available"] = true
			analysis["bank_count"] = len(banksResp.Data)

			// Get a sample of bank capabilities
			var capabilities []string
			for _, bank := range banksResp.Data {
				if bank.PayWithBank {
					capabilities = append(capabilities, "direct_payment")
					break
				}
			}
			analysis["capabilities"] = capabilities

			fmt.Printf("  🏦 Banks available: %d\n", len(banksResp.Data))
			fmt.Printf("  💰 Currency: %s\n", country.DefaultCurrencyCode)
		}

		countryAnalysis[country.Name] = analysis
	}

	// Generate setup recommendations
	fmt.Printf("\n📋 SETUP RECOMMENDATIONS:\n")
	for countryName, analysis := range countryAnalysis {
		fmt.Printf("\n%s:\n", countryName)
		fmt.Printf("  • Currency: %s\n", analysis["currency"])
		fmt.Printf("  • ISO Code: %s\n", analysis["iso_code"])

		if banksAvailable, ok := analysis["banks_available"].(bool); ok && banksAvailable {
			fmt.Printf("  • Banks: %d available\n", analysis["bank_count"])
			fmt.Printf("  • Status: ✅ Ready for integration\n")
		} else {
			fmt.Printf("  • Status: ⚠️  Limited banking integration\n")
		}
	}
}

func buildGeographicInformation(client *paystack.Client) {
	// Build address verification capabilities
	fmt.Println("🗺️  Building geographic information system...")

	testCountries := []string{"US", "CA", "NG", "GH"}
	geoData := make(map[string][]miscellaneous.State)

	for _, countryCode := range testCountries {
		fmt.Printf("\n📍 Fetching states for %s...\n", countryCode)

		statesReq := miscellaneous.NewStateListRequest(countryCode)

		statesResp, err := client.Miscellaneous.ListStates(context.Background(), statesReq)
		if err != nil {
			fmt.Printf("  ❌ No state data available for %s: %v\n", countryCode, err)
			continue
		}

		geoData[countryCode] = statesResp.Data
		fmt.Printf("  ✅ Found %d states/provinces\n", len(statesResp.Data))

		// Show sample states
		fmt.Printf("  📋 Sample states:\n")
		for i, state := range statesResp.Data {
			if i >= 5 {
				break
			}
			abbreviation := state.Abbreviation
			if abbreviation == "" {
				abbreviation = "N/A"
			}
			fmt.Printf("    • %s (%s)\n", state.Name, abbreviation)
		}
	}

	// Summary
	fmt.Printf("\n🌐 GEOGRAPHIC INFORMATION SUMMARY:\n")
	totalStates := 0
	for country, states := range geoData {
		fmt.Printf("%s: %d states/provinces\n", country, len(states))
		totalStates += len(states)
	}
	fmt.Printf("\nTotal geographic entries: %d\n", totalStates)
	fmt.Printf("Address verification available for %d countries\n", len(geoData))
}

// Helper function to find common banks between two slices
func findCommonBanks(banks1, banks2 []miscellaneous.Bank) []miscellaneous.Bank {
	bankMap := make(map[string]miscellaneous.Bank)
	for _, bank := range banks1 {
		bankMap[bank.Code] = bank
	}

	var common []miscellaneous.Bank
	for _, bank := range banks2 {
		if _, exists := bankMap[bank.Code]; exists {
			common = append(common, bank)
		}
	}

	return common
}
