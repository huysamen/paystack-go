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

	// 1. Initiate a single transfer
	fmt.Println("=== Initiating Transfer ===")
	initiateReq := &transfers.TransferInitiateRequest{
		Source:    "balance",
		Amount:    2000000,               // ₦20,000 in kobo
		Recipient: "RCP_gx2wn530m0i3w3m", // Transfer recipient code
		Reason:    stringPtr("Payment for services rendered"),
		Currency:  stringPtr("NGN"),
		Reference: stringPtr("my-unique-transfer-ref-001"),
	}

	initiateResp, err := client.Transfers.Initiate(context.Background(), initiateReq)
	if err != nil {
		log.Fatalf("Failed to initiate transfer: %v", err)
	}

	fmt.Printf("Transfer initiated: %s\n", initiateResp.Data.TransferCode)
	fmt.Printf("Status: %s\n", initiateResp.Data.Status)
	fmt.Printf("Amount: ₦%.2f\n", float64(initiateResp.Data.Amount)/100)
	fmt.Printf("Reason: %s\n", initiateResp.Data.Reason)
	fmt.Printf("Reference: %s\n", initiateResp.Data.Reference)

	// Check if OTP is required
	if initiateResp.Data.Status == "otp" {
		fmt.Println("\n⚠️  OTP Required for this transfer")
		fmt.Printf("Transfer Code: %s\n", initiateResp.Data.TransferCode)
		fmt.Println("You need to:")
		fmt.Println("1. Check your business phone for OTP")
		fmt.Println("2. Use the Finalize endpoint with the OTP")

		// Example finalize request (commented out as it requires actual OTP)
		/*
			finalizeReq := &transfers.TransferFinalizeRequest{
				TransferCode: initiateResp.Data.TransferCode,
				OTP:          "123456", // OTP from phone
			}

			finalizeResp, err := client.Transfers.Finalize(context.Background(), finalizeReq)
			if err != nil {
				log.Fatalf("Failed to finalize transfer: %v", err)
			}

			fmt.Printf("Transfer finalized: %s\n", finalizeResp.Data.Status)
		*/
	} else {
		fmt.Println("✅ Transfer completed without OTP")
	}

	// 2. List transfers
	fmt.Println("\n=== Listing Transfers ===")
	listReq := &transfers.TransferListRequest{
		PerPage: intPtr(10),
		Page:    intPtr(1),
	}

	listResp, err := client.Transfers.List(context.Background(), listReq)
	if err != nil {
		log.Fatalf("Failed to list transfers: %v", err)
	}

	fmt.Printf("Found %d transfers\n", len(listResp.Data.Data))
	for i, transfer := range listResp.Data.Data {
		fmt.Printf("%d. %s - ₦%.2f (%s)\n",
			i+1,
			transfer.TransferCode,
			float64(transfer.Amount)/100,
			transfer.Status)
		fmt.Printf("   To: %s (%s)\n",
			transfer.Recipient.Name,
			transfer.Recipient.RecipientCode)
		fmt.Printf("   Reason: %s\n", transfer.Reason)
	}

	// 3. Fetch specific transfer details
	fmt.Println("\n=== Fetching Transfer Details ===")
	fetchResp, err := client.Transfers.Fetch(context.Background(), initiateResp.Data.TransferCode)
	if err != nil {
		log.Fatalf("Failed to fetch transfer: %v", err)
	}

	fmt.Printf("Transfer: %s\n", fetchResp.Data.TransferCode)
	fmt.Printf("Amount: ₦%.2f %s\n",
		float64(fetchResp.Data.Amount)/100,
		fetchResp.Data.Currency)
	fmt.Printf("Status: %s\n", fetchResp.Data.Status)
	fmt.Printf("Source: %s\n", fetchResp.Data.Source)
	fmt.Printf("Created: %s\n", fetchResp.Data.CreatedAt.Format("2006-01-02 15:04:05"))

	if fetchResp.Data.TransferredAt != nil {
		fmt.Printf("Transferred: %s\n", fetchResp.Data.TransferredAt.Format("2006-01-02 15:04:05"))
	}

	fmt.Printf("\nRecipient Details:\n")
	fmt.Printf("  Name: %s\n", fetchResp.Data.Recipient.Name)
	fmt.Printf("  Code: %s\n", fetchResp.Data.Recipient.RecipientCode)
	fmt.Printf("  Type: %s\n", fetchResp.Data.Recipient.Type)
	fmt.Printf("  Bank: %s (%s)\n",
		fetchResp.Data.Recipient.Details.BankName,
		fetchResp.Data.Recipient.Details.BankCode)
	fmt.Printf("  Account: %s (%s)\n",
		fetchResp.Data.Recipient.Details.AccountNumber,
		fetchResp.Data.Recipient.Details.AccountName)

	// 4. Verify transfer by reference
	fmt.Println("\n=== Verifying Transfer by Reference ===")
	verifyResp, err := client.Transfers.Verify(context.Background(), fetchResp.Data.Reference)
	if err != nil {
		log.Fatalf("Failed to verify transfer: %v", err)
	}

	fmt.Printf("Verified transfer: %s\n", verifyResp.Data.Reference)
	fmt.Printf("Status: %s\n", verifyResp.Data.Status)
	fmt.Printf("Amount: ₦%.2f\n", float64(verifyResp.Data.Amount)/100)

	// 5. Bulk transfer example (requires OTP to be disabled)
	fmt.Println("\n=== Bulk Transfer Example ===")
	bulkReq := &transfers.BulkTransferRequest{
		Source:   "balance",
		Currency: stringPtr("NGN"),
		Transfers: []transfers.BulkTransferItem{
			{
				Amount:    500000, // ₦5,000
				Reference: "bulk-transfer-001",
				Reason:    "Monthly salary - Employee 1",
				Recipient: "RCP_recipient_code_1",
			},
			{
				Amount:    750000, // ₦7,500
				Reference: "bulk-transfer-002",
				Reason:    "Monthly salary - Employee 2",
				Recipient: "RCP_recipient_code_2",
			},
			{
				Amount:    600000, // ₦6,000
				Reference: "bulk-transfer-003",
				Reason:    "Monthly salary - Employee 3",
				Recipient: "RCP_recipient_code_3",
			},
		},
	}

	fmt.Printf("Preparing bulk transfer for %d recipients\n", len(bulkReq.Transfers))
	totalAmount := 0
	for _, transfer := range bulkReq.Transfers {
		totalAmount += transfer.Amount
		fmt.Printf("- %s: ₦%.2f (%s)\n",
			transfer.Reference,
			float64(transfer.Amount)/100,
			transfer.Reason)
	}
	fmt.Printf("Total: ₦%.2f\n", float64(totalAmount)/100)

	// Note: Bulk transfers require OTP to be disabled in your Paystack settings
	fmt.Println("\n⚠️  Note: Bulk transfers require OTP to be disabled in your Paystack dashboard")
	fmt.Println("Uncomment the code below to execute bulk transfer:")

	/*
		bulkResp, err := client.Transfers.Bulk(context.Background(), bulkReq)
		if err != nil {
			log.Fatalf("Failed to execute bulk transfer: %v", err)
		}

		fmt.Printf("Bulk transfer completed: %d transfers processed\n", len(bulkResp.Data))
		for _, transfer := range bulkResp.Data {
			fmt.Printf("- %s: %s (₦%.2f)\n",
				transfer.Reference,
				transfer.Status,
				float64(transfer.Amount)/100)
		}
	*/

	// 6. List transfers with date filtering
	fmt.Println("\n=== Recent Transfers (Last 7 Days) ===")
	from := time.Now().AddDate(0, 0, -7)
	to := time.Now()

	recentReq := &transfers.TransferListRequest{
		PerPage: intPtr(20),
		From:    &from,
		To:      &to,
	}

	recentResp, err := client.Transfers.List(context.Background(), recentReq)
	if err != nil {
		log.Fatalf("Failed to list recent transfers: %v", err)
	}

	fmt.Printf("Found %d transfers in the last 7 days\n", len(recentResp.Data.Data))

	statusCount := make(map[string]int)
	totalTransferred := 0

	for _, transfer := range recentResp.Data.Data {
		statusCount[transfer.Status]++
		if transfer.Status == "success" {
			totalTransferred += transfer.Amount
		}
	}

	fmt.Printf("\nTransfer Summary:\n")
	for status, count := range statusCount {
		fmt.Printf("- %s: %d transfers\n", status, count)
	}
	fmt.Printf("Total successfully transferred: ₦%.2f\n", float64(totalTransferred)/100)

	fmt.Println("\nTransfers API example completed!")
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
