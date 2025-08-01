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

	fmt.Println("=== Products Fixed-Price Items Example ===")
	fmt.Println()

	// Example 1: Create fixed-price digital products
	fmt.Println("1. Creating fixed-price digital products...")

	// E-book at fixed price
	ebook, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"Complete JavaScript Course",
			"Master JavaScript from basics to advanced concepts with hands-on projects",
			2500000, // Fixed at ₦25,000.00
			"NGN",
		).
			Unlimited(true).
			Metadata(&types.Metadata{
				"type":       "ebook",
				"format":     "PDF + Video",
				"duration":   "40 hours",
				"level":      "beginner_to_advanced",
				"price_type": "fixed",
				"downloads":  "unlimited",
			}),
	)
	if err != nil {
		log.Printf("Error creating e-book: %v", err)
		return
	}
	fmt.Printf("✓ E-book created: %s - Fixed Price: ₦%.2f\n",
		ebook.Name, float64(ebook.Price)/100)

	// Software license at fixed price
	software, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"Professional Design Suite License",
			"One-year license for complete design software suite",
			15000000, // Fixed at ₦150,000.00 per year
			"NGN",
		).
			Unlimited(true).
			Metadata(&types.Metadata{
				"type":         "software_license",
				"duration":     "1 year",
				"license_type": "single_user",
				"price_type":   "fixed_annual",
				"renewals":     "manual",
			}),
	)
	if err != nil {
		log.Printf("Error creating software license: %v", err)
	} else {
		fmt.Printf("✓ Software license created: %s - Fixed Price: ₦%.2f/year\n",
			software.Name, float64(software.Price)/100)
	}

	// Example 2: Create fixed-inventory physical products
	fmt.Println("\n2. Creating fixed-inventory physical products...")

	// Limited edition collectible
	collectible, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"Limited Edition Art Print",
			"Signed and numbered art print - only 50 copies available",
			5000000, // Fixed at ₦50,000.00
			"NGN",
		).
			Unlimited(false).
			Quantity(50). // Fixed inventory
			Metadata(&types.Metadata{
				"type":         "collectible",
				"edition_size": "50",
				"signed":       "true",
				"numbered":     "true",
				"price_type":   "fixed_limited",
				"artist":       "John Doe",
			}),
	)
	if err != nil {
		log.Printf("Error creating collectible: %v", err)
	} else {
		fmt.Printf("✓ Collectible created: %s - Fixed Price: ₦%.2f, Limited to %d units\n",
			collectible.Name, float64(collectible.Price)/100, *collectible.Quantity)
	}

	// Fixed-price hardware device
	device, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"Smart Home Security Camera",
			"Professional-grade security camera with cloud storage",
			7500000, // Fixed at ₦75,000.00
			"NGN",
		).
			Unlimited(false).
			Quantity(200).
			Metadata(&types.Metadata{
				"type":       "hardware",
				"category":   "security",
				"features":   "4K,Night Vision,Motion Detection,Cloud Storage",
				"warranty":   "2 years",
				"price_type": "fixed_retail",
				"brand":      "SecureTech",
			}),
	)
	if err != nil {
		log.Printf("Error creating device: %v", err)
	} else {
		fmt.Printf("✓ Device created: %s - Fixed Price: ₦%.2f, Stock: %d units\n",
			device.Name, float64(device.Price)/100, *device.Quantity)
	}

	// Example 3: Create fixed-price service packages
	fmt.Println("\n3. Creating fixed-price service packages...")

	// Website design package
	webDesign, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"Professional Website Design Package",
			"Complete website design and development service with 3 revisions",
			25000000, // Fixed at ₦250,000.00
			"NGN",
		).
			Unlimited(true).
			Metadata(&types.Metadata{
				"type":       "service_package",
				"delivery":   "14 days",
				"revisions":  "3",
				"includes":   "design,development,hosting_setup,seo_basic",
				"price_type": "fixed_package",
				"support":    "30 days",
			}),
	)
	if err != nil {
		log.Printf("Error creating web design service: %v", err)
	} else {
		fmt.Printf("✓ Service created: %s - Fixed Package Price: ₦%.2f\n",
			webDesign.Name, float64(webDesign.Price)/100)
	}

	// Photography session package
	photography, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"Professional Portrait Session",
			"2-hour professional portrait photography session with edited photos",
			8000000, // Fixed at ₦80,000.00
			"NGN",
		).
			Unlimited(false).
			Quantity(30). // Limited sessions per month
			Metadata(&types.Metadata{
				"type":       "service_session",
				"duration":   "2 hours",
				"photos":     "50 edited photos",
				"location":   "studio or outdoor",
				"price_type": "fixed_session",
				"booking":    "advance_required",
			}),
	)
	if err != nil {
		log.Printf("Error creating photography service: %v", err)
	} else {
		fmt.Printf("✓ Photography service created: %s - Fixed Price: ₦%.2f, %d sessions available\n",
			photography.Name, float64(photography.Price)/100, *photography.Quantity)
	}

	// Example 4: List and manage fixed-price products
	fmt.Println("\n4. Managing fixed-price products...")

	// List all products
	productsResp, err := client.Products.List(ctx,
		products.NewListProductsRequest().
			PerPage(10).
			Page(1),
	)
	if err != nil {
		log.Printf("Error listing products: %v", err)
	} else {
		fmt.Printf("Found %d total products:\n", len(productsResp.Data))

		// Categorize by price type
		fixedPriceProducts := 0
		digitalProducts := 0
		physicalProducts := 0
		serviceProducts := 0

		for _, product := range productsResp.Data {
			if product.Metadata != nil {
				if priceType, exists := (*product.Metadata)["price_type"]; exists {
					if priceTypeStr, ok := priceType.(string); ok && priceTypeStr != "" {
						fixedPriceProducts++
					}
				}

				if productType, exists := (*product.Metadata)["type"]; exists {
					if typeStr, ok := productType.(string); ok {
						switch typeStr {
						case "ebook", "software_license":
							digitalProducts++
						case "collectible", "hardware":
							physicalProducts++
						case "service_package", "service_session":
							serviceProducts++
						}
					}
				}
			}
		}

		fmt.Printf("  Fixed-price items: %d\n", fixedPriceProducts)
		fmt.Printf("  Digital products: %d\n", digitalProducts)
		fmt.Printf("  Physical products: %d\n", physicalProducts)
		fmt.Printf("  Service packages: %d\n", serviceProducts)
	}

	// Example 5: Update fixed-price product (maintaining price structure)
	if ebook != nil {
		fmt.Println("\n5. Updating fixed-price product...")
		updatedEbook, err := client.Products.Update(ctx, ebook.ProductCode,
			products.NewUpdateProductRequest().
				Description("Master JavaScript from basics to advanced concepts with hands-on projects + BONUS: React fundamentals").
				Metadata(&types.Metadata{
					"type":       "ebook",
					"format":     "PDF + Video",
					"duration":   "45 hours", // Updated duration
					"level":      "beginner_to_advanced",
					"price_type": "fixed",
					"downloads":  "unlimited",
					"bonus":      "React fundamentals course",
					"updated":    "true",
				}),
		)
		if err != nil {
			log.Printf("Error updating e-book: %v", err)
		} else {
			fmt.Printf("✓ E-book updated with bonus content - Price remains fixed at ₦%.2f\n",
				float64(updatedEbook.Price)/100)
		}
	}

	// Example 6: Inventory management for fixed-quantity items
	if collectible != nil {
		fmt.Println("\n6. Inventory management for limited items...")

		// Simulate selling some units
		newQuantity := *collectible.Quantity - 5 // Sold 5 units
		updatedCollectible, err := client.Products.Update(ctx, collectible.ProductCode,
			products.NewUpdateProductRequest().
				Quantity(newQuantity).
				Metadata(&types.Metadata{
					"type":         "collectible",
					"edition_size": "50",
					"signed":       "true",
					"numbered":     "true",
					"price_type":   "fixed_limited",
					"artist":       "John Doe",
					"sold":         "5",
					"remaining":    fmt.Sprintf("%d", newQuantity),
				}),
		)
		if err != nil {
			log.Printf("Error updating inventory: %v", err)
		} else {
			fmt.Printf("✓ Inventory updated: %s - %d units remaining (sold 5)\n",
				updatedCollectible.Name, *updatedCollectible.Quantity)
		}
	}

	fmt.Println("\n=== Fixed-Price Products Management Complete ===")
	fmt.Println("✓ Created various fixed-price digital products")
	fmt.Printf("✓ Managed limited inventory physical products\n")
	fmt.Println("✓ Established fixed-price service packages")
	fmt.Println("✓ Demonstrated inventory management for limited items")
	fmt.Println("✓ Updated products while maintaining fixed pricing structure")

	fmt.Println("\nFixed-price products example completed successfully!")
}
