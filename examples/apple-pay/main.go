package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huysamen/paystack-go"
	applepay "github.com/huysamen/paystack-go/api/apple-pay"
)

func main() {
	// Initialize client with test secret key
	client := paystack.DefaultClient("sk_test_your_secret_key")
	ctx := context.Background()

	fmt.Println("=== Paystack Apple Pay API Examples ===")
	fmt.Println()

	// Example 1: Register a domain for Apple Pay using builder
	fmt.Println("1. Registering domain for Apple Pay integration...")
	registerBuilder := applepay.NewRegisterDomainRequest("example.com")

	registerResp, err := client.ApplePay.RegisterDomain(ctx, registerBuilder)
	if err != nil {
		log.Printf("Error registering domain: %v", err)
		// Continue with other examples
	} else {
		fmt.Printf("✓ Domain registered: %s\n", registerResp.Message)
	}

	// Example 2: Register a subdomain for Apple Pay using builder
	fmt.Println("\n2. Registering subdomain for Apple Pay integration...")
	registerSubdomainBuilder := applepay.NewRegisterDomainRequest("checkout.example.com")

	registerSubdomainResp, err := client.ApplePay.RegisterDomain(ctx, registerSubdomainBuilder)
	if err != nil {
		log.Printf("Error registering subdomain: %v", err)
	} else {
		fmt.Printf("✓ Subdomain registered: %s\n", registerSubdomainResp.Message)
	}

	// Example 3: List all registered domains using builder
	fmt.Println("\n3. Listing all registered Apple Pay domains...")
	listBuilder := applepay.NewListDomainsRequest()

	domainsResp, err := client.ApplePay.ListDomains(ctx, listBuilder)
	if err != nil {
		log.Printf("Error listing domains: %v", err)
		return
	}

	fmt.Printf("Registered domains (%d):\n", len(domainsResp.Data.DomainNames))
	for i, domain := range domainsResp.Data.DomainNames {
		fmt.Printf("  %d. %s\n", i+1, domain)
	}

	// Example 4: List domains with cursor pagination using builder
	fmt.Println("\n4. Listing domains with cursor pagination...")
	paginatedListBuilder := applepay.NewListDomainsRequest().
		UseCursor(true)

	paginatedDomainsResp, err := client.ApplePay.ListDomains(ctx, paginatedListBuilder)
	if err != nil {
		log.Printf("Error listing domains with pagination: %v", err)
	} else {
		fmt.Printf("Domains with pagination: %d found\n", len(paginatedDomainsResp.Data.DomainNames))
		if paginatedDomainsResp.Meta != nil {
			fmt.Printf("Meta information available: %+v\n", paginatedDomainsResp.Meta)
		}
	}

	// Example 5: Unregister a domain using builder (commented out to prevent accidental removal)
	/*
		fmt.Println("\n5. Unregistering domain from Apple Pay...")
		unregisterBuilder := applepay.NewUnregisterDomainRequest("example.com")

		unregisterResp, err := client.ApplePay.UnregisterDomain(ctx, unregisterBuilder)
		if err != nil {
			log.Printf("Error unregistering domain: %v", err)
		} else if unregisterResp.Data.Status {
			fmt.Printf("✓ Domain unregistered: %s\n", unregisterResp.Data.Message)
		} else {
			fmt.Printf("× Domain unregistration failed: %s\n", unregisterResp.Data.Message)
		}
	*/
	fmt.Println("\n5. Domain unregistration example skipped (uncomment to test)")

	// Example 6: Error handling with invalid domain using builder
	fmt.Println("\n6. Testing error handling with invalid domain...")
	invalidRegisterBuilder := applepay.NewRegisterDomainRequest("") // Empty domain should cause validation error

	_, err = client.ApplePay.RegisterDomain(ctx, invalidRegisterBuilder)
	if err != nil {
		fmt.Printf("✓ Validation error caught: %v\n", err)
	} else {
		fmt.Println("Unexpected: No error occurred with empty domain")
	}

	// Example 7: Register multiple domains using builders
	fmt.Println("\n7. Registering multiple domains...")
	domains := []string{
		"pay.example.com",
		"secure.example.com",
		"app.example.com",
	}

	for _, domain := range domains {
		regBuilder := applepay.NewRegisterDomainRequest(domain)

		regResp, err := client.ApplePay.RegisterDomain(ctx, regBuilder)
		if err != nil {
			log.Printf("Error registering %s: %v", domain, err)
		} else {
			fmt.Printf("✓ %s: %s\n", domain, regResp.Message)
		}
	}

	// Example 8: Final domain list using builder with fluent chaining
	fmt.Println("\n8. Final list of registered domains...")
	finalListBuilder := applepay.NewListDomainsRequest().UseCursor(false)

	finalListResp, err := client.ApplePay.ListDomains(ctx, finalListBuilder)
	if err != nil {
		log.Printf("Error getting final domain list: %v", err)
	} else {
		fmt.Printf("Total registered domains: %d\n", len(finalListResp.Data.DomainNames))
		for _, domain := range finalListResp.Data.DomainNames {
			fmt.Printf("  - %s\n", domain)
		}
	}

	fmt.Println("\n=== Apple Pay API Examples Completed ===")
}
