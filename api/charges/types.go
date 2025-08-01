package charges

import (
	"github.com/huysamen/paystack-go/types"
)

// ChargeData represents the charge data in API responses
type ChargeData struct {
	ID              int                  `json:"id"`
	Domain          string               `json:"domain"`
	Status          string               `json:"status"`
	Reference       string               `json:"reference"`
	Amount          int                  `json:"amount"`
	Message         string               `json:"message"`
	GatewayResponse string               `json:"gateway_response"`
	PaidAt          *types.DateTime      `json:"paid_at"`
	CreatedAt       *types.DateTime      `json:"created_at"`
	Channel         string               `json:"channel"`
	Currency        string               `json:"currency"`
	IPAddress       string               `json:"ip_address"`
	Metadata        map[string]any       `json:"metadata"`
	Log             any                  `json:"log"`
	Fees            int                  `json:"fees"`
	RequestedAmount int                  `json:"requested_amount"`
	TransactionDate *types.DateTime      `json:"transaction_date"`
	Plan            any                  `json:"plan"`
	Authorization   *types.Authorization `json:"authorization"`
	Customer        *types.Customer      `json:"customer"`
}
