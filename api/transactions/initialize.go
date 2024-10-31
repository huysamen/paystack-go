package transactions

import (
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionInitializeRequest struct {
	Amount            int             `json:"amount,omitempty"`
	Email             string          `json:"email,omitempty"`
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

func (c *Client) Initialize(req *TransactionInitializeRequest) (*types.Response[TransactionInitializeResponse], error) {
	return net.Post[TransactionInitializeRequest, TransactionInitializeResponse](
		c.client,
		c.secret,
		fmt.Sprintf("%s%s", transactionBasePath, transactionInitializePath),
		req,
	)
}
