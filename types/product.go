package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Product represents a product on the integration
type Product struct {
	ID                 int                `json:"id,omitempty"`
	Integration        int                `json:"integration,omitempty"`
	Name               string             `json:"name"`
	Description        string             `json:"description"`
	ProductCode        string             `json:"product_code,omitempty"`
	Price              int                `json:"price"`
	Currency           enums.Currency     `json:"currency"`
	Quantity           data.NullInt       `json:"quantity,omitempty"`
	QuantitySold       data.NullInt       `json:"quantity_sold,omitempty"`
	Type               string             `json:"type,omitempty"` // good, service
	ImagePath          data.NullString    `json:"image_path,omitempty"`
	FilePath           data.NullString    `json:"file_path,omitempty"`
	Files              *Metadata          `json:"files,omitempty"`
	IsShippable        bool               `json:"is_shippable,omitempty"`
	ShippingFields     *Metadata          `json:"shipping_fields,omitempty"`
	Unlimited          bool               `json:"unlimited,omitempty"`
	Domain             string             `json:"domain,omitempty"`
	Active             bool               `json:"active,omitempty"`
	Features           *Metadata          `json:"features,omitempty"`
	InStock            bool               `json:"in_stock,omitempty"`
	Metadata           *Metadata          `json:"metadata,omitempty"`
	Slug               string             `json:"slug,omitempty"`
	SuccessMessage     data.NullString    `json:"success_message,omitempty"`
	RedirectURL        data.NullString    `json:"redirect_url,omitempty"`
	SplitCode          data.NullString    `json:"split_code,omitempty"`
	NotificationEmails []string           `json:"notification_emails,omitempty"`
	MinimumOrderable   data.NullInt       `json:"minimum_orderable,omitempty"`
	MaximumOrderable   data.NullInt       `json:"maximum_orderable,omitempty"`
	LowStockAlert      bool               `json:"low_stock_alert,omitempty"`
	StockThreshold     data.NullInt       `json:"stock_threshold,omitempty"`
	ExpiresIn          data.NullInt       `json:"expires_in,omitempty"`
	DigitalAssets      []any              `json:"digital_assets,omitempty"`
	CreatedAt          data.MultiDateTime `json:"createdAt,omitempty"`
	UpdatedAt          data.MultiDateTime `json:"updatedAt,omitempty"`
}
