package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/subaccounts"
)

func main() {
	// Create a client with your secret key
	client := paystack.DefaultClient("sk_test_your_secret_key_here")

	fmt.Println("=== Subaccounts API Examples ===")
	fmt.Println()

	// Example 1: Create a subaccount
	fmt.Println("1. Creating a subaccount...")
	createReq := &subaccounts.SubaccountCreateRequest{
		BusinessName:        "Oasis Marketplace",
		BankCode:            "058", // GTBank
		AccountNumber:       "0123456789",
		PercentageCharge:    15.5, // 15.5% goes to main account
		Description:         &[]string{"Online marketplace for local vendors"}[0],
		PrimaryContactName:  &[]string{"John Doe"}[0],
		PrimaryContactEmail: &[]string{"john.doe@oasis.com"}[0],
		PrimaryContactPhone: &[]string{"+2348123456789"}[0],
		Metadata: map[string]any{
			"business_type":   "marketplace",
			"vendor_category": "electronics",
			"location":        "Lagos, Nigeria",
		},
	}

	createResp, err := client.Subaccounts.Create(context.Background(), createReq)
	if err != nil {
		log.Printf("Error creating subaccount: %v", err)
		return
	}

	fmt.Printf("✅ Subaccount created: %s\n", createResp.Data.SubaccountCode)
	fmt.Printf("   Business: %s\n", createResp.Data.BusinessName)
	fmt.Printf("   Bank: %s\n", createResp.Data.SettlementBank)
	fmt.Printf("   Account: %s\n", createResp.Data.AccountNumber)
	fmt.Printf("   Percentage Charge: %.1f%%\n", createResp.Data.PercentageCharge)
	fmt.Printf("   Active: %t\n\n", createResp.Data.Active)

	// Store the subaccount code for other operations
	subaccountCode := createResp.Data.SubaccountCode

	// Example 2: Fetch the created subaccount
	fmt.Println("2. Fetching the subaccount...")
	fetchResp, err := client.Subaccounts.Fetch(context.Background(), subaccountCode)
	if err != nil {
		log.Printf("Error fetching subaccount: %v", err)
		return
	}

	fmt.Printf("✅ Subaccount fetched: %s\n", fetchResp.Data.SubaccountCode)
	fmt.Printf("   Business: %s\n", fetchResp.Data.BusinessName)
	if fetchResp.Data.Description != nil {
		fmt.Printf("   Description: %s\n", *fetchResp.Data.Description)
	}
	if fetchResp.Data.PrimaryContactEmail != nil {
		fmt.Printf("   Contact: %s\n", *fetchResp.Data.PrimaryContactEmail)
	}
	fmt.Printf("   Verified: %t\n", fetchResp.Data.IsVerified)
	fmt.Printf("   Created: %s\n\n", fetchResp.Data.CreatedAt.Format("2006-01-02 15:04:05"))

	// Example 3: Update the subaccount
	fmt.Println("3. Updating subaccount...")
	updateReq := &subaccounts.SubaccountUpdateRequest{
		BusinessName:        "Oasis Global Marketplace", // Updated name
		Description:         "Enhanced online marketplace for global vendors",
		PercentageCharge:    &[]float64{12.0}[0], // Reduced percentage
		PrimaryContactName:  &[]string{"Jane Smith"}[0],
		PrimaryContactEmail: &[]string{"jane.smith@oasis.com"}[0],
		SettlementSchedule:  &[]subaccounts.SettlementSchedule{subaccounts.SettlementScheduleWeekly}[0],
		Metadata: map[string]any{
			"business_type":   "marketplace",
			"vendor_category": "electronics",
			"location":        "Lagos, Nigeria",
			"updated_at":      time.Now().Format(time.RFC3339),
			"plan":            "premium",
		},
	}

	updateResp, err := client.Subaccounts.Update(context.Background(), subaccountCode, updateReq)
	if err != nil {
		log.Printf("Error updating subaccount: %v", err)
		return
	}

	fmt.Printf("✅ Subaccount updated: %s\n", updateResp.Data.BusinessName)
	if updateResp.Data.Description != nil {
		fmt.Printf("   Description: %s\n", *updateResp.Data.Description)
	}
	fmt.Printf("   New Percentage: %.1f%%\n", updateResp.Data.PercentageCharge)
	if updateResp.Data.SettlementSchedule != nil {
		fmt.Printf("   Settlement Schedule: %s\n", *updateResp.Data.SettlementSchedule)
	}
	fmt.Println()

	// Example 4: List all subaccounts with pagination
	fmt.Println("4. Listing subaccounts...")
	listReq := &subaccounts.SubaccountListRequest{
		PerPage: &[]int{10}[0],
		Page:    &[]int{1}[0],
	}

	listResp, err := client.Subaccounts.List(context.Background(), listReq)
	if err != nil {
		log.Printf("Error listing subaccounts: %v", err)
		return
	}

	fmt.Printf("✅ Found %d subaccounts:\n", len(listResp.Data))
	for i, sub := range listResp.Data {
		if i >= 3 { // Show only first 3 for brevity
			fmt.Printf("   ... and %d more\n", len(listResp.Data)-3)
			break
		}
		status := "✅ Active"
		if !sub.Active {
			status = "❌ Inactive"
		}
		fmt.Printf("   %d. %s (%s) - %.1f%% - %s\n",
			i+1,
			sub.BusinessName,
			sub.SubaccountCode,
			sub.PercentageCharge,
			status)
	}
	fmt.Println()

	// Example 5: List subaccounts with date filtering
	fmt.Println("5. Listing recent subaccounts (last 30 days)...")
	from := time.Now().AddDate(0, 0, -30) // Last 30 days
	to := time.Now()

	listFilteredReq := &subaccounts.SubaccountListRequest{
		PerPage: &[]int{5}[0],
		From:    &from,
		To:      &to,
	}

	listFilteredResp, err := client.Subaccounts.List(context.Background(), listFilteredReq)
	if err != nil {
		log.Printf("Error listing filtered subaccounts: %v", err)
		return
	}

	fmt.Printf("✅ Recent subaccounts: %d\n", len(listFilteredResp.Data))
	for _, sub := range listFilteredResp.Data {
		fmt.Printf("   - %s (%s) - Created: %s\n",
			sub.BusinessName,
			sub.SubaccountCode,
			sub.CreatedAt.Format("Jan 2, 2006"))
	}
	fmt.Println()

	// Example 6: Demonstrate business scenarios
	fmt.Println("6. Business scenario demonstrations...")

	// Scenario A: Marketplace vendor onboarding
	fmt.Println("   Scenario A: Onboarding a new marketplace vendor...")
	vendorReq := &subaccounts.SubaccountCreateRequest{
		BusinessName:        "Tech Gadgets Store",
		BankCode:            "044", // Access Bank
		AccountNumber:       "9876543210",
		PercentageCharge:    20.0, // 20% marketplace fee
		Description:         &[]string{"Electronics and gadgets vendor"}[0],
		PrimaryContactName:  &[]string{"Mike Johnson"}[0],
		PrimaryContactEmail: &[]string{"mike@techgadgets.com"}[0],
		PrimaryContactPhone: &[]string{"+2348098765432"}[0],
		Metadata: map[string]any{
			"vendor_type":     "electronics",
			"onboarding_date": time.Now().Format("2006-01-02"),
			"tier":            "standard",
			"commission_rate": 20.0,
		},
	}

	vendorResp, err := client.Subaccounts.Create(context.Background(), vendorReq)
	if err != nil {
		log.Printf("   Error onboarding vendor: %v", err)
	} else {
		fmt.Printf("   ✅ Vendor onboarded: %s\n", vendorResp.Data.SubaccountCode)
		fmt.Printf("      Commission: %.1f%% to marketplace\n", vendorResp.Data.PercentageCharge)
	}

	// Scenario B: Affiliate partner setup
	fmt.Println("   Scenario B: Setting up affiliate partner...")
	affiliateReq := &subaccounts.SubaccountCreateRequest{
		BusinessName:        "Digital Marketing Agency",
		BankCode:            "033", // UBA
		AccountNumber:       "1122334455",
		PercentageCharge:    8.0, // 8% affiliate commission to main account
		Description:         &[]string{"Digital marketing and lead generation"}[0],
		PrimaryContactName:  &[]string{"Sarah Wilson"}[0],
		PrimaryContactEmail: &[]string{"sarah@digitalagency.com"}[0],
		Metadata: map[string]any{
			"partner_type":    "affiliate",
			"commission_rate": 8.0,
			"territory":       "West Africa",
			"contract_start":  time.Now().Format("2006-01-02"),
		},
	}

	affiliateResp, err := client.Subaccounts.Create(context.Background(), affiliateReq)
	if err != nil {
		log.Printf("   Error setting up affiliate: %v", err)
	} else {
		fmt.Printf("   ✅ Affiliate setup: %s\n", affiliateResp.Data.SubaccountCode)
		fmt.Printf("      Revenue share: %.1f%% to main account\n", affiliateResp.Data.PercentageCharge)
	}

	fmt.Println()

	// Example 7: Error handling demonstration
	fmt.Println("7. Error handling demonstration...")

	// Try to create subaccount with invalid data
	invalidReq := &subaccounts.SubaccountCreateRequest{
		BusinessName:     "", // Invalid: empty business name
		BankCode:         "058",
		AccountNumber:    "0123456789",
		PercentageCharge: 150.0, // Invalid: over 100%
	}

	_, err = client.Subaccounts.Create(context.Background(), invalidReq)
	if err != nil {
		fmt.Printf("✅ Validation error caught: %v\n", err)
	}

	// Try to fetch non-existent subaccount
	_, err = client.Subaccounts.Fetch(context.Background(), "ACCT_nonexistent")
	if err != nil {
		fmt.Printf("✅ Not found error handled: %v\n", err)
	}

	fmt.Println()
	fmt.Println("=== Subaccounts Examples Complete ===")
}
