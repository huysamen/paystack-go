package types

import (
	"github.com/huysamen/paystack-go/enums"
)

// Plan represents a Paystack plan
type Plan struct {
	Domain                   string         `json:"domain"`
	Name                     string         `json:"name"`
	PlanCode                 string         `json:"plan_code"`
	Description              string         `json:"description"`
	Amount                   int            `json:"amount"`
	Interval                 enums.Interval `json:"interval"`
	InvoiceLimit             int            `json:"invoice_limit"`
	SendInvoices             bool           `json:"send_invoices"`
	SendSms                  bool           `json:"send_sms"`
	HostedPage               bool           `json:"hosted_page"`
	HostedPageURL            string         `json:"hosted_page_url"`
	HostedPageSummary        string         `json:"hosted_page_summary"`
	Currency                 enums.Currency `json:"currency"`
	Migrate                  bool           `json:"migrate"`
	IsDeleted                bool           `json:"is_deleted"`
	IsArchived               bool           `json:"is_archived"`
	ID                       uint64         `json:"id"`
	Integration              int            `json:"integration"`
	CreatedAt                DateTime       `json:"createdAt"`
	UpdatedAt                DateTime       `json:"updatedAt"`
	PagesCount               int            `json:"pages_count"`
	SubscribersCount         int            `json:"subscribers_count"`
	ActiveSubscriptionsCount int            `json:"active_subscriptions_count"`
	TotalRevenue             int            `json:"total_revenue"`
	Subscriptions            []Subscription `json:"subscriptions,omitempty"`
	Pages                    []Page         `json:"pages,omitempty"`
	Subscribers              []Subscriber   `json:"subscribers,omitempty"`
}

// Subscription represents a Paystack subscription
type Subscription struct {
	ID                uint64           `json:"id,omitempty"`
	Domain            string           `json:"domain"`
	Status            string           `json:"status"`
	SubscriptionCode  string           `json:"subscription_code"`
	EmailToken        string           `json:"email_token"`
	Amount            int              `json:"amount"`
	CronExpression    string           `json:"cron_expression"`
	NextPaymentDate   DateTime         `json:"next_payment_date"`
	OpenInvoice       *Invoice         `json:"open_invoice,omitempty"`
	CreatedAt         DateTime         `json:"createdAt"`
	CancelledAt       *DateTime        `json:"cancelledAt,omitempty"`
	Integration       int              `json:"integration"`
	Plan              Plan             `json:"plan"`
	Authorization     Authorization    `json:"authorization"`
	Customer          Customer         `json:"customer"`
	Invoices          []Invoice        `json:"invoices,omitempty"`
	InvoicesHistory   []InvoiceHistory `json:"invoices_history,omitempty"`
	InvoiceLimit      int              `json:"invoice_limit"`
	SplitCode         string           `json:"split_code,omitempty"`
	MostRecentInvoice *Invoice         `json:"most_recent_invoice,omitempty"`
	PaymentsCount     int              `json:"payments_count"`

	// Returned as part of fetching related resources.
	Quantity           *int      `json:"quantity,omitempty"`
	SuccessfulPayments *int      `json:"successful_payments,omitempty"`
	Start              *DateTime `json:"start,omitempty"`
}

// Subscriber represents a plan subscriber
type Subscriber struct {
	CustomerCode            string         `json:"customer_code"`
	CustomerFirstName       string         `json:"customer_first_name"`
	CustomerLastName        string         `json:"customer_last_name"`
	CustomerEmail           string         `json:"customer_email"`
	SubscriptionStatus      string         `json:"subscription_status"`
	Currency                enums.Currency `json:"currency"`
	CustomerTotalAmountPaid int            `json:"customer_total_amount_paid"`
}

// Invoice represents a Paystack invoice
type Invoice struct {
	ID                 uint64              `json:"id"`
	Domain             string              `json:"domain"`
	Status             string              `json:"status"`
	Reference          string              `json:"reference"`
	ReceiptNumber      string              `json:"receipt_number"`
	Amount             int                 `json:"amount"`
	Message            string              `json:"message"`
	GatewayResponse    string              `json:"gateway_response"`
	PaidAt             *DateTime           `json:"paid_at,omitempty"`
	CreatedAt          DateTime            `json:"created_at"`
	Channel            enums.Channel       `json:"channel"`
	Currency           enums.Currency      `json:"currency"`
	IPAddress          string              `json:"ip_address"`
	Metadata           Metadata            `json:"metadata"`
	Log                Log                 `json:"log"`
	Fees               int                 `json:"fees"`
	FeesSplit          *FeesSplit          `json:"fees_split,omitempty"`
	Plan               *Plan               `json:"plan,omitempty"`
	Subaccount         *Subaccount         `json:"subaccount,omitempty"`
	Split              *TransactionSplit   `json:"split,omitempty"`
	OrderID            *string             `json:"order_id,omitempty"`
	RequestedAmount    int                 `json:"requested_amount"`
	PosTransactionData *POSTransactionData `json:"pos_transaction_data,omitempty"`
	Source             *Source             `json:"source,omitempty"`
	FeesBreakdown      *FeesSplit          `json:"fees_breakdown,omitempty"`
	Connect            *ConnectData        `json:"connect,omitempty"`
	Authorization      Authorization       `json:"authorization"`
	Customer           Customer            `json:"customer"`
}

// InvoiceHistory represents invoice history
type InvoiceHistory struct {
	ID            uint64        `json:"id"`
	Domain        string        `json:"domain"`
	InvoiceCode   string        `json:"invoice_code"`
	Amount        int           `json:"amount"`
	PeriodStart   DateTime      `json:"period_start"`
	PeriodEnd     DateTime      `json:"period_end"`
	Status        string        `json:"status"`
	Paid          bool          `json:"paid"`
	PaidAt        *DateTime     `json:"paid_at,omitempty"`
	Description   string        `json:"description"`
	CreatedAt     DateTime      `json:"createdAt"`
	Authorization Authorization `json:"authorization"`
	Subscription  Subscription  `json:"subscription"`
	Customer      Customer      `json:"customer"`
	Transaction   *Transaction  `json:"transaction,omitempty"`
}
