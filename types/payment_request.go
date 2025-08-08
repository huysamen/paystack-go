package types
import "github.com/huysamen/paystack-go/types/data"
import "github.com/huysamen/paystack-go/enums"

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

// Notification represents a notification for a payment request
type Notification struct {
	SentAt  *data.MultiDateTime `json:"sent_at,omitempty"`
	Channel string    `json:"channel"`
}

// Source represents the source of a payment request
type Source struct {
	Type       string `json:"type"`
	Source     string `json:"source"`
	Identifier string `json:"identifier"`
}

// Invoice represents an invoice associated with a payment request
type Invoice struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	Amount int    `json:"amount"`
}

// PaymentRequest represents a payment request
type PaymentRequest struct {
	ID               int               `json:"id"`
	Domain           string            `json:"domain"`
	Amount           int               `json:"amount"`
	Currency         enums.Currency    `json:"currency"`
	DueDate          data.MultiDateTime          `json:"due_date"`
	HasInvoice       bool              `json:"has_invoice"`
	InvoiceNumber    *int              `json:"invoice_number"`
	Description      string            `json:"description"`
	LineItems        []LineItem        `json:"line_items"`
	Tax              []Tax             `json:"tax"`
	RequestCode      string            `json:"request_code"`
	Status           string            `json:"status"`
	Paid             bool              `json:"paid"`
	PaidAt           *data.MultiDateTime         `json:"paid_at"`
	Metadata         *Metadata         `json:"metadata"`
	Notifications    []Notification    `json:"notifications"`
	OfflineReference string            `json:"offline_reference"`
	Customer         Customer          `json:"customer"`
	CreatedAt        data.MultiDateTime          `json:"created_at"`
	UpdatedAt        data.MultiDateTime          `json:"updated_at"`
	PendingAmount    int               `json:"pending_amount"`
	Split            *TransactionSplit `json:"split"`
	Integration      int               `json:"integration"`
	SplitCode        string            `json:"split_code"`
	Archived         bool              `json:"archived"`
	Source           *Source           `json:"source"`
	Invoice          *Invoice          `json:"invoice"`
	Plan             *Plan             `json:"plan"`
	Transaction      *Transaction      `json:"transaction"`
}

// PaymentRequestTotals represents totals for payment requests
type PaymentRequestTotals struct {
	PendingPaymentRequests int `json:"pending"`
	SuccessfulPayments     int `json:"successful"`
	TotalPaymentRequests   int `json:"total"`
}
