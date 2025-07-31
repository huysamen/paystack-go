package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/huysamen/paystack-go/api/webhook"
)

func main() {
	fmt.Println("Webhook Event Data Parsing Examples")
	fmt.Println("===================================")

	// Example 1: Charge Success Event
	chargeSuccessJSON := `{
		"event": "charge.success",
		"data": {
			"id": 302961,
			"domain": "live",
			"status": "success",
			"reference": "qTPrJoy9Bx",
			"amount": 10000,
			"message": null,
			"gateway_response": "Approved",
			"paid_at": "2016-09-30T21:10:19.000Z",
			"created_at": "2016-09-30T21:09:56.000Z",
			"channel": "card",
			"currency": "NGN",
			"ip_address": "41.242.49.37",
			"metadata": {},
			"fees": 1500,
			"customer": {
				"id": 84312,
				"customer_code": "CUS_hdhye17yj8qd2tx",
				"email": "test@paystack.com",
				"phone": "+2348123456789"
			}
		}
	}`

	var chargeEvent webhook.Event
	if err := json.Unmarshal([]byte(chargeSuccessJSON), &chargeEvent); err != nil {
		log.Fatal(err)
	}

	// Using the convenience method
	chargeSuccessEvent, err := chargeEvent.AsChargeSuccess()
	if err != nil {
		log.Printf("Error parsing charge success event: %v", err)
	} else {
		fmt.Printf("✅ Charge Success:\n")
		fmt.Printf("   Reference: %s\n", chargeSuccessEvent.Reference)
		fmt.Printf("   Amount: ₦%.2f\n", float64(chargeSuccessEvent.Amount)/100)
		fmt.Printf("   Status: %s\n", chargeSuccessEvent.Status)
		fmt.Printf("   Customer: %s\n", chargeSuccessEvent.Customer.Email)
		fmt.Println()
	}

	// Example 2: Customer Identification Failed Event
	customerIDFailedJSON := `{
		"event": "customeridentification.failed",
		"data": {
			"customer_id": 82796315,
			"customer_code": "CUS_XXXXXXXXXXXXXXX",
			"email": "email@email.com",
			"identification": {
				"country": "NG",
				"type": "bank_account",
				"bvn": "123*****456",
				"account_number": "012****345",
				"bank_code": "999991"
			},
			"reason": "Account number or BVN is incorrect"
		}
	}`

	var customerEvent webhook.Event
	if err := json.Unmarshal([]byte(customerIDFailedJSON), &customerEvent); err != nil {
		log.Fatal(err)
	}

	// Using the convenience method
	customerFailedEvent, err := customerEvent.AsCustomerIdentificationFailed()
	if err != nil {
		log.Printf("Error parsing customer identification event: %v", err)
	} else {
		fmt.Printf("❌ Customer Identification Failed:\n")
		fmt.Printf("   Customer: %s (%s)\n", customerFailedEvent.Email, customerFailedEvent.CustomerCode)
		fmt.Printf("   Reason: %s\n", customerFailedEvent.Reason)
		fmt.Printf("   BVN: %s\n", customerFailedEvent.Identification.BVN)
		fmt.Printf("   Account: %s\n", customerFailedEvent.Identification.AccountNumber)
		fmt.Println()
	}

	// Example 3: Using generic ParseEventData function
	genericChargeEvent, err := webhook.ParseEventData[webhook.ChargeSuccessEvent](&chargeEvent)
	if err != nil {
		log.Printf("Error using generic parser: %v", err)
	} else {
		fmt.Printf("🔧 Using Generic Parser:\n")
		fmt.Printf("   Gateway Response: %s\n", genericChargeEvent.GatewayResponse)
		fmt.Printf("   Channel: %s\n", genericChargeEvent.Channel)
		fmt.Println()
	}

	// Example 4: Event type checking with constants
	fmt.Printf("📝 Event Type Examples:\n")
	events := []webhook.Event{chargeEvent, customerEvent}

	for i, event := range events {
		fmt.Printf("   Event %d: %s\n", i+1, event.Event)

		switch event.Event {
		case webhook.EventChargeSuccess:
			fmt.Printf("     -> This is a successful charge event\n")
		case webhook.EventCustomerIdentificationFailed:
			fmt.Printf("     -> This is a failed customer identification event\n")
		case webhook.EventTransferSuccess:
			fmt.Printf("     -> This is a successful transfer event\n")
		default:
			fmt.Printf("     -> Unknown or unhandled event type\n")
		}
	}
}
