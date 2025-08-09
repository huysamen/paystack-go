package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Bank represents a bank in the system
type Bank struct {
	ID          data.Int           `json:"id"`
	Name        data.String        `json:"name"`
	Slug        data.String        `json:"slug"`
	Code        data.String        `json:"code"`
	LongCode    data.String        `json:"longcode"`
	Gateway     data.NullString    `json:"gateway"`
	PayWithBank data.Bool          `json:"pay_with_bank"`
	Active      data.Bool          `json:"active"`
	IsDeleted   data.Bool          `json:"is_deleted"`
	Country     data.String        `json:"country"`
	Currency    enums.Currency     `json:"currency"`
	Type        data.String        `json:"type"`
	CreatedAt   data.MultiDateTime `json:"createdAt"`
	UpdatedAt   data.MultiDateTime `json:"updatedAt"`
}

// BankProvider represents a bank provider
type BankProvider struct {
	ID           data.Int    `json:"id"`
	ProviderSlug data.String `json:"provider_slug"`
	BankID       data.Int    `json:"bank_id"`
	BankName     data.String `json:"bank_name"`
}
