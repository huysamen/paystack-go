package types

import (
	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/types/data"
)

// Charge represents the charge data in API responses
type Charge struct {
	ID              data.Int            `json:"id"`
	Domain          data.String         `json:"domain"`
	Status          data.String         `json:"status"`
	Reference       data.String         `json:"reference"`
	Amount          data.Int            `json:"amount"`
	Message         data.String         `json:"message"`
	GatewayResponse data.String         `json:"gateway_response"`
	PaidAt          *data.MultiDateTime `json:"paid_at"`
	CreatedAt       *data.MultiDateTime `json:"created_at"`
	Channel         enums.Channel       `json:"channel"`
	Currency        enums.Currency      `json:"currency"`
	IPAddress       data.String         `json:"ip_address"`
	Metadata        *Metadata           `json:"metadata"`
	Log             *Metadata           `json:"log"`
	Fees            data.Int            `json:"fees"`
	RequestedAmount data.Int            `json:"requested_amount"`
	TransactionDate *data.MultiDateTime `json:"transaction_date"`
	Plan            *Plan               `json:"plan"`
	Authorization   *Authorization      `json:"authorization"`
	Customer        *Customer           `json:"customer"`
}
