package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/subaccounts"
)

func main() {
	// Initialize the Paystack client
	secret := os.Getenv("PAYSTACK_SECRET_KEY")
	if secret == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable is required")
	}

	client := paystack.DefaultClient(secret)
	ctx := context.Background()

	fmt.Println("=== Subaccounts API Examples ===")
	fmt.Println()

	// Example 1: Create a subaccount using the builder pattern
	fmt.Println("1. Creating a subaccount...")
	createReq := subaccounts.NewSubaccountCreateRequest(
		"Oasis Marketplace", // business name
		"058",               // bank code (GTBank)
		"0123456789",        // account number
		15.5,                // percentage charge (15.5% goes to main account)
	).
		Description("Online marketplace for local vendors").
		PrimaryContactName("John Doe").
		PrimaryContactEmail("john.doe@oasis.com").
		PrimaryContactPhone("+2348123456789").
		Metadata(map[string]any{
			"business_type":   "marketplace",
			"vendor_category": "electronics",
			"location":        "Lagos, Nigeria",
		})

	createResp, err := client.Subaccounts.Create(ctx, createReq)
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
	fetchResp, err := client.Subaccounts.Fetch(ctx, subaccountCode)
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

	// Example 3: Update the subaccount using the builder pattern
	fmt.Println("3. Updating subaccount...")
	updateReq := subaccounts.NewSubaccountUpdateRequest().
		BusinessName("Oasis Global Marketplace").
		Description("Enhanced online marketplace for global vendors").
		PercentageCharge(12.0).
		PrimaryContactName("Jane Smith").
		PrimaryContactEmail("jane.smith@oasis.com").
		PrimaryContactPhone("+2348123456790").
		Metadata(map[string]any{
			"business_type":       "global_marketplace",
			"vendor_category":     "electronics,fashion,home",
			"location":            "Lagos, Nigeria",
			"international":       true,
			"supported_countries": []string{"NG", "GH", "KE"},
		})

	updateResp, err := client.Subaccounts.Update(ctx, subaccountCode, updateReq)
	if err != nil {
		log.Printf("Error updating subaccount: %v", err)
		return
	}

	fmt.Printf("✅ Subaccount updated: %s\n", updateResp.Data.SubaccountCode)
	fmt.Printf("   Updated Business: %s\n", updateResp.Data.BusinessName)
	if updateResp.Data.Description != nil {
		fmt.Printf("   Updated Description: %s\n", *updateResp.Data.Description)
	}
	fmt.Printf("   Updated Percentage Charge: %.1f%%\n", updateResp.Data.PercentageCharge)
	if updateResp.Data.PrimaryContactEmail != nil {
		fmt.Printf("   Updated Contact: %s\n", *updateResp.Data.PrimaryContactEmail)
	}
	fmt.Printf("   Updated At: %s\n\n", updateResp.Data.UpdatedAt.Format("2006-01-02 15:04:05"))

	// Example 4: List subaccounts with pagination using the builder pattern
	fmt.Println("4. Listing subaccounts...")
	listReq := subaccounts.NewSubaccountListRequest().
		PerPage(10).
		Page(1)

	listResp, err := client.Subaccounts.List(ctx, listReq)
	if err != nil {
		log.Printf("Error listing subaccounts: %v", err)
		return
	}

	fmt.Printf("✅ Found %d subaccounts (showing first 10)\n", len(listResp.Data))
	for i, subaccount := range listResp.Data {
		fmt.Printf("   %d. %s (%s) - %.1f%%\n",
			i+1,
			subaccount.BusinessName,
			subaccount.SubaccountCode,
			subaccount.PercentageCharge,
		)
	}

	// Example 5: List subaccounts with date range filtering
	fmt.Println("\n5. Listing subaccounts with date range...")
	from := time.Now().AddDate(0, -1, 0) // 1 month ago
	to := time.Now()

	listWithDateReq := subaccounts.NewSubaccountListRequest().
		PerPage(5).
		Page(1).
		DateRange(from, to)

	listWithDateResp, err := client.Subaccounts.List(ctx, listWithDateReq)
	if err != nil {
		log.Printf("Error listing subaccounts with date range: %v", err)
		return
	}

	fmt.Printf("✅ Found %d subaccounts in the last month\n", len(listWithDateResp.Data))
	for i, subaccount := range listWithDateResp.Data {
		fmt.Printf("   %d. %s - Created: %s\n",
			i+1,
			subaccount.BusinessName,
			subaccount.CreatedAt.Format("2006-01-02"),
		)
	}

	fmt.Println("\n🎉 All subaccounts operations completed successfully!")
}
