package transactions

import (
	"context"
	"fmt"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionPartialDebitRequest struct {
	// Required
	AuthorizationCode string         `json:"authorization_code"`
	Currency          types.Currency `json:"currency"`
	Amount            int            `json:"amount"`
	Email             string         `json:"email"`

	// Optional
	Reference string `json:"reference,omitempty"`
	AtLeast   string `json:"at_least,omitempty"`
}

type TransactionPartialDebitResponse struct {
	ID              uint64              `json:"id"`
	Amount          int                 `json:"amount"`
	Currency        types.Currency      `json:"currency"`
	TransactionDate time.Time           `json:"transaction_date"`
	Status          string              `json:"status"`
	Reference       string              `json:"reference"`
	Domain          string              `json:"domain"`
	Metadata        types.Metadata      `json:"metadata"`
	GatewayResponse string              `json:"gateway_response"`
	Message         string              `json:"message"`
	Channel         types.Channel       `json:"channel"`
	IPAddress       string              `json:"ip_address"`
	Log             types.Log           `json:"log"`
	Fees            int                 `json:"fees"`
	Authorization   types.Authorization `json:"authorization"`
	Customer        types.Customer      `json:"customer"`
	Plan            uint64              `json:"plan"`
	RequestedAmount int                 `json:"requested_amount"`
}

func (c *Client) PartialDebit(ctx context.Context, req *TransactionPartialDebitRequest) (*types.Response[TransactionPartialDebitResponse], error) {
	return net.Post[TransactionPartialDebitRequest, TransactionPartialDebitResponse](
		ctx,
		c.client,
		c.secret,
		fmt.Sprintf("%s%s", transactionBasePath, transactionPartialDebitPath),
		req,
		c.baseURL,
	)
}
