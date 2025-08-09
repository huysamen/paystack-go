package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// CustomFilters represents filters for customizing payment options
type CustomFilters struct {
	Recurring                     bool         `json:"recurring,omitempty"`
	Banks                         []string     `json:"banks,omitempty"`
	CardBrands                    []string     `json:"card_brands,omitempty"`
	SupportedMobileMoneyProviders []enums.MoMo `json:"supported_mobile_money_providers,omitempty"`
}

// PaymentPage represents a payment page
type PaymentPage struct {
	ID                int                `json:"id"`
	Integration       int                `json:"integration"`
	Domain            string             `json:"domain"`
	Name              string             `json:"name"`
	Description       data.NullString    `json:"description,omitempty"`
	Amount            data.NullInt       `json:"amount,omitempty"`
	Currency          enums.Currency     `json:"currency"`
	Slug              string             `json:"slug"`
	Type              data.NullString    `json:"type,omitempty"`
	FixedAmount       data.NullBool      `json:"fixed_amount,omitempty"`
	RedirectURL       data.NullString    `json:"redirect_url,omitempty"`
	SuccessMessage    data.NullString    `json:"success_message,omitempty"`
	NotificationEmail data.NullString    `json:"notification_email,omitempty"`
	CollectPhone      data.NullBool      `json:"collect_phone,omitempty"`
	Active            bool               `json:"active"`
	Published         data.NullBool      `json:"published,omitempty"`
	Migrate           data.NullBool      `json:"migrate,omitempty"`
	CustomFields      []CustomField      `json:"custom_fields,omitempty"`
	SplitCode         data.NullString    `json:"split_code,omitempty"`
	Plan              data.NullInt       `json:"plan,omitempty"`
	Products          []PageProduct      `json:"products,omitempty"`
	Metadata          *Metadata          `json:"metadata,omitempty"`
	CreatedAt         data.MultiDateTime `json:"createdAt,omitempty"`
	UpdatedAt         data.MultiDateTime `json:"updatedAt,omitempty"`
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
