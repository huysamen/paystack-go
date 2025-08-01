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
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("PAYSTACK_SECRET_KEY environment variable is required")
	}

	client := paystack.DefaultClient(secretKey)
	ctx := context.Background()

	fmt.Println("=== Advanced Products Management Example ===")

	// Scenario 1: E-commerce Store Product Catalog
	fmt.Println("\n🛒 Scenario 1: E-commerce Store Product Catalog")

	// Create electronics category products
	fmt.Println("\n📱 Creating electronics products...")
	smartphone, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"iPhone 15 Pro Max",
			"Latest iPhone with advanced camera system and A17 Pro chip",
			180000000, // ₦1,800,000.00 in kobo
			"NGN",
		).
			Unlimited(false).
			Quantity(25).
			Metadata(&types.Metadata{
				"category":    "electronics",
				"subcategory": "smartphones",
				"brand":       "Apple",
				"model":       "iPhone 15 Pro Max",
				"storage":     "256GB",
				"color":       "Natural Titanium",
				"warranty":    "1 year",
				"features":    "Face ID,Wireless Charging,5G,A17 Pro",
				"weight":      "221g",
				"dimensions":  "159.9 x 76.7 x 8.3 mm",
			}),
	)
	if err != nil {
		log.Fatalf("Failed to create smartphone: %v", err)
	}
	fmt.Printf("✓ Smartphone: %s (₦%.2f) - Stock: %d\n",
		smartphone.Name, float64(smartphone.Price)/100, *smartphone.Quantity)

	laptop, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"MacBook Pro 16-inch M3 Max",
			"Professional laptop with M3 Max chip for creative professionals",
			450000000, // ₦4,500,000.00 in kobo
			"NGN",
		).
			Unlimited(false).
			Quantity(10).
			Metadata(&types.Metadata{
				"category":     "electronics",
				"subcategory":  "laptops",
				"brand":        "Apple",
				"processor":    "M3 Max",
				"memory":       "36GB",
				"storage":      "1TB SSD",
				"display":      "16.2-inch Liquid Retina XDR",
				"warranty":     "1 year",
				"target_users": "professionals,creators,developers",
			}),
	)
	if err != nil {
		log.Printf("Warning: Could not create laptop: %v", err)
	} else {
		fmt.Printf("✓ Laptop: %s (₦%.2f) - Stock: %d\n",
			laptop.Name, float64(laptop.Price)/100, *laptop.Quantity)
	}

	// Create fashion products
	fmt.Println("\n👕 Creating fashion products...")
	tshirt, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"Premium Organic Cotton T-Shirt",
			"Comfortable, eco-friendly t-shirt made from 100% organic cotton",
			1500000, // ₦15,000.00 in kobo
			"NGN",
		).
			Unlimited(false).
			Quantity(100).
			Metadata(&types.Metadata{
				"category":      "fashion",
				"subcategory":   "clothing",
				"type":          "t-shirt",
				"material":      "100% organic cotton",
				"sizes":         "XS,S,M,L,XL,XXL",
				"colors":        "White,Black,Navy,Grey",
				"care":          "Machine wash cold, tumble dry low",
				"eco_friendly":  "true",
				"certification": "GOTS certified",
			}),
	)
	if err != nil {
		log.Printf("Warning: Could not create t-shirt: %v", err)
	} else {
		fmt.Printf("✓ T-shirt: %s (₦%.2f) - Stock: %d\n",
			tshirt.Name, float64(tshirt.Price)/100, *tshirt.Quantity)
	}

	sneakers, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"Premium Athletic Sneakers",
			"High-performance sneakers designed for comfort and style",
			4500000, // ₦45,000.00 in kobo
			"NGN",
		).
			Unlimited(false).
			Quantity(50).
			Metadata(&types.Metadata{
				"category":     "fashion",
				"subcategory":  "footwear",
				"type":         "sneakers",
				"sizes":        "37,38,39,40,41,42,43,44,45",
				"colors":       "White/Black,All Black,Navy/White",
				"material":     "mesh upper, rubber sole",
				"features":     "breathable,lightweight,shock absorption",
				"suitable_for": "running,walking,casual wear",
			}),
	)
	if err != nil {
		log.Printf("Warning: Could not create sneakers: %v", err)
	} else {
		fmt.Printf("✓ Sneakers: %s (₦%.2f) - Stock: %d\n",
			sneakers.Name, float64(sneakers.Price)/100, *sneakers.Quantity)
	}

	// Scenario 2: Digital Products and Services
	fmt.Println("\n💻 Scenario 2: Digital Products and Services")

	// Software and licenses
	softwareLicense, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"Adobe Creative Suite License",
			"Annual license for Adobe Creative Suite including Photoshop, Illustrator, InDesign",
			12000000, // ₦120,000.00 in kobo
			"NGN",
		).
			Unlimited(true).
			Metadata(&types.Metadata{
				"type":         "software_license",
				"duration":     "1 year",
				"applications": "Photoshop,Illustrator,InDesign,Premiere Pro,After Effects",
				"license_type": "single_user",
				"platforms":    "Windows,macOS",
				"updates":      "included",
				"support":      "email_chat",
			}),
	)
	if err != nil {
		log.Printf("Warning: Could not create software license: %v", err)
	} else {
		fmt.Printf("✓ Software License: %s (₦%.2f/year)\n",
			softwareLicense.Name, float64(softwareLicense.Price)/100)
	}

	// Online courses
	webDevCourse, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"Full Stack Web Development Bootcamp",
			"Comprehensive 6-month program covering frontend and backend development",
			50000000, // ₦500,000.00 in kobo
			"NGN",
		).
			Unlimited(true).
			Metadata(&types.Metadata{
				"type":            "online_course",
				"duration":        "26 weeks",
				"format":          "live_sessions,recorded_videos,hands_on_projects",
				"technologies":    "HTML,CSS,JavaScript,React,Node.js,MongoDB,PostgreSQL",
				"level":           "beginner_to_advanced",
				"certificate":     "included",
				"support":         "1_on_1_mentoring,community_access",
				"career_services": "resume_review,interview_prep,job_placement_assistance",
			}),
	)
	if err != nil {
		log.Printf("Warning: Could not create web dev course: %v", err)
	} else {
		fmt.Printf("✓ Bootcamp: %s (₦%.2f)\n",
			webDevCourse.Name, float64(webDevCourse.Price)/100)
	}

	// E-books and digital content
	ebook, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"The Complete Guide to Nigerian Fintech",
			"Comprehensive guide covering the Nigerian fintech ecosystem, regulations, and opportunities",
			2500000, // ₦25,000.00 in kobo
			"NGN",
		).
			Unlimited(true).
			Metadata(&types.Metadata{
				"type":     "ebook",
				"format":   "PDF,EPUB,MOBI",
				"pages":    "450",
				"language": "English",
				"author":   "Fintech Expert",
				"topics":   "payments,banking,regulations,investment,blockchain",
				"updated":  "2024",
				"bonus":    "case_studies,templates,resource_links",
			}),
	)
	if err != nil {
		log.Printf("Warning: Could not create ebook: %v", err)
	} else {
		fmt.Printf("✓ E-book: %s (₦%.2f)\n",
			ebook.Name, float64(ebook.Price)/100)
	}

	// Scenario 3: Subscription and Service Products
	fmt.Println("\n📋 Scenario 3: Subscription and Service Products")

	// SaaS subscription
	saasSubscription, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"Business Analytics Platform Pro",
			"Advanced analytics platform with real-time dashboards and AI insights",
			15000000, // ₦150,000.00 in kobo per month
			"NGN",
		).
			Unlimited(true).
			Metadata(&types.Metadata{
				"type":          "saas_subscription",
				"billing_cycle": "monthly",
				"features":      "unlimited_dashboards,ai_insights,api_access,white_labeling",
				"users":         "unlimited",
				"data_storage":  "1TB",
				"support":       "24_7_priority",
				"integrations":  "100+",
				"uptime_sla":    "99.9%",
			}),
	)
	if err != nil {
		log.Printf("Warning: Could not create SaaS subscription: %v", err)
	} else {
		fmt.Printf("✓ SaaS Platform: %s (₦%.2f/month)\n",
			saasSubscription.Name, float64(saasSubscription.Price)/100)
	}

	// Consulting services
	consultingService, err := client.Products.Create(ctx,
		products.NewCreateProductRequest(
			"Strategic Business Consulting Package",
			"Comprehensive business strategy consultation with market analysis and implementation roadmap",
			25000000, // ₦250,000.00 in kobo
			"NGN",
		).
			Unlimited(true).
			Metadata(&types.Metadata{
				"type":         "consulting_service",
				"duration":     "4 weeks",
				"deliverables": "market_analysis,business_strategy,implementation_plan,financial_projections",
				"meetings":     "8 sessions (2 per week)",
				"format":       "virtual_in_person",
				"follow_up":    "3 months support",
				"expertise":    "business_strategy,market_research,financial_planning",
			}),
	)
	if err != nil {
		log.Printf("Warning: Could not create consulting service: %v", err)
	} else {
		fmt.Printf("✓ Consulting: %s (₦%.2f)\n",
			consultingService.Name, float64(consultingService.Price)/100)
	}

	// Scenario 4: Advanced Product Management
	fmt.Println("\n⚙️ Scenario 4: Advanced Product Management")

	// List products by category (using date filtering if available)
	fmt.Println("\nListing products with advanced filtering...")
	allProducts, err := client.Products.List(ctx,
		products.NewListProductsRequest().
			PerPage(20).
			Page(1),
	)
	if err != nil {
		log.Printf("Warning: Could not list products: %v", err)
	} else {
		fmt.Printf("Found %d products total\n", len(allProducts.Data))

		// Categorize products
		categories := make(map[string][]products.Product)
		for _, product := range allProducts.Data {
			category := "other"
			if product.Metadata != nil {
				if cat, exists := (*product.Metadata)["category"]; exists {
					if catStr, ok := cat.(string); ok {
						category = catStr
					}
				}
			}
			categories[category] = append(categories[category], product)
		}

		fmt.Println("\nProducts by category:")
		for category, prods := range categories {
			fmt.Printf("  %s: %d products\n", category, len(prods))
		}
	}

	// Update product with seasonal pricing
	if smartphone != nil {
		fmt.Println("\nUpdating product with seasonal pricing...")
		updatedSmartphone, err := client.Products.Update(ctx, smartphone.ProductCode,
			products.NewUpdateProductRequest().
				Price(smartphone.Price-10000000). // ₦100,000 discount
				Metadata(&types.Metadata{
					"category":    "electronics",
					"subcategory": "smartphones",
					"brand":       "Apple",
					"model":       "iPhone 15 Pro Max",
					"storage":     "256GB",
					"color":       "Natural Titanium",
					"warranty":    "1 year",
					"promotion":   "holiday_sale",
					"discount":    "100000", // ₦1,000.00
					"sale_ends":   time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
					"updated_at":  time.Now().Format("2006-01-02 15:04:05"),
				}),
		)
		if err != nil {
			log.Printf("Warning: Could not update smartphone: %v", err)
		} else {
			fmt.Printf("✓ Updated %s with holiday pricing: ₦%.2f (was ₦%.2f)\n",
				updatedSmartphone.Name,
				float64(updatedSmartphone.Price)/100,
				float64(smartphone.Price)/100)
		}
	}

	// Inventory management simulation
	fmt.Println("\n📦 Inventory management operations...")
	if tshirt != nil {
		// Simulate selling some inventory
		newQuantity := *tshirt.Quantity - 20 // Sold 20 units
		updatedTshirt, err := client.Products.Update(ctx, tshirt.ProductCode,
			products.NewUpdateProductRequest().
				Quantity(newQuantity).
				Metadata(&types.Metadata{
					"category":    "fashion",
					"subcategory": "clothing",
					"type":        "t-shirt",
					"material":    "100% organic cotton",
					"sizes":       "XS,S,M,L,XL,XXL",
					"colors":      "White,Black,Navy,Grey",
					"last_sale":   time.Now().Format("2006-01-02"),
					"units_sold":  "20",
					"stock_status": func() string {
						if newQuantity < 30 {
							return "low_stock"
						}
						return "in_stock"
					}(),
				}),
		)
		if err != nil {
			log.Printf("Warning: Could not update inventory: %v", err)
		} else {
			fmt.Printf("✓ Inventory updated: %s - Stock: %d (sold 20 units)\n",
				updatedTshirt.Name, *updatedTshirt.Quantity)
		}
	}

	// Product analytics summary
	fmt.Println("\n📊 Product Analytics Summary:")
	productCount := 0
	totalValue := 0.0
	digitalProducts := 0
	physicalProducts := 0

	if allProducts != nil {
		for _, product := range allProducts.Data {
			productCount++
			totalValue += float64(product.Price) / 100

			if product.Unlimited {
				digitalProducts++
			} else {
				physicalProducts++
			}
		}
	}

	fmt.Printf("  Total Products: %d\n", productCount)
	fmt.Printf("  Average Price: ₦%.2f\n", totalValue/float64(productCount))
	fmt.Printf("  Digital Products: %d\n", digitalProducts)
	fmt.Printf("  Physical Products: %d\n", physicalProducts)

	fmt.Println("\n=== Advanced Products Management Complete ===")
	fmt.Println("✓ E-commerce product catalog created")
	fmt.Println("✓ Digital products and services configured")
	fmt.Println("✓ Subscription services established")
	fmt.Println("✓ Advanced product management demonstrated")
	fmt.Println("✓ Inventory management operations performed")
	fmt.Println("✓ Product analytics and categorization implemented")
	fmt.Println("✓ Seasonal pricing and promotions applied")

	fmt.Println("\nAdvanced products example completed successfully!")
}
