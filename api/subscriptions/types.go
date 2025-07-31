package subscriptions

import (
	"time"

	"github.com/huysamen/paystack-go/types"
)

// Subscription represents a Paystack subscription
type Subscription struct {
	ID               int           `json:"id"`
	Domain           string        `json:"domain"`
	Status           string        `json:"status"`
	SubscriptionCode string        `json:"subscription_code"`
	EmailToken       string        `json:"email_token"`
	Amount           int           `json:"amount"`
	CronExpression   string        `json:"cron_expression"`
	NextPaymentDate  *time.Time    `json:"next_payment_date"`
	OpenInvoice      *string       `json:"open_invoice"`
	Integration      int           `json:"integration"`
	Plan             Plan          `json:"plan"`
	Customer         Customer      `json:"customer"`
	Authorization    Authorization `json:"authorization"`
	CreatedAt        time.Time     `json:"createdAt"`
	UpdatedAt        time.Time     `json:"updatedAt"`
}

// SubscriptionWithInvoices represents a subscription with related invoices (used in fetch endpoint)
type SubscriptionWithInvoices struct {
	Subscription
	Invoices []Invoice `json:"invoices"`
}

// Plan represents a subscription plan (simplified for subscription context)
type Plan struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	PlanCode     string         `json:"plan_code"`
	Description  string         `json:"description"`
	Amount       int            `json:"amount"`
	Interval     types.Interval `json:"interval"`
	SendInvoices bool           `json:"send_invoices"`
	SendSMS      bool           `json:"send_sms"`
	Currency     types.Currency `json:"currency"`
}

// Customer represents a customer (simplified for subscription context)
type Customer struct {
	ID           int            `json:"id"`
	FirstName    *string        `json:"first_name"`
	LastName     *string        `json:"last_name"`
	Email        string         `json:"email"`
	CustomerCode string         `json:"customer_code"`
	Phone        *string        `json:"phone"`
	Metadata     map[string]any `json:"metadata"`
	RiskAction   string         `json:"risk_action"`
	Domain       string         `json:"domain"`
	Integration  int            `json:"integration"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
}

// Authorization represents a payment authorization (simplified for subscription context)
type Authorization struct {
	AuthorizationCode string          `json:"authorization_code"`
	Bin               string          `json:"bin"`
	Last4             string          `json:"last4"`
	ExpMonth          string          `json:"exp_month"`
	ExpYear           string          `json:"exp_year"`
	Channel           types.Channel   `json:"channel"`
	CardType          string          `json:"card_type"`
	Bank              string          `json:"bank"`
	CountryCode       string          `json:"country_code"`
	Brand             types.CardBrand `json:"brand"`
	Reusable          bool            `json:"reusable"`
	Signature         string          `json:"signature"`
	AccountName       *string         `json:"account_name"`
}

// Invoice represents a subscription invoice
type Invoice struct {
	ID           int       `json:"id"`
	Domain       string    `json:"domain"`
	Status       string    `json:"status"`
	InvoiceCode  string    `json:"invoice_code"`
	Amount       int       `json:"amount"`
	PeriodStart  time.Time `json:"period_start"`
	PeriodEnd    time.Time `json:"period_end"`
	Subscription int       `json:"subscription"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
