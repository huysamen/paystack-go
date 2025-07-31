package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/products"
	"github.com/huysamen/paystack-go/types"
)

func main() {
	// Initialize client with test secret key
	client := paystack.DefaultClient("sk_test_your_secret_key")
	ctx := context.Background()

	fmt.Println("=== Paystack Products API Examples ===")
	fmt.Println()

	// Example 1: Create a physical product with limited stock
	fmt.Println("1. Creating a physical product with limited stock...")
	unlimited := false
	quantity := 100
	createReq := &products.CreateProductRequest{
		Name:        "Wireless Headphones",
		Description: "High-quality wireless headphones with noise cancellation",
		Price:       25000, // ₦250.00
		Currency:    "NGN",
		Unlimited:   &unlimited,
		Quantity:    &quantity,
	}

	product, err := client.Products.Create(ctx, createReq)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		return
	}
	fmt.Printf("✓ Product created: %s (Code: %s)\n", product.Name, product.ProductCode)
	fmt.Printf("  Price: ₦%.2f, Stock: %d, In Stock: %t\n",
		float64(product.Price)/100, *product.Quantity, product.InStock)

	// Example 2: Create a digital product with unlimited stock
	fmt.Println("\n2. Creating a digital product with unlimited stock...")
	unlimitedTrue := true
	digitalCreateReq := &products.CreateProductRequest{
		Name:        "E-book: Go Programming Guide",
		Description: "Comprehensive guide to Go programming for beginners",
		Price:       5000, // ₦50.00
		Currency:    "NGN",
		Unlimited:   &unlimitedTrue,
	}

	digitalProduct, err := client.Products.Create(ctx, digitalCreateReq)
	if err != nil {
		log.Printf("Error creating digital product: %v", err)
	} else {
		fmt.Printf("✓ Digital product created: %s (Code: %s)\n",
			digitalProduct.Name, digitalProduct.ProductCode)
		fmt.Printf("  Price: ₦%.2f, Unlimited: %t\n",
			float64(digitalProduct.Price)/100, digitalProduct.Unlimited)
	}

	// Example 3: List all products
	fmt.Println("\n3. Listing all products...")
	listReq := &products.ListProductsRequest{}

	productsResp, err := client.Products.List(ctx, listReq)
	if err != nil {
		log.Printf("Error listing products: %v", err)
		return
	}

	fmt.Printf("Found %d products:\n", len(productsResp.Data))
	for i, prod := range productsResp.Data {
		stockInfo := "Unlimited"
		if !prod.Unlimited && prod.Quantity != nil {
			stockInfo = fmt.Sprintf("%d units", *prod.Quantity)
		}
		fmt.Printf("  %d. %s - ₦%.2f (%s)\n",
			i+1, prod.Name, float64(prod.Price)/100, stockInfo)
	}

	// Example 4: List products with pagination
	fmt.Println("\n4. Listing products with pagination...")
	perPage := 5
	page := 1
	paginatedListReq := &products.ListProductsRequest{
		PerPage: &perPage,
		Page:    &page,
	}

	paginatedResp, err := client.Products.List(ctx, paginatedListReq)
	if err != nil {
		log.Printf("Error listing products with pagination: %v", err)
	} else {
		fmt.Printf("Page %d: %d products (max %d per page)\n",
			page, len(paginatedResp.Data), perPage)
		if paginatedResp.Meta != nil {
			fmt.Printf("Meta: %+v\n", paginatedResp.Meta)
		}
	}

	// Example 5: Fetch product details
	if product != nil && product.ProductCode != "" {
		fmt.Println("\n5. Fetching product details...")
		fetchedProduct, err := client.Products.Fetch(ctx, product.ProductCode)
		if err != nil {
			log.Printf("Error fetching product: %v", err)
		} else {
			fmt.Printf("✓ Product details: %s\n", fetchedProduct.Name)
			fmt.Printf("  Description: %s\n", fetchedProduct.Description)
			fmt.Printf("  Price: ₦%.2f %s\n",
				float64(fetchedProduct.Price)/100, fetchedProduct.Currency)
			fmt.Printf("  Active: %t, In Stock: %t\n",
				fetchedProduct.Active, fetchedProduct.InStock)
			if fetchedProduct.Quantity != nil {
				fmt.Printf("  Quantity: %d", *fetchedProduct.Quantity)
				if fetchedProduct.QuantitySold != nil {
					fmt.Printf(" (Sold: %d)", *fetchedProduct.QuantitySold)
				}
				fmt.Println()
			}
		}
	}

	// Example 6: Update product
	if product != nil && product.ProductCode != "" {
		fmt.Println("\n6. Updating product...")
		newPrice := 30000 // ₦300.00
		newDesc := "Premium wireless headphones with advanced noise cancellation and 30-hour battery life"
		updateReq := &products.UpdateProductRequest{
			Price:       &newPrice,
			Description: &newDesc,
		}

		updatedProduct, err := client.Products.Update(ctx, product.ProductCode, updateReq)
		if err != nil {
			log.Printf("Error updating product: %v", err)
		} else {
			fmt.Printf("✓ Product updated: %s\n", updatedProduct.Name)
			fmt.Printf("  New price: ₦%.2f\n", float64(updatedProduct.Price)/100)
			fmt.Printf("  Updated description: %s\n", updatedProduct.Description)
		}
	}

	// Example 7: Create product with metadata
	fmt.Println("\n7. Creating product with metadata...")
	metadata := types.Metadata{
		"category":  "electronics",
		"brand":     "TechCorp",
		"warranty":  "2 years",
		"color":     "black",
		"weight_kg": 0.3,
		"dimensions": map[string]interface{}{
			"length": 20,
			"width":  15,
			"height": 8,
		},
	}

	metadataCreateReq := &products.CreateProductRequest{
		Name:        "Bluetooth Speaker",
		Description: "Portable bluetooth speaker with deep bass",
		Price:       15000, // ₦150.00
		Currency:    "NGN",
		Unlimited:   &unlimited,
		Quantity:    &[]int{50}[0],
		Metadata:    &metadata,
	}

	metadataProduct, err := client.Products.Create(ctx, metadataCreateReq)
	if err != nil {
		log.Printf("Error creating product with metadata: %v", err)
	} else {
		fmt.Printf("✓ Product with metadata created: %s\n", metadataProduct.Name)
		if metadataProduct.Metadata != nil {
			fmt.Printf("  Metadata: %+v\n", *metadataProduct.Metadata)
		}
	}

	// Example 8: Error handling with invalid data
	fmt.Println("\n8. Testing error handling...")
	invalidReq := &products.CreateProductRequest{
		Name: "", // Empty name should cause validation error
	}

	_, err = client.Products.Create(ctx, invalidReq)
	if err != nil {
		fmt.Printf("✓ Validation error caught: %v\n", err)
	} else {
		fmt.Println("Unexpected: No error occurred with invalid data")
	}

	fmt.Println("\n=== Products API Examples Completed ===")
}
