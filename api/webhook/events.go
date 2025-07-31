package webhook

import (
	"time"

	transfer_recipients "github.com/huysamen/paystack-go/api/transfer-recipients"
	"github.com/huysamen/paystack-go/types"
)

// Event data structures for specific webhook events

// ChargeSuccessEvent represents the data for charge.success events
type ChargeSuccessEvent struct {
	ID                 int64                `json:"id"`
	Domain             string               `json:"domain"`
	Status             string               `json:"status"`
	Reference          string               `json:"reference"`
	Amount             int64                `json:"amount"`
	Message            *string              `json:"message"`
	GatewayResponse    string               `json:"gateway_response"`
	PaidAt             *time.Time           `json:"paid_at"`
	CreatedAt          time.Time            `json:"created_at"`
	Channel            string               `json:"channel"`
	Currency           string               `json:"currency"`
	IPAddress          *string              `json:"ip_address"`
	Metadata           map[string]any       `json:"metadata"`
	Log                *types.Log           `json:"log"`
	Fees               *int64               `json:"fees"`
	FeesSplit          *any                 `json:"fees_split"`
	Authorization      *types.Authorization `json:"authorization"`
	Customer           *types.Customer      `json:"customer"`
	Plan               *types.Plan          `json:"plan"`
	Split              *any                 `json:"split"`
	OrderID            *any                 `json:"order_id"`
	PaidAmount         *int64               `json:"paidAmount"`
	RequestedAmount    *int64               `json:"requested_amount"`
	POSTransactionData *any                 `json:"pos_transaction_data"`
	Source             *any                 `json:"source"`
	FeesBreakdown      *any                 `json:"fees_breakdown"`
}

// CustomerIdentificationFailedEvent represents data for customeridentification.failed events
type CustomerIdentificationFailedEvent struct {
	CustomerID     int64                  `json:"customer_id"`
	CustomerCode   string                 `json:"customer_code"`
	Email          string                 `json:"email"`
	Identification CustomerIdentification `json:"identification"`
	Reason         string                 `json:"reason"`
}

// CustomerIdentificationSuccessEvent represents data for customeridentification.success events
type CustomerIdentificationSuccessEvent struct {
	CustomerID     int64                  `json:"customer_id"`
	CustomerCode   string                 `json:"customer_code"`
	Email          string                 `json:"email"`
	Identification CustomerIdentification `json:"identification"`
}

// TransferSuccessEvent represents data for transfer.success events
type TransferSuccessEvent struct {
	Amount        int64                                  `json:"amount"`
	Currency      string                                 `json:"currency"`
	Domain        string                                 `json:"domain"`
	Failures      *any                                   `json:"failures"`
	ID            int64                                  `json:"id"`
	Integration   int64                                  `json:"integration"`
	Reason        string                                 `json:"reason"`
	Reference     string                                 `json:"reference"`
	Source        string                                 `json:"source"`
	SourceDetails *any                                   `json:"source_details"`
	Status        string                                 `json:"status"`
	Titan         *any                                   `json:"titan"`
	TransferCode  string                                 `json:"transfer_code"`
	TransferredAt *time.Time                             `json:"transferred_at"`
	CreatedAt     time.Time                              `json:"created_at"`
	UpdatedAt     time.Time                              `json:"updated_at"`
	Recipient     *transfer_recipients.TransferRecipient `json:"recipient"`
}

// TransferFailedEvent represents data for transfer.failed events
type TransferFailedEvent struct {
	Amount        int64                                  `json:"amount"`
	Currency      string                                 `json:"currency"`
	Domain        string                                 `json:"domain"`
	Failures      *any                                   `json:"failures"`
	ID            int64                                  `json:"id"`
	Integration   int64                                  `json:"integration"`
	Reason        string                                 `json:"reason"`
	Reference     string                                 `json:"reference"`
	Source        string                                 `json:"source"`
	SourceDetails *any                                   `json:"source_details"`
	Status        string                                 `json:"status"`
	TransferCode  string                                 `json:"transfer_code"`
	CreatedAt     time.Time                              `json:"created_at"`
	UpdatedAt     time.Time                              `json:"updated_at"`
	Recipient     *transfer_recipients.TransferRecipient `json:"recipient"`
}

// TransferReversedEvent represents data for transfer.reversed events
type TransferReversedEvent struct {
	Amount        int64                                  `json:"amount"`
	Currency      string                                 `json:"currency"`
	Domain        string                                 `json:"domain"`
	Failures      *any                                   `json:"failures"`
	ID            int64                                  `json:"id"`
	Integration   int64                                  `json:"integration"`
	Reason        string                                 `json:"reason"`
	Reference     string                                 `json:"reference"`
	Source        string                                 `json:"source"`
	SourceDetails *any                                   `json:"source_details"`
	Status        string                                 `json:"status"`
	TransferCode  string                                 `json:"transfer_code"`
	ReversedAt    *time.Time                             `json:"reversed_at"`
	CreatedAt     time.Time                              `json:"created_at"`
	UpdatedAt     time.Time                              `json:"updated_at"`
	Recipient     *transfer_recipients.TransferRecipient `json:"recipient"`
}

// SubscriptionCreateEvent represents data for subscription.create events
type SubscriptionCreateEvent struct {
	ID               int64                `json:"id"`
	Domain           string               `json:"domain"`
	Status           string               `json:"status"`
	SubscriptionCode string               `json:"subscription_code"`
	EmailToken       string               `json:"email_token"`
	Amount           int64                `json:"amount"`
	CronExpression   string               `json:"cron_expression"`
	NextPaymentDate  time.Time            `json:"next_payment_date"`
	OpenInvoice      *string              `json:"open_invoice"`
	Integration      int64                `json:"integration"`
	Plan             *types.Plan          `json:"plan"`
	Authorization    *types.Authorization `json:"authorization"`
	Customer         *types.Customer      `json:"customer"`
	CreatedAt        time.Time            `json:"created_at"`
	UpdatedAt        time.Time            `json:"updated_at"`
}

// InvoiceCreateEvent represents data for invoice.create events
type InvoiceCreateEvent struct {
	ID            int64                `json:"id"`
	Domain        string               `json:"domain"`
	InvoiceCode   string               `json:"invoice_code"`
	Amount        int64                `json:"amount"`
	PeriodStart   time.Time            `json:"period_start"`
	PeriodEnd     time.Time            `json:"period_end"`
	Status        string               `json:"status"`
	Paid          bool                 `json:"paid"`
	PaidAt        *time.Time           `json:"paid_at"`
	Description   *string              `json:"description"`
	Authorization *types.Authorization `json:"authorization"`
	Subscription  *types.Subscription  `json:"subscription"`
	Customer      *types.Customer      `json:"customer"`
	Transaction   *any                 `json:"transaction"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
}

// RefundProcessedEvent represents data for refund.processed events
type RefundProcessedEvent struct {
	ID            int64     `json:"id"`
	Integration   int64     `json:"integration"`
	Domain        string    `json:"domain"`
	Transaction   int64     `json:"transaction"`
	Dispute       *int64    `json:"dispute"`
	Amount        int64     `json:"amount"`
	Currency      string    `json:"currency"`
	Channel       string    `json:"channel"`
	FullyDeducted bool      `json:"fully_deducted"`
	RefundedBy    string    `json:"refunded_by"`
	RefundedAt    time.Time `json:"refunded_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// CustomerIdentification represents customer identification details
// This struct is specific to webhook events and doesn't exist in the main types package
type CustomerIdentification struct {
	Country       string `json:"country"`
	Type          string `json:"type"`
	BVN           string `json:"bvn"`
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
}
