package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/huysamen/paystack-go"
	applepay "github.com/huysamen/paystack-go/api/apple-pay"
)

func main() {
	// Get secret key from environment variable
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable is required")
	}

	// Initialize client
	client := paystack.DefaultClient(secretKey)
	ctx := context.Background()

	fmt.Println("=== Advanced Apple Pay API Examples ===")
	fmt.Println()

	// Example 1: Domain Management Workflow
	err := manageDomainWorkflow(ctx, client)
	if err != nil {
		log.Printf("Domain management workflow failed: %v", err)
	}

	// Example 2: Bulk Domain Operations
	err = bulkDomainOperations(ctx, client)
	if err != nil {
		log.Printf("Bulk domain operations failed: %v", err)
	}

	// Example 3: Domain Validation and Error Handling
	err = validateDomainOperations(ctx, client)
	if err != nil {
		log.Printf("Domain validation examples failed: %v", err)
	}

	// Example 4: Production-Ready Domain Setup
	err = productionDomainSetup(ctx, client)
	if err != nil {
		log.Printf("Production domain setup failed: %v", err)
	}

	fmt.Println("\n=== Advanced Examples Completed ===")
}

func manageDomainWorkflow(ctx context.Context, client *paystack.Client) error {
	fmt.Println("1. Domain Management Workflow...")

	// Check current registered domains
	currentDomains, err := client.ApplePay.ListDomains(ctx, &applepay.ListDomainsRequest{})
	if err != nil {
		return fmt.Errorf("failed to list current domains: %w", err)
	}

	fmt.Printf("Current registered domains: %d\n", len(currentDomains.Data.DomainNames))

	// Register a new domain if not already registered
	newDomain := "checkout.mystore.com"
	domainExists := false
	for _, domain := range currentDomains.Data.DomainNames {
		if domain == newDomain {
			domainExists = true
			break
		}
	}

	if !domainExists {
		registerReq := &applepay.RegisterDomainRequest{
			DomainName: newDomain,
		}

		registerResp, err := client.ApplePay.RegisterDomain(ctx, registerReq)
		if err != nil {
			return fmt.Errorf("failed to register domain %s: %w", newDomain, err)
		}

		fmt.Printf("✓ New domain registered: %s - %s\n", newDomain, registerResp.Message)
	} else {
		fmt.Printf("✓ Domain %s already registered\n", newDomain)
	}

	return nil
}

func bulkDomainOperations(ctx context.Context, client *paystack.Client) error {
	fmt.Println("\n2. Bulk Domain Operations...")

	// List of domains to register for a multi-environment setup
	domains := []string{
		"dev-pay.example.com",
		"staging-pay.example.com",
		"prod-pay.example.com",
	}

	successCount := 0
	failCount := 0

	for _, domain := range domains {
		registerReq := &applepay.RegisterDomainRequest{
			DomainName: domain,
		}

		registerResp, err := client.ApplePay.RegisterDomain(ctx, registerReq)
		if err != nil {
			log.Printf("Failed to register %s: %v", domain, err)
			failCount++
		} else {
			fmt.Printf("✓ %s: %s\n", domain, registerResp.Message)
			successCount++
		}
	}

	fmt.Printf("Bulk registration results: %d successful, %d failed\n", successCount, failCount)

	// List all domains to verify
	allDomains, err := client.ApplePay.ListDomains(ctx, &applepay.ListDomainsRequest{})
	if err != nil {
		return fmt.Errorf("failed to list domains after bulk registration: %w", err)
	}

	fmt.Printf("Total registered domains after bulk operation: %d\n", len(allDomains.Data.DomainNames))

	return nil
}

func validateDomainOperations(ctx context.Context, client *paystack.Client) error {
	fmt.Println("\n3. Domain Validation and Error Handling...")

	// Test cases for validation
	testCases := []struct {
		domain      string
		expectError bool
		description string
	}{
		{"", true, "empty domain"},
		{"invalid-domain", false, "domain without TLD"},
		{"valid.example.com", false, "valid subdomain"},
		{"another-valid-domain.co.uk", false, "valid domain with country code"},
	}

	for _, testCase := range testCases {
		fmt.Printf("Testing %s: %s\n", testCase.description, testCase.domain)

		registerReq := &applepay.RegisterDomainRequest{
			DomainName: testCase.domain,
		}

		_, err := client.ApplePay.RegisterDomain(ctx, registerReq)
		if testCase.expectError && err != nil {
			fmt.Printf("  ✓ Expected error occurred: %v\n", err)
		} else if !testCase.expectError && err != nil {
			fmt.Printf("  ⚠ Unexpected error: %v\n", err)
		} else if testCase.expectError && err == nil {
			fmt.Printf("  ⚠ Expected error but registration succeeded\n")
		} else {
			fmt.Printf("  ✓ Registration successful\n")
		}
	}

	return nil
}

func productionDomainSetup(ctx context.Context, client *paystack.Client) error {
	fmt.Println("\n4. Production-Ready Domain Setup...")

	// Typical production domains for an e-commerce platform
	productionDomains := []string{
		"payments.mystore.com",
		"checkout.mystore.com",
		"api.mystore.com",
	}

	// Get current domains
	currentDomains, err := client.ApplePay.ListDomains(ctx, &applepay.ListDomainsRequest{})
	if err != nil {
		return fmt.Errorf("failed to list current domains: %w", err)
	}

	currentDomainMap := make(map[string]bool)
	for _, domain := range currentDomains.Data.DomainNames {
		currentDomainMap[domain] = true
	}

	// Register only new domains
	for _, domain := range productionDomains {
		if currentDomainMap[domain] {
			fmt.Printf("✓ %s already registered\n", domain)
			continue
		}

		registerReq := &applepay.RegisterDomainRequest{
			DomainName: domain,
		}

		registerResp, err := client.ApplePay.RegisterDomain(ctx, registerReq)
		if err != nil {
			log.Printf("Failed to register production domain %s: %v", domain, err)
		} else {
			fmt.Printf("✓ Production domain registered: %s - %s\n", domain, registerResp.Message)
		}
	}

	// Final verification
	finalDomains, err := client.ApplePay.ListDomains(ctx, &applepay.ListDomainsRequest{})
	if err != nil {
		return fmt.Errorf("failed to verify final domain list: %w", err)
	}

	fmt.Printf("\nProduction setup complete. Total domains: %d\n", len(finalDomains.Data.DomainNames))

	// Group domains by environment for reporting
	environments := map[string][]string{
		"production":  {},
		"staging":     {},
		"development": {},
		"other":       {},
	}

	for _, domain := range finalDomains.Data.DomainNames {
		if strings.Contains(domain, "prod") || (!strings.Contains(domain, "dev") && !strings.Contains(domain, "staging")) {
			environments["production"] = append(environments["production"], domain)
		} else if strings.Contains(domain, "staging") {
			environments["staging"] = append(environments["staging"], domain)
		} else if strings.Contains(domain, "dev") {
			environments["development"] = append(environments["development"], domain)
		} else {
			environments["other"] = append(environments["other"], domain)
		}
	}

	fmt.Println("\nDomains by environment:")
	for env, domains := range environments {
		if len(domains) > 0 {
			fmt.Printf("  %s (%d): %v\n", env, len(domains), domains)
		}
	}

	return nil
}
