package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/transactions"
	"github.com/huysamen/paystack-go/api/webhook"
	"github.com/huysamen/paystack-go/types"
)

func main() {
	// Example 1: Basic client usage
	basicExample()

	// Example 2: Advanced configuration
	advancedConfigExample()

	// Example 3: Transaction workflow
	transactionWorkflowExample()

	// Example 4: Export example
	exportExample()

	// Example 5: Webhook handling
	webhookExample()
}

func basicExample() {
	fmt.Println("=== Basic Example ===")

	// Create a client with your secret key
	client := paystack.DefaultClient("sk_test_your_secret_key_here")

	// Initialize a transaction
	req := &transactions.TransactionInitializeRequest{
		Amount:   50000, // Amount in kobo (500.00 NGN)
		Email:    "customer@example.com",
		Currency: types.CurrencyNGN,
	}

	resp, err := client.Transactions.Initialize(context.Background(), req)
	if err != nil {
		log.Printf("Error initializing transaction: %v", err)
		return
	}

	fmt.Printf("Authorization URL: %s\n", resp.Data.AuthorizationURL)
	fmt.Printf("Reference: %s\n", resp.Data.Reference)
}

func advancedConfigExample() {
	fmt.Println("\n=== Advanced Configuration Example ===")

	// Create a custom HTTP client
	httpClient := &http.Client{
		Timeout: 15 * time.Second,
		// Add custom transport, retry logic, etc.
	}

	// Create a custom configuration
	config := paystack.NewConfig("sk_test_your_secret_key_here").
		WithTimeout(30 * time.Second).
		WithEnvironment(paystack.EnvironmentProduction).
		WithHTTPClient(httpClient)

	// Create client with custom config
	client := paystack.NewClient(config)

	// Use the client with the new builder pattern
	builder := transactions.NewTransactionListRequest().PerPage(10)
	resp, err := client.Transactions.List(context.Background(), builder)
	if err != nil {
		log.Printf("Error listing transactions: %v", err)
		return
	}

	fmt.Printf("Found %d transactions\n", len(resp.Data))
}

func transactionWorkflowExample() {
	fmt.Println("\n=== Transaction Workflow Example ===")

	client := paystack.DefaultClient("sk_test_your_secret_key_here")

	// 1. Initialize transaction
	initReq := &transactions.TransactionInitializeRequest{
		Amount:      100000, // 1000.00 NGN
		Email:       "customer@example.com",
		Currency:    types.CurrencyNGN,
		CallbackURL: "https://yourapp.com/callback",
		Metadata: types.Metadata{
			"order_id":    "12345",
			"customer_id": "67890",
		},
	}

	initResp, err := client.Transactions.Initialize(context.Background(), initReq)
	if err != nil {
		log.Printf("Error initializing transaction: %v", err)
		return
	}

	fmt.Printf("1. Transaction initialized\n")
	fmt.Printf("   Reference: %s\n", initResp.Data.Reference)
	fmt.Printf("   Pay here: %s\n", initResp.Data.AuthorizationURL)

	// 2. Simulate payment completion and verify
	// In real usage, this would happen after customer pays
	fmt.Printf("2. Verifying transaction...\n")

	verifyResp, err := client.Transactions.Verify(context.Background(), initResp.Data.Reference)
	if err != nil {
		log.Printf("Error verifying transaction: %v", err)
		return
	}

	fmt.Printf("   Status: %s\n", verifyResp.Data.Status)
	fmt.Printf("   Amount: %d kobo\n", verifyResp.Data.Amount)

	// 3. Fetch transaction details
	if verifyResp.Data.Status == "success" {
		fetchResp, err := client.Transactions.Fetch(context.Background(), uint64(verifyResp.Data.ID))
		if err != nil {
			log.Printf("Error fetching transaction: %v", err)
			return
		}

		fmt.Printf("3. Transaction details:\n")
		fmt.Printf("   Gateway Response: %s\n", fetchResp.Data.GatewayResponse)
		fmt.Printf("   Channel: %s\n", fetchResp.Data.Channel.String())
	}
}

func webhookExample() {
	fmt.Println("\n=== Webhook Example ===")

	// Create webhook validator
	validator := webhook.NewValidator("sk_test_your_secret_key_here")

	// Example webhook handler
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		// Validate the webhook
		event, err := validator.ValidateRequest(r)
		if err != nil {
			log.Printf("Invalid webhook: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Handle different event types
		switch event.Event {
		case webhook.EventChargeSuccess:
			// Parse transaction data using convenience method
			transaction, err := event.AsChargeSuccess()
			if err != nil {
				log.Printf("Error parsing charge success data: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			fmt.Printf("Payment successful: %s - ₦%.2f\n",
				transaction.Reference, float64(transaction.Amount)/100)

		case webhook.EventTransferSuccess:
			// Parse transfer data using convenience method
			transfer, err := event.AsTransferSuccess()
			if err != nil {
				log.Printf("Error parsing transfer data: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			fmt.Printf("Transfer completed: %s - ₦%.2f\n",
				transfer.Reference, float64(transfer.Amount)/100)

		case webhook.EventRefundProcessed:
			// Parse refund data using convenience method
			refund, err := event.AsRefundProcessed()
			if err != nil {
				log.Printf("Error parsing refund data: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			fmt.Printf("Refund processed: ₦%.2f\n", float64(refund.Amount)/100)

		default:
			fmt.Printf("Unhandled event: %s\n", event.Event)
		}

		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("Webhook handler example created")
	fmt.Println("In production, start server with: http.ListenAndServe(\":8080\", nil)")
}

func exportExample() {
	fmt.Println("\n=== Export Example ===")

	client := paystack.DefaultClient("sk_test_your_secret_key_here")

	// Export transactions from last month
	from := time.Now().AddDate(0, -1, 0)
	to := time.Now()
	perPage := 100

	exportReq := transactions.NewTransactionExportRequest().
		DateRange(from, to).
		PerPage(perPage)

	resp, err := client.Transactions.Export(context.Background(), exportReq)
	if err != nil {
		log.Printf("Error exporting transactions: %v", err)
		return
	}

	fmt.Printf("Export file: %s\n", resp.Data.Path)
	fmt.Printf("Expires at: %s\n", resp.Data.ExpiresAt.Time)
}
