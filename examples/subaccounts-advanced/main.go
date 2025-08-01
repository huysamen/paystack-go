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

// Business scenarios demonstration with builder patterns
func main() {
	// Initialize the Paystack client
	secret := os.Getenv("PAYSTACK_SECRET_KEY")
	if secret == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable is required")
	}

	client := paystack.DefaultClient(secret)
	ctx := context.Background()

	fmt.Println("=== Advanced Subaccounts Examples ===")
	fmt.Println()

	// Scenario 1: E-commerce Marketplace with Multiple Vendors
	fmt.Println("📦 Scenario 1: E-commerce Marketplace Platform")
	fmt.Println("Creating subaccounts for different vendor categories with varying commission rates...")

	// Electronics Vendor - Higher commission due to higher margins
	electronicsReq := subaccounts.NewSubaccountCreateRequest(
		"TechHub Electronics",
		"044", // Access Bank
		"0123456789",
		20.0, // 20% platform commission
	).
		Description("Premium electronics and gadgets marketplace vendor").
		PrimaryContactName("Alex Tech").
		PrimaryContactEmail("alex@techhub.com").
		PrimaryContactPhone("+2348123456789").
		Metadata(map[string]any{
			"vendor_tier":    "premium",
			"category":       "electronics",
			"min_order":      10000,
			"shipping_zones": []string{"Lagos", "Abuja", "Port Harcourt"},
			"verified":       true,
		})

	electronicsResp, err := client.Subaccounts.Create(ctx, electronicsReq)
	if err != nil {
		log.Printf("Error creating electronics vendor: %v", err)
	} else {
		fmt.Printf("✅ Electronics Vendor: %s (%.1f%% commission)\n",
			electronicsResp.Data.BusinessName, electronicsResp.Data.PercentageCharge)
	}

	// Fashion Vendor - Standard commission
	fashionReq := subaccounts.NewSubaccountCreateRequest(
		"StyleCo Fashion",
		"058", // GTBank
		"9876543210",
		15.0, // 15% platform commission
	).
		Description("Trendy fashion and accessories vendor").
		PrimaryContactName("Sarah Style").
		PrimaryContactEmail("sarah@styleco.com").
		PrimaryContactPhone("+2348987654321").
		Metadata(map[string]any{
			"vendor_tier":    "standard",
			"category":       "fashion",
			"seasonal_sales": true,
			"return_policy":  "30_days",
			"size_guide":     true,
		})

	fashionResp, err := client.Subaccounts.Create(ctx, fashionReq)
	if err != nil {
		log.Printf("Error creating fashion vendor: %v", err)
	} else {
		fmt.Printf("✅ Fashion Vendor: %s (%.1f%% commission)\n",
			fashionResp.Data.BusinessName, fashionResp.Data.PercentageCharge)
	}

	// Performance-based commission update example
	if electronicsResp != nil {
		fmt.Printf("\n📈 Performance-Based Commission Update\n")
		fmt.Printf("Updating commission for top-performing vendor: %s\n", electronicsResp.Data.BusinessName)

		performanceUpdateReq := subaccounts.NewSubaccountUpdateRequest().
			PercentageCharge(18.0). // Reduced from 20% to 18% as reward
			Description("Premium electronics vendor - Performance reward tier").
			Metadata(map[string]any{
				"vendor_tier":      "premium_plus",
				"performance_tier": "top_performer",
				"commission_history": []map[string]any{
					{"rate": 20.0, "period": "2024-01-01 to 2024-06-30", "reason": "initial"},
					{"rate": 18.0, "period": "2024-07-01 onwards", "reason": "performance_reward"},
				},
				"metrics": map[string]any{
					"monthly_sales":   450000,
					"customer_rating": 4.8,
					"return_rate":     2.1,
					"delivery_score":  96.5,
				},
			})

		updateResp, err := client.Subaccounts.Update(ctx, electronicsResp.Data.SubaccountCode, performanceUpdateReq)
		if err != nil {
			log.Printf("Error updating performance commission: %v", err)
		} else {
			fmt.Printf("✅ Commission updated: %s now pays %.1f%% (reduced from 20.0%%)\n",
				updateResp.Data.BusinessName, updateResp.Data.PercentageCharge)
		}
	}

	// Analytics and reporting
	fmt.Println("\n📊 Platform Analytics and Reporting")

	// List all subaccounts for analytics
	analyticsReq := subaccounts.NewSubaccountListRequest().
		PerPage(50).
		Page(1)

	analyticsResp, err := client.Subaccounts.List(ctx, analyticsReq)
	if err != nil {
		log.Printf("Error fetching analytics data: %v", err)
	} else {
		fmt.Printf("📈 Platform Analytics:\n")
		fmt.Printf("   Total Subaccounts: %d\n", len(analyticsResp.Data))

		// Calculate average commission
		var totalCommission float64
		activeCount := 0
		for _, sub := range analyticsResp.Data {
			if sub.Active {
				totalCommission += sub.PercentageCharge
				activeCount++
			}
		}

		if activeCount > 0 {
			avgCommission := totalCommission / float64(activeCount)
			fmt.Printf("   Active Subaccounts: %d\n", activeCount)
			fmt.Printf("   Average Commission Rate: %.2f%%\n", avgCommission)
		}
	}

	// Date-based filtering for reporting
	fmt.Println("\n📅 Time-Based Reporting")

	// Get recent subaccounts (last 7 days)
	weekAgo := time.Now().AddDate(0, 0, -7)
	now := time.Now()

	recentReq := subaccounts.NewSubaccountListRequest().
		PerPage(10).
		DateRange(weekAgo, now)

	recentResp, err := client.Subaccounts.List(ctx, recentReq)
	if err != nil {
		log.Printf("Error fetching recent subaccounts: %v", err)
	} else {
		fmt.Printf("📊 Recent Activity (Last 7 Days):\n")
		fmt.Printf("   New Subaccounts: %d\n", len(recentResp.Data))

		for i, sub := range recentResp.Data {
			if i < 3 { // Show first 3
				fmt.Printf("   %d. %s - %.1f%% - %s\n",
					i+1,
					sub.BusinessName,
					sub.PercentageCharge,
					sub.CreatedAt.Format("Jan 2, 2006"),
				)
			}
		}

		if len(recentResp.Data) > 3 {
			fmt.Printf("   ... and %d more\n", len(recentResp.Data)-3)
		}
	}

	fmt.Println("\n🎯 Advanced Subaccounts Key Insights:")
	fmt.Println("   • Different business models require different commission structures")
	fmt.Println("   • Performance-based commission adjustments incentivize quality")
	fmt.Println("   • Rich metadata enables detailed analytics and segmentation")
	fmt.Println("   • Date-based filtering supports comprehensive reporting")
	fmt.Println("   • Builder patterns make complex configurations readable and maintainable")

	fmt.Println("\n🎉 Advanced subaccounts operations completed successfully!")
}
