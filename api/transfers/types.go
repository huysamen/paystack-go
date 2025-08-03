package transfers

import (
	"github.com/huysamen/paystack-go/types"
)

// BulkTransferItem represents a single transfer in a bulk transfer request
type BulkTransferItem struct {
	Amount    int    `json:"amount"`
	Reference string `json:"reference"`
	Reason    string `json:"reason"`
	Recipient string `json:"recipient"`
}

// BulkTransferResponse represents a single transfer result in a bulk transfer response
type BulkTransferResponse struct {
	Reference    string         `json:"reference"`
	Recipient    string         `json:"recipient"`
	Amount       int            `json:"amount"`
	TransferCode string         `json:"transfer_code"`
	Currency     types.Currency `json:"currency"`
	Status       string         `json:"status"`
}
