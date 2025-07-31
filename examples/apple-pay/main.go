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

	// Example 1: Register a domain for Apple Pay
	fmt.Println("1. Registering domain for Apple Pay integration...")
	registerReq := &applepay.RegisterDomainRequest{
		DomainName: "example.com",
	}

	registerResp, err := client.ApplePay.RegisterDomain(ctx, registerReq)
	if err != nil {
		log.Printf("Error registering domain: %v", err)
		// Continue with other examples
	} else {
		fmt.Printf("✓ Domain registered: %s\n", registerResp.Message)
	}

	// Example 2: Register a subdomain for Apple Pay
	fmt.Println("\n2. Registering subdomain for Apple Pay integration...")
	registerSubdomainReq := &applepay.RegisterDomainRequest{
		DomainName: "checkout.example.com",
	}

	registerSubdomainResp, err := client.ApplePay.RegisterDomain(ctx, registerSubdomainReq)
	if err != nil {
		log.Printf("Error registering subdomain: %v", err)
	} else {
		fmt.Printf("✓ Subdomain registered: %s\n", registerSubdomainResp.Message)
	}

	// Example 3: List all registered domains
	fmt.Println("\n3. Listing all registered Apple Pay domains...")
	listReq := &applepay.ListDomainsRequest{}

	domainsResp, err := client.ApplePay.ListDomains(ctx, listReq)
	if err != nil {
		log.Printf("Error listing domains: %v", err)
		return
	}

	fmt.Printf("Registered domains (%d):\n", len(domainsResp.Data.DomainNames))
	for i, domain := range domainsResp.Data.DomainNames {
		fmt.Printf("  %d. %s\n", i+1, domain)
	}

	// Example 4: List domains with cursor pagination
	fmt.Println("\n4. Listing domains with cursor pagination...")
	useCursor := true
	paginatedListReq := &applepay.ListDomainsRequest{
		UseCursor: &useCursor,
	}

	paginatedDomainsResp, err := client.ApplePay.ListDomains(ctx, paginatedListReq)
	if err != nil {
		log.Printf("Error listing domains with pagination: %v", err)
	} else {
		fmt.Printf("Domains with pagination: %d found\n", len(paginatedDomainsResp.Data.DomainNames))
		if paginatedDomainsResp.Meta != nil {
			fmt.Printf("Meta information available: %+v\n", paginatedDomainsResp.Meta)
		}
	}

	// Example 5: Unregister a domain (commented out to prevent accidental removal)
	/*
		fmt.Println("\n5. Unregistering domain from Apple Pay...")
		unregisterReq := &applepay.UnregisterDomainRequest{
			DomainName: "example.com",
		}

		unregisterResp, err := client.ApplePay.UnregisterDomain(ctx, unregisterReq)
		if err != nil {
			log.Printf("Error unregistering domain: %v", err)
		} else {
			fmt.Printf("✓ Domain unregistered: %s\n", unregisterResp.Message)
		}
	*/
	fmt.Println("\n5. Domain unregistration example skipped (uncomment to test)")

	// Example 6: Error handling with invalid domain
	fmt.Println("\n6. Testing error handling with invalid domain...")
	invalidRegisterReq := &applepay.RegisterDomainRequest{
		DomainName: "", // Empty domain should cause validation error
	}

	_, err = client.ApplePay.RegisterDomain(ctx, invalidRegisterReq)
	if err != nil {
		fmt.Printf("✓ Validation error caught: %v\n", err)
	} else {
		fmt.Println("Unexpected: No error occurred with empty domain")
	}

	// Example 7: Register multiple domains
	fmt.Println("\n7. Registering multiple domains...")
	domains := []string{
		"pay.example.com",
		"secure.example.com",
		"app.example.com",
	}

	for _, domain := range domains {
		regReq := &applepay.RegisterDomainRequest{
			DomainName: domain,
		}

		regResp, err := client.ApplePay.RegisterDomain(ctx, regReq)
		if err != nil {
			log.Printf("Error registering %s: %v", domain, err)
		} else {
			fmt.Printf("✓ %s: %s\n", domain, regResp.Message)
		}
	}

	// Example 8: Final domain list
	fmt.Println("\n8. Final list of registered domains...")
	finalListResp, err := client.ApplePay.ListDomains(ctx, &applepay.ListDomainsRequest{})
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
