package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// LineItem represents a line item in a payment request
type LineItem struct {
	Name     data.String `json:"name"`
	Amount   data.Int    `json:"amount"`
	Quantity data.Int    `json:"quantity,omitempty"`
}

// Tax represents tax information for a payment request
type Tax struct {
	Name   data.String `json:"name"`
	Amount data.Int    `json:"amount"`
}

// Notification represents a notification for a payment request
type Notification struct {
	SentAt  data.NullTime `json:"sent_at,omitempty"`
	Channel data.String   `json:"channel"`
}

// Source represents the source of a payment request
type Source struct {
	Type       data.String `json:"type"`
	Source     data.String `json:"source"`
	Identifier data.String `json:"identifier"`
}

// Invoice represents an invoice associated with a payment request
type Invoice struct {
	ID     data.Int    `json:"id"`
	Code   data.String `json:"code"`
	Amount data.Int    `json:"amount"`
}

// PaymentRequest represents a payment request
type PaymentRequest struct {
	ID               data.Int          `json:"id"`
	Domain           data.String       `json:"domain"`
	Amount           data.Int          `json:"amount"`
	Currency         enums.Currency    `json:"currency"`
	DueDate          data.Time         `json:"due_date"`
	HasInvoice       data.Bool         `json:"has_invoice"`
	InvoiceNumber    data.NullInt      `json:"invoice_number"`
	Description      data.String       `json:"description"`
	LineItems        []LineItem        `json:"line_items"`
	Tax              []Tax             `json:"tax"`
	RequestCode      data.String       `json:"request_code"`
	Status           data.String       `json:"status"`
	Paid             data.Bool         `json:"paid"`
	PaidAt           data.NullTime     `json:"paid_at"`
	Metadata         Metadata          `json:"metadata"`
	Notifications    []Notification    `json:"notifications"`
	OfflineReference data.String       `json:"offline_reference"`
	Customer         Customer          `json:"customer"`
	CreatedAt        data.Time         `json:"created_at"`
	UpdatedAt        data.Time         `json:"updated_at"`
	PendingAmount    data.Int          `json:"pending_amount"`
	Split            *TransactionSplit `json:"split"`
	Integration      data.Int          `json:"integration"`
	SplitCode        data.String       `json:"split_code"`
	Archived         data.Bool         `json:"archived"`
	Source           *Source           `json:"source"`
	Invoice          *Invoice          `json:"invoice"`
	Plan             *Plan             `json:"plan"`
	Transaction      *Transaction      `json:"transaction"`
}

// PaymentRequestTotals represents totals for payment requests
type PaymentRequestTotals struct {
	PendingPaymentRequests data.Int `json:"pending"`
	SuccessfulPayments     data.Int `json:"successful"`
	TotalPaymentRequests   data.Int `json:"total"`
}
