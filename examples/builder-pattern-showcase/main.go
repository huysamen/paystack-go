package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/huysamen/paystack-go"
	bulkcharges "github.com/huysamen/paystack-go/api/bulk-charges"
	"github.com/huysamen/paystack-go/api/customers"
	"github.com/huysamen/paystack-go/api/plans"
	"github.com/huysamen/paystack-go/api/products"
	"github.com/huysamen/paystack-go/api/refunds"
	"github.com/huysamen/paystack-go/api/subscriptions"
	"github.com/huysamen/paystack-go/api/transactions"
	"github.com/huysamen/paystack-go/api/transfers"
	"github.com/huysamen/paystack-go/types"
)

func main() {
	// Initialize client
	client := paystack.DefaultClient("sk_test_your_secret_key")

	fmt.Println("=== Builder Pattern Showcase ===")
	fmt.Println("Clean, fluent API across all modules!")

	// 1. Transactions - Complex filtering made simple
	fmt.Println("\n🔹 Transactions:")
	txBuilder := transactions.NewTransactionListRequest().
		PerPage(50).
		Customer(12345).
		Status("success").
		DateRange(time.Now().AddDate(0, -1, 0), time.Now())

	txResp, err := client.Transactions.List(context.Background(), txBuilder)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("   Found %d transactions\n", len(txResp.Data))
	}

	// 2. Customers - Simple pagination
	fmt.Println("\n🔹 Customers:")
	customerBuilder := customers.NewCustomerListRequest().
		PerPage(20).
		Page(1)

	customerResp, err := client.Customers.List(context.Background(), customerBuilder)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("   Found %d customers\n", len(customerResp.Data.Data))
	}

	// 3. Plans - Filter by status and amount
	fmt.Println("\n🔹 Plans:")
	planBuilder := plans.NewPlanListRequest().
		Status("active").
		Interval(types.IntervalMonthly).
		Amount(50000) // ₦500.00

	planResp, err := client.Plans.List(context.Background(), planBuilder)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("   Found %d plans\n", len(planResp.Data))
	}

	// 4. Products - Date range filtering
	fmt.Println("\n🔹 Products:")
	productBuilder := products.NewListProductsRequest().
		PerPage(25).
		DateRange("2024-01-01", "2024-12-31")

	productResp, err := client.Products.List(context.Background(), productBuilder)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("   Found %d products\n", len(productResp.Data))
	}

	// 5. Subscriptions - Filter by customer and plan
	fmt.Println("\n🔹 Subscriptions:")
	subBuilder := subscriptions.NewSubscriptionListRequest().
		Customer(12345).
		Plan(67890).
		PerPage(10)

	subResp, err := client.Subscriptions.List(context.Background(), subBuilder)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("   Found %d subscriptions\n", len(subResp.Data.Data))
	}

	// 6. Transfers - Recent transfers with recipient filter
	fmt.Println("\n🔹 Transfers:")
	transferBuilder := transfers.NewTransferListRequest().
		Recipient(54321).
		From(time.Now().AddDate(0, 0, -7)). // Last week
		PerPage(15)

	transferResp, err := client.Transfers.List(context.Background(), transferBuilder)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("   Found %d transfers\n", len(transferResp.Data.Data))
	}

	// 7. Refunds - List with transaction filter
	fmt.Println("\n🔹 Refunds:")
	refundBuilder := refunds.NewRefundListRequest().
		Transaction("TXN_123456789").
		Currency("NGN").
		PerPage(5)

	refundResp, err := client.Refunds.List(context.Background(), refundBuilder)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("   Listed refunds successfully\n")
		_ = refundResp // Use the response to avoid unused variable warning
	}

	// 8. Bulk Charges - List recent batches
	fmt.Println("\n🔹 Bulk Charges:")
	bulkBuilder := bulkcharges.NewListBulkChargeBatchesRequest().
		PerPage(10).
		From("2024-01-01").
		To("2024-12-31")

	bulkResp, err := client.BulkCharges.List(context.Background(), bulkBuilder)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("   Listed bulk charge batches\n")
		_ = bulkResp // Use the response to avoid unused variable warning
	}

	// 9. Transaction Export - Complex export with many filters
	fmt.Println("\n🔹 Transaction Export:")
	exportBuilder := transactions.NewTransactionExportRequest().
		PerPage(1000).
		Status("success").
		Currency(types.CurrencyNGN).
		Settled(true).
		DateRange(time.Now().AddDate(0, -3, 0), time.Now())

	exportResp, err := client.Transactions.Export(context.Background(), exportBuilder)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("   Export created: %s\n", exportResp.Data.Path)
	}

	// 10. Creating a refund with builder pattern
	fmt.Println("\n🔹 Refund Creation:")
	refundCreateBuilder := refunds.NewRefundCreateRequest("TXN_123456789").
		Amount(25000). // Partial refund of ₦250.00
		Currency("NGN").
		CustomerNote("Refund for damaged item").
		MerchantNote("Item damage confirmed")

	createResp, err := client.Refunds.Create(context.Background(), refundCreateBuilder)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("   Refund created successfully\n")
		_ = createResp // Use the response to avoid unused variable warning
	}

	fmt.Println("\n✅ Builder Pattern Benefits Demonstrated:")
	fmt.Println("   🚀 No more &[]int{50}[0] syntax!")
	fmt.Println("   🔗 Fluent, chainable method calls")
	fmt.Println("   📝 Self-documenting and readable")
	fmt.Println("   🛡️  Type-safe at compile time")
	fmt.Println("   🎯 Method names show exactly what each parameter does")
	fmt.Println("   🧩 No naming conflicts between modules")
	fmt.Println("   🔧 Easy to extend without breaking changes")
	fmt.Println("   🏗️  Each request type has its own dedicated builder")
}
