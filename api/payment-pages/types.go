package paymentpages

import (
	"github.com/huysamen/paystack-go/types"
)

// CustomField represents a custom field for payment pages
type CustomField struct {
	DisplayName  string `json:"display_name"`
	VariableName string `json:"variable_name"`
	Required     bool   `json:"required"`
}

// Product represents a product associated with a payment page
type Product struct {
	ProductID   int    `json:"product_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ProductCode string `json:"product_code"`
	Price       int    `json:"price"`
	Currency    string `json:"currency"`
	Quantity    int    `json:"quantity"`
	Type        string `json:"type"`
}

// PaymentPage represents a payment page
type PaymentPage struct {
	ID                int             `json:"id"`
	Name              string          `json:"name"`
	Description       string          `json:"description,omitempty"`
	Amount            *int            `json:"amount,omitempty"`
	Currency          string          `json:"currency"`
	Slug              string          `json:"slug"`
	Type              string          `json:"type"`
	FixedAmount       bool            `json:"fixed_amount"`
	RedirectURL       string          `json:"redirect_url,omitempty"`
	SuccessMessage    string          `json:"success_message,omitempty"`
	NotificationEmail string          `json:"notification_email,omitempty"`
	CollectPhone      bool            `json:"collect_phone"`
	Active            bool            `json:"active"`
	Published         bool            `json:"published"`
	Migrate           bool            `json:"migrate"`
	CustomFields      []CustomField   `json:"custom_fields,omitempty"`
	SplitCode         string          `json:"split_code,omitempty"`
	Plan              *int            `json:"plan,omitempty"`
	Products          []Product       `json:"products,omitempty"`
	Integration       int             `json:"integration"`
	Domain            string          `json:"domain"`
	Metadata          *types.Metadata `json:"metadata,omitempty"`
	CreatedAt         string          `json:"createdAt,omitempty"`
	UpdatedAt         string          `json:"updatedAt,omitempty"`
}

// CreatePaymentPageRequest represents the request to create a payment page
type CreatePaymentPageRequest struct {
	Name              string          `json:"name"`
	Description       string          `json:"description,omitempty"`
	Amount            *int            `json:"amount,omitempty"`
	Currency          string          `json:"currency,omitempty"`
	Slug              string          `json:"slug,omitempty"`
	Type              string          `json:"type,omitempty"`
	Plan              string          `json:"plan,omitempty"`
	FixedAmount       *bool           `json:"fixed_amount,omitempty"`
	SplitCode         string          `json:"split_code,omitempty"`
	Metadata          *types.Metadata `json:"metadata,omitempty"`
	RedirectURL       string          `json:"redirect_url,omitempty"`
	SuccessMessage    string          `json:"success_message,omitempty"`
	NotificationEmail string          `json:"notification_email,omitempty"`
	CollectPhone      *bool           `json:"collect_phone,omitempty"`
	CustomFields      []CustomField   `json:"custom_fields,omitempty"`
}

// UpdatePaymentPageRequest represents the request to update a payment page
type UpdatePaymentPageRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Amount      *int   `json:"amount,omitempty"`
	Active      *bool  `json:"active,omitempty"`
}

// ListPaymentPagesRequest represents the request to list payment pages
type ListPaymentPagesRequest struct {
	PerPage int    `json:"perPage,omitempty"`
	Page    int    `json:"page,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
}

// AddProductsToPageRequest represents the request to add products to a payment page
type AddProductsToPageRequest struct {
	Product []int `json:"product"`
}

// CreatePaymentPageResponse represents the response from creating a payment page
type CreatePaymentPageResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    PaymentPage `json:"data"`
}

// ListPaymentPagesResponse represents the response from listing payment pages
type ListPaymentPagesResponse struct {
	Status  bool          `json:"status"`
	Message string        `json:"message"`
	Data    []PaymentPage `json:"data"`
	Meta    *types.Meta   `json:"meta,omitempty"`
}

// FetchPaymentPageResponse represents the response from fetching a payment page
type FetchPaymentPageResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    PaymentPage `json:"data"`
}

// UpdatePaymentPageResponse represents the response from updating a payment page
type UpdatePaymentPageResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    PaymentPage `json:"data"`
}

// CheckSlugAvailabilityResponse represents the response from checking slug availability
type CheckSlugAvailabilityResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// AddProductsToPageResponse represents the response from adding products to a payment page
type AddProductsToPageResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    PaymentPage `json:"data"`
}
