package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/huysamen/paystack-go"
	transfer_recipients "github.com/huysamen/paystack-go/api/transfer-recipients"
)

func main() {
	// Create a client with your secret key
	client := paystack.DefaultClient("sk_test_your_secret_key_here")

	fmt.Println("=== Transfer Recipients API Examples ===")
	fmt.Println()

	// Example 1: Create a transfer recipient
	fmt.Println("1. Creating a transfer recipient...")
	createReq := &transfer_recipients.TransferRecipientCreateRequest{
		Type:          transfer_recipients.RecipientTypeNuban,
		Name:          "John Doe",
		AccountNumber: "0123456789",
		BankCode:      "058", // GTBank
		Currency:      &[]string{"NGN"}[0],
		Description:   &[]string{"Company employee"}[0],
		Metadata: map[string]any{
			"employee_id": "EMP001",
			"department":  "Engineering",
		},
	}

	createResp, err := client.TransferRecipients.Create(context.Background(), createReq)
	if err != nil {
		log.Printf("Error creating recipient: %v", err)
		return
	}

	fmt.Printf("✅ Recipient created: %s\n", createResp.Data.RecipientCode)
	fmt.Printf("   Name: %s\n", createResp.Data.Name)
	fmt.Printf("   Bank: %s\n", *createResp.Data.Details.BankName)
	fmt.Printf("   Status: Active=%t\n\n", createResp.Data.Active)

	// Store the recipient code for other operations
	recipientCode := createResp.Data.RecipientCode

	// Example 2: Fetch the created recipient
	fmt.Println("2. Fetching the recipient...")
	fetchResp, err := client.TransferRecipients.Fetch(context.Background(), recipientCode)
	if err != nil {
		log.Printf("Error fetching recipient: %v", err)
		return
	}

	fmt.Printf("✅ Recipient fetched: %s\n", fetchResp.Data.RecipientCode)
	fmt.Printf("   Account: %s (%s)\n", *fetchResp.Data.Details.AccountNumber, *fetchResp.Data.Details.BankName)
	fmt.Printf("   Created: %s\n\n", fetchResp.Data.CreatedAt.Format("2006-01-02 15:04:05"))

	// Example 3: Update the recipient
	fmt.Println("3. Updating recipient name...")
	updateReq := &transfer_recipients.TransferRecipientUpdateRequest{
		Name:  "John Smith Doe",
		Email: &[]string{"john.doe@company.com"}[0],
	}

	updateResp, err := client.TransferRecipients.Update(context.Background(), recipientCode, updateReq)
	if err != nil {
		log.Printf("Error updating recipient: %v", err)
		return
	}

	fmt.Printf("✅ Recipient updated: %s\n", updateResp.Data.Name)
	if updateResp.Data.Email != nil {
		fmt.Printf("   Email: %s\n\n", *updateResp.Data.Email)
	}

	// Example 4: Create multiple recipients with bulk create
	fmt.Println("4. Bulk creating recipients...")
	bulkReq := &transfer_recipients.BulkCreateTransferRecipientRequest{
		Batch: []transfer_recipients.BulkRecipientItem{
			{
				Type:          transfer_recipients.RecipientTypeNuban,
				Name:          "Alice Johnson",
				AccountNumber: "9876543210",
				BankCode:      "044", // Access Bank
				Currency:      &[]string{"NGN"}[0],
				Description:   &[]string{"Marketing team"}[0],
				Metadata: map[string]any{
					"employee_id": "EMP002",
					"department":  "Marketing",
				},
			},
			{
				Type:          transfer_recipients.RecipientTypeNuban,
				Name:          "Bob Wilson",
				AccountNumber: "1122334455",
				BankCode:      "033", // UBA
				Currency:      &[]string{"NGN"}[0],
				Description:   &[]string{"Finance team"}[0],
				Metadata: map[string]any{
					"employee_id": "EMP003",
					"department":  "Finance",
				},
			},
		},
	}

	bulkResp, err := client.TransferRecipients.BulkCreate(context.Background(), bulkReq)
	if err != nil {
		log.Printf("Error bulk creating recipients: %v", err)
		return
	}

	fmt.Printf("✅ Bulk create completed:\n")
	fmt.Printf("   Successful: %d recipients\n", len(bulkResp.Data.Success))
	fmt.Printf("   Failed: %d recipients\n", len(bulkResp.Data.Errors))

	for _, success := range bulkResp.Data.Success {
		fmt.Printf("   + %s (%s)\n", success.Name, success.RecipientCode)
	}

	for _, failure := range bulkResp.Data.Errors {
		fmt.Printf("   - %s: %s\n", failure.Name, failure.Message)
	}
	fmt.Println()

	// Example 5: List all recipients with pagination
	fmt.Println("5. Listing transfer recipients...")
	listReq := &transfer_recipients.TransferRecipientListRequest{
		PerPage: &[]int{10}[0],
		Page:    &[]int{1}[0],
	}

	listResp, err := client.TransferRecipients.List(context.Background(), listReq)
	if err != nil {
		log.Printf("Error listing recipients: %v", err)
		return
	}

	fmt.Printf("✅ Found %d recipients:\n", len(listResp.Data))
	for i, recipient := range listResp.Data {
		if i >= 3 { // Show only first 3 for brevity
			fmt.Printf("   ... and %d more\n", len(listResp.Data)-3)
			break
		}
		fmt.Printf("   %d. %s (%s) - %s\n",
			i+1,
			recipient.Name,
			recipient.RecipientCode,
			recipient.Type.String())
	}
	fmt.Println()

	// Example 6: List recipients with date filtering
	fmt.Println("6. Listing recent recipients (last 7 days)...")
	from := time.Now().AddDate(0, 0, -7) // Last 7 days
	to := time.Now()

	listFilteredReq := &transfer_recipients.TransferRecipientListRequest{
		PerPage: &[]int{5}[0],
		From:    &from,
		To:      &to,
	}

	listFilteredResp, err := client.TransferRecipients.List(context.Background(), listFilteredReq)
	if err != nil {
		log.Printf("Error listing filtered recipients: %v", err)
		return
	}

	fmt.Printf("✅ Recent recipients: %d\n", len(listFilteredResp.Data))
	for _, recipient := range listFilteredResp.Data {
		fmt.Printf("   - %s (%s)\n", recipient.Name, recipient.CreatedAt.Format("Jan 2, 2006"))
	}
	fmt.Println()

	// Example 7: Delete (deactivate) a recipient
	fmt.Printf("7. Deleting recipient: %s...\n", recipientCode)
	deleteResp, err := client.TransferRecipients.Delete(context.Background(), recipientCode)
	if err != nil {
		log.Printf("Error deleting recipient: %v", err)
		return
	}

	fmt.Printf("✅ %s\n", deleteResp.Message)

	// Verify deletion by fetching again
	fmt.Println("   Verifying deletion...")
	verifyResp, err := client.TransferRecipients.Fetch(context.Background(), recipientCode)
	if err != nil {
		log.Printf("Error verifying deletion: %v", err)
		return
	}

	fmt.Printf("   Status after deletion: Active=%t\n", verifyResp.Data.Active)

	fmt.Println("\n=== Transfer Recipients Examples Complete ===")
}
