package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/huysamen/paystack-go"
	"github.com/huysamen/paystack-go/api/products"
	"github.com/huysamen/paystack-go/types"
)

func main() {
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable is required")
	}

	client := paystack.DefaultClient(secretKey)
	ctx := context.Background()

	fmt.Println("=== Paystack Products API Examples ===")
	fmt.Println()

	// Example 1: Create a physical product with limited stock using builder pattern
	fmt.Println("1. Creating a physical product with limited stock...")
	product, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"Wireless Headphones",
			"High-quality wireless headphones with noise cancellation",
			2500000, // ₦25,000.00 in kobo
			"NGN",
		).
			Unlimited(false).
			Quantity(100).
			Metadata(&types.Metadata{
				"category": "electronics",
				"brand":    "TechSound",
				"model":    "WH-1000X",
			}),
	)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		return
	}
	fmt.Printf("✓ Product created: %s (Code: %s)\n", product.Name, product.ProductCode)
	fmt.Printf("  Price: ₦%.2f, Stock: %d, In Stock: %t\n",
		float64(product.Price)/100, *product.Quantity, product.InStock)

	// Example 2: Create a digital product with unlimited stock using builder pattern
	fmt.Println("\n2. Creating a digital product with unlimited stock...")
	digitalProduct, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"E-book: Go Programming Guide",
			"Comprehensive guide to Go programming for beginners",
			500000, // ₦5,000.00 in kobo
			"NGN",
		).
			Unlimited(true).
			Metadata(&types.Metadata{
				"type":   "ebook",
				"format": "PDF",
				"pages":  "250",
				"level":  "beginner",
			}),
	)
	if err != nil {
		log.Printf("Error creating digital product: %v", err)
	} else {
		fmt.Printf("✓ Digital product created: %s (Code: %s)\n",
			digitalProduct.Name, digitalProduct.ProductCode)
		fmt.Printf("  Price: ₦%.2f, Unlimited: %t\n",
			float64(digitalProduct.Price)/100, digitalProduct.Unlimited)
	}

	// Example 3: List all products using builder pattern
	fmt.Println("\n3. Listing all products...")
	productsResp, err := client.Products.List(ctx,
		products.NewListProductsRequest().
			PerPage(10).
			Page(1),
	)
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

	// Example 4: Paginated listing
	fmt.Println("\n4. Paginated product listing...")
	paginatedResp, err := client.Products.List(ctx,
		products.NewListProductsRequest().
			PerPage(5).
			Page(2),
	)
	if err != nil {
		log.Printf("Error with paginated listing: %v", err)
	} else {
		fmt.Printf("Page 2 - Found %d products\n", len(paginatedResp.Data))
	}

	// Example 5: Fetch specific product
	if len(productsResp.Data) > 0 {
		fmt.Println("\n5. Fetching specific product details...")
		firstProduct := productsResp.Data[0]

		fetchedProduct, err := client.Products.Fetch(ctx, firstProduct.ProductCode)
		if err != nil {
			log.Printf("Error fetching product: %v", err)
		} else {
			fmt.Printf("✓ Fetched product: %s\n", fetchedProduct.Name)
			fmt.Printf("  Description: %s\n", fetchedProduct.Description)
			fmt.Printf("  Price: ₦%.2f\n", float64(fetchedProduct.Price)/100)
			if fetchedProduct.Metadata != nil {
				fmt.Printf("  Metadata: %+v\n", *fetchedProduct.Metadata)
			}
		}

		// Example 6: Update product using builder pattern
		fmt.Println("\n6. Updating product...")
		updatedProduct, err := client.Products.Update(ctx, firstProduct.ProductCode,
			products.NewUpdateProductRequest().
				Name(firstProduct.Name+" (Updated)").
				Description("Updated description with enhanced features").
				Price(firstProduct.Price+500000). // Increase price by ₦5,000
				Metadata(&types.Metadata{
					"updated":    "true",
					"version":    "2.0",
					"updated_by": "system",
				}),
		)
		if err != nil {
			log.Printf("Error updating product: %v", err)
		} else {
			fmt.Printf("✓ Product updated: %s\n", updatedProduct.Name)
			fmt.Printf("  New price: ₦%.2f\n", float64(updatedProduct.Price)/100)
		}
	}

	// Example 7: Create product with comprehensive metadata
	fmt.Println("\n7. Creating product with comprehensive metadata...")
	metadataProduct, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"Premium Software License",
			"Professional software license with 1-year support",
			5000000, // ₦50,000.00 in kobo
			"NGN",
		).
			Unlimited(true).
			Metadata(&types.Metadata{
				"type":         "software",
				"license_type": "professional",
				"support":      "1_year",
				"features":     "advanced_analytics,priority_support,custom_integrations",
				"valid_until":  "2025-12-31",
			}),
	)
	if err != nil {
		log.Printf("Error creating metadata product: %v", err)
	} else {
		fmt.Printf("✓ Metadata product created: %s\n", metadataProduct.Name)
		fmt.Printf("  Product Code: %s\n", metadataProduct.ProductCode)
		if metadataProduct.Metadata != nil {
			fmt.Printf("  Rich metadata included\n")
		}
	}

	// Example 8: Create a course/service product
	fmt.Println("\n8. Creating course/service product...")
	courseProduct, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"Advanced React Development Course",
			"Complete course on advanced React patterns and best practices",
			12000000, // ₦120,000.00 in kobo
			"NGN",
		).
			Unlimited(true).
			Metadata(&types.Metadata{
				"type":       "course",
				"duration":   "30_hours",
				"difficulty": "advanced",
				"includes":   "video_lessons,code_samples,certificate",
				"instructor": "Jane Smith",
			}),
	)
	if err != nil {
		log.Printf("Error creating course product: %v", err)
	} else {
		fmt.Printf("✓ Course product created: %s\n", courseProduct.Name)
		fmt.Printf("  Price: ₦%.2f\n", float64(courseProduct.Price)/100)
	}

	fmt.Println("\n=== Examples completed successfully! ===")
	fmt.Println("✓ Created physical products with inventory")
	fmt.Println("✓ Created digital products with unlimited stock")
	fmt.Println("✓ Listed products with pagination")
	fmt.Println("✓ Fetched individual product details")
	fmt.Println("✓ Updated product information")
	fmt.Println("✓ Used comprehensive metadata")
	fmt.Println("✓ Demonstrated builder pattern usage")
}
