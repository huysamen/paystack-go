package products

import (
	"github.com/huysamen/paystack-go/types"
)

// Product represents a product on the integration
type Product struct {
	ID             int                    `json:"id,omitempty"`
	Name           string                 `json:"name"`
	Description    string                 `json:"description"`
	ProductCode    string                 `json:"product_code,omitempty"`
	Price          int                    `json:"price"`
	Currency       string                 `json:"currency"`
	Quantity       *int                   `json:"quantity,omitempty"`
	QuantitySold   *int                   `json:"quantity_sold,omitempty"`
	Type           string                 `json:"type,omitempty"`
	ImagePath      string                 `json:"image_path,omitempty"`
	FilePath       string                 `json:"file_path,omitempty"`
	IsShippable    bool                   `json:"is_shippable,omitempty"`
	Unlimited      bool                   `json:"unlimited,omitempty"`
	Domain         string                 `json:"domain,omitempty"`
	Active         bool                   `json:"active,omitempty"`
	Features       interface{}            `json:"features,omitempty"`
	InStock        bool                   `json:"in_stock,omitempty"`
	Metadata       *types.Metadata        `json:"metadata,omitempty"`
	Slug           string                 `json:"slug,omitempty"`
	Integration    int                    `json:"integration,omitempty"`
	CreatedAt      string                 `json:"created_at,omitempty"`
	UpdatedAt      string                 `json:"updated_at,omitempty"`
	DigitalAssets  []interface{}          `json:"digital_assets,omitempty"`
	Files          interface{}            `json:"files,omitempty"`
	ShippingFields map[string]interface{} `json:"shipping_fields,omitempty"`
}

// CreateProductRequest represents the request to create a product
type CreateProductRequest struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       int             `json:"price"`
	Currency    string          `json:"currency"`
	Unlimited   *bool           `json:"unlimited,omitempty"`
	Quantity    *int            `json:"quantity,omitempty"`
	Metadata    *types.Metadata `json:"metadata,omitempty"`
}

// UpdateProductRequest represents the request to update a product
type UpdateProductRequest struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Price       *int            `json:"price,omitempty"`
	Currency    *string         `json:"currency,omitempty"`
	Unlimited   *bool           `json:"unlimited,omitempty"`
	Quantity    *int            `json:"quantity,omitempty"`
	Metadata    *types.Metadata `json:"metadata,omitempty"`
}

// ListProductsRequest represents the request to list products
type ListProductsRequest struct {
	PerPage *int    `json:"perPage,omitempty"`
	Page    *int    `json:"page,omitempty"`
	From    *string `json:"from,omitempty"`
	To      *string `json:"to,omitempty"`
}

// CreateProductResponse represents the response from creating a product
type CreateProductResponse struct {
	Status  bool    `json:"status"`
	Message string  `json:"message"`
	Data    Product `json:"data"`
}

// ListProductsResponse represents the response from listing products
type ListProductsResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    []Product   `json:"data"`
	Meta    *types.Meta `json:"meta,omitempty"`
}

// FetchProductResponse represents the response from fetching a product
type FetchProductResponse struct {
	Status  bool    `json:"status"`
	Message string  `json:"message"`
	Data    Product `json:"data"`
}

// UpdateProductResponse represents the response from updating a product
type UpdateProductResponse struct {
	Status  bool    `json:"status"`
	Message string  `json:"message"`
	Data    Product `json:"data"`
}
