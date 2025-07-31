package main

import (
	"fmt"

	"github.com/huysamen/paystack-go/api/webhook"
)

func main() {
	fmt.Println("Paystack Webhook Event Constants:")
	fmt.Println()

	// Charge events
	fmt.Println("Charge Events:")
	fmt.Printf("  Success: %s\n", webhook.EventChargeSuccess)
	fmt.Println()

	// Dispute events
	fmt.Println("Dispute Events:")
	fmt.Printf("  Create:  %s\n", webhook.EventChargeDisputeCreate)
	fmt.Printf("  Remind:  %s\n", webhook.EventChargeDisputeRemind)
	fmt.Printf("  Resolve: %s\n", webhook.EventChargeDisputeResolve)
	fmt.Println()

	// Transfer events
	fmt.Println("Transfer Events:")
	fmt.Printf("  Success:  %s\n", webhook.EventTransferSuccess)
	fmt.Printf("  Failed:   %s\n", webhook.EventTransferFailed)
	fmt.Printf("  Reversed: %s\n", webhook.EventTransferReversed)
	fmt.Println()

	// Subscription events
	fmt.Println("Subscription Events:")
	fmt.Printf("  Create:        %s\n", webhook.EventSubscriptionCreate)
	fmt.Printf("  Disable:       %s\n", webhook.EventSubscriptionDisable)
	fmt.Printf("  Not Renew:     %s\n", webhook.EventSubscriptionNotRenew)
	fmt.Printf("  Expiring Cards: %s\n", webhook.EventSubscriptionExpiringCards)
	fmt.Println()

	// Example of using constants in switch statement
	eventType := webhook.EventChargeSuccess

	switch eventType {
	case webhook.EventChargeSuccess:
		fmt.Println("✅ Payment was successful!")
	case webhook.EventChargeDisputeCreate:
		fmt.Println("⚠️ Dispute created!")
	case webhook.EventTransferSuccess:
		fmt.Println("💸 Transfer completed!")
	case webhook.EventRefundProcessed:
		fmt.Println("💰 Refund processed!")
	default:
		fmt.Printf("🔔 Received event: %s\n", eventType)
	}
}
