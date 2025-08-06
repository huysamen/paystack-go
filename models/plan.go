package models

import (
	"github.com/huysamen/paystack-go/enums"
)

// Plan represents a Paystack subscription plan
type Plan struct {
	ID                       uint64         `json:"id"`
	Domain                   string         `json:"domain"`
	Name                     string         `json:"name"`
	PlanCode                 string         `json:"plan_code"`
	Description              *string        `json:"description"`
	Amount                   int            `json:"amount"`
	Interval                 enums.Interval `json:"interval"`
	InvoiceLimit             int            `json:"invoice_limit"`
	SendInvoices             bool           `json:"send_invoices"`
	SendSms                  bool           `json:"send_sms"`
	HostedPage               bool           `json:"hosted_page"`
	HostedPageURL            *string        `json:"hosted_page_url,omitempty"`
	HostedPageSummary        *string        `json:"hosted_page_summary,omitempty"`
	Currency                 enums.Currency `json:"currency"`
	Integration              int            `json:"integration"`
	Migrate                  bool           `json:"migrate"`
	IsDeleted                bool           `json:"is_deleted"`
	IsArchived               bool           `json:"is_archived"`
	CreatedAt                DateTime       `json:"createdAt"`
	UpdatedAt                DateTime       `json:"updatedAt"`
	PagesCount               int            `json:"pages_count"`
	SubscribersCount         int            `json:"subscribers_count"`
	ActiveSubscriptionsCount int            `json:"active_subscriptions_count"`
	TotalRevenue             int            `json:"total_revenue"`
	Subscriptions            []Subscription `json:"subscriptions,omitempty"`
	Pages                    []Page         `json:"pages,omitempty"`
}

// Subscription represents a Paystack subscription
type Subscription struct {
	ID               uint64        `json:"id"`
	Customer         *Customer     `json:"customer,omitempty"`
	Plan             *Plan         `json:"plan,omitempty"`
	Integration      int           `json:"integration"`
	Domain           string        `json:"domain"`
	Start            *int64        `json:"start,omitempty"` // Unix timestamp
	Status           string        `json:"status"`
	Quantity         int           `json:"quantity"`
	Amount           int           `json:"amount"`
	SubscriptionCode string        `json:"subscription_code"`
	EmailToken       string        `json:"email_token"`
	Authorization    Authorization `json:"authorization"`
	EasyCronID       *string       `json:"easy_cron_id"`
	CronExpression   string        `json:"cron_expression"`
	NextPaymentDate  *DateTime     `json:"next_payment_date"`
	OpenInvoice      *string       `json:"open_invoice"`
	CreatedAt        DateTime      `json:"createdAt"`
	UpdatedAt        DateTime      `json:"updatedAt"`

	// Additional fields from customer subscription data
	CustomerCode            string         `json:"customer_code,omitempty"`
	CustomerFirstName       *string        `json:"customer_first_name,omitempty"`
	CustomerLastName        *string        `json:"customer_last_name,omitempty"`
	CustomerEmail           string         `json:"customer_email,omitempty"`
	SubscriptionStatus      string         `json:"subscription_status,omitempty"`
	Currency                enums.Currency `json:"currency,omitempty"`
	CustomerTotalAmountPaid int            `json:"customer_total_amount_paid,omitempty"`
}

// Page represents a payment page
type Page struct {
	ID                uint64         `json:"id"`
	Integration       int            `json:"integration"`
	Domain            string         `json:"domain"`
	Name              string         `json:"name"`
	Description       *string        `json:"description"`
	Amount            int            `json:"amount"`
	Currency          enums.Currency `json:"currency"`
	Slug              string         `json:"slug"`
	CustomFields      []CustomField  `json:"custom_fields"`
	Type              enums.PageType `json:"type"`
	RedirectURL       *string        `json:"redirect_url"`
	SuccessMessage    *string        `json:"success_message"`
	CollectPhone      bool           `json:"collect_phone"`
	Active            bool           `json:"active"`
	Published         bool           `json:"published"`
	Migrate           bool           `json:"migrate"`
	NotificationEmail *string        `json:"notification_email"`
	Metadata          Metadata       `json:"metadata"`
	SplitCode         *string        `json:"split_code"`
	CreatedAt         DateTime       `json:"createdAt"`
	UpdatedAt         DateTime       `json:"updatedAt"`
}
