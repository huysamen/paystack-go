package types

import (
	"time"
)

type Authorization struct {
	AuthorizationCode         string    `json:"authorization_code"`
	Bin                       string    `json:"bin"`
	Last4                     string    `json:"last4"`
	Description               string    `json:"description"`
	ExpMonth                  string    `json:"exp_month"`
	ExpYear                   string    `json:"exp_year"`
	Channel                   Channel   `json:"channel"`
	CardType                  string    `json:"card_type"`
	Brand                     CardBrand `json:"brand"`
	Bank                      string    `json:"bank"`
	Reusable                  bool      `json:"reusable"`
	Signature                 string    `json:"signature"`
	AccountName               string    `json:"account_name"`
	CountryCode               string    `json:"country_code"`
	CountryName               string    `json:"country_name"`
	ReceiverBankAccountNumber string    `json:"receiver_bank_account_number"`
	ReceiverBank              string    `json:"receiver_bank"`
}

type Customer struct {
	ID                       uint64   `json:"id"`
	FirstName                string   `json:"first_name"`
	LastName                 string   `json:"last_name"`
	Email                    string   `json:"email"`
	CustomerCode             string   `json:"customer_code"`
	Phone                    string   `json:"phone"`
	Metadata                 Metadata `json:"metadata"`
	RiskAction               string   `json:"risk_action"`
	InternationalFormatPhone string   `json:"international_format_phone"`
}

type Invoice struct {
	ID                 uint64        `json:"id"`
	Domain             string        `json:"domain"`
	Status             string        `json:"status"`
	Reference          string        `json:"reference"`
	ReceiptNumber      string        `json:"receipt_number"`
	Amount             int           `json:"amount"`
	Message            string        `json:"message"`
	GatewayResponse    string        `json:"gateway_response"`
	PaidAt             time.Time     `json:"paid_at"`
	CreatedAt          time.Time     `json:"created_at"`
	Channel            Channel       `json:"channel"`
	Currency           Currency      `json:"currency"`
	IPAddress          string        `json:"ip_address"`
	Metadata           Metadata      `json:"metadata"`
	Log                Log           `json:"log"`
	Fees               int           `json:"fees"`
	FeesSplit          any           `json:"fees_split"`
	Plan               Plan          `json:"plan"`
	Subaccount         Subaccount    `json:"subaccount"`
	Split              Split         `json:"split"`
	OrderID            any           `json:"order_id"`
	RequestedAmount    int           `json:"requested_amount"`
	PosTransactionData any           `json:"pos_transaction_data"`
	Source             any           `json:"source"`
	FeesBreakdown      any           `json:"fees_breakdown"`
	Connect            any           `json:"connect"`
	Authorization      Authorization `json:"authorization"`
	Customer           Customer      `json:"customer"`
}

type InvoiceHistory struct {
	ID            uint64        `json:"id"`
	Domain        string        `json:"domain"`
	InvoiceCode   string        `json:"invoice_code"`
	Amount        int           `json:"amount"`
	PeriodStart   time.Time     `json:"period_start"`
	PeriodEnd     time.Time     `json:"period_end"`
	Status        string        `json:"status"`
	Paid          bool          `json:"paid"`
	PaidAt        time.Time     `json:"paid_at"`
	Description   string        `json:"description"`
	CreatedAt     time.Time     `json:"createdAt"`
	Authorization Authorization `json:"authorization"`
	Subscription  Subscription  `json:"subscription"`
	Customer      Customer      `json:"customer"`
	Transaction   Transaction   `json:"transaction"`
}

type Log struct {
	StartTime int  `json:"start_time"`
	TimeSpent int  `json:"time_spent"`
	Attempts  int  `json:"attempts"`
	Errors    int  `json:"errors"`
	Success   bool `json:"success"`
	Mobile    bool `json:"mobile"`
	Input     any  `json:"input"`
	History   []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
		Time    int    `json:"time"`
	} `json:"history"`
}

type Page struct {
	Integration       int           `json:"integration"`
	Domain            string        `json:"domain"`
	Name              string        `json:"name"`
	Description       string        `json:"description"`
	Amount            int           `json:"amount"`
	Currency          Currency      `json:"currency"`
	Slug              string        `json:"slug"`
	CustomFields      []CustomField `json:"custom_fields"`
	Type              PageType      `json:"type"`
	RedirectURL       string        `json:"redirect_url"`
	SuccessMessage    string        `json:"success_message"`
	CollectPhone      bool          `json:"collect_phone"`
	Active            bool          `json:"active"`
	Published         bool          `json:"published"`
	Migrate           bool          `json:"migrate"`
	NotificationEmail string        `json:"notification_email"`
	Metadata          Metadata      `json:"metadata"`
	SplitCode         string        `json:"split_code"`
	ID                uint64        `json:"id"`
	CreatedAt         time.Time     `json:"createdAt"`
	UpdatedAt         time.Time     `json:"updatedAt"`
}

type Plan struct {
	Domain                   string         `json:"domain"`
	Name                     string         `json:"name"`
	PlanCode                 string         `json:"plan_code"`
	Description              string         `json:"description"`
	Amount                   int            `json:"amount"`
	Interval                 Interval       `json:"interval"`
	InvoiceLimit             int            `json:"invoice_limit"`
	SendInvoices             bool           `json:"send_invoices"`
	SendSms                  bool           `json:"send_sms"`
	HostedPage               bool           `json:"hosted_page"`
	HostedPageURL            string         `json:"hosted_page_url"`
	HostedPageSummary        string         `json:"hosted_page_summary"`
	Currency                 Currency       `json:"currency"`
	Migrate                  bool           `json:"migrate"`
	IsDeleted                bool           `json:"is_deleted"`
	IsArchived               bool           `json:"is_archived"`
	ID                       uint64         `json:"id"`
	Integration              int            `json:"integration"`
	CreatedAt                time.Time      `json:"createdAt"`
	UpdatedAt                time.Time      `json:"updatedAt"`
	PagesCount               int            `json:"pages_count"`
	SubscribersCount         int            `json:"subscribers_count"`
	ActiveSubscriptionsCount int            `json:"active_subscriptions_count"`
	TotalRevenue             int            `json:"total_revenue"`
	Subscriptions            []Subscription `json:"subscriptions"`
	Pages                    []Page         `json:"pages"`
	Subscribers              []Subscriber   `json:"subscribers"`
}

type Source struct {
	Type       string `json:"type"`
	Source     string `json:"source"`
	EntryPoint string `json:"entry_point"`
	Identifier string `json:"identifier"`
}

type Split struct {
	// todo
}

type Subscriber struct {
	CustomerCode            string   `json:"customer_code"`
	CustomerFirstName       string   `json:"customer_first_name"`
	CustomerLastName        string   `json:"customer_last_name"`
	CustomerEmail           string   `json:"customer_email"`
	SubscriptionStatus      string   `json:"subscription_status"`
	Currency                Currency `json:"currency"`
	CustomerTotalAmountPaid int      `json:"customer_total_amount_paid"`
}

type Subaccount struct {
	Integration          int       `json:"integration"`
	Bank                 int       `json:"bank"`
	ManagedByIntegration int       `json:"managed_by_integration"`
	Domain               string    `json:"domain"`
	SubaccountCode       string    `json:"subaccount_code"`
	BusinessName         string    `json:"business_name"`
	Description          string    `json:"description"`
	PrimaryContactName   string    `json:"primary_contact_name"`
	PrimaryContactEmail  string    `json:"primary_contact_email"`
	PrimaryContactPhone  string    `json:"primary_contact_phone"`
	Metadata             Metadata  `json:"metadata"`
	PercentageCharge     int       `json:"percentage_charge"`
	IsVerified           bool      `json:"is_verified"`
	SettlementBank       string    `json:"settlement_bank"`
	AccountNumber        string    `json:"account_number"`
	SettlementSchedule   string    `json:"settlement_schedule"`
	Active               bool      `json:"active"`
	Migrate              bool      `json:"migrate"`
	Currency             Currency  `json:"currency"`
	AccountName          string    `json:"account_name"`
	Product              string    `json:"product"`
	ID                   uint64    `json:"id"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}

type Subscription struct {
	ID                uint64           `json:"id,omitempty"`
	Domain            string           `json:"domain"`
	Status            string           `json:"status"`
	SubscriptionCode  string           `json:"subscription_code"`
	EmailToken        string           `json:"email_token"`
	Amount            int              `json:"amount"`
	CronExpression    string           `json:"cron_expression"`
	NextPaymentDate   time.Time        `json:"next_payment_date"`
	OpenInvoice       *Invoice         `json:"open_invoice,omitempty"`
	CreatedAt         time.Time        `json:"createdAt"`
	CancelledAt       *time.Time       `json:"cancelledAt"`
	Integration       int              `json:"integration"`
	Plan              Plan             `json:"plan"`
	Authorization     Authorization    `json:"authorization"`
	Customer          Customer         `json:"customer"`
	Invoices          []Invoice        `json:"invoices"`
	InvoicesHistory   []InvoiceHistory `json:"invoices_history"`
	InvoiceLimit      int              `json:"invoice_limit"`
	SplitCode         string           `json:"split_code"`
	MostRecentInvoice Invoice          `json:"most_recent_invoice"`
	PaymentsCount     int              `json:"payments_count"`

	// Returned as part of fetching related resources.
	Quantity           *int       `json:"quantity,omitempty"`
	SuccessfulPayments *int       `json:"successful_payments,omitempty"`
	Start              *time.Time `json:"start,omitempty"`
}

type Transaction struct {
	ID                 uint64        `json:"id"`
	Domain             string        `json:"domain"`
	Status             string        `json:"status"`
	Reference          string        `json:"reference"`
	Amount             int           `json:"amount"`
	Message            string        `json:"message"`
	GatewayResponse    string        `json:"gateway_response"`
	PaidAt             time.Time     `json:"paid_at"`
	CreatedAt          time.Time     `json:"created_at"`
	Channel            Channel       `json:"channel"`
	Currency           Currency      `json:"currency"`
	IPAddress          string        `json:"ip_address"`
	Metadata           Metadata      `json:"metadata"`
	Log                Log           `json:"log"`
	Fees               int           `json:"fees"`
	FeesSplit          any           `json:"fees_split"`
	Customer           Customer      `json:"customer"`
	Authorization      Authorization `json:"authorization"`
	Plan               Plan          `json:"plan"`
	Split              Split         `json:"split"`
	Subaccount         Subaccount    `json:"subsccount"`
	OrderID            any           `json:"order_id"`
	RequestedAmount    int           `json:"requested_amount"`
	Source             Source        `json:"source"`
	Connect            any           `json:"connect"`
	POSTransactionData any           `json:"pos_transaction_data"`
}
