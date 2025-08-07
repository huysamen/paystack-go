package types

import "github.com/huysamen/paystack-go/enums"

// Product represents a product on the integration
type Product struct {
	ID                 int            `json:"id,omitempty"`
	Integration        int            `json:"integration,omitempty"`
	Name               string         `json:"name"`
	Description        string         `json:"description"`
	ProductCode        string         `json:"product_code,omitempty"`
	Price              int            `json:"price"`
	Currency           enums.Currency `json:"currency"`
	Quantity           *int           `json:"quantity,omitempty"`
	QuantitySold       *int           `json:"quantity_sold,omitempty"`
	Type               string         `json:"type,omitempty"` // good, service
	ImagePath          *string        `json:"image_path,omitempty"`
	FilePath           *string        `json:"file_path,omitempty"`
	Files              *Metadata      `json:"files,omitempty"`
	IsShippable        bool           `json:"is_shippable,omitempty"`
	ShippingFields     *Metadata      `json:"shipping_fields,omitempty"`
	Unlimited          bool           `json:"unlimited,omitempty"`
	Domain             string         `json:"domain,omitempty"`
	Active             bool           `json:"active,omitempty"`
	Features           *Metadata      `json:"features,omitempty"`
	InStock            bool           `json:"in_stock,omitempty"`
	Metadata           *Metadata      `json:"metadata,omitempty"`
	Slug               string         `json:"slug,omitempty"`
	SuccessMessage     *string        `json:"success_message,omitempty"`
	RedirectURL        *string        `json:"redirect_url,omitempty"`
	SplitCode          *string        `json:"split_code,omitempty"`
	NotificationEmails []string       `json:"notification_emails,omitempty"`
	MinimumOrderable   *int           `json:"minimum_orderable,omitempty"`
	MaximumOrderable   *int           `json:"maximum_orderable,omitempty"`
	LowStockAlert      bool           `json:"low_stock_alert,omitempty"`
	StockThreshold     *int           `json:"stock_threshold,omitempty"`
	ExpiresIn          *int           `json:"expires_in,omitempty"`
	DigitalAssets      []any          `json:"digital_assets,omitempty"`
	CreatedAt          DateTime       `json:"createdAt,omitempty"`
	UpdatedAt          DateTime       `json:"updatedAt,omitempty"`
}
