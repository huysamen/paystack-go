package types
import "github.com/huysamen/paystack-go/types/data"
import "github.com/huysamen/paystack-go/enums"

// Charge represents the charge data in API responses
type Charge struct {
	ID              int            `json:"id"`
	Domain          string         `json:"domain"`
	Status          string         `json:"status"`
	Reference       string         `json:"reference"`
	Amount          int            `json:"amount"`
	Message         string         `json:"message"`
	GatewayResponse string         `json:"gateway_response"`
	PaidAt          *data.MultiDateTime      `json:"paid_at"`
	CreatedAt       *data.MultiDateTime      `json:"created_at"`
	Channel         enums.Channel  `json:"channel"`
	Currency        enums.Currency `json:"currency"`
	IPAddress       string         `json:"ip_address"`
	Metadata        *Metadata      `json:"metadata"`
	Log             *Metadata      `json:"log"`
	Fees            int            `json:"fees"`
	RequestedAmount int            `json:"requested_amount"`
	TransactionDate *data.MultiDateTime      `json:"transaction_date"`
	Plan            *Plan          `json:"plan"`
	Authorization   *Authorization `json:"authorization"`
	Customer        *Customer      `json:"customer"`
}
