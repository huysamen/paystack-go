package webhook

import (
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

type DisputeHistoryEntry struct {
	Status    data.String `json:"status"`
	By        data.String `json:"by"`
	CreatedAt data.Time   `json:"created_at"`
}

type DisputeMessage struct {
	Sender    data.String `json:"sender"`
	Body      data.String `json:"body"`
	CreatedAt data.Time   `json:"created_at"`
}

type ChargeSuccessEvent struct {
	ID                 data.Int             `json:"id"`
	Domain             data.String          `json:"domain"`
	Status             data.String          `json:"status"`
	Reference          data.String          `json:"reference"`
	ReceiptNumber      data.NullString      `json:"receipt_number"`
	Amount             data.Int             `json:"amount"`
	Message            data.NullString      `json:"message"`
	GatewayResponse    data.String          `json:"gateway_response"`
	PaidAt             data.NullTime        `json:"paid_at"`
	CreatedAt          data.Time            `json:"created_at"`
	Channel            data.String          `json:"channel"`
	Currency           data.String          `json:"currency"`
	IPAddress          data.String          `json:"ip_address"`
	Metadata           types.Metadata       `json:"metadata"`
	Log                *any                 `json:"log"`
	Fees               data.Int             `json:"fees"`
	FeesSplit          *any                 `json:"fees_split"`
	Authorization      *types.Authorization `json:"authorization"`
	Customer           *types.Customer      `json:"customer"`
	Plan               *any                 `json:"plan"`
	Split              *any                 `json:"split"`
	OrderID            *any                 `json:"order_id"`
	PaidAt2            data.NullTime        `json:"paidAt"`
	RequestedAmount    data.Int             `json:"requested_amount"`
	PosTransactionData *any                 `json:"pos_transaction_data"`
	Source             *any                 `json:"source"`
	FeesBreakdown      *any                 `json:"fees_breakdown"`
}

type ChargeDisputeEvent struct {
	ID                   data.Int              `json:"id"`
	RefundAmount         data.Int              `json:"refund_amount"`
	Currency             data.String           `json:"currency"`
	Status               data.String           `json:"status"`
	Resolution           data.NullString       `json:"resolution"`
	Domain               data.String           `json:"domain"`
	Transaction          *types.Transaction    `json:"transaction"`
	TransactionReference data.NullString       `json:"transaction_reference"`
	Category             data.String           `json:"category"`
	Customer             *types.Customer       `json:"customer"`
	BIN                  data.String           `json:"bin"`
	Last4                data.String           `json:"last4"`
	DueAt                data.NullTime         `json:"dueAt"`
	ResolvedAt           data.NullTime         `json:"resolvedAt"`
	Evidence             *any                  `json:"evidence"`
	Attachments          *any                  `json:"attachments"`
	Note                 data.NullString       `json:"note"`
	History              []DisputeHistoryEntry `json:"history"`
	Messages             []DisputeMessage      `json:"messages"`
	CreatedAt            data.Time             `json:"created_at"`
	UpdatedAt            data.Time             `json:"updated_at"`
}

type DedicatedAccountEvent struct {
	ID             data.Int        `json:"id"`
	Domain         data.String     `json:"domain"`
	Status         data.String     `json:"status"`
	AccountName    data.String     `json:"account_name"`
	AccountNumber  data.String     `json:"account_number"`
	BankCode       data.String     `json:"bank_code"`
	BankName       data.String     `json:"bank_name"`
	Customer       *types.Customer `json:"customer"`
	CustomerCode   data.String     `json:"customer_code"`
	ExpiresAt      data.NullTime   `json:"expires_at"`
	CreatedAt      data.Time       `json:"created_at"`
	UpdatedAt      data.Time       `json:"updated_at"`
	Identification types.Metadata  `json:"identification"`
}

type InvoicePaymentFailedEvent struct {
	ID          data.Int        `json:"id"`
	Domain      data.String     `json:"domain"`
	InvoiceCode data.String     `json:"invoice_code"`
	Amount      data.Int        `json:"amount"`
	PeriodStart data.Time       `json:"period_start"`
	PeriodEnd   data.Time       `json:"period_end"`
	Status      data.String     `json:"status"`
	Paid        data.Bool       `json:"paid"`
	Currency    data.String     `json:"currency"`
	Customer    *types.Customer `json:"customer"`
	Transaction *any            `json:"transaction"`
	CreatedAt   data.Time       `json:"created_at"`
	UpdatedAt   data.Time       `json:"updated_at"`
}

type InvoiceUpdateEvent struct {
	ID          data.Int        `json:"id"`
	Domain      data.String     `json:"domain"`
	InvoiceCode data.String     `json:"invoice_code"`
	Amount      data.Int        `json:"amount"`
	PeriodStart data.Time       `json:"period_start"`
	PeriodEnd   data.Time       `json:"period_end"`
	Status      data.String     `json:"status"`
	Paid        data.Bool       `json:"paid"`
	Currency    data.String     `json:"currency"`
	Customer    *types.Customer `json:"customer"`
	Transaction *any            `json:"transaction"`
	CreatedAt   data.Time       `json:"created_at"`
	UpdatedAt   data.Time       `json:"updated_at"`
}

type PaymentRequestEvent struct {
	ID               data.Int        `json:"id"`
	Domain           data.String     `json:"domain"`
	Amount           data.Int        `json:"amount"`
	Currency         data.String     `json:"currency"`
	DueDate          data.NullTime   `json:"due_date"`
	HasInvoice       data.Bool       `json:"has_invoice"`
	InvoiceNumber    data.NullString `json:"invoice_number"`
	Description      data.String     `json:"description"`
	PDF_URL          data.NullString `json:"pdf_url"`
	LineItems        []any           `json:"line_items"`
	Tax              []any           `json:"tax"`
	RequestCode      data.String     `json:"request_code"`
	Status           data.String     `json:"status"`
	Paid             data.Bool       `json:"paid"`
	PaidAt           data.NullTime   `json:"paid_at"`
	Metadata         types.Metadata  `json:"metadata"`
	Notifications    []any           `json:"notifications"`
	OfflineReference data.NullString `json:"offline_reference"`
	// In webhooks this can be ID or object; allow raw JSON in types.Customer by using Metadata workaround
	Customer  types.Metadata `json:"customer"`
	CreatedAt data.Time      `json:"created_at"`
}

type RefundFailedEvent struct {
	ID             data.Int           `json:"id"`
	Integration    data.Int           `json:"integration"`
	Domain         data.String        `json:"domain"`
	Transaction    *types.Transaction `json:"transaction"`
	Dispute        *any               `json:"dispute"`
	Amount         data.Int           `json:"amount"`
	DeductedAmount data.Int           `json:"deducted_amount"`
	FullyDeducted  data.Bool          `json:"fully_deducted"`
	Currency       data.String        `json:"currency"`
	Status         data.String        `json:"status"`
	RefundedBy     data.String        `json:"refunded_by"`
	RefundedAt     data.NullTime      `json:"refunded_at"`
	ExpectedAt     data.Time          `json:"expected_at"`
	CreatedAt      data.Time          `json:"created_at"`
	UpdatedAt      data.Time          `json:"updated_at"`
}

type RefundPendingEvent struct {
	ID             data.Int           `json:"id"`
	Integration    data.Int           `json:"integration"`
	Domain         data.String        `json:"domain"`
	Transaction    *types.Transaction `json:"transaction"`
	Dispute        *any               `json:"dispute"`
	Amount         data.Int           `json:"amount"`
	DeductedAmount data.Int           `json:"deducted_amount"`
	FullyDeducted  data.Bool          `json:"fully_deducted"`
	Currency       data.String        `json:"currency"`
	Status         data.String        `json:"status"`
	RefundedBy     data.String        `json:"refunded_by"`
	RefundedAt     data.NullTime      `json:"refunded_at"`
	ExpectedAt     data.Time          `json:"expected_at"`
	CreatedAt      data.Time          `json:"created_at"`
	UpdatedAt      data.Time          `json:"updated_at"`
}

type SubscriptionDisableEvent struct {
	ID               data.Int             `json:"id"`
	Domain           data.String          `json:"domain"`
	Status           data.String          `json:"status"`
	SubscriptionCode data.String          `json:"subscription_code"`
	EmailToken       data.String          `json:"email_token"`
	Amount           data.Int             `json:"amount"`
	CronExpression   data.String          `json:"cron_expression"`
	NextPaymentDate  data.NullTime        `json:"next_payment_date"`
	OpenInvoice      data.NullString      `json:"open_invoice"`
	Customer         *types.Customer      `json:"customer"`
	Plan             *any                 `json:"plan"`
	Authorization    *types.Authorization `json:"authorization"`
	Invoices         []any                `json:"invoices"`
	CreatedAt        data.Time            `json:"created_at"`
	UpdatedAt        data.Time            `json:"updated_at"`
}

type SubscriptionNotRenewEvent struct {
	ID               data.Int             `json:"id"`
	Domain           data.String          `json:"domain"`
	Status           data.String          `json:"status"`
	SubscriptionCode data.String          `json:"subscription_code"`
	EmailToken       data.String          `json:"email_token"`
	Amount           data.Int             `json:"amount"`
	CronExpression   data.String          `json:"cron_expression"`
	NextPaymentDate  data.NullTime        `json:"next_payment_date"`
	OpenInvoice      data.NullString      `json:"open_invoice"`
	Customer         *types.Customer      `json:"customer"`
	Plan             *any                 `json:"plan"`
	Authorization    *types.Authorization `json:"authorization"`
	Invoices         []any                `json:"invoices"`
	CreatedAt        data.Time            `json:"created_at"`
	UpdatedAt        data.Time            `json:"updated_at"`
}

type SubscriptionExpiringCardsEvent struct {
	// Expiring cards payload is an array of entries; we accept Metadata for flexibility
	Entries types.Metadata `json:"-"`
}

type CustomerIdentificationFailedEvent struct {
	ID             data.Int               `json:"id"`
	CustomerID     data.String            `json:"customer_id"`
	CustomerCode   data.String            `json:"customer_code"`
	Email          data.String            `json:"email"`
	Identification CustomerIdentification `json:"identification"`
	CreatedAt      data.Time              `json:"created_at"`
	UpdatedAt      data.Time              `json:"updated_at"`
}

type CustomerIdentificationSuccessEvent struct {
	ID             data.Int               `json:"id"`
	CustomerID     data.String            `json:"customer_id"`
	CustomerCode   data.String            `json:"customer_code"`
	Email          data.String            `json:"email"`
	Identification CustomerIdentification `json:"identification"`
	CreatedAt      data.Time              `json:"created_at"`
	UpdatedAt      data.Time              `json:"updated_at"`
}

type CustomerIdentification struct {
	Type          data.String `json:"type"`
	Value         data.String `json:"value"`
	Country       data.String `json:"country"`
	BVN           data.String `json:"bvn"`
	AccountNumber data.String `json:"account_number"`
	BankCode      data.String `json:"bank_code"`
}

type TransferSuccessEvent struct {
	Amount        data.Int        `json:"amount"`
	Currency      data.String     `json:"currency"`
	Domain        data.String     `json:"domain"`
	Failures      *any            `json:"failures"`
	ID            data.Int        `json:"id"`
	Integration   types.Metadata  `json:"integration"`
	Reason        data.String     `json:"reason"`
	Reference     data.String     `json:"reference"`
	Source        data.String     `json:"source"`
	SourceDetails *any            `json:"source_details"`
	Status        data.String     `json:"status"`
	TitanCode     data.NullString `json:"titan_code"`
	TransferCode  data.String     `json:"transfer_code"`
	TransferredAt data.NullTime   `json:"transferred_at"`
	Recipient     types.Recipient `json:"recipient"`
	Session       *any            `json:"session"`
	CreatedAt     data.Time       `json:"created_at"`
	UpdatedAt     data.Time       `json:"updated_at"`
}

type TransferFailedEvent struct {
	Amount        data.Int        `json:"amount"`
	Currency      data.String     `json:"currency"`
	Domain        data.String     `json:"domain"`
	Failures      *any            `json:"failures"`
	ID            data.Int        `json:"id"`
	Integration   types.Metadata  `json:"integration"`
	Reason        data.String     `json:"reason"`
	Reference     data.String     `json:"reference"`
	Source        data.String     `json:"source"`
	SourceDetails *any            `json:"source_details"`
	Status        data.String     `json:"status"`
	TitanCode     data.NullString `json:"titan_code"`
	TransferCode  data.String     `json:"transfer_code"`
	TransferredAt data.NullTime   `json:"transferred_at"`
	Recipient     types.Recipient `json:"recipient"`
	Session       *any            `json:"session"`
	CreatedAt     data.Time       `json:"created_at"`
	UpdatedAt     data.Time       `json:"updated_at"`
}

type TransferReversedEvent struct {
	Amount        data.Int        `json:"amount"`
	Currency      data.String     `json:"currency"`
	Domain        data.String     `json:"domain"`
	Failures      *any            `json:"failures"`
	ID            data.Int        `json:"id"`
	Integration   types.Metadata  `json:"integration"`
	Reason        data.String     `json:"reason"`
	Reference     data.String     `json:"reference"`
	Source        data.String     `json:"source"`
	SourceDetails *any            `json:"source_details"`
	Status        data.String     `json:"status"`
	TitanCode     data.NullString `json:"titan_code"`
	TransferCode  data.String     `json:"transfer_code"`
	TransferredAt data.NullTime   `json:"transferred_at"`
	Recipient     types.Recipient `json:"recipient"`
	Session       *any            `json:"session"`
	CreatedAt     data.Time       `json:"created_at"`
	UpdatedAt     data.Time       `json:"updated_at"`
}

type SubscriptionCreateEvent struct {
	Domain           data.String          `json:"domain"`
	Status           data.String          `json:"status"`
	SubscriptionCode data.String          `json:"subscription_code"`
	Amount           data.Int             `json:"amount"`
	CronExpression   data.String          `json:"cron_expression"`
	NextPaymentDate  data.Time            `json:"next_payment_date"`
	OpenInvoice      data.NullString      `json:"open_invoice"`
	Integration      data.Int             `json:"integration"`
	Plan             *any                 `json:"plan"`
	Authorization    *types.Authorization `json:"authorization"`
	Customer         *types.Customer      `json:"customer"`
	ID               data.Int             `json:"id"`
	CreatedAt        data.Time            `json:"created_at"`
	UpdatedAt        data.Time            `json:"updated_at"`
}

type InvoiceCreateEvent struct {
	Domain       data.String     `json:"domain"`
	InvoiceCode  data.String     `json:"invoice_code"`
	Amount       data.Int        `json:"amount"`
	PeriodStart  data.Time       `json:"period_start"`
	PeriodEnd    data.Time       `json:"period_end"`
	Status       data.String     `json:"status"`
	Paid         data.Bool       `json:"paid"`
	Currency     data.String     `json:"currency"`
	Customer     *types.Customer `json:"customer"`
	Subscription *any            `json:"subscription"`
	ID           data.Int        `json:"id"`
	CreatedAt    data.Time       `json:"created_at"`
	UpdatedAt    data.Time       `json:"updated_at"`
}

type RefundProcessedEvent struct {
	ID             data.Int           `json:"id"`
	Integration    data.Int           `json:"integration"`
	Domain         data.String        `json:"domain"`
	Transaction    *types.Transaction `json:"transaction"`
	Dispute        *any               `json:"dispute"`
	Amount         data.Int           `json:"amount"`
	DeductedAmount data.Int           `json:"deducted_amount"`
	FullyDeducted  data.Bool          `json:"fully_deducted"`
	Currency       data.String        `json:"currency"`
	Status         data.String        `json:"status"`
	RefundedBy     data.String        `json:"refunded_by"`
	RefundedAt     data.NullTime      `json:"refunded_at"`
	ExpectedAt     data.Time          `json:"expected_at"`
	CreatedAt      data.Time          `json:"created_at"`
	UpdatedAt      data.Time          `json:"updated_at"`
}
