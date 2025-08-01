package paymentrequests

import (
	"github.com/huysamen/paystack-go/types"
)

// LineItem represents an item in a payment request
type LineItem struct {
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	Quantity int    `json:"quantity,omitempty"`
}

// Tax represents a tax item in a payment request
type Tax struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

// Customer represents customer information in a payment request
type Customer struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email"`
	CustomerCode string `json:"customer_code"`
	Phone        string `json:"phone,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

// Transaction represents a transaction related to a payment request
type Transaction struct {
	ID            int            `json:"id"`
	Reference     string         `json:"reference"`
	Amount        int            `json:"amount"`
	Currency      string         `json:"currency"`
	Status        string         `json:"status"`
	Gateway       string         `json:"gateway"`
	Channel       string         `json:"channel"`
	PaidAt        string         `json:"paid_at,omitempty"`
	CreatedAt     string         `json:"created_at"`
	Authorization map[string]any `json:"authorization,omitempty"`
	Customer      *Customer      `json:"customer,omitempty"`
}

// PaymentRequest represents a payment request
type PaymentRequest struct {
	ID               int             `json:"id"`
	Domain           string          `json:"domain"`
	Amount           int             `json:"amount"`
	Currency         string          `json:"currency"`
	DueDate          string          `json:"due_date,omitempty"`
	HasInvoice       bool            `json:"has_invoice"`
	InvoiceNumber    *int            `json:"invoice_number,omitempty"`
	Description      string          `json:"description"`
	PDFURL           *string         `json:"pdf_url,omitempty"`
	LineItems        []LineItem      `json:"line_items,omitempty"`
	Tax              []Tax           `json:"tax,omitempty"`
	RequestCode      string          `json:"request_code"`
	Status           string          `json:"status"`
	Paid             bool            `json:"paid"`
	PaidAt           *string         `json:"paid_at,omitempty"`
	Metadata         *types.Metadata `json:"metadata,omitempty"`
	Notifications    []any           `json:"notifications,omitempty"`
	OfflineReference *string         `json:"offline_reference,omitempty"`
	Customer         *Customer       `json:"customer,omitempty"`
	CreatedAt        string          `json:"created_at"`
	UpdatedAt        string          `json:"updated_at,omitempty"`
	Transactions     []Transaction   `json:"transactions,omitempty"`
	SplitCode        *string         `json:"split_code,omitempty"`
	Integration      int             `json:"integration,omitempty"`
	Archived         bool            `json:"archived,omitempty"`
}

// CreatePaymentRequestRequest represents the request to create a payment request
type CreatePaymentRequestRequest struct {
	Customer         string     `json:"customer"`
	Amount           *int       `json:"amount,omitempty"`
	DueDate          string     `json:"due_date,omitempty"`
	Description      string     `json:"description,omitempty"`
	LineItems        []LineItem `json:"line_items,omitempty"`
	Tax              []Tax      `json:"tax,omitempty"`
	Currency         string     `json:"currency,omitempty"`
	SendNotification *bool      `json:"send_notification,omitempty"`
	Draft            *bool      `json:"draft,omitempty"`
	HasInvoice       *bool      `json:"has_invoice,omitempty"`
	InvoiceNumber    *int       `json:"invoice_number,omitempty"`
	SplitCode        string     `json:"split_code,omitempty"`
}

// UpdatePaymentRequestRequest represents the request to update a payment request
type UpdatePaymentRequestRequest struct {
	Customer         string     `json:"customer,omitempty"`
	Amount           *int       `json:"amount,omitempty"`
	DueDate          string     `json:"due_date,omitempty"`
	Description      string     `json:"description,omitempty"`
	LineItems        []LineItem `json:"line_items,omitempty"`
	Tax              []Tax      `json:"tax,omitempty"`
	Currency         string     `json:"currency,omitempty"`
	SendNotification *bool      `json:"send_notification,omitempty"`
	Draft            *bool      `json:"draft,omitempty"`
	InvoiceNumber    *int       `json:"invoice_number,omitempty"`
	SplitCode        string     `json:"split_code,omitempty"`
}

// ListPaymentRequestsRequest represents the request to list payment requests
type ListPaymentRequestsRequest struct {
	PerPage        int    `json:"perPage,omitempty"`
	Page           int    `json:"page,omitempty"`
	Customer       string `json:"customer,omitempty"`
	Status         string `json:"status,omitempty"`
	Currency       string `json:"currency,omitempty"`
	IncludeArchive string `json:"include_archive,omitempty"`
	From           string `json:"from,omitempty"`
	To             string `json:"to,omitempty"`
}

// FinalizePaymentRequestRequest represents the request to finalize a payment request
type FinalizePaymentRequestRequest struct {
	SendNotification *bool `json:"send_notification,omitempty"`
}

// PaymentRequestTotal represents payment request totals by status and currency
type PaymentRequestTotal struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

// PaymentRequestTotals represents the totals data structure
type PaymentRequestTotals struct {
	Pending    []PaymentRequestTotal `json:"pending"`
	Successful []PaymentRequestTotal `json:"successful"`
	Total      []PaymentRequestTotal `json:"total"`
}

// CreatePaymentRequestResponse represents the response from creating a payment request
type CreatePaymentRequestResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    PaymentRequest `json:"data"`
}

// ListPaymentRequestsResponse represents the response from listing payment requests
type ListPaymentRequestsResponse struct {
	Status  bool             `json:"status"`
	Message string           `json:"message"`
	Data    []PaymentRequest `json:"data"`
	Meta    *types.Meta      `json:"meta,omitempty"`
}

// FetchPaymentRequestResponse represents the response from fetching a payment request
type FetchPaymentRequestResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    PaymentRequest `json:"data"`
}

// VerifyPaymentRequestResponse represents the response from verifying a payment request
type VerifyPaymentRequestResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    PaymentRequest `json:"data"`
}

// UpdatePaymentRequestResponse represents the response from updating a payment request
type UpdatePaymentRequestResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    PaymentRequest `json:"data"`
}

// FinalizePaymentRequestResponse represents the response from finalizing a payment request
type FinalizePaymentRequestResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    PaymentRequest `json:"data"`
}

// ArchivePaymentRequestResponse represents the response from archiving a payment request
type ArchivePaymentRequestResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// SendNotificationResponse represents the response from sending a notification
type SendNotificationResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// PaymentRequestTotalsResponse represents the response from getting payment request totals
type PaymentRequestTotalsResponse struct {
	Status  bool                 `json:"status"`
	Message string               `json:"message"`
	Data    PaymentRequestTotals `json:"data"`
}
