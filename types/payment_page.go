package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// CustomFilters represents filters for customizing payment options
type CustomFilters struct {
	Recurring                     data.Bool     `json:"recurring,omitempty"`
	Banks                         []data.String `json:"banks,omitempty"`
	CardBrands                    []data.String `json:"card_brands,omitempty"`
	SupportedMobileMoneyProviders []enums.MoMo  `json:"supported_mobile_money_providers,omitempty"`
}

// PaymentPage represents a payment page
type PaymentPage struct {
	ID                data.Int        `json:"id"`
	Integration       data.Int        `json:"integration"`
	Domain            data.String     `json:"domain"`
	Name              data.String     `json:"name"`
	Description       data.NullString `json:"description,omitempty"`
	Amount            data.NullInt    `json:"amount,omitempty"`
	Currency          enums.Currency  `json:"currency"`
	Slug              data.String     `json:"slug"`
	Type              data.NullString `json:"type,omitempty"`
	FixedAmount       data.NullBool   `json:"fixed_amount,omitempty"`
	RedirectURL       data.NullString `json:"redirect_url,omitempty"`
	SuccessMessage    data.NullString `json:"success_message,omitempty"`
	NotificationEmail data.NullString `json:"notification_email,omitempty"`
	CollectPhone      data.NullBool   `json:"collect_phone,omitempty"`
	Active            data.Bool       `json:"active"`
	Published         data.NullBool   `json:"published,omitempty"`
	Migrate           data.NullBool   `json:"migrate,omitempty"`
	CustomFields      []CustomField   `json:"custom_fields,omitempty"`
	SplitCode         data.NullString `json:"split_code,omitempty"`
	Plan              data.NullInt    `json:"plan,omitempty"`
	Products          []PageProduct   `json:"products,omitempty"`
	Metadata          Metadata        `json:"metadata,omitempty"`
	CreatedAt         data.NullTime   `json:"createdAt,omitempty"`
	UpdatedAt         data.NullTime   `json:"updatedAt,omitempty"`
}

// PageProduct represents a product within a payment page
type PageProduct struct {
	ProductID   data.Int       `json:"product_id"`
	Name        data.String    `json:"name"`
	Description data.String    `json:"description"`
	ProductCode data.String    `json:"product_code"`
	Page        data.Int       `json:"page"`
	Price       data.Int       `json:"price"`
	Currency    enums.Currency `json:"currency"`
	Quantity    data.Int       `json:"quantity"`
	Type        data.String    `json:"type"`
	Features    Metadata       `json:"features"`
	IsShippable data.Int       `json:"is_shippable"` // 0 or 1
	Domain      data.String    `json:"domain"`
	Integration data.Int       `json:"integration"`
	Active      data.Int       `json:"active"`   // 0 or 1
	InStock     data.Int       `json:"in_stock"` // 0 or 1
}
