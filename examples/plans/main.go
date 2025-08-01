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

	fmt.Println("📋 Plans Management Example - Builder Pattern")
	fmt.Println("=============================================")

	// Example 1: Create a basic subscription plan using builder pattern
	fmt.Println("\n🎯 Example 1: Creating Basic Subscription Plan")

	basicPlan, err := client.Plans.Create(ctx,
		plans.NewCreatePlanRequest("Basic Plan", 500000, types.IntervalMonthly). // ₦5,000.00 monthly
												Description("Basic subscription plan with essential features").
												Currency(types.CurrencyNGN).
												SendInvoices(true).
												SendSMS(false).
												InvoiceLimit(12), // 12 months maximum
	)
	if err != nil {
		log.Printf("Failed to create basic plan: %v", err)
		// Continue with other examples
	} else {
		fmt.Printf("✅ Basic Plan Created:\n")
		fmt.Printf("   Plan Code: %s\n", basicPlan.PlanCode)
		fmt.Printf("   Name: %s\n", basicPlan.Name)
		fmt.Printf("   Amount: ₦%.2f %s\n", float64(basicPlan.Amount)/100, basicPlan.Currency)
		fmt.Printf("   Interval: %s\n", basicPlan.Interval)
		fmt.Printf("   Description: %s\n", basicPlan.Description)
	}

	// Example 2: Create a premium plan with different settings
	fmt.Println("\n⚡ Example 2: Creating Premium Plan")

	premiumPlan, err := client.Plans.Create(ctx,
		plans.NewCreatePlanRequest("Premium Plan", 1500000, types.IntervalMonthly). // ₦15,000.00 monthly
												Description("Premium subscription with advanced features and priority support").
												Currency(types.CurrencyNGN).
												SendInvoices(true).
												SendSMS(true).    // Enable SMS notifications for premium
												InvoiceLimit(24), // 24 months maximum
	)
	if err != nil {
		log.Printf("Failed to create premium plan: %v", err)
	} else {
		fmt.Printf("✅ Premium Plan Created:\n")
		fmt.Printf("   Plan Code: %s\n", premiumPlan.PlanCode)
		fmt.Printf("   Name: %s\n", premiumPlan.Name)
		fmt.Printf("   Amount: ₦%.2f %s\n", float64(premiumPlan.Amount)/100, premiumPlan.Currency)
		fmt.Printf("   SMS Enabled: %t\n", premiumPlan.SendSms)
	}

	// Example 3: Create an annual plan
	fmt.Println("\n📅 Example 3: Creating Annual Plan")

	annualPlan, err := client.Plans.Create(ctx,
		plans.NewCreatePlanRequest("Annual Pro", 10000000, types.IntervalAnnually). // ₦100,000.00 annually
												Description("Annual professional plan with significant savings").
												Currency(types.CurrencyNGN).
												SendInvoices(true).
												SendSMS(true).
												InvoiceLimit(5), // 5 years maximum
	)
	if err != nil {
		log.Printf("Failed to create annual plan: %v", err)
	} else {
		fmt.Printf("✅ Annual Plan Created:\n")
		fmt.Printf("   Plan Code: %s\n", annualPlan.PlanCode)
		fmt.Printf("   Name: %s\n", annualPlan.Name)
		fmt.Printf("   Amount: ₦%.2f %s\n", float64(annualPlan.Amount)/100, annualPlan.Currency)
		fmt.Printf("   Interval: %s\n", annualPlan.Interval)
	}

	// Store plan codes for later operations
	var planCodes []string
	if basicPlan != nil {
		planCodes = append(planCodes, basicPlan.PlanCode)
	}
	if premiumPlan != nil {
		planCodes = append(planCodes, premiumPlan.PlanCode)
	}
	if annualPlan != nil {
		planCodes = append(planCodes, annualPlan.PlanCode)
	}

	// Example 4: List plans with filtering using builder pattern
	fmt.Println("\n📊 Example 4: Listing Plans with Filters")

	// List all monthly plans
	monthlyPlansResp, err := client.Plans.List(ctx,
		plans.NewListPlansRequest().
			Interval(types.IntervalMonthly).
			PerPage(10).
			Page(1),
	)
	if err != nil {
		log.Printf("Failed to list monthly plans: %v", err)
	} else {
		fmt.Printf("📋 Monthly Plans (%d found):\n", len(monthlyPlansResp.Data))
		for i, plan := range monthlyPlansResp.Data {
			if i >= 5 { // Show only first 5
				fmt.Printf("   ... and %d more\n", len(monthlyPlansResp.Data)-5)
				break
			}
			fmt.Printf("   %d. %s - ₦%.2f %s (%s)\n",
				i+1, plan.Name, float64(plan.Amount)/100, plan.Currency, plan.Interval)
		}
	}

	// Example 5: List plans by status
	fmt.Println("\n✅ Example 5: Listing Active Plans")

	activePlansResp, err := client.Plans.List(ctx,
		plans.NewListPlansRequest().
			Status("active").
			PerPage(20),
	)
	if err != nil {
		log.Printf("Failed to list active plans: %v", err)
	} else {
		fmt.Printf("📋 Active Plans (%d found):\n", len(activePlansResp.Data))
		for i, plan := range activePlansResp.Data {
			if i >= 3 { // Show only first 3
				fmt.Printf("   ... and %d more\n", len(activePlansResp.Data)-3)
				break
			}
			fmt.Printf("   %d. %s - ₦%.2f (%d subscribers)\n",
				i+1, plan.Name, float64(plan.Amount)/100, plan.SubscribersCount)
		}
	}

	// Example 6: Fetch specific plan details
	if len(planCodes) > 0 {
		fmt.Println("\n🔍 Example 6: Fetching Plan Details")

		planDetails, err := client.Plans.Fetch(ctx, planCodes[0])
		if err != nil {
			log.Printf("Failed to fetch plan details: %v", err)
		} else {
			fmt.Printf("📄 Plan Details:\n")
			fmt.Printf("   Name: %s\n", planDetails.Name)
			fmt.Printf("   Plan Code: %s\n", planDetails.PlanCode)
			fmt.Printf("   Amount: ₦%.2f %s\n", float64(planDetails.Amount)/100, planDetails.Currency)
			fmt.Printf("   Interval: %s\n", planDetails.Interval)
			fmt.Printf("   Description: %s\n", planDetails.Description)
			fmt.Printf("   Send Invoices: %t\n", planDetails.SendInvoices)
			fmt.Printf("   Send SMS: %t\n", planDetails.SendSms)
			fmt.Printf("   Invoice Limit: %d\n", planDetails.InvoiceLimit)
			fmt.Printf("   Subscribers: %d\n", planDetails.SubscribersCount)
			fmt.Printf("   Active Subscriptions: %d\n", planDetails.ActiveSubscriptionsCount)
			fmt.Printf("   Total Revenue: ₦%.2f\n", float64(planDetails.TotalRevenue)/100)
		}
	}

	// Example 7: Update a plan using builder pattern
	if len(planCodes) > 0 {
		fmt.Println("\n✏️ Example 7: Updating Plan")

		updateResp, err := client.Plans.Update(ctx, planCodes[0],
			plans.NewUpdatePlanRequest("Basic Plan (Updated)", 600000, types.IntervalMonthly). // Increased to ₦6,000.00
														Description("Updated basic plan with enhanced features and new pricing").
														Currency(types.CurrencyNGN).
														SendInvoices(true).
														SendSMS(true).                      // Now enable SMS
														InvoiceLimit(18).                   // Increased limit
														UpdateExistingSubscriptions(false), // Don't update existing subscribers
		)
		if err != nil {
			log.Printf("Failed to update plan: %v", err)
		} else {
			fmt.Printf("✅ Plan Updated Successfully:\n")
			fmt.Printf("   Status: %s\n", updateResp.Message)

			// Fetch the updated plan to show changes
			updatedPlan, err := client.Plans.Fetch(ctx, planCodes[0])
			if err == nil {
				fmt.Printf("   New Amount: ₦%.2f\n", float64(updatedPlan.Amount)/100)
				fmt.Printf("   SMS Now Enabled: %t\n", updatedPlan.SendSms)
			}
		}
	}

	// Example 8: Create plans for different use cases
	fmt.Println("\n🎨 Example 8: Creating Specialized Plans")

	// Starter plan
	starterPlan, err := client.Plans.Create(ctx,
		plans.NewCreatePlanRequest("Starter", 200000, types.IntervalMonthly). // ₦2,000.00
											Description("Perfect for individuals and small teams just getting started").
											Currency(types.CurrencyNGN).
											SendInvoices(true).
											InvoiceLimit(6), // 6 months trial period
	)
	if err != nil {
		log.Printf("Failed to create starter plan: %v", err)
	} else {
		fmt.Printf("🚀 Starter Plan: %s (₦%.2f)\n", starterPlan.PlanCode, float64(starterPlan.Amount)/100)
	}

	// Enterprise plan
	enterprisePlan, err := client.Plans.Create(ctx,
		plans.NewCreatePlanRequest("Enterprise", 5000000, types.IntervalMonthly). // ₦50,000.00
												Description("Enterprise-grade solution with unlimited access and dedicated support").
												Currency(types.CurrencyNGN).
												SendInvoices(true).
												SendSMS(true).
												InvoiceLimit(60), // 5 years
	)
	if err != nil {
		log.Printf("Failed to create enterprise plan: %v", err)
	} else {
		fmt.Printf("🏢 Enterprise Plan: %s (₦%.2f)\n", enterprisePlan.PlanCode, float64(enterprisePlan.Amount)/100)
	}

	// Summary
	fmt.Println("\n📋 Plans Management Summary")
	fmt.Println("===========================")
	fmt.Printf("✅ Demonstrated builder pattern usage for:\n")
	fmt.Printf("   • Creating plans with different intervals and features\n")
	fmt.Printf("   • Listing plans with advanced filtering\n")
	fmt.Printf("   • Fetching detailed plan information\n")
	fmt.Printf("   • Updating existing plans\n")
	fmt.Printf("   • Creating specialized plans for different use cases\n")

	fmt.Printf("\n💡 Builder Pattern Benefits:\n")
	fmt.Printf("   • Fluent method chaining for intuitive API usage\n")
	fmt.Printf("   • Type-safe request construction\n")
	fmt.Printf("   • Required parameters enforced at creation time\n")
	fmt.Printf("   • Optional parameters easily added with method calls\n")
	fmt.Printf("   • Clear separation between required and optional fields\n")

	// Export one plan as JSON for reference
	if basicPlan != nil {
		fmt.Println("\n📄 Sample Plan JSON:")
		planJSON, _ := json.MarshalIndent(basicPlan, "", "  ")
		fmt.Println(string(planJSON))
	}

	fmt.Println("\n🎉 Plans management examples completed!")
}
