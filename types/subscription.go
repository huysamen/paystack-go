package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Subscription represents a Paystack subscription
type Subscription struct {
	ID               uint64              `json:"id"`
	Customer         *Customer           `json:"customer,omitempty"`
	Plan             *Plan               `json:"plan,omitempty"`
	Integration      int                 `json:"integration"`
	Domain           string              `json:"domain"`
	Start            data.NullInt        `json:"start,omitempty"` // Unix timestamp
	Status           string              `json:"status"`
	Quantity         int                 `json:"quantity"`
	Amount           int                 `json:"amount"`
	SubscriptionCode string              `json:"subscription_code"`
	EmailToken       string              `json:"email_token"`
	Authorization    Authorization       `json:"authorization"`
	EasyCronID       data.NullString     `json:"easy_cron_id"`
	CronExpression   string              `json:"cron_expression"`
	NextPaymentDate  *data.MultiDateTime `json:"next_payment_date"`
	OpenInvoice      data.NullString     `json:"open_invoice"`
	CreatedAt        data.MultiDateTime  `json:"createdAt"`
	UpdatedAt        data.MultiDateTime  `json:"updatedAt"`

	// Additional fields from customer subscription data
	CustomerCode            string          `json:"customer_code,omitempty"`
	CustomerFirstName       data.NullString `json:"customer_first_name,omitempty"`
	CustomerLastName        data.NullString `json:"customer_last_name,omitempty"`
	CustomerEmail           string          `json:"customer_email,omitempty"`
	SubscriptionStatus      string          `json:"subscription_status,omitempty"`
	Currency                enums.Currency  `json:"currency,omitempty"`
	CustomerTotalAmountPaid int             `json:"customer_total_amount_paid,omitempty"`
}
