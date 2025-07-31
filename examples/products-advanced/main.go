package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/products"
	"github.com/huysamen/paystack-go/types"
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

	fmt.Println("=== Advanced Products API Examples ===")
	fmt.Println()

	// Example 1: E-commerce inventory management
	err := manageEcommerceInventory(ctx, client)
	if err != nil {
		log.Printf("E-commerce inventory example failed: %v", err)
	}

	// Example 2: Bulk product operations
	err = bulkProductOperations(ctx, client)
	if err != nil {
		log.Printf("Bulk operations example failed: %v", err)
	}

	// Example 3: Product lifecycle management
	err = productLifecycleManagement(ctx, client)
	if err != nil {
		log.Printf("Product lifecycle example failed: %v", err)
	}

	// Example 4: Product analytics and reporting
	err = productAnalytics(ctx, client)
	if err != nil {
		log.Printf("Product analytics example failed: %v", err)
	}

	fmt.Println("\n=== Advanced Examples Completed ===")
}

func manageEcommerceInventory(ctx context.Context, client *paystack.Client) error {
	fmt.Println("1. E-commerce Inventory Management...")

	// Create different product categories
	categories := []struct {
		name        string
		description string
		price       int
		unlimited   bool
		quantity    *int
		metadata    types.Metadata
	}{
		{
			name:        "Premium Coffee Beans (1kg)",
			description: "Organic single-origin coffee beans from Colombian highlands",
			price:       8500,
			unlimited:   false,
			quantity:    &[]int{200}[0],
			metadata: types.Metadata{
				"category":        "food",
				"subcategory":     "coffee",
				"organic":         true,
				"origin":          "Colombia",
				"roast_level":     "medium",
				"weight_kg":       1.0,
				"shelf_life_days": 365,
			},
		},
		{
			name:        "Digital Marketing Course",
			description: "Complete digital marketing course with lifetime access",
			price:       25000,
			unlimited:   true,
			quantity:    nil,
			metadata: types.Metadata{
				"category":       "education",
				"type":           "digital",
				"duration_hours": 40,
				"level":          "intermediate",
				"language":       "english",
				"includes":       []string{"videos", "ebooks", "templates", "support"},
			},
		},
		{
			name:        "Handcrafted Leather Wallet",
			description: "Premium genuine leather wallet with RFID protection",
			price:       15000,
			unlimited:   false,
			quantity:    &[]int{50}[0],
			metadata: types.Metadata{
				"category":        "accessories",
				"material":        "genuine leather",
				"color":           "brown",
				"rfid_protection": true,
				"dimensions": map[string]interface{}{
					"length_cm": 11,
					"width_cm":  9,
					"height_cm": 2,
				},
				"warranty_months": 12,
			},
		},
	}

	var createdProducts []string

	for _, cat := range categories {
		createReq := &products.CreateProductRequest{
			Name:        cat.name,
			Description: cat.description,
			Price:       cat.price,
			Currency:    "NGN",
			Unlimited:   &cat.unlimited,
			Quantity:    cat.quantity,
			Metadata:    &cat.metadata,
		}

		product, err := client.Products.Create(ctx, createReq)
		if err != nil {
			log.Printf("Failed to create product %s: %v", cat.name, err)
			continue
		}

		createdProducts = append(createdProducts, product.ProductCode)
		fmt.Printf("✓ Created: %s (₦%.2f) - %s\n",
			product.Name, float64(product.Price)/100, product.ProductCode)
	}

	fmt.Printf("Created %d products for inventory management\n", len(createdProducts))
	return nil
}

func bulkProductOperations(ctx context.Context, client *paystack.Client) error {
	fmt.Println("\n2. Bulk Product Operations...")

	// List products with pagination to handle large inventories
	perPage := 10
	page := 1

	for {
		listReq := &products.ListProductsRequest{
			PerPage: &perPage,
			Page:    &page,
		}

		productsResp, err := client.Products.List(ctx, listReq)
		if err != nil {
			return fmt.Errorf("failed to list products page %d: %w", page, err)
		}

		fmt.Printf("Page %d: %d products\n", page, len(productsResp.Data))

		// Process each product
		for _, product := range productsResp.Data {
			// Example: Update products with missing metadata
			if product.Metadata == nil || len(*product.Metadata) == 0 {
				fmt.Printf("  Updating metadata for: %s\n", product.Name)

				defaultMetadata := types.Metadata{
					"updated_at": time.Now().Format(time.RFC3339),
					"status":     "active",
					"reviewed":   true,
				}

				updateReq := &products.UpdateProductRequest{
					Metadata: &defaultMetadata,
				}

				_, err := client.Products.Update(ctx, product.ProductCode, updateReq)
				if err != nil {
					log.Printf("Failed to update product %s: %v", product.ProductCode, err)
				}
			}

			// Example: Flag low stock items
			if !product.Unlimited && product.Quantity != nil && *product.Quantity < 10 {
				fmt.Printf("  ⚠ Low stock warning: %s (%d units)\n",
					product.Name, *product.Quantity)
			}
		}

		// Check if there are more pages
		if productsResp.Meta == nil || len(productsResp.Data) < perPage {
			break
		}
		page++
	}

	return nil
}

func productLifecycleManagement(ctx context.Context, client *paystack.Client) error {
	fmt.Println("\n3. Product Lifecycle Management...")

	// Create a product with lifecycle tracking
	metadata := types.Metadata{
		"lifecycle_stage": "development",
		"created_by":      "inventory_manager",
		"version":         "1.0",
		"status":          "draft",
	}

	unlimited := false
	quantity := 100

	createReq := &products.CreateProductRequest{
		Name:        "Smart Fitness Tracker",
		Description: "Advanced fitness tracker with heart rate monitoring",
		Price:       35000,
		Currency:    "NGN",
		Unlimited:   &unlimited,
		Quantity:    &quantity,
		Metadata:    &metadata,
	}

	product, err := client.Products.Create(ctx, createReq)
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}

	fmt.Printf("✓ Created product: %s\n", product.Name)

	// Simulate lifecycle stages
	stages := []struct {
		stage       string
		description string
		price       *int
		status      string
	}{
		{"testing", "Product in testing phase", nil, "testing"},
		{"launch", "Product launched to market", &[]int{32000}[0], "active"},
		{"promotion", "Product on promotion", &[]int{28000}[0], "promoted"},
		{"mature", "Stable product in market", &[]int{35000}[0], "stable"},
	}

	for _, stage := range stages {
		fmt.Printf("Moving to stage: %s\n", stage.stage)

		updatedMetadata := types.Metadata{
			"lifecycle_stage": stage.stage,
			"updated_at":      time.Now().Format(time.RFC3339),
			"status":          stage.status,
			"updated_by":      "system",
		}

		updateReq := &products.UpdateProductRequest{
			Metadata: &updatedMetadata,
		}

		if stage.price != nil {
			updateReq.Price = stage.price
		}

		updatedProduct, err := client.Products.Update(ctx, product.ProductCode, updateReq)
		if err != nil {
			log.Printf("Failed to update product to stage %s: %v", stage.stage, err)
			continue
		}

		fmt.Printf("  ✓ Stage: %s, Price: ₦%.2f\n",
			stage.stage, float64(updatedProduct.Price)/100)
	}

	return nil
}

func productAnalytics(ctx context.Context, client *paystack.Client) error {
	fmt.Println("\n4. Product Analytics and Reporting...")

	// Get all products for analysis
	allProducts, err := getAllProducts(ctx, client)
	if err != nil {
		return fmt.Errorf("failed to get all products: %w", err)
	}

	fmt.Printf("Analyzing %d products...\n", len(allProducts))

	// Analytics
	stats := struct {
		totalProducts    int
		digitalProducts  int
		physicalProducts int
		unlimitedStock   int
		lowStock         int
		totalValue       int64
		averagePrice     float64
		categories       map[string]int
	}{
		categories: make(map[string]int),
	}

	var totalPrice int64
	for _, product := range allProducts {
		stats.totalProducts++
		totalPrice += int64(product.Price)

		// Classify by stock type
		if product.Unlimited {
			stats.unlimitedStock++
			stats.digitalProducts++ // Assume unlimited = digital
		} else {
			stats.physicalProducts++
			if product.Quantity != nil && *product.Quantity < 10 {
				stats.lowStock++
			}
		}

		// Category analysis from metadata
		if product.Metadata != nil {
			if category, exists := (*product.Metadata)["category"]; exists {
				if catStr, ok := category.(string); ok {
					stats.categories[catStr]++
				}
			}
		}
	}

	stats.totalValue = totalPrice
	if stats.totalProducts > 0 {
		stats.averagePrice = float64(totalPrice) / float64(stats.totalProducts)
	}

	// Print analytics report
	fmt.Println("\n📊 Product Analytics Report:")
	fmt.Printf("  Total Products: %d\n", stats.totalProducts)
	fmt.Printf("  Digital Products: %d (%.1f%%)\n",
		stats.digitalProducts,
		float64(stats.digitalProducts)/float64(stats.totalProducts)*100)
	fmt.Printf("  Physical Products: %d (%.1f%%)\n",
		stats.physicalProducts,
		float64(stats.physicalProducts)/float64(stats.totalProducts)*100)
	fmt.Printf("  Low Stock Items: %d\n", stats.lowStock)
	fmt.Printf("  Total Inventory Value: ₦%.2f\n", float64(stats.totalValue)/100)
	fmt.Printf("  Average Product Price: ₦%.2f\n", stats.averagePrice/100)

	if len(stats.categories) > 0 {
		fmt.Println("  Categories:")
		for category, count := range stats.categories {
			fmt.Printf("    %s: %d products\n", category, count)
		}
	}

	return nil
}

func getAllProducts(ctx context.Context, client *paystack.Client) ([]products.Product, error) {
	var allProducts []products.Product
	perPage := 50
	page := 1

	for {
		listReq := &products.ListProductsRequest{
			PerPage: &perPage,
			Page:    &page,
		}

		productsResp, err := client.Products.List(ctx, listReq)
		if err != nil {
			return nil, err
		}

		allProducts = append(allProducts, productsResp.Data...)

		if len(productsResp.Data) < perPage {
			break
		}
		page++
	}

	return allProducts, nil
}
