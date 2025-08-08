package types

import "github.com/huysamen/paystack-go/enums"

// CustomFilters represents filters for customizing payment options
type CustomFilters struct {
	Recurring                     bool         `json:"recurring,omitempty"`
	Banks                         []string     `json:"banks,omitempty"`
	CardBrands                    []string     `json:"card_brands,omitempty"`
	SupportedMobileMoneyProviders []enums.MoMo `json:"supported_mobile_money_providers,omitempty"`
}

// PaymentPage represents a payment page
type PaymentPage struct {
	ID                int            `json:"id"`
	Integration       int            `json:"integration"`
	Domain            string         `json:"domain"`
	Name              string         `json:"name"`
	Description       *string        `json:"description,omitempty"`
	Amount            *int           `json:"amount,omitempty"`
	Currency          enums.Currency `json:"currency"`
	Slug              string         `json:"slug"`
	Type              *string        `json:"type,omitempty"`
	FixedAmount       *bool          `json:"fixed_amount,omitempty"`
	RedirectURL       *string        `json:"redirect_url,omitempty"`
	SuccessMessage    *string        `json:"success_message,omitempty"`
	NotificationEmail *string        `json:"notification_email,omitempty"`
	CollectPhone      *bool          `json:"collect_phone,omitempty"`
	Active            bool           `json:"active"`
	Published         *bool          `json:"published,omitempty"`
	Migrate           *bool          `json:"migrate,omitempty"`
	CustomFields      []CustomField  `json:"custom_fields,omitempty"`
	SplitCode         *string        `json:"split_code,omitempty"`
	Plan              *int           `json:"plan,omitempty"`
	Products          []PageProduct  `json:"products,omitempty"`
	Metadata          *Metadata      `json:"metadata,omitempty"`
	CreatedAt         DateTime       `json:"createdAt,omitempty"`
	UpdatedAt         DateTime       `json:"updatedAt,omitempty"`
}

// PageProduct represents a product within a payment page
type PageProduct struct {
	ProductID   int            `json:"product_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ProductCode string         `json:"product_code"`
	Page        int            `json:"page"`
	Price       int            `json:"price"`
	Currency    enums.Currency `json:"currency"`
	Quantity    int            `json:"quantity"`
	Type        string         `json:"type"`
	Features    *Metadata      `json:"features"`
	IsShippable int            `json:"is_shippable"` // 0 or 1
	Domain      string         `json:"domain"`
	Integration int            `json:"integration"`
	Active      int            `json:"active"`   // 0 or 1
	InStock     int            `json:"in_stock"` // 0 or 1
}
