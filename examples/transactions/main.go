package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/transactions"
)

func main() {
	// Initialize client
	client := paystack.DefaultClient("sk_test_your_secret_key")

	fmt.Println("=== Transaction List Examples ===")

	// Example 1: Simple pagination
	fmt.Println("\n1. Simple pagination:")

	builder1 := transactions.NewTransactionListRequest().
		PerPage(10).
		Page(1)

	resp1, err := client.Transactions.List(context.Background(), builder1)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   Found %d transactions (page 1, 10 per page)\n", len(resp1.Data))

	// Example 2: Advanced filtering - one fluent chain
	fmt.Println("\n2. Advanced filtering:")

	lastWeek := time.Now().AddDate(0, 0, -7)
	builder2 := transactions.NewTransactionListRequest().
		PerPage(50).
		Status("success").
		From(lastWeek).
		Amount(100000) // 1000.00 NGN in kobo

	resp2, err := client.Transactions.List(context.Background(), builder2)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   Found %d successful transactions from last week with amount 1000 NGN\n", len(resp2.Data))

	// Example 3: Date range filtering
	fmt.Println("\n3. Date range filtering:")

	startDate := time.Now().AddDate(0, -1, 0) // 1 month ago
	endDate := time.Now()
	builder3 := transactions.NewTransactionListRequest().
		DateRange(startDate, endDate).
		Customer(12345)

	resp3, err := client.Transactions.List(context.Background(), builder3)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   Found %d transactions for customer 12345 in the last month\n", len(resp3.Data))

	// Example 4: Complex filtering with terminal and multiple conditions
	fmt.Println("\n4. Complex filtering:")

	complexBuilder := transactions.NewTransactionListRequest().
		PerPage(25).
		Page(1).
		Customer(67890).
		TerminalID("terminal_123").
		Status("success").
		From(time.Now().AddDate(0, 0, -30)). // Last 30 days
		Amount(50000)                        // 500.00 NGN

	resp4, err := client.Transactions.List(context.Background(), complexBuilder)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   Found %d transactions matching complex criteria\n", len(resp4.Data))

	fmt.Println("\n✅ All examples completed successfully!")
	fmt.Println("\n💡 Benefits of builder pattern:")
	fmt.Println("   - No more &[]int{50}[0] or pointer helpers needed")
	fmt.Println("   - Fluent, chainable interface")
	fmt.Println("   - Type-safe at compile time")
	fmt.Println("   - Self-documenting method names")
	fmt.Println("   - No naming conflicts between different request types")
	fmt.Println("   - Easy to add new methods without breaking changes")
	fmt.Println("   - Clear ownership - each request type has its own builder")
}
