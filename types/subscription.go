package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Subscription represents a Paystack subscription
type Subscription struct {
	ID data.Uint `json:"id"`
	// Standard shape when fetched returns full objects
	Customer         *Customer       `json:"customer,omitempty"`
	Plan             *Plan           `json:"plan,omitempty"`
	Integration      data.Int        `json:"integration"`
	Domain           data.String     `json:"domain"`
	Start            data.NullInt    `json:"start,omitempty"` // Unix timestamp
	Status           data.String     `json:"status"`
	Quantity         data.Int        `json:"quantity"`
	Amount           data.Int        `json:"amount"`
	SubscriptionCode data.String     `json:"subscription_code"`
	EmailToken       data.String     `json:"email_token"`
	Authorization    Authorization   `json:"authorization"`
	EasyCronID       data.NullString `json:"easy_cron_id"`
	CronExpression   data.String     `json:"cron_expression"`
	NextPaymentDate  data.NullTime   `json:"next_payment_date"`
	OpenInvoice      data.NullString `json:"open_invoice"`
	CreatedAt        data.Time       `json:"createdAt"`
	UpdatedAt        data.Time       `json:"updatedAt"`

	// Additional fields from customer subscription data
	CustomerCode            data.String     `json:"customer_code,omitempty"`
	CustomerFirstName       data.NullString `json:"customer_first_name,omitempty"`
	CustomerLastName        data.NullString `json:"customer_last_name,omitempty"`
	CustomerEmail           data.String     `json:"customer_email,omitempty"`
	SubscriptionStatus      data.String     `json:"subscription_status,omitempty"`
	Currency                enums.Currency  `json:"currency,omitempty"`
	CustomerTotalAmountPaid data.Int        `json:"customer_total_amount_paid,omitempty"`
}
