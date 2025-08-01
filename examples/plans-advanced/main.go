package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/plans"
	"github.com/huysamen/paystack-go/types"
)

func main() {
	// Initialize client
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable is required")
	}

	client := paystack.DefaultClient(secretKey)
	ctx := context.Background()

	fmt.Println("🚀 Advanced Plans Management - Builder Pattern Edition")
	fmt.Println("======================================================")

	// Scenario 1: SaaS Company with Tiered Pricing
	fmt.Println("\n💼 Scenario 1: SaaS Company Tiered Plans")
	fmt.Println("Creating a complete pricing structure for a SaaS platform...")

	// Free Tier (Free trial with invoice limit)
	freeTier, err := client.Plans.Create(ctx,
		plans.NewCreatePlanRequest("Free Tier", 0, types.IntervalMonthly).
			Description("Free trial with limited features - perfect for testing our platform").
			Currency(types.CurrencyNGN).
			SendInvoices(false). // Don't invoice for free
			SendSMS(false).
			InvoiceLimit(3), // 3 month trial
	)
	if err != nil {
		log.Printf("Failed to create free tier: %v", err)
	} else {
		fmt.Printf("✅ Free Tier: %s (₦%.2f/month)\n", freeTier.PlanCode, float64(freeTier.Amount)/100)
	}

	// Startup Plan
	startupPlan, err := client.Plans.Create(ctx,
		plans.NewCreatePlanRequest("Startup", 1500000, types.IntervalMonthly). // ₦15,000/month
											Description("Perfect for startups and small businesses with growing needs").
											Currency(types.CurrencyNGN).
											SendInvoices(true).
											SendSMS(true).
											InvoiceLimit(36), // 3 years
	)
	if err != nil {
		log.Printf("Failed to create startup plan: %v", err)
	} else {
		fmt.Printf("✅ Startup Plan: %s (₦%.2f/month)\n", startupPlan.PlanCode, float64(startupPlan.Amount)/100)
	}

	// Business Plan
	businessPlan, err := client.Plans.Create(ctx,
		plans.NewCreatePlanRequest("Business", 4500000, types.IntervalMonthly). // ₦45,000/month
											Description("Advanced features for growing businesses with team collaboration").
											Currency(types.CurrencyNGN).
											SendInvoices(true).
											SendSMS(true).
											InvoiceLimit(60), // 5 years
	)
	if err != nil {
		log.Printf("Failed to create business plan: %v", err)
	} else {
		fmt.Printf("✅ Business Plan: %s (₦%.2f/month)\n", businessPlan.PlanCode, float64(businessPlan.Amount)/100)
	}

	// Enterprise Plan
	enterprisePlan, err := client.Plans.Create(ctx,
		plans.NewCreatePlanRequest("Enterprise", 15000000, types.IntervalMonthly). // ₦150,000/month
												Description("Enterprise-grade solution with dedicated support and custom integrations").
												Currency(types.CurrencyNGN).
												SendInvoices(true).
												SendSMS(true).
												InvoiceLimit(120), // 10 years
	)
	if err != nil {
		log.Printf("Failed to create enterprise plan: %v", err)
	} else {
		fmt.Printf("✅ Enterprise Plan: %s (₦%.2f/month)\n", enterprisePlan.PlanCode, float64(enterprisePlan.Amount)/100)
	}

	// Scenario 2: Educational Platform with Multiple Intervals
	fmt.Println("\n🎓 Scenario 2: Educational Platform Plans")
	fmt.Println("Creating flexible pricing for different learning schedules...")

	// Weekly Course Access
	weeklyAccess, err := client.Plans.Create(ctx,
		plans.NewCreatePlanRequest("Weekly Course Access", 250000, types.IntervalWeekly). // ₦2,500/week
													Description("Week-by-week access to premium courses and materials").
													Currency(types.CurrencyNGN).
													SendInvoices(true).
													SendSMS(true).
													InvoiceLimit(52), // 1 year max
	)
	if err != nil {
		log.Printf("Failed to create weekly access: %v", err)
	} else {
		fmt.Printf("📚 Weekly Access: %s (₦%.2f/week)\n", weeklyAccess.PlanCode, float64(weeklyAccess.Amount)/100)
	}

	// Monthly Learning Plan
	monthlyLearning, err := client.Plans.Create(ctx,
		plans.NewCreatePlanRequest("Monthly Learning", 800000, types.IntervalMonthly). // ₦8,000/month
												Description("Full monthly access to all courses, workshops, and live sessions").
												Currency(types.CurrencyNGN).
												SendInvoices(true).
												SendSMS(true).
												InvoiceLimit(24), // 2 years
	)
	if err != nil {
		log.Printf("Failed to create monthly learning: %v", err)
	} else {
		fmt.Printf("📖 Monthly Learning: %s (₦%.2f/month)\n", monthlyLearning.PlanCode, float64(monthlyLearning.Amount)/100)
	}

	// Annual Student Discount
	annualStudent, err := client.Plans.Create(ctx,
		plans.NewCreatePlanRequest("Annual Student", 7200000, types.IntervalAnnually). // ₦72,000/year (save ₦24,000)
												Description("Annual student plan with significant savings and exclusive perks").
												Currency(types.CurrencyNGN).
												SendInvoices(true).
												SendSMS(true).
												InvoiceLimit(4), // 4 years max
	)
	if err != nil {
		log.Printf("Failed to create annual student: %v", err)
	} else {
		fmt.Printf("🎓 Annual Student: %s (₦%.2f/year - Save ₦24,000!)\n", annualStudent.PlanCode, float64(annualStudent.Amount)/100)
	}

	// Scenario 3: Subscription Box Service
	fmt.Println("\n📦 Scenario 3: Subscription Box Service")
	fmt.Println("Creating plans for a monthly subscription box service...")

	// Basic Box
	basicBox, err := client.Plans.Create(ctx,
		plans.NewCreatePlanRequest("Basic Box", 1200000, types.IntervalMonthly). // ₦12,000/month
												Description("Curated selection of 3-5 premium items delivered monthly").
												Currency(types.CurrencyNGN).
												SendInvoices(true).
												SendSMS(true).
												InvoiceLimit(12), // 1 year commitment
	)
	if err != nil {
		log.Printf("Failed to create basic box: %v", err)
	} else {
		fmt.Printf("📦 Basic Box: %s (₦%.2f/month)\n", basicBox.PlanCode, float64(basicBox.Amount)/100)
	}

	// Premium Box
	premiumBox, err := client.Plans.Create(ctx,
		plans.NewCreatePlanRequest("Premium Box", 2500000, types.IntervalMonthly). // ₦25,000/month
												Description("Deluxe selection of 6-8 premium items plus exclusive limited editions").
												Currency(types.CurrencyNGN).
												SendInvoices(true).
												SendSMS(true).
												InvoiceLimit(24), // 2 years
	)
	if err != nil {
		log.Printf("Failed to create premium box: %v", err)
	} else {
		fmt.Printf("🎁 Premium Box: %s (₦%.2f/month)\n", premiumBox.PlanCode, float64(premiumBox.Amount)/100)
	}

	// Quarterly Special
	quarterlySpecial, err := client.Plans.Create(ctx,
		plans.NewCreatePlanRequest("Quarterly Special", 10000000, types.IntervalQuarterly). // ₦100,000/quarter
													Description("Seasonal mega-box with exclusive items and special collaborations").
													Currency(types.CurrencyNGN).
													SendInvoices(true).
													SendSMS(true).
													InvoiceLimit(20), // 5 years
	)
	if err != nil {
		log.Printf("Failed to create quarterly special: %v", err)
	} else {
		fmt.Printf("🌟 Quarterly Special: %s (₦%.2f/quarter)\n", quarterlySpecial.PlanCode, float64(quarterlySpecial.Amount)/100)
	}

	// Advanced Filtering and Analytics
	fmt.Println("\n📊 Advanced Plan Analytics")
	fmt.Println("===========================")

	// List all monthly plans with different amounts
	fmt.Println("\n💰 Monthly Plans by Price Range:")

	monthlyPlans, err := client.Plans.List(ctx,
		plans.NewListPlansRequest().
			Interval(types.IntervalMonthly).
			PerPage(50),
	)
	if err != nil {
		log.Printf("Failed to list monthly plans: %v", err)
	} else {
		// Group by price ranges
		var budget, standard, premium, enterprise []types.Plan

		for _, plan := range monthlyPlans.Data {
			switch {
			case plan.Amount == 0:
				// Free plans - don't categorize by price
			case plan.Amount <= 1000000: // Up to ₦10,000
				budget = append(budget, plan)
			case plan.Amount <= 5000000: // ₦10,001 - ₦50,000
				standard = append(standard, plan)
			case plan.Amount <= 10000000: // ₦50,001 - ₦100,000
				premium = append(premium, plan)
			default: // Above ₦100,000
				enterprise = append(enterprise, plan)
			}
		}

		fmt.Printf("   🏠 Budget (≤₦10,000): %d plans\n", len(budget))
		for _, plan := range budget {
			fmt.Printf("      • %s: ₦%.2f\n", plan.Name, float64(plan.Amount)/100)
		}

		fmt.Printf("   🏢 Standard (₦10,001-₦50,000): %d plans\n", len(standard))
		for _, plan := range standard {
			fmt.Printf("      • %s: ₦%.2f\n", plan.Name, float64(plan.Amount)/100)
		}

		fmt.Printf("   💎 Premium (₦50,001-₦100,000): %d plans\n", len(premium))
		for _, plan := range premium {
			fmt.Printf("      • %s: ₦%.2f\n", plan.Name, float64(plan.Amount)/100)
		}

		fmt.Printf("   🏰 Enterprise (>₦100,000): %d plans\n", len(enterprise))
		for _, plan := range enterprise {
			fmt.Printf("      • %s: ₦%.2f\n", plan.Name, float64(plan.Amount)/100)
		}
	}

	// List plans by different intervals
	intervals := []types.Interval{types.IntervalWeekly, types.IntervalMonthly, types.IntervalQuarterly, types.IntervalAnnually}
	intervalNames := []string{"Weekly", "Monthly", "Quarterly", "Annual"}

	fmt.Println("\n⏰ Plans by Billing Interval:")
	for i, interval := range intervals {
		plans, err := client.Plans.List(ctx,
			plans.NewListPlansRequest().
				Interval(interval).
				PerPage(20),
		)
		if err != nil {
			log.Printf("Failed to list %s plans: %v", intervalNames[i], err)
			continue
		}

		fmt.Printf("   %s Plans: %d\n", intervalNames[i], len(plans.Data))
		for j, plan := range plans.Data {
			if j >= 3 { // Show only first 3
				fmt.Printf("      ... and %d more\n", len(plans.Data)-3)
				break
			}
			fmt.Printf("      %d. %s - ₦%.2f\n", j+1, plan.Name, float64(plan.Amount)/100)
		}
	}

	// Plan Management Operations
	fmt.Println("\n🔧 Plan Management Operations")
	fmt.Println("=============================")

	// Update a plan (if we created any successfully)
	var planToUpdate string
	if startupPlan != nil {
		planToUpdate = startupPlan.PlanCode
	} else if businessPlan != nil {
		planToUpdate = businessPlan.PlanCode
	}

	if planToUpdate != "" {
		fmt.Printf("\n✏️ Updating Plan: %s\n", planToUpdate)

		updateResp, err := client.Plans.Update(ctx, planToUpdate,
			plans.NewUpdatePlanRequest("Startup Pro", 1800000, types.IntervalMonthly). // Increased price
													Description("Enhanced startup plan with additional features and priority support").
													Currency(types.CurrencyNGN).
													SendInvoices(true).
													SendSMS(true).
													InvoiceLimit(48).                   // Increased to 4 years
													UpdateExistingSubscriptions(false), // Keep existing subscribers on old pricing
		)
		if err != nil {
			log.Printf("Failed to update plan: %v", err)
		} else {
			fmt.Printf("✅ Plan updated: %s\n", updateResp.Message)

			// Fetch updated plan to show changes
			updatedPlan, err := client.Plans.Fetch(ctx, planToUpdate)
			if err == nil {
				fmt.Printf("   New Name: %s\n", updatedPlan.Name)
				fmt.Printf("   New Price: ₦%.2f\n", float64(updatedPlan.Amount)/100)
				fmt.Printf("   New Invoice Limit: %d\n", updatedPlan.InvoiceLimit)
			}
		}
	}

	// Detailed plan analysis
	fmt.Println("\n📈 Detailed Plan Analysis")
	if businessPlan != nil {
		detailedPlan, err := client.Plans.Fetch(ctx, businessPlan.PlanCode)
		if err != nil {
			log.Printf("Failed to fetch detailed plan: %v", err)
		} else {
			fmt.Printf("📊 Business Plan Analytics:\n")
			fmt.Printf("   Plan Code: %s\n", detailedPlan.PlanCode)
			fmt.Printf("   Current Subscribers: %d\n", detailedPlan.SubscribersCount)
			fmt.Printf("   Active Subscriptions: %d\n", detailedPlan.ActiveSubscriptionsCount)
			fmt.Printf("   Total Revenue: ₦%.2f\n", float64(detailedPlan.TotalRevenue)/100)
			fmt.Printf("   Revenue per Subscriber: ₦%.2f\n",
				func() float64 {
					if detailedPlan.SubscribersCount > 0 {
						return float64(detailedPlan.TotalRevenue) / float64(detailedPlan.SubscribersCount) / 100
					}
					return 0
				}())
			fmt.Printf("   Created: %s\n", detailedPlan.CreatedAt.Format("2006-01-02 15:04:05"))
			fmt.Printf("   Last Updated: %s\n", detailedPlan.UpdatedAt.Format("2006-01-02 15:04:05"))
		}
	}

	// Summary and Builder Pattern Benefits
	fmt.Println("\n📋 Advanced Plans Management Summary")
	fmt.Println("====================================")

	createdPlans := 0
	if freeTier != nil {
		createdPlans++
	}
	if startupPlan != nil {
		createdPlans++
	}
	if businessPlan != nil {
		createdPlans++
	}
	if enterprisePlan != nil {
		createdPlans++
	}
	if weeklyAccess != nil {
		createdPlans++
	}
	if monthlyLearning != nil {
		createdPlans++
	}
	if annualStudent != nil {
		createdPlans++
	}
	if basicBox != nil {
		createdPlans++
	}
	if premiumBox != nil {
		createdPlans++
	}
	if quarterlySpecial != nil {
		createdPlans++
	}

	fmt.Printf("✅ Successfully demonstrated:\n")
	fmt.Printf("   • Created %d different subscription plans\n", createdPlans)
	fmt.Printf("   • SaaS tiered pricing (Free → Startup → Business → Enterprise)\n")
	fmt.Printf("   • Educational platform with multiple intervals\n")
	fmt.Printf("   • Subscription box service pricing\n")
	fmt.Printf("   • Advanced filtering and analytics\n")
	fmt.Printf("   • Plan updates and management\n")

	fmt.Printf("\n🚀 Builder Pattern Advanced Features:\n")
	fmt.Printf("   • Fluent method chaining for complex configurations\n")
	fmt.Printf("   • Type-safe interval and currency handling\n")
	fmt.Printf("   • Required parameters enforced at build time\n")
	fmt.Printf("   • Optional parameters with sensible defaults\n")
	fmt.Printf("   • Easy plan updates without breaking existing code\n")
	fmt.Printf("   • Advanced filtering with multiple criteria\n")

	// Export sample plans as JSON
	fmt.Println("\n📄 Sample Plans (JSON):")
	if enterprisePlan != nil {
		fmt.Println("Enterprise Plan:")
		planJSON, _ := json.MarshalIndent(enterprisePlan, "", "  ")
		fmt.Println(string(planJSON))
	}

	fmt.Println("\n🎯 Advanced plans management completed successfully!")
}
