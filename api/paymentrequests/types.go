package paymentrequests

import "github.com/huysamen/paystack-go/types"

// LineItem represents a line item in a payment request
type LineItem struct {
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	Quantity int    `json:"quantity,omitempty"`
}

// Tax represents tax information for a payment request
type Tax struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

// PaymentRequest represents a payment request
type PaymentRequest struct {
	ID                 int                  `json:"id"`
	Domain             string               `json:"domain"`
	Amount             int                  `json:"amount"`
	Currency           string               `json:"currency"`
	DueDate            string               `json:"due_date"`
	HasInvoice         bool                 `json:"has_invoice"`
	InvoiceNumber      *int                 `json:"invoice_number"`
	Description        string               `json:"description"`
	LineItems          []LineItem           `json:"line_items"`
	Tax                []Tax                `json:"tax"`
	RequestCode        string               `json:"request_code"`
	Status             string               `json:"status"`
	Paid               bool                 `json:"paid"`
	PaidAt             *types.DateTime      `json:"paid_at"`
	Metadata           types.Metadata       `json:"metadata"`
	Notifications      []Notification       `json:"notifications"`
	OfflineReference   string               `json:"offline_reference"`
	Customer           types.Customer       `json:"customer"`
	CreatedAt          types.DateTime       `json:"created_at"`
	UpdatedAt          types.DateTime       `json:"updated_at"`
	PendingAmount      int                  `json:"pending_amount"`
	Split              *types.Split         `json:"split"`
	Integration        int                  `json:"integration"`
	SplitCode          string               `json:"split_code"`
	Archived           bool                 `json:"archived"`
	Source             *types.Source        `json:"source"`
	Invoice            *types.Invoice       `json:"invoice"`
	Plan               *types.Plan          `json:"plan"`
	Transaction        *types.Transaction   `json:"transaction"`
}

// Notification represents a notification for a payment request
type Notification struct {
	SentAt  *types.DateTime `json:"sent_at"`
	Channel string          `json:"channel"`
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
