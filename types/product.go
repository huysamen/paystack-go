package types

// Product represents a product on the integration
type Product struct {
	ID             int            `json:"id,omitempty"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	ProductCode    string         `json:"product_code,omitempty"`
	Price          int            `json:"price"`
	Currency       string         `json:"currency"`
	Quantity       *int           `json:"quantity,omitempty"`
	QuantitySold   *int           `json:"quantity_sold,omitempty"`
	Type           string         `json:"type,omitempty"`
	ImagePath      string         `json:"image_path,omitempty"`
	FilePath       string         `json:"file_path,omitempty"`
	IsShippable    bool           `json:"is_shippable,omitempty"`
	Unlimited      bool           `json:"unlimited,omitempty"`
	Domain         string         `json:"domain,omitempty"`
	Active         bool           `json:"active,omitempty"`
	Features       any            `json:"features,omitempty"`
	InStock        bool           `json:"in_stock,omitempty"`
	Metadata       *Metadata      `json:"metadata,omitempty"`
	Slug           string         `json:"slug,omitempty"`
	Integration    int            `json:"integration,omitempty"`
	CreatedAt      string         `json:"created_at,omitempty"`
	UpdatedAt      string         `json:"updated_at,omitempty"`
	DigitalAssets  []any          `json:"digital_assets,omitempty"`
	Files          any            `json:"files,omitempty"`
	ShippingFields map[string]any `json:"shipping_fields,omitempty"`
}
