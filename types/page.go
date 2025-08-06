package types

import (
	"time"

	"github.com/huysamen/paystack-go/enums"
)

// Page represents a Paystack payment page
type Page struct {
	Integration       int            `json:"integration"`
	Domain            string         `json:"domain"`
	Name              string         `json:"name"`
	Description       string         `json:"description"`
	Amount            int            `json:"amount"`
	Currency          enums.Currency `json:"currency"`
	Slug              string         `json:"slug"`
	CustomFields      []CustomField  `json:"custom_fields"`
	Type              enums.PageType `json:"type"`
	RedirectURL       string         `json:"redirect_url"`
	SuccessMessage    string         `json:"success_message"`
	CollectPhone      bool           `json:"collect_phone"`
	Active            bool           `json:"active"`
	Published         bool           `json:"published"`
	Migrate           bool           `json:"migrate"`
	NotificationEmail string         `json:"notification_email"`
	Metadata          Metadata       `json:"metadata"`
	SplitCode         string         `json:"split_code"`
	ID                uint64         `json:"id"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
}
