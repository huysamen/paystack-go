package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/settlements"
)

func main() {
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable is required")
	}

	// Create client
	client := paystack.DefaultClient(secretKey)

	// Advanced Example 1: Settlement Analytics Dashboard
	fmt.Println("=== Settlement Analytics Dashboard ===")
	settlementAnalytics(client)

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Advanced Example 2: Transaction Reconciliation
	fmt.Println("=== Transaction Reconciliation ===")
	reconcileTransactions(client)

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Advanced Example 3: Settlement Performance Analysis
	fmt.Println("=== Settlement Performance Analysis ===")
	analyzeSettlementPerformance(client)

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Advanced Example 4: Subaccount Revenue Analysis
	fmt.Println("=== Subaccount Revenue Analysis ===")
	analyzeSubaccountRevenue(client)
}

// settlementAnalytics provides a comprehensive view of settlement metrics
func settlementAnalytics(client *paystack.Client) {
	// Get last 30 days of settlements
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	now := time.Now()

	request := settlements.NewSettlementListRequest().
		DateRange(thirtyDaysAgo, now).
		PerPage(100).
		Page(1)

	response, err := client.Settlements.List(context.Background(), request)
	if err != nil {
		log.Printf("Error fetching settlements: %v", err)
		return
	}

	// Calculate metrics
	var totalSettled, totalFees int64
	var successCount, failedCount, processingCount, pendingCount int

	for _, settlement := range response.Data {
		totalSettled += settlement.EffectiveAmount
		totalFees += settlement.TotalFees

		switch settlement.Status {
		case settlements.SettlementStatusSuccess:
			successCount++
		case settlements.SettlementStatusFailed:
			failedCount++
		case settlements.SettlementStatusProcessing:
			processingCount++
		case settlements.SettlementStatusPending:
			pendingCount++
		}
	}

	fmt.Printf("📊 Settlement Summary (Last 30 Days)\n")
	fmt.Printf("   Total Settlements: %d\n", len(response.Data))
	fmt.Printf("   Total Settled Amount: ₦%.2f\n", float64(totalSettled)/100)
	fmt.Printf("   Total Fees: ₦%.2f\n", float64(totalFees)/100)
	fmt.Printf("   Success Rate: %.1f%%\n", float64(successCount)/float64(len(response.Data))*100)
	fmt.Printf("\n📈 Status Breakdown:\n")
	fmt.Printf("   ✅ Successful: %d\n", successCount)
	fmt.Printf("   ⚠️  Failed: %d\n", failedCount)
	fmt.Printf("   🔄 Processing: %d\n", processingCount)
	fmt.Printf("   ⏳ Pending: %d\n", pendingCount)

	// Average settlement amount
	if len(response.Data) > 0 {
		avgAmount := totalSettled / int64(len(response.Data))
		fmt.Printf("   💰 Average Settlement: ₦%.2f\n", float64(avgAmount)/100)
	}
}

// reconcileTransactions helps reconcile transactions within settlements
func reconcileTransactions(client *paystack.Client) {
	// Get recent settlements
	request := settlements.NewSettlementListRequest().
		PerPage(10).
		Page(1)

	settlementsData, err := client.Settlements.List(context.Background(), request)
	if err != nil {
		log.Printf("Error fetching settlements: %v", err)
		return
	}

	fmt.Println("🔍 Reconciling recent settlements...")

	for i, settlement := range settlementsData.Data {
		if i >= 3 { // Limit to first 3 for demo
			break
		}

		settlementID := fmt.Sprintf("%d", settlement.ID)

		// Get transactions for this settlement
		txRequest := settlements.NewSettlementTransactionListRequest().
			PerPage(50).
			Page(1)

		transactions, err := client.Settlements.ListTransactions(context.Background(), settlementID, txRequest)
		if err != nil {
			log.Printf("Error fetching transactions for settlement %s: %v", settlementID, err)
			continue
		}

		// Calculate totals
		var txTotalAmount, txTotalFees int64
		statusCounts := make(map[string]int)

		for _, tx := range transactions.Data {
			txTotalAmount += tx.Amount
			txTotalFees += tx.Fees
			statusCounts[tx.Status]++
		}

		fmt.Printf("\n📋 Settlement ID: %d\n", settlement.ID)
		fmt.Printf("   Settlement Amount: ₦%.2f\n", float64(settlement.EffectiveAmount)/100)
		fmt.Printf("   Settlement Fees: ₦%.2f\n", float64(settlement.TotalFees)/100)
		fmt.Printf("   Transaction Count: %d\n", len(transactions.Data))
		fmt.Printf("   Calculated Tx Total: ₦%.2f\n", float64(txTotalAmount)/100)
		fmt.Printf("   Calculated Tx Fees: ₦%.2f\n", float64(txTotalFees)/100)

		// Status breakdown
		fmt.Printf("   Transaction Status:\n")
		for status, count := range statusCounts {
			fmt.Printf("     %s: %d\n", status, count)
		}

		// Reconciliation check
		if settlement.TotalProcessed == txTotalAmount {
			fmt.Printf("   ✅ Reconciliation: MATCHED\n")
		} else {
			fmt.Printf("   ❌ Reconciliation: MISMATCH (Settlement: ₦%.2f vs Calculated: ₦%.2f)\n",
				float64(settlement.TotalProcessed)/100, float64(txTotalAmount)/100)
		}
	}
}

// analyzeSettlementPerformance analyzes settlement timing and performance
func analyzeSettlementPerformance(client *paystack.Client) {
	// Get settlements from different time periods
	last7Days := time.Now().AddDate(0, 0, -7)
	last30Days := time.Now().AddDate(0, 0, -30)
	now := time.Now()

	// Last 7 days
	weekRequest := settlements.NewSettlementListRequest().
		DateRange(last7Days, now).
		PerPage(100).
		Page(1)

	weekSettlements, err := client.Settlements.List(context.Background(), weekRequest)
	if err != nil {
		log.Printf("Error fetching week settlements: %v", err)
		return
	}

	// Last 30 days
	monthRequest := settlements.NewSettlementListRequest().
		DateRange(last30Days, now).
		PerPage(100).
		Page(1)

	monthSettlements, err := client.Settlements.List(context.Background(), monthRequest)
	if err != nil {
		log.Printf("Error fetching month settlements: %v", err)
		return
	}

	fmt.Println("📊 Settlement Performance Analysis")

	// Calculate weekly metrics
	var weekTotal, weekSuccessful int64
	for _, s := range weekSettlements.Data {
		weekTotal += s.EffectiveAmount
		if s.Status == settlements.SettlementStatusSuccess {
			weekSuccessful += s.EffectiveAmount
		}
	}

	// Calculate monthly metrics
	var monthTotal, monthSuccessful int64
	for _, s := range monthSettlements.Data {
		monthTotal += s.EffectiveAmount
		if s.Status == settlements.SettlementStatusSuccess {
			monthSuccessful += s.EffectiveAmount
		}
	}

	fmt.Printf("\n📈 Volume Comparison:\n")
	fmt.Printf("   Last 7 Days:  %d settlements, ₦%.2f total\n",
		len(weekSettlements.Data), float64(weekTotal)/100)
	fmt.Printf("   Last 30 Days: %d settlements, ₦%.2f total\n",
		len(monthSettlements.Data), float64(monthTotal)/100)

	// Daily average
	if len(monthSettlements.Data) > 0 {
		dailyAvg := monthTotal / 30
		fmt.Printf("   Daily Average: ₦%.2f\n", float64(dailyAvg)/100)
	}

	// Success rate comparison
	weekSuccessRate := float64(0)
	if weekTotal > 0 {
		weekSuccessRate = float64(weekSuccessful) / float64(weekTotal) * 100
	}

	monthSuccessRate := float64(0)
	if monthTotal > 0 {
		monthSuccessRate = float64(monthSuccessful) / float64(monthTotal) * 100
	}

	fmt.Printf("\n🎯 Success Rate Analysis:\n")
	fmt.Printf("   Last 7 Days:  %.1f%%\n", weekSuccessRate)
	fmt.Printf("   Last 30 Days: %.1f%%\n", monthSuccessRate)

	if weekSuccessRate > monthSuccessRate {
		fmt.Printf("   📈 Trend: IMPROVING\n")
	} else if weekSuccessRate < monthSuccessRate {
		fmt.Printf("   📉 Trend: DECLINING\n")
	} else {
		fmt.Printf("   📊 Trend: STABLE\n")
	}
}

// analyzeSubaccountRevenue analyzes revenue across different subaccounts
func analyzeSubaccountRevenue(client *paystack.Client) {
	fmt.Println("💼 Subaccount Revenue Analysis")

	// Get all settlements (limited for demo)
	request := settlements.NewSettlementListRequest().
		PerPage(50).
		Page(1)

	allSettlements, err := client.Settlements.List(context.Background(), request)
	if err != nil {
		log.Printf("Error fetching settlements: %v", err)
		return
	}

	// Group by subaccount (simplified - in real scenario you'd need to identify subaccounts)
	fmt.Printf("\n📊 Settlement Distribution:\n")
	fmt.Printf("   Total Settlements Analyzed: %d\n", len(allSettlements.Data))

	// Calculate settlement frequency by day of week
	dayFrequency := make(map[time.Weekday]int)
	for _, settlement := range allSettlements.Data {
		dayFrequency[settlement.SettlementDate.Weekday()]++
	}

	fmt.Printf("\n📅 Settlement Pattern by Day:\n")
	days := []time.Weekday{
		time.Monday, time.Tuesday, time.Wednesday,
		time.Thursday, time.Friday, time.Saturday, time.Sunday,
	}

	for _, day := range days {
		count := dayFrequency[day]
		fmt.Printf("   %s: %d settlements\n", day.String(), count)
	}

	// Find peak settlement day
	maxDay := time.Monday
	maxCount := 0
	for day, count := range dayFrequency {
		if count > maxCount {
			maxCount = count
			maxDay = day
		}
	}

	fmt.Printf("\n🏆 Peak Settlement Day: %s (%d settlements)\n", maxDay.String(), maxCount)

	// Calculate average processing time (mock calculation)
	fmt.Printf("\n⏱️  Settlement Insights:\n")
	fmt.Printf("   Most settlements occur on %s\n", maxDay.String())
	fmt.Printf("   Consider scheduling important operations around this pattern\n")
}
