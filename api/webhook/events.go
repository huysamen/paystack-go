package webhook

import (
	"time"

	"github.com/huysamen/paystack-go/types"
)

type DisputeHistoryEntry struct {
	Status    string    `json:"status"`
	By        string    `json:"by"`
	CreatedAt time.Time `json:"created_at"`
}

type DisputeMessage struct {
	Sender    string    `json:"sender"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type ChargeSuccessEvent struct {
	ID                 int64                `json:"id"`
	Domain             string               `json:"domain"`
	Status             string               `json:"status"`
	Reference          string               `json:"reference"`
	ReceiptNumber      *string              `json:"receipt_number"`
	Amount             int64                `json:"amount"`
	Message            *string              `json:"message"`
	GatewayResponse    string               `json:"gateway_response"`
	PaidAt             *time.Time           `json:"paid_at"`
	CreatedAt          time.Time            `json:"created_at"`
	Channel            string               `json:"channel"`
	Currency           string               `json:"currency"`
	IPAddress          string               `json:"ip_address"`
	Metadata           types.Metadata       `json:"metadata"`
	Log                *any                 `json:"log"`
	Fees               int64                `json:"fees"`
	FeesSplit          *any                 `json:"fees_split"`
	Authorization      *types.Authorization `json:"authorization"`
	Customer           *types.Customer      `json:"customer"`
	Plan               *any                 `json:"plan"`
	Split              *any                 `json:"split"`
	OrderID            *any                 `json:"order_id"`
	PaidAt2            *time.Time           `json:"paidAt"`
	RequestedAmount    int64                `json:"requested_amount"`
	PosTransactionData *any                 `json:"pos_transaction_data"`
	Source             *any                 `json:"source"`
	FeesBreakdown      *any                 `json:"fees_breakdown"`
}

type ChargeDisputeEvent struct {
	ID                   int64                 `json:"id"`
	RefundAmount         int64                 `json:"refund_amount"`
	Currency             string                `json:"currency"`
	Status               string                `json:"status"`
	Resolution           *string               `json:"resolution"`
	Domain               string                `json:"domain"`
	Transaction          *types.Transaction    `json:"transaction"`
	TransactionReference *string               `json:"transaction_reference"`
	Category             string                `json:"category"`
	Customer             *types.Customer       `json:"customer"`
	BIN                  string                `json:"bin"`
	Last4                string                `json:"last4"`
	DueAt                *time.Time            `json:"dueAt"`
	ResolvedAt           *time.Time            `json:"resolvedAt"`
	Evidence             *any                  `json:"evidence"`
	Attachments          *any                  `json:"attachments"`
	Note                 *string               `json:"note"`
	History              []DisputeHistoryEntry `json:"history"`
	Messages             []DisputeMessage      `json:"messages"`
	CreatedAt            time.Time             `json:"created_at"`
	UpdatedAt            time.Time             `json:"updated_at"`
}

type DedicatedAccountEvent struct {
	ID            int64           `json:"id"`
	Domain        string          `json:"domain"`
	Status        string          `json:"status"`
	AccountName   string          `json:"account_name"`
	AccountNumber string          `json:"account_number"`
	BankCode      string          `json:"bank_code"`
	BankName      string          `json:"bank_name"`
	Customer      *types.Customer `json:"customer"`
	CustomerCode  string          `json:"customer_code"`
	ExpiresAt     *time.Time      `json:"expires_at"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

type InvoicePaymentFailedEvent struct {
	ID          int64           `json:"id"`
	Domain      string          `json:"domain"`
	InvoiceCode string          `json:"invoice_code"`
	Amount      int64           `json:"amount"`
	PeriodStart time.Time       `json:"period_start"`
	PeriodEnd   time.Time       `json:"period_end"`
	Status      string          `json:"status"`
	Paid        bool            `json:"paid"`
	Currency    string          `json:"currency"`
	Customer    *types.Customer `json:"customer"`
	Transaction *any            `json:"transaction"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type InvoiceUpdateEvent struct {
	ID          int64           `json:"id"`
	Domain      string          `json:"domain"`
	InvoiceCode string          `json:"invoice_code"`
	Amount      int64           `json:"amount"`
	PeriodStart time.Time       `json:"period_start"`
	PeriodEnd   time.Time       `json:"period_end"`
	Status      string          `json:"status"`
	Paid        bool            `json:"paid"`
	Currency    string          `json:"currency"`
	Customer    *types.Customer `json:"customer"`
	Transaction *any            `json:"transaction"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type PaymentRequestEvent struct {
	ID               int64           `json:"id"`
	Domain           string          `json:"domain"`
	Amount           int64           `json:"amount"`
	Currency         string          `json:"currency"`
	DueDate          *time.Time      `json:"due_date"`
	HasInvoice       bool            `json:"has_invoice"`
	InvoiceNumber    *string         `json:"invoice_number"`
	Description      string          `json:"description"`
	PDF_URL          *string         `json:"pdf_url"`
	LineItems        []any           `json:"line_items"`
	Tax              []any           `json:"tax"`
	RequestCode      string          `json:"request_code"`
	Status           string          `json:"status"`
	Paid             bool            `json:"paid"`
	PaidAt           *time.Time      `json:"paid_at"`
	Metadata         types.Metadata  `json:"metadata"`
	Notifications    []any           `json:"notifications"`
	OfflineReference *string         `json:"offline_reference"`
	Customer         *types.Customer `json:"customer"`
	CreatedAt        time.Time       `json:"created_at"`
}

type RefundFailedEvent struct {
	ID             int64              `json:"id"`
	Integration    int64              `json:"integration"`
	Domain         string             `json:"domain"`
	Transaction    *types.Transaction `json:"transaction"`
	Dispute        *any               `json:"dispute"`
	Amount         int64              `json:"amount"`
	DeductedAmount int64              `json:"deducted_amount"`
	FullyDeducted  bool               `json:"fully_deducted"`
	Currency       string             `json:"currency"`
	Status         string             `json:"status"`
	RefundedBy     string             `json:"refunded_by"`
	RefundedAt     *time.Time         `json:"refunded_at"`
	ExpectedAt     time.Time          `json:"expected_at"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}

type RefundPendingEvent struct {
	ID             int64              `json:"id"`
	Integration    int64              `json:"integration"`
	Domain         string             `json:"domain"`
	Transaction    *types.Transaction `json:"transaction"`
	Dispute        *any               `json:"dispute"`
	Amount         int64              `json:"amount"`
	DeductedAmount int64              `json:"deducted_amount"`
	FullyDeducted  bool               `json:"fully_deducted"`
	Currency       string             `json:"currency"`
	Status         string             `json:"status"`
	RefundedBy     string             `json:"refunded_by"`
	RefundedAt     *time.Time         `json:"refunded_at"`
	ExpectedAt     time.Time          `json:"expected_at"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}

type SubscriptionDisableEvent struct {
	ID               int64                `json:"id"`
	Domain           string               `json:"domain"`
	Status           string               `json:"status"`
	SubscriptionCode string               `json:"subscription_code"`
	EmailToken       string               `json:"email_token"`
	Amount           int64                `json:"amount"`
	CronExpression   string               `json:"cron_expression"`
	NextPaymentDate  *time.Time           `json:"next_payment_date"`
	OpenInvoice      *string              `json:"open_invoice"`
	Customer         *types.Customer      `json:"customer"`
	Plan             *any                 `json:"plan"`
	Authorization    *types.Authorization `json:"authorization"`
	Invoices         []any                `json:"invoices"`
	CreatedAt        time.Time            `json:"created_at"`
	UpdatedAt        time.Time            `json:"updated_at"`
}

type SubscriptionNotRenewEvent struct {
	ID               int64                `json:"id"`
	Domain           string               `json:"domain"`
	Status           string               `json:"status"`
	SubscriptionCode string               `json:"subscription_code"`
	EmailToken       string               `json:"email_token"`
	Amount           int64                `json:"amount"`
	CronExpression   string               `json:"cron_expression"`
	NextPaymentDate  *time.Time           `json:"next_payment_date"`
	OpenInvoice      *string              `json:"open_invoice"`
	Customer         *types.Customer      `json:"customer"`
	Plan             *any                 `json:"plan"`
	Authorization    *types.Authorization `json:"authorization"`
	Invoices         []any                `json:"invoices"`
	CreatedAt        time.Time            `json:"created_at"`
	UpdatedAt        time.Time            `json:"updated_at"`
}

type SubscriptionExpiringCardsEvent struct {
	ID               int64                `json:"id"`
	Domain           string               `json:"domain"`
	Status           string               `json:"status"`
	SubscriptionCode string               `json:"subscription_code"`
	EmailToken       string               `json:"email_token"`
	Amount           int64                `json:"amount"`
	CronExpression   string               `json:"cron_expression"`
	NextPaymentDate  *time.Time           `json:"next_payment_date"`
	OpenInvoice      *string              `json:"open_invoice"`
	Customer         *types.Customer      `json:"customer"`
	Plan             *any                 `json:"plan"`
	Authorization    *types.Authorization `json:"authorization"`
	Invoices         []any                `json:"invoices"`
	CreatedAt        time.Time            `json:"created_at"`
	UpdatedAt        time.Time            `json:"updated_at"`
}

type CustomerIdentificationFailedEvent struct {
	ID             int64                  `json:"id"`
	CustomerID     string                 `json:"customer_id"`
	CustomerCode   string                 `json:"customer_code"`
	Email          string                 `json:"email"`
	Identification CustomerIdentification `json:"identification"`
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
}

type CustomerIdentificationSuccessEvent struct {
	ID             int64                  `json:"id"`
	CustomerID     string                 `json:"customer_id"`
	CustomerCode   string                 `json:"customer_code"`
	Email          string                 `json:"email"`
	Identification CustomerIdentification `json:"identification"`
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
}

type CustomerIdentification struct {
	Type          string `json:"type"`
	Value         string `json:"value"`
	Country       string `json:"country"`
	BVN           string `json:"bvn"`
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
}

type TransferSuccessEvent struct {
	Amount        int64           `json:"amount"`
	Currency      string          `json:"currency"`
	Domain        string          `json:"domain"`
	Failures      *any            `json:"failures"`
	ID            int64           `json:"id"`
	Integration   int64           `json:"integration"`
	Reason        string          `json:"reason"`
	Reference     string          `json:"reference"`
	Source        string          `json:"source"`
	SourceDetails *any            `json:"source_details"`
	Status        string          `json:"status"`
	TitanCode     *string         `json:"titan_code"`
	TransferCode  string          `json:"transfer_code"`
	TransferredAt *time.Time      `json:"transferred_at"`
	Recipient     types.Recipient `json:"recipient"`
	Session       *any            `json:"session"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

type TransferFailedEvent struct {
	Amount        int64           `json:"amount"`
	Currency      string          `json:"currency"`
	Domain        string          `json:"domain"`
	Failures      *any            `json:"failures"`
	ID            int64           `json:"id"`
	Integration   int64           `json:"integration"`
	Reason        string          `json:"reason"`
	Reference     string          `json:"reference"`
	Source        string          `json:"source"`
	SourceDetails *any            `json:"source_details"`
	Status        string          `json:"status"`
	TitanCode     *string         `json:"titan_code"`
	TransferCode  string          `json:"transfer_code"`
	TransferredAt *time.Time      `json:"transferred_at"`
	Recipient     types.Recipient `json:"recipient"`
	Session       *any            `json:"session"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

type TransferReversedEvent struct {
	Amount        int64           `json:"amount"`
	Currency      string          `json:"currency"`
	Domain        string          `json:"domain"`
	Failures      *any            `json:"failures"`
	ID            int64           `json:"id"`
	Integration   int64           `json:"integration"`
	Reason        string          `json:"reason"`
	Reference     string          `json:"reference"`
	Source        string          `json:"source"`
	SourceDetails *any            `json:"source_details"`
	Status        string          `json:"status"`
	TitanCode     *string         `json:"titan_code"`
	TransferCode  string          `json:"transfer_code"`
	TransferredAt *time.Time      `json:"transferred_at"`
	Recipient     types.Recipient `json:"recipient"`
	Session       *any            `json:"session"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

type SubscriptionCreateEvent struct {
	Domain           string               `json:"domain"`
	Status           string               `json:"status"`
	SubscriptionCode string               `json:"subscription_code"`
	Amount           int64                `json:"amount"`
	CronExpression   string               `json:"cron_expression"`
	NextPaymentDate  time.Time            `json:"next_payment_date"`
	OpenInvoice      *string              `json:"open_invoice"`
	Integration      int64                `json:"integration"`
	Plan             *any                 `json:"plan"`
	Authorization    *types.Authorization `json:"authorization"`
	Customer         *types.Customer      `json:"customer"`
	ID               int64                `json:"id"`
	CreatedAt        time.Time            `json:"created_at"`
	UpdatedAt        time.Time            `json:"updated_at"`
}

type InvoiceCreateEvent struct {
	Domain       string          `json:"domain"`
	InvoiceCode  string          `json:"invoice_code"`
	Amount       int64           `json:"amount"`
	PeriodStart  time.Time       `json:"period_start"`
	PeriodEnd    time.Time       `json:"period_end"`
	Status       string          `json:"status"`
	Paid         bool            `json:"paid"`
	Currency     string          `json:"currency"`
	Customer     *types.Customer `json:"customer"`
	Subscription *any            `json:"subscription"`
	ID           int64           `json:"id"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

type RefundProcessedEvent struct {
	ID             int64              `json:"id"`
	Integration    int64              `json:"integration"`
	Domain         string             `json:"domain"`
	Transaction    *types.Transaction `json:"transaction"`
	Dispute        *any               `json:"dispute"`
	Amount         int64              `json:"amount"`
	DeductedAmount int64              `json:"deducted_amount"`
	FullyDeducted  bool               `json:"fully_deducted"`
	Currency       string             `json:"currency"`
	Status         string             `json:"status"`
	RefundedBy     string             `json:"refunded_by"`
	RefundedAt     *time.Time         `json:"refunded_at"`
	ExpectedAt     time.Time          `json:"expected_at"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}
