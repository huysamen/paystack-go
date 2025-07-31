package customers

import (
	"time"

	"github.com/huysamen/paystack-go/types"
)

// Customer represents a Paystack customer
type Customer struct {
	ID              int            `json:"id"`
	CustomerCode    string         `json:"customer_code"`
	Email           string         `json:"email"`
	FirstName       *string        `json:"first_name"`
	LastName        *string        `json:"last_name"`
	Phone           *string        `json:"phone"`
	Metadata        map[string]any `json:"metadata"`
	Domain          string         `json:"domain"`
	Integration     int            `json:"integration"`
	Identified      bool           `json:"identified"`
	Identifications any            `json:"identifications"`
	RiskAction      string         `json:"risk_action"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
}

// CustomerWithRelations represents a customer with related data (used in fetch endpoint)
type CustomerWithRelations struct {
	Customer
	Transactions   []any           `json:"transactions"`
	Subscriptions  []any           `json:"subscriptions"`
	Authorizations []Authorization `json:"authorizations"`
}

// Authorization represents a customer's payment authorization
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

// Account represents customer account details for direct debit
type Account struct {
	Number   string `json:"number"`
	BankCode string `json:"bank_code"`
}

// Address represents customer address information
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}
