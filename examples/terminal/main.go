package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/terminal"
)

func main() {
	// Initialize client
	client := paystack.DefaultClient("sk_test_your_secret_key")

	fmt.Println("=== Terminal API Examples ===")

	// Example 1: List terminals
	fmt.Println("\n=== Listing Terminals ===")
	listReq := &terminal.TerminalListRequest{
		PerPage: &[]int{10}[0],
	}

	listResp, err := client.Terminal.List(context.Background(), listReq)
	if err != nil {
		log.Printf("Error listing terminals: %v", err)
	} else {
		fmt.Printf("Found %d terminals\n", len(listResp.Data))
		for i, term := range listResp.Data {
			if i >= 3 { // Show only first 3 for demo
				break
			}
			fmt.Printf("Terminal %d: %s (ID: %s) - Status: %s\n",
				i+1, term.Name, term.TerminalID, term.Status)
		}
	}

	// Example 2: Fetch a specific terminal (use first terminal from list if available)
	var terminalID string
	if listResp != nil && len(listResp.Data) > 0 {
		terminalID = listResp.Data[0].TerminalID

		fmt.Println("\n=== Fetching Terminal Details ===")
		fetchResp, err := client.Terminal.Fetch(context.Background(), terminalID)
		if err != nil {
			log.Printf("Error fetching terminal: %v", err)
		} else {
			fmt.Printf("Terminal Name: %s\n", fetchResp.Data.Name)
			fmt.Printf("Serial Number: %s\n", fetchResp.Data.SerialNumber)
			fmt.Printf("Status: %s\n", fetchResp.Data.Status)
			if fetchResp.Data.Address != nil {
				fmt.Printf("Address: %s\n", *fetchResp.Data.Address)
			}
		}

		// Example 3: Check terminal status/presence
		fmt.Println("\n=== Checking Terminal Status ===")
		statusResp, err := client.Terminal.FetchTerminalStatus(context.Background(), terminalID)
		if err != nil {
			log.Printf("Error checking terminal status: %v", err)
		} else {
			fmt.Printf("Terminal online: %v\n", statusResp.Data.Online)
			fmt.Printf("Terminal available: %v\n", statusResp.Data.Available)
		}

		// Example 4: Update terminal details
		fmt.Println("\n=== Updating Terminal ===")
		updateReq := &terminal.TerminalUpdateRequest{
			Name:    &[]string{"Updated Terminal Name"}[0],
			Address: &[]string{"123 Updated Street, Lagos, Nigeria"}[0],
		}

		updateResp, err := client.Terminal.Update(context.Background(), terminalID, updateReq)
		if err != nil {
			log.Printf("Error updating terminal: %v", err)
		} else {
			fmt.Printf("Terminal updated: %s\n", updateResp.Message)
		}

		// Example 5: Send an invoice event to terminal
		fmt.Println("\n=== Sending Invoice Event to Terminal ===")
		eventReq := &terminal.TerminalSendEventRequest{
			Type:   terminal.TerminalEventTypeInvoice,
			Action: terminal.TerminalEventActionProcess,
			Data: terminal.TerminalEventData{
				"id":        123456,        // Invoice ID
				"reference": 4634337895939, // Offline reference
			},
		}

		eventResp, err := client.Terminal.SendEvent(context.Background(), terminalID, eventReq)
		if err != nil {
			log.Printf("Error sending event to terminal: %v", err)
		} else {
			fmt.Printf("Event sent to terminal: %s\n", eventResp.Message)
			fmt.Printf("Event ID: %s\n", eventResp.Data.ID)

			// Example 6: Check event status
			fmt.Println("\n=== Checking Event Status ===")
			eventStatusResp, err := client.Terminal.FetchEventStatus(context.Background(), terminalID, eventResp.Data.ID)
			if err != nil {
				log.Printf("Error fetching event status: %v", err)
			} else {
				fmt.Printf("Event delivered: %v\n", eventStatusResp.Data.Delivered)
			}
		}

		// Example 7: Send a transaction event to terminal
		fmt.Println("\n=== Sending Transaction Event to Terminal ===")
		transactionEventReq := &terminal.TerminalSendEventRequest{
			Type:   terminal.TerminalEventTypeTransaction,
			Action: terminal.TerminalEventActionProcess,
			Data: terminal.TerminalEventData{
				"id": 789012, // Transaction ID
			},
		}

		transactionEventResp, err := client.Terminal.SendEvent(context.Background(), terminalID, transactionEventReq)
		if err != nil {
			log.Printf("Error sending transaction event to terminal: %v", err)
		} else {
			fmt.Printf("Transaction event sent: %s\n", transactionEventResp.Message)
		}
	}

	// Example 8: Commission a device (this would typically fail unless you have an actual device)
	fmt.Println("\n=== Commissioning Device (Demo) ===")
	commissionReq := &terminal.TerminalCommissionRequest{
		SerialNumber: "1111150412230003899", // Example serial number
	}

	commissionResp, err := client.Terminal.CommissionDevice(context.Background(), commissionReq)
	if err != nil {
		log.Printf("Expected error commissioning device: %v", err)
	} else {
		fmt.Printf("Device commissioned: %s\n", commissionResp.Message)
	}

	// Example 9: Decommission a device (this would typically fail unless you have an actual device)
	fmt.Println("\n=== Decommissioning Device (Demo) ===")
	decommissionReq := &terminal.TerminalDecommissionRequest{
		SerialNumber: "1111150412230003899", // Example serial number
	}

	decommissionResp, err := client.Terminal.DecommissionDevice(context.Background(), decommissionReq)
	if err != nil {
		log.Printf("Expected error decommissioning device: %v", err)
	} else {
		fmt.Printf("Device decommissioned: %s\n", decommissionResp.Message)
	}

	fmt.Println("\n=== Terminal API Examples Complete ===")
}
