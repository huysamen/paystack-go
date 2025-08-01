// go:build example
//go:build example
// +build example

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/charges"
)

func main() {
	// Create client with custom configuration
	config := paystack.NewConfig("sk_test_your_secret_key_here").
		WithTimeout(30 * time.Second)

	client := paystack.NewClient(config)
	ctx := context.Background()

	fmt.Println("=== Advanced Charges API Demo ===\n")

	// Different payment channel examples
	examples := []struct {
		name        string
		description string
		createFn    func() *charges.CreateChargeRequest
	}{
		{
			name:        "Bank Account Charge",
			description: "Direct charge from customer's bank account",
			createFn: func() *charges.CreateChargeRequest {
				return &charges.CreateChargeRequest{
					Email:     "bank-customer@example.com",
					Amount:    "25000", // ₦250.00
					Reference: stringPtr("bank-charge-" + generateReference()),
					Bank: &charges.BankDetails{
						Code:          "044", // Access Bank
						AccountNumber: "0123456789",
					},
					Metadata: map[string]any{
						"channel":    "bank",
						"customer":   "premium",
						"product_id": "PROD_001",
					},
				}
			},
		},
		{
			name:        "Authorization Code Charge",
			description: "Charge using saved authorization code",
			createFn: func() *charges.CreateChargeRequest {
				return &charges.CreateChargeRequest{
					Email:             "auth-customer@example.com",
					Amount:            "15000", // ₦150.00
					Reference:         stringPtr("auth-charge-" + generateReference()),
					AuthorizationCode: stringPtr("AUTH_example_code"),
					PIN:               stringPtr("1234"), // For non-reusable auth codes
					Metadata: map[string]any{
						"channel":      "card",
						"subscription": "monthly",
						"auto_charge":  true,
					},
				}
			},
		},
		{
			name:        "USSD Charge",
			description: "USSD-based payment",
			createFn: func() *charges.CreateChargeRequest {
				return &charges.CreateChargeRequest{
					Email:     "ussd-customer@example.com",
					Amount:    "5000", // ₦50.00
					Reference: stringPtr("ussd-charge-" + generateReference()),
					USSD: &charges.USSDDetails{
						Type: "737", // GTBank USSD
					},
					Metadata: map[string]any{
						"channel": "ussd",
						"service": "mobile_recharge",
					},
				}
			},
		},
		{
			name:        "Mobile Money Charge",
			description: "Mobile money payment (Ghana/Kenya)",
			createFn: func() *charges.CreateChargeRequest {
				return &charges.CreateChargeRequest{
					Email:     "momo-customer@example.com",
					Amount:    "10000", // ₦100.00
					Reference: stringPtr("momo-charge-" + generateReference()),
					MobileMoney: &charges.MobileMoneyDetails{
						Phone:    "+233244000000",
						Provider: "mtn",
					},
					Metadata: map[string]any{
						"channel": "mobile_money",
						"country": "ghana",
					},
				}
			},
		},
		{
			name:        "QR Code Charge",
			description: "QR code-based payment",
			createFn: func() *charges.CreateChargeRequest {
				return &charges.CreateChargeRequest{
					Email:     "qr-customer@example.com",
					Amount:    "7500", // ₦75.00
					Reference: stringPtr("qr-charge-" + generateReference()),
					QR: &charges.QRDetails{
						Provider: "visa",
					},
					Metadata: map[string]any{
						"channel":  "qr",
						"location": "pos_terminal",
					},
				}
			},
		},
		{
			name:        "Bank Transfer Charge",
			description: "Pay with Transfer (PwT) channel",
			createFn: func() *charges.CreateChargeRequest {
				expiresAt := time.Now().Add(24 * time.Hour) // Expire in 24 hours
				return &charges.CreateChargeRequest{
					Email:     "transfer-customer@example.com",
					Amount:    "100000", // ₦1,000.00
					Reference: stringPtr("transfer-charge-" + generateReference()),
					BankTransfer: &charges.BankTransferDetails{
						AccountExpiresAt: &expiresAt,
					},
					Metadata: map[string]any{
						"channel":    "bank_transfer",
						"expires_in": "24_hours",
					},
				}
			},
		},
	}

	// Process each charge example
	for i, example := range examples {
		fmt.Printf("%d. %s\n", i+1, example.name)
		fmt.Printf("   %s\n", example.description)

		req := example.createFn()
		charge, err := client.Charges.Create(ctx, req)
		if err != nil {
			fmt.Printf("   ❌ Error: %v\n", err)
			fmt.Println()
			continue
		}

		fmt.Printf("   ✅ Created: %s\n", charge.Data.Reference)
		fmt.Printf("   Status: %s\n", charge.Data.Status)
		fmt.Printf("   Amount: ₦%.2f\n", float64(charge.Data.Amount)/100)
		fmt.Printf("   Channel: %s\n", charge.Data.Channel)

		// Handle specific responses based on status
		if charge.Data.Status == "pending" {
			fmt.Println("   ⏳ Checking pending status...")

			// In real scenario, wait at least 10 seconds before checking
			time.Sleep(2 * time.Second)

			pendingCharge, err := client.Charges.CheckPending(ctx, charge.Data.Reference)
			if err != nil {
				fmt.Printf("   Failed to check status: %v\n", err)
			} else {
				fmt.Printf("   Updated status: %s\n", pendingCharge.Data.Status)
				if pendingCharge.Data.GatewayResponse != "" {
					fmt.Printf("   Gateway response: %s\n", pendingCharge.Data.GatewayResponse)
				}
			}
		}

		fmt.Println()
		time.Sleep(1 * time.Second) // Brief pause between examples
	}

	// Advanced charge scenarios
	fmt.Println("=== Advanced Scenarios ===\n")

	// 1. Charge with subaccount and transaction splitting
	fmt.Println("1. Charge with Revenue Splitting")
	subaccountReq := &charges.CreateChargeRequest{
		Email:      "split-customer@example.com",
		Amount:     "50000", // ₦500.00
		Reference:  stringPtr("split-charge-" + generateReference()),
		Subaccount: stringPtr("ACCT_subaccount_code"),
		SplitCode:  stringPtr("SPL_split_code"),
		Bearer:     stringPtr("subaccount"), // Subaccount pays fees
		Bank: &charges.BankDetails{
			Code:          "058", // GTBank
			AccountNumber: "1234567890",
		},
		TransactionCharge: intPtr(5000), // Override split with ₦50 to main account
		Metadata: map[string]any{
			"merchant_share": 10.0, // 10% to main account
			"vendor_share":   90.0, // 90% to subaccount
			"split_type":     "percentage",
		},
	}

	splitCharge, err := client.Charges.Create(ctx, subaccountReq)
	if err != nil {
		fmt.Printf("   ❌ Split charge error: %v\n", err)
	} else {
		fmt.Printf("   ✅ Split charge created: %s\n", splitCharge.Data.Reference)
		fmt.Printf("   Status: %s\n", splitCharge.Data.Status)
	}

	fmt.Println()

	// 2. Error handling and validation examples
	fmt.Println("2. Error Handling Examples")

	errorExamples := []struct {
		name string
		req  *charges.CreateChargeRequest
	}{
		{
			name: "Missing Email",
			req: &charges.CreateChargeRequest{
				Amount: "10000",
				Bank: &charges.BankDetails{
					Code:          "044",
					AccountNumber: "1234567890",
				},
			},
		},
		{
			name: "Invalid Email",
			req: &charges.CreateChargeRequest{
				Email:  "invalid-email",
				Amount: "10000",
				Bank: &charges.BankDetails{
					Code:          "044",
					AccountNumber: "1234567890",
				},
			},
		},
		{
			name: "No Payment Method",
			req: &charges.CreateChargeRequest{
				Email:  "test@example.com",
				Amount: "10000",
			},
		},
		{
			name: "Invalid QR Provider",
			req: &charges.CreateChargeRequest{
				Email:  "test@example.com",
				Amount: "10000",
				QR: &charges.QRDetails{
					Provider: "invalid-provider",
				},
			},
		},
	}

	for _, example := range errorExamples {
		fmt.Printf("   • %s: ", example.name)
		_, err := client.Charges.Create(ctx, example.req)
		if err != nil {
			fmt.Printf("✅ Correctly caught error: %v\n", err)
		} else {
			fmt.Printf("❌ Should have failed but didn't\n")
		}
	}

	fmt.Println()

	// 3. Best practices demonstration
	fmt.Println("=== Best Practices ===")
	fmt.Println("• Always validate input before creating charges")
	fmt.Println("• Use meaningful references for tracking")
	fmt.Println("• Handle all possible charge statuses")
	fmt.Println("• Wait at least 10 seconds before checking pending charges")
	fmt.Println("• Store charge references for reconciliation")
	fmt.Println("• Use metadata for additional context")
	fmt.Println("• Implement proper error handling")
	fmt.Println("• Consider using webhooks for status updates")

	fmt.Println("\n=== Advanced Charges Demo Complete ===")
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func generateReference() string {
	return fmt.Sprintf("%d", time.Now().UnixNano()%1000000000)
}
