package transactions

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionInitializeRequest struct {
	// Required fields
	Amount int    `json:"amount"`
	Email  string `json:"email"`

	// Optional fields
	Currency          types.Currency  `json:"currency,omitempty"`
	Reference         string          `json:"reference,omitempty"`
	CallbackURL       string          `json:"callback_url,omitempty"`
	Plan              string          `json:"plan,omitempty"`
	InvoiceLimit      int             `json:"invoice_limit,omitempty"`
	Metadata          types.Metadata  `json:"metadata,omitempty"`
	Channels          []types.Channel `json:"channels,omitempty"`
	SplitCode         []string        `json:"split_code,omitempty"`
	Subaccount        string          `json:"subaccount,omitempty"`
	TransactionCharge int             `json:"transaction_charge,omitempty"`
	Bearer            types.Bearer    `json:"bearer,omitempty"`
}

type TransactionInitializeResponse struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

func (c *Client) Initialize(ctx context.Context, req *TransactionInitializeRequest) (*types.Response[TransactionInitializeResponse], error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	return net.Post[TransactionInitializeRequest, TransactionInitializeResponse](
		ctx,
		c.client,
		c.secret,
		fmt.Sprintf("%s%s", transactionBasePath, transactionInitializePath),
		req,
		c.baseURL,
	)
}
