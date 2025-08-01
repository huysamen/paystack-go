package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/transfers"
)

func main() {
	// Initialize client
	client := paystack.DefaultClient("sk_test_your_secret_key")

	fmt.Println("=== Advanced Transfer Management Workflows ===")

	// Scenario 1: Process payroll for a small company
	fmt.Println("\n🏢 Scenario 1: Monthly Payroll Processing")

	employees := []struct {
		Name          string
		RecipientCode string
		Salary        int // in kobo
		Department    string
	}{
		{"Alice Johnson", "RCP_alice_001", 15000000, "Engineering"}, // ₦150,000
		{"Bob Smith", "RCP_bob_002", 12000000, "Marketing"},         // ₦120,000
		{"Carol Davis", "RCP_carol_003", 13500000, "Sales"},         // ₦135,000
		{"David Wilson", "RCP_david_004", 11000000, "Support"},      // ₦110,000
		{"Eve Brown", "RCP_eve_005", 16000000, "Product"},           // ₦160,000
	}

	fmt.Printf("Processing payroll for %d employees...\n", len(employees))

	var payrollTransfers []transfers.BulkTransferItem
	totalPayroll := 0

	for i, emp := range employees {
		reference := fmt.Sprintf("payroll-2024-01-%03d", i+1)
		reason := fmt.Sprintf("January 2024 Salary - %s (%s)", emp.Name, emp.Department)

		payrollTransfers = append(payrollTransfers, transfers.BulkTransferItem{
			Amount:    emp.Salary,
			Reference: reference,
			Reason:    reason,
			Recipient: emp.RecipientCode,
		})

		totalPayroll += emp.Salary
		fmt.Printf("✓ %s (%s): ₦%.2f\n",
			emp.Name,
			emp.Department,
			float64(emp.Salary)/100)
	}

	fmt.Printf("\nTotal Payroll: ₦%.2f\n", float64(totalPayroll)/100)

	// Create bulk transfer request (for demonstration - commented out as it requires OTP disabled)
	fmt.Println("\n📤 Preparing bulk payroll transfer...")
	fmt.Printf("Would transfer to %d employees\n", len(payrollTransfers))
	fmt.Println("Note: Actual bulk transfer requires OTP to be disabled in Paystack settings")

	// payrollReq := &transfers.BulkTransferRequest{
	//	Source:   "balance",
	//	Currency: stringPtr("NGN"),
	//	Transfers: payrollTransfers,
	// }

	// Scenario 2: Individual high-value transfer with OTP
	fmt.Println("\n💰 Scenario 2: High-Value Transfer (with OTP)")

	highValueReq := &transfers.TransferInitiateRequest{
		Source:    "balance",
		Amount:    50000000, // ₦500,000
		Recipient: "RCP_vendor_001",
		Reason:    stringPtr("Q1 2024 Vendor Payment - Office Equipment"),
		Currency:  stringPtr("NGN"),
		Reference: stringPtr("vendor-payment-q1-2024-001"),
	}

	highValueResp, err := client.Transfers.Initiate(context.Background(), highValueReq)
	if err != nil {
		log.Fatalf("Failed to initiate high-value transfer: %v", err)
	}

	fmt.Printf("Transfer Code: %s\n", highValueResp.Data.TransferCode)
	fmt.Printf("Amount: ₦%.2f\n", float64(highValueResp.Data.Amount)/100)
	fmt.Printf("Status: %s\n", highValueResp.Data.Status)

	if highValueResp.Data.Status == "otp" {
		fmt.Println("\n🔐 OTP Required - Security Check")
		fmt.Println("1. Check your registered business phone")
		fmt.Println("2. Enter OTP to finalize transfer")
		fmt.Printf("3. Use transfer code: %s\n", highValueResp.Data.TransferCode)

		// Simulate OTP finalization workflow
		fmt.Println("\nSimulating OTP finalization...")

		// In real scenario, you'd get OTP from user input
		// finalizeReq := &transfers.TransferFinalizeRequest{
		//     TransferCode: highValueResp.Data.TransferCode,
		//     OTP:          "123456", // From phone
		// }

		fmt.Println("⏳ Waiting for OTP entry and finalization...")
	}

	// Scenario 3: Transfer monitoring and reconciliation
	fmt.Println("\n📊 Scenario 3: Transfer Monitoring & Reconciliation")

	// Get transfers from last 24 hours
	yesterday := time.Now().AddDate(0, 0, -1)
	today := time.Now()

	monitorReq := transfers.NewTransferListRequest().
		PerPage(100).
		DateRange(yesterday, today)

	monitorResp, err := client.Transfers.List(context.Background(), monitorReq)
	if err != nil {
		log.Fatalf("Failed to get transfer monitoring data: %v", err)
	}

	fmt.Printf("Analyzing %d transfers from last 24 hours...\n", len(monitorResp.Data.Data))

	// Analyze transfer patterns
	statusStats := make(map[string]int)
	currencyStats := make(map[string]int)
	totalByStatus := make(map[string]int)
	hourlyVolume := make(map[int]int)

	for _, transfer := range monitorResp.Data.Data {
		// Status analysis
		statusStats[transfer.Status]++
		totalByStatus[transfer.Status] += transfer.Amount

		// Currency analysis
		currencyStats[transfer.Currency.String()]++

		// Hourly volume analysis
		hour := transfer.CreatedAt.Hour()
		hourlyVolume[hour] += transfer.Amount
	}

	fmt.Println("\n📈 Transfer Analytics:")
	fmt.Println("\nBy Status:")
	for status, count := range statusStats {
		amount := totalByStatus[status]
		fmt.Printf("  %s: %d transfers (₦%.2f)\n",
			status, count, float64(amount)/100)
	}

	fmt.Println("\nBy Currency:")
	for currency, count := range currencyStats {
		fmt.Printf("  %s: %d transfers\n", currency, count)
	}

	fmt.Println("\nPeak Hours (by volume):")
	type hourStat struct {
		Hour   int
		Volume int
	}

	var hours []hourStat
	for hour, volume := range hourlyVolume {
		hours = append(hours, hourStat{Hour: hour, Volume: volume})
	}

	// Simple sorting for top 3 hours
	for i := 0; i < len(hours)-1; i++ {
		for j := i + 1; j < len(hours); j++ {
			if hours[i].Volume < hours[j].Volume {
				hours[i], hours[j] = hours[j], hours[i]
			}
		}
	}

	for i := 0; i < 3 && i < len(hours); i++ {
		fmt.Printf("  %02d:00 - ₦%.2f\n",
			hours[i].Hour,
			float64(hours[i].Volume)/100)
	}

	// Scenario 4: Failed transfer investigation
	fmt.Println("\n🔍 Scenario 4: Failed Transfer Investigation")

	// Find any failed transfers
	var failedTransfers []transfers.Transfer
	for _, transfer := range monitorResp.Data.Data {
		if transfer.Status == "failed" || transfer.Status == "reversed" {
			failedTransfers = append(failedTransfers, transfer)
		}
	}

	if len(failedTransfers) > 0 {
		fmt.Printf("Found %d failed/reversed transfers:\n", len(failedTransfers))

		for _, transfer := range failedTransfers {
			fmt.Printf("\n❌ Transfer: %s\n", transfer.TransferCode)
			fmt.Printf("   Amount: ₦%.2f\n", float64(transfer.Amount)/100)
			fmt.Printf("   Status: %s\n", transfer.Status)
			fmt.Printf("   Reason: %s\n", transfer.Reason)
			fmt.Printf("   Recipient: %s\n", transfer.Recipient.Name)
			fmt.Printf("   Created: %s\n", transfer.CreatedAt.Format("2006-01-02 15:04:05"))

			if transfer.Failures != nil {
				fmt.Printf("   Failure Info: %v\n", transfer.Failures)
			}

			// Fetch detailed information
			details, err := client.Transfers.Fetch(context.Background(), transfer.TransferCode)
			if err == nil {
				fmt.Printf("   Reference: %s\n", details.Data.Reference)
				fmt.Printf("   Bank: %s (%s)\n",
					details.Data.Recipient.Details.BankName,
					details.Data.Recipient.Details.BankCode)
			}
		}

		fmt.Println("\n💡 Recommended Actions:")
		fmt.Println("1. Verify recipient bank details")
		fmt.Println("2. Check account balance sufficiency")
		fmt.Println("3. Retry transfer if issue resolved")
		fmt.Println("4. Contact recipient for account verification")
	} else {
		fmt.Println("✅ No failed transfers in the last 24 hours")
	}

	// Scenario 5: Transfer verification workflow
	fmt.Println("\n✅ Scenario 5: Transfer Verification Workflow")

	if len(monitorResp.Data.Data) > 0 {
		// Verify the most recent transfer
		recentTransfer := monitorResp.Data.Data[0]

		fmt.Printf("Verifying recent transfer: %s\n", recentTransfer.Reference)

		verifyResp, err := client.Transfers.Verify(context.Background(), recentTransfer.Reference)
		if err != nil {
			fmt.Printf("❌ Verification failed: %v\n", err)
		} else {
			fmt.Printf("✅ Verification successful\n")
			fmt.Printf("   Status: %s\n", verifyResp.Data.Status)
			fmt.Printf("   Amount: ₦%.2f\n", float64(verifyResp.Data.Amount)/100)
			fmt.Printf("   Created: %s\n", verifyResp.Data.CreatedAt.Format("2006-01-02 15:04:05"))

			if verifyResp.Data.TransferredAt != nil {
				fmt.Printf("   Completed: %s\n", verifyResp.Data.TransferredAt.Format("2006-01-02 15:04:05"))
				duration := verifyResp.Data.TransferredAt.Sub(verifyResp.Data.CreatedAt)
				fmt.Printf("   Processing Time: %v\n", duration.Round(time.Second))
			}
		}
	}

	fmt.Println("\n🎯 Summary & Best Practices:")
	fmt.Println("1. Use bulk transfers for batch payments (payroll, dividends)")
	fmt.Println("2. Enable OTP for high-value transfers (security)")
	fmt.Println("3. Monitor transfer status regularly")
	fmt.Println("4. Investigate failed transfers promptly")
	fmt.Println("5. Verify critical transfers using reference")
	fmt.Println("6. Keep detailed logs for reconciliation")

	fmt.Println("\nAdvanced transfers example completed!")
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
