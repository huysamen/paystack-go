package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Product represents a product on the integration
type Product struct {
	ID                 data.Int        `json:"id,omitempty"`
	Integration        data.Int        `json:"integration,omitempty"`
	Name               data.String     `json:"name"`
	Description        data.String     `json:"description"`
	ProductCode        data.String     `json:"product_code,omitempty"`
	Price              data.Int        `json:"price"`
	Currency           enums.Currency  `json:"currency"`
	Quantity           data.NullInt    `json:"quantity,omitempty"`
	QuantitySold       data.NullInt    `json:"quantity_sold,omitempty"`
	Type               data.String     `json:"type,omitempty"` // good, service
	ImagePath          data.NullString `json:"image_path,omitempty"`
	FilePath           data.NullString `json:"file_path,omitempty"`
	Files              *Metadata       `json:"files,omitempty"`
	IsShippable        data.Bool       `json:"is_shippable,omitempty"`
	ShippingFields     *Metadata       `json:"shipping_fields,omitempty"`
	Unlimited          data.Bool       `json:"unlimited,omitempty"`
	Domain             data.String     `json:"domain,omitempty"`
	Active             data.Bool       `json:"active,omitempty"`
	Features           *Metadata       `json:"features,omitempty"`
	InStock            data.Bool       `json:"in_stock,omitempty"`
	Metadata           *Metadata       `json:"metadata,omitempty"`
	Slug               data.String     `json:"slug,omitempty"`
	SuccessMessage     data.NullString `json:"success_message,omitempty"`
	RedirectURL        data.NullString `json:"redirect_url,omitempty"`
	SplitCode          data.NullString `json:"split_code,omitempty"`
	NotificationEmails []data.String   `json:"notification_emails,omitempty"`
	MinimumOrderable   data.NullInt    `json:"minimum_orderable,omitempty"`
	MaximumOrderable   data.NullInt    `json:"maximum_orderable,omitempty"`
	LowStockAlert      data.Bool       `json:"low_stock_alert,omitempty"`
	StockThreshold     data.NullInt    `json:"stock_threshold,omitempty"`
	ExpiresIn          data.NullInt    `json:"expires_in,omitempty"`
	DigitalAssets      []any           `json:"digital_assets,omitempty"`
	CreatedAt          data.NullTime   `json:"createdAt,omitempty"`
	UpdatedAt          data.NullTime   `json:"updatedAt,omitempty"`
}
