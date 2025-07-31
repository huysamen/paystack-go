package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/subaccounts"
)

// Vendor represents a marketplace vendor
type Vendor struct {
	ID             string
	BusinessName   string
	Category       string
	BankCode       string
	AccountNumber  string
	ContactName    string
	ContactEmail   string
	ContactPhone   string
	CommissionRate float64
	Tier           string
	Status         string
}

// Partner represents an affiliate partner
type Partner struct {
	ID            string
	BusinessName  string
	PartnerType   string
	BankCode      string
	AccountNumber string
	ContactName   string
	ContactEmail  string
	RevenueShare  float64
	Territory     string
	ContractStart time.Time
}

// SubaccountManager manages subaccounts for a marketplace platform
type SubaccountManager struct {
	client *paystack.Client
}

func NewSubaccountManager(client *paystack.Client) *SubaccountManager {
	return &SubaccountManager{client: client}
}

// OnboardVendor creates a subaccount for a new marketplace vendor
func (sm *SubaccountManager) OnboardVendor(vendor Vendor) (*subaccounts.Subaccount, error) {
	req := &subaccounts.SubaccountCreateRequest{
		BusinessName:        vendor.BusinessName,
		BankCode:            vendor.BankCode,
		AccountNumber:       vendor.AccountNumber,
		PercentageCharge:    vendor.CommissionRate,
		Description:         &[]string{fmt.Sprintf("%s vendor - %s tier", vendor.Category, vendor.Tier)}[0],
		PrimaryContactName:  &vendor.ContactName,
		PrimaryContactEmail: &vendor.ContactEmail,
		PrimaryContactPhone: &vendor.ContactPhone,
		Metadata: map[string]any{
			"vendor_id":       vendor.ID,
			"category":        vendor.Category,
			"tier":            vendor.Tier,
			"commission_rate": vendor.CommissionRate,
			"onboarding_date": time.Now().Format(time.RFC3339),
			"status":          vendor.Status,
			"managed_by":      "marketplace_platform",
		},
	}

	resp, err := sm.client.Subaccounts.Create(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to onboard vendor %s: %w", vendor.ID, err)
	}

	return &resp.Data, nil
}

// SetupPartner creates a subaccount for an affiliate partner
func (sm *SubaccountManager) SetupPartner(partner Partner) (*subaccounts.Subaccount, error) {
	req := &subaccounts.SubaccountCreateRequest{
		BusinessName:        partner.BusinessName,
		BankCode:            partner.BankCode,
		AccountNumber:       partner.AccountNumber,
		PercentageCharge:    partner.RevenueShare,
		Description:         &[]string{fmt.Sprintf("%s partner - %s", partner.PartnerType, partner.Territory)}[0],
		PrimaryContactName:  &partner.ContactName,
		PrimaryContactEmail: &partner.ContactEmail,
		Metadata: map[string]any{
			"partner_id":     partner.ID,
			"partner_type":   partner.PartnerType,
			"territory":      partner.Territory,
			"revenue_share":  partner.RevenueShare,
			"contract_start": partner.ContractStart.Format(time.RFC3339),
			"managed_by":     "affiliate_program",
		},
	}

	resp, err := sm.client.Subaccounts.Create(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to setup partner %s: %w", partner.ID, err)
	}

	return &resp.Data, nil
}

// UpdateVendorTier updates a vendor's tier and commission rate
func (sm *SubaccountManager) UpdateVendorTier(subaccountCode, newTier string, newCommissionRate float64) error {
	// First fetch current details
	current, err := sm.client.Subaccounts.Fetch(context.Background(), subaccountCode)
	if err != nil {
		return fmt.Errorf("failed to fetch subaccount: %w", err)
	}

	// Update with new tier information
	req := &subaccounts.SubaccountUpdateRequest{
		BusinessName:     current.Data.BusinessName,
		Description:      fmt.Sprintf("Updated to %s tier", newTier),
		PercentageCharge: &newCommissionRate,
		Metadata: map[string]any{
			"vendor_id":       current.Data.Metadata["vendor_id"],
			"category":        current.Data.Metadata["category"],
			"tier":            newTier,
			"commission_rate": newCommissionRate,
			"last_updated":    time.Now().Format(time.RFC3339),
			"updated_by":      "tier_management_system",
		},
	}

	_, err = sm.client.Subaccounts.Update(context.Background(), subaccountCode, req)
	if err != nil {
		return fmt.Errorf("failed to update vendor tier: %w", err)
	}

	return nil
}

// DeactivateSubaccount deactivates a subaccount
func (sm *SubaccountManager) DeactivateSubaccount(subaccountCode, reason string) error {
	// Fetch current details first
	current, err := sm.client.Subaccounts.Fetch(context.Background(), subaccountCode)
	if err != nil {
		return fmt.Errorf("failed to fetch subaccount: %w", err)
	}

	// Update to deactivate
	active := false
	req := &subaccounts.SubaccountUpdateRequest{
		BusinessName: current.Data.BusinessName,
		Description:  fmt.Sprintf("Deactivated: %s", reason),
		Active:       &active,
		Metadata: map[string]any{
			"status":              "deactivated",
			"deactivated_at":      time.Now().Format(time.RFC3339),
			"deactivation_reason": reason,
		},
	}

	_, err = sm.client.Subaccounts.Update(context.Background(), subaccountCode, req)
	if err != nil {
		return fmt.Errorf("failed to deactivate subaccount: %w", err)
	}

	return nil
}

// GenerateMarketplaceReport generates a comprehensive marketplace report
func (sm *SubaccountManager) GenerateMarketplaceReport() error {
	req := &subaccounts.SubaccountListRequest{
		PerPage: &[]int{100}[0],
	}

	resp, err := sm.client.Subaccounts.List(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to list subaccounts for report: %w", err)
	}

	// Analyze data
	vendors := make(map[string][]subaccounts.Subaccount)
	partners := make(map[string][]subaccounts.Subaccount)
	totalRevenue := make(map[string]float64)
	activeCount := 0

	for _, sub := range resp.Data {
		if sub.Active {
			activeCount++
		}

		if sub.Metadata != nil {
			if managedBy, ok := sub.Metadata["managed_by"].(string); ok {
				switch managedBy {
				case "marketplace_platform":
					category := "Unknown"
					if cat, ok := sub.Metadata["category"].(string); ok {
						category = cat
					}
					vendors[category] = append(vendors[category], sub)
					totalRevenue[category] += sub.PercentageCharge
				case "affiliate_program":
					partnerType := "Unknown"
					if pt, ok := sub.Metadata["partner_type"].(string); ok {
						partnerType = pt
					}
					partners[partnerType] = append(partners[partnerType], sub)
				}
			}
		}
	}

	// Generate report
	fmt.Println("\n=== MARKETPLACE REPORT ===")
	fmt.Printf("Generated: %s\n", time.Now().Format("January 2, 2006 15:04:05"))
	fmt.Printf("Total Subaccounts: %d\n", len(resp.Data))
	fmt.Printf("Active Subaccounts: %d\n\n", activeCount)

	// Vendor analysis
	fmt.Println("📊 VENDOR ANALYSIS")
	fmt.Printf("Total Vendor Categories: %d\n", len(vendors))
	for category, vendorList := range vendors {
		avgCommission := 0.0
		if len(vendorList) > 0 {
			avgCommission = totalRevenue[category] / float64(len(vendorList))
		}

		fmt.Printf("\n🏪 %s Category\n", category)
		fmt.Printf("   Vendors: %d\n", len(vendorList))
		fmt.Printf("   Avg Commission: %.1f%%\n", avgCommission)
		fmt.Printf("   Active: %d\n", countActiveSubaccounts(vendorList))

		// Show top vendors by commission
		if len(vendorList) > 0 {
			fmt.Printf("   Top Vendors:\n")
			for i, vendor := range vendorList {
				if i >= 3 { // Show top 3
					break
				}
				status := "✅"
				if !vendor.Active {
					status = "❌"
				}
				fmt.Printf("   %s %s (%.1f%%)\n", status, vendor.BusinessName, vendor.PercentageCharge)
			}
		}
	}

	// Partner analysis
	if len(partners) > 0 {
		fmt.Println("\n🤝 PARTNER ANALYSIS")
		fmt.Printf("Total Partner Types: %d\n", len(partners))
		for partnerType, partnerList := range partners {
			fmt.Printf("\n📈 %s Partners\n", partnerType)
			fmt.Printf("   Partners: %d\n", len(partnerList))
			fmt.Printf("   Active: %d\n", countActiveSubaccounts(partnerList))

			for _, partner := range partnerList {
				status := "✅"
				if !partner.Active {
					status = "❌"
				}
				territory := "Unknown"
				if partner.Metadata != nil {
					if t, ok := partner.Metadata["territory"].(string); ok {
						territory = t
					}
				}
				fmt.Printf("   %s %s (%.1f%% - %s)\n",
					status, partner.BusinessName, partner.PercentageCharge, territory)
			}
		}
	}

	return nil
}

// Helper functions
func countActiveSubaccounts(subaccounts []subaccounts.Subaccount) int {
	count := 0
	for _, sub := range subaccounts {
		if sub.Active {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("=== Advanced Subaccounts Management ===")
	fmt.Println()

	// Create client and manager
	client := paystack.DefaultClient("sk_test_your_secret_key_here")
	manager := NewSubaccountManager(client)

	// Sample vendors data
	vendors := []Vendor{
		{
			ID:             "VND001",
			BusinessName:   "TechWorld Electronics",
			Category:       "Electronics",
			BankCode:       "058", // GTBank
			AccountNumber:  "1234567890",
			ContactName:    "Alice Johnson",
			ContactEmail:   "alice@techworld.com",
			ContactPhone:   "+2348123456789",
			CommissionRate: 18.0,
			Tier:           "Premium",
			Status:         "active",
		},
		{
			ID:             "VND002",
			BusinessName:   "Fashion Hub",
			Category:       "Fashion",
			BankCode:       "044", // Access Bank
			AccountNumber:  "2345678901",
			ContactName:    "Bob Smith",
			ContactEmail:   "bob@fashionhub.com",
			ContactPhone:   "+2348098765432",
			CommissionRate: 22.0,
			Tier:           "Standard",
			Status:         "active",
		},
		{
			ID:             "VND003",
			BusinessName:   "HomeDecor Paradise",
			Category:       "Home & Garden",
			BankCode:       "033", // UBA
			AccountNumber:  "3456789012",
			ContactName:    "Carol Davis",
			ContactEmail:   "carol@homedecor.com",
			ContactPhone:   "+2348087654321",
			CommissionRate: 15.0,
			Tier:           "Premium",
			Status:         "active",
		},
	}

	// Sample partners data
	partners := []Partner{
		{
			ID:            "PTR001",
			BusinessName:  "Digital Marketing Pro",
			PartnerType:   "affiliate",
			BankCode:      "058",
			AccountNumber: "4567890123",
			ContactName:   "David Wilson",
			ContactEmail:  "david@digitalmarketing.com",
			RevenueShare:  8.0,
			Territory:     "West Africa",
			ContractStart: time.Now().AddDate(0, -6, 0), // 6 months ago
		},
		{
			ID:            "PTR002",
			BusinessName:  "SocialMedia Experts",
			PartnerType:   "referral",
			BankCode:      "044",
			AccountNumber: "5678901234",
			ContactName:   "Eva Brown",
			ContactEmail:  "eva@socialmedia.com",
			RevenueShare:  12.0,
			Territory:     "East Africa",
			ContractStart: time.Now().AddDate(0, -3, 0), // 3 months ago
		},
	}

	// Scenario 1: Vendor onboarding campaign
	fmt.Println("🚀 Scenario 1: Vendor onboarding campaign...")
	var vendorSubaccounts []string

	for _, vendor := range vendors {
		subaccount, err := manager.OnboardVendor(vendor)
		if err != nil {
			log.Printf("Failed to onboard vendor %s: %v", vendor.ID, err)
			continue
		}

		vendorSubaccounts = append(vendorSubaccounts, subaccount.SubaccountCode)
		fmt.Printf("✅ Onboarded: %s (%s) - %s Category - %.1f%% commission\n",
			subaccount.BusinessName,
			subaccount.SubaccountCode,
			vendor.Category,
			subaccount.PercentageCharge)
	}
	fmt.Println()

	// Scenario 2: Partner program setup
	fmt.Println("🤝 Scenario 2: Partner program setup...")

	for _, partner := range partners {
		subaccount, err := manager.SetupPartner(partner)
		if err != nil {
			log.Printf("Failed to setup partner %s: %v", partner.ID, err)
			continue
		}

		fmt.Printf("✅ Partner setup: %s (%s) - %s - %.1f%% revenue share\n",
			subaccount.BusinessName,
			subaccount.SubaccountCode,
			partner.PartnerType,
			subaccount.PercentageCharge)
	}
	fmt.Println()

	// Scenario 3: Vendor tier management
	if len(vendorSubaccounts) > 0 {
		fmt.Println("📈 Scenario 3: Vendor tier upgrade...")
		vendorCode := vendorSubaccounts[0]

		err := manager.UpdateVendorTier(vendorCode, "Platinum", 12.0) // Lower commission for higher tier
		if err != nil {
			log.Printf("Failed to upgrade vendor: %v", err)
		} else {
			fmt.Printf("✅ Vendor upgraded to Platinum tier (12%% commission)\n")

			// Verify the update
			updated, err := client.Subaccounts.Fetch(context.Background(), vendorCode)
			if err != nil {
				log.Printf("Error fetching updated vendor: %v", err)
			} else {
				fmt.Printf("   Current commission: %.1f%%\n", updated.Data.PercentageCharge)
				if updated.Data.Metadata != nil {
					if tier, ok := updated.Data.Metadata["tier"].(string); ok {
						fmt.Printf("   Current tier: %s\n", tier)
					}
				}
			}
		}
		fmt.Println()
	}

	// Scenario 4: Compliance and account management
	fmt.Println("⚖️ Scenario 4: Compliance management...")

	// Create a test subaccount for deactivation
	testVendor := Vendor{
		ID:             "VND999",
		BusinessName:   "Test Vendor For Deactivation",
		Category:       "Test",
		BankCode:       "058",
		AccountNumber:  "9999999999",
		ContactName:    "Test User",
		ContactEmail:   "test@example.com",
		ContactPhone:   "+2348000000000",
		CommissionRate: 25.0,
		Tier:           "Standard",
		Status:         "active",
	}

	testSubaccount, err := manager.OnboardVendor(testVendor)
	if err != nil {
		log.Printf("Failed to create test vendor: %v", err)
	} else {
		fmt.Printf("✅ Test vendor created: %s\n", testSubaccount.SubaccountCode)

		// Deactivate for policy violation
		err = manager.DeactivateSubaccount(testSubaccount.SubaccountCode, "Policy violation - fake products")
		if err != nil {
			log.Printf("Failed to deactivate: %v", err)
		} else {
			fmt.Printf("✅ Vendor deactivated for policy violation\n")

			// Verify deactivation
			deactivated, err := client.Subaccounts.Fetch(context.Background(), testSubaccount.SubaccountCode)
			if err != nil {
				log.Printf("Error fetching deactivated vendor: %v", err)
			} else {
				fmt.Printf("   Status: Active=%t\n", deactivated.Data.Active)
				if deactivated.Data.Metadata != nil {
					if reason, ok := deactivated.Data.Metadata["deactivation_reason"].(string); ok {
						fmt.Printf("   Reason: %s\n", reason)
					}
				}
			}
		}
	}
	fmt.Println()

	// Scenario 5: Settlement schedule optimization
	if len(vendorSubaccounts) > 1 {
		fmt.Println("💰 Scenario 5: Settlement schedule optimization...")
		vendorCode := vendorSubaccounts[1]

		// Fetch current vendor
		current, err := client.Subaccounts.Fetch(context.Background(), vendorCode)
		if err != nil {
			log.Printf("Error fetching vendor: %v", err)
		} else {
			// Update to weekly settlement for better cash flow
			req := &subaccounts.SubaccountUpdateRequest{
				BusinessName:       current.Data.BusinessName,
				Description:        "Updated settlement schedule for better cash flow",
				SettlementSchedule: &[]subaccounts.SettlementSchedule{subaccounts.SettlementScheduleWeekly}[0],
				Metadata: map[string]any{
					"settlement_optimized": true,
					"optimization_date":    time.Now().Format(time.RFC3339),
					"optimization_reason":  "improved_cash_flow",
				},
			}

			_, err = client.Subaccounts.Update(context.Background(), vendorCode, req)
			if err != nil {
				log.Printf("Failed to update settlement schedule: %v", err)
			} else {
				fmt.Printf("✅ Settlement schedule optimized to weekly for %s\n", current.Data.BusinessName)
			}
		}
		fmt.Println()
	}

	// Scenario 6: Comprehensive marketplace analytics
	fmt.Println("📊 Scenario 6: Marketplace analytics...")
	err = manager.GenerateMarketplaceReport()
	if err != nil {
		log.Printf("Error generating report: %v", err)
	}

	// Scenario 7: Error handling and validation
	fmt.Println("⚠️ Scenario 7: Error handling demonstration...")

	// Invalid vendor data
	invalidVendor := Vendor{
		ID:             "INVALID",
		BusinessName:   "", // Invalid: empty name
		Category:       "Test",
		BankCode:       "999", // Invalid bank code
		AccountNumber:  "123",
		CommissionRate: 150.0, // Invalid: over 100%
	}

	_, err = manager.OnboardVendor(invalidVendor)
	if err != nil {
		fmt.Printf("✅ Validation error caught: %v\n", err)
	}

	// Try to update non-existent subaccount
	err = manager.UpdateVendorTier("ACCT_nonexistent", "Premium", 15.0)
	if err != nil {
		fmt.Printf("✅ Not found error handled: %v\n", err)
	}

	fmt.Println()
	fmt.Println("=== Advanced Subaccounts Management Complete ===")
}
