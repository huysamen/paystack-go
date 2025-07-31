package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huysamen/paystack-go"
	virtualterminal "github.com/huysamen/paystack-go/api/virtual-terminal"
	"github.com/huysamen/paystack-go/types"
)

func main() {
	// Initialize client with test secret key
	client := paystack.DefaultClient("sk_test_your_secret_key")
	ctx := context.Background()

	fmt.Println("=== Paystack Virtual Terminal API Examples ===")
	fmt.Println()

	// Example 1: Create Virtual Terminal
	fmt.Println("1. Creating a virtual terminal...")
	createReq := &virtualterminal.CreateVirtualTerminalRequest{
		Name: "My Virtual Terminal",
		Destinations: []virtualterminal.VirtualTerminalDestination{
			{
				Target: "+2347012345678",
				Name:   "Primary Contact",
			},
		},
		Currency: "NGN",
		CustomFields: []virtualterminal.CustomField{
			{
				DisplayName:  "Customer ID",
				VariableName: "customer_id",
			},
		},
		Metadata: &types.Metadata{
			"department": "sales",
			"location":   "lagos",
		},
	}

	terminal, err := client.VirtualTerminal.Create(ctx, createReq)
	if err != nil {
		log.Printf("Error creating virtual terminal: %v", err)
		return
	}
	fmt.Printf("Created virtual terminal: %s (Code: %s)\n", terminal.Name, terminal.Code)

	// Example 2: List Virtual Terminals
	fmt.Println("\n2. Listing virtual terminals...")
	listReq := &virtualterminal.ListVirtualTerminalsRequest{
		Status:  "active",
		PerPage: 10,
	}

	terminals, err := client.VirtualTerminal.List(ctx, listReq)
	if err != nil {
		log.Printf("Error listing virtual terminals: %v", err)
		return
	}
	fmt.Printf("Found %d virtual terminals\n", len(terminals.Data))
	for _, term := range terminals.Data {
		fmt.Printf("- %s (Code: %s, Active: %t)\n", term.Name, term.Code, term.Active)
	}

	// Example 3: Fetch Virtual Terminal
	fmt.Println("\n3. Fetching virtual terminal...")
	terminalCode := terminal.Code
	fetchedTerminal, err := client.VirtualTerminal.Fetch(ctx, terminalCode)
	if err != nil {
		log.Printf("Error fetching virtual terminal: %v", err)
		return
	}
	fmt.Printf("Fetched terminal: %s (ID: %d)\n", fetchedTerminal.Name, fetchedTerminal.ID)
	fmt.Printf("Destinations: %d\n", len(fetchedTerminal.Destinations))

	// Example 4: Update Virtual Terminal
	fmt.Println("\n4. Updating virtual terminal...")
	updateReq := &virtualterminal.UpdateVirtualTerminalRequest{
		Name: "Updated Virtual Terminal",
	}

	updatedTerminal, err := client.VirtualTerminal.Update(ctx, terminalCode, updateReq)
	if err != nil {
		log.Printf("Error updating virtual terminal: %v", err)
		return
	}
	fmt.Printf("Updated terminal name: %s\n", updatedTerminal.Name)

	// Example 5: Assign Destination
	fmt.Println("\n5. Assigning destination to virtual terminal...")
	assignReq := &virtualterminal.AssignDestinationRequest{
		Destinations: []virtualterminal.VirtualTerminalDestination{
			{
				Target: "+2348012345678",
				Name:   "Secondary Contact",
			},
		},
	}

	destinations, err := client.VirtualTerminal.AssignDestination(ctx, terminalCode, assignReq)
	if err != nil {
		log.Printf("Error assigning destination: %v", err)
		return
	}
	fmt.Printf("Assigned %d destinations\n", len(*destinations))

	// Example 6: Add Split Code
	fmt.Println("\n6. Adding split code to virtual terminal...")
	addSplitReq := &virtualterminal.AddSplitCodeRequest{
		SplitCode: "SPL_98WF13Zu8w5", // Replace with actual split code
	}

	splitResult, err := client.VirtualTerminal.AddSplitCode(ctx, terminalCode, addSplitReq)
	if err != nil {
		log.Printf("Error adding split code: %v", err)
		// This might fail if split code doesn't exist, continue with example
	} else {
		fmt.Printf("Split code added successfully: %+v\n", *splitResult)
	}

	// Example 7: Remove Split Code
	fmt.Println("\n7. Removing split code from virtual terminal...")
	removeSplitReq := &virtualterminal.RemoveSplitCodeRequest{
		SplitCode: "SPL_98WF13Zu8w5", // Replace with actual split code
	}

	removeResult, err := client.VirtualTerminal.RemoveSplitCode(ctx, terminalCode, removeSplitReq)
	if err != nil {
		log.Printf("Error removing split code: %v", err)
		// This might fail if split code doesn't exist, continue with example
	} else {
		fmt.Printf("Split code removed: %s\n", removeResult.Message)
	}

	// Example 8: Unassign Destination
	fmt.Println("\n8. Unassigning destination from virtual terminal...")
	unassignReq := &virtualterminal.UnassignDestinationRequest{
		Targets: []string{"+2348012345678"},
	}

	unassignResult, err := client.VirtualTerminal.UnassignDestination(ctx, terminalCode, unassignReq)
	if err != nil {
		log.Printf("Error unassigning destination: %v", err)
		return
	}
	fmt.Printf("Destination unassigned: %s\n", unassignResult.Message)

	// Example 9: Deactivate Virtual Terminal
	fmt.Println("\n9. Deactivating virtual terminal...")
	deactivateResult, err := client.VirtualTerminal.Deactivate(ctx, terminalCode)
	if err != nil {
		log.Printf("Error deactivating virtual terminal: %v", err)
		return
	}
	fmt.Printf("Terminal deactivated: %s\n", deactivateResult.Message)

	fmt.Println("\n=== Virtual Terminal API Examples Completed ===")
}
