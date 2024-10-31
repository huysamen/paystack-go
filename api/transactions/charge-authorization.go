package transactions

import (
	"fmt"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionChargeAuthorizationRequest struct {
	// Required
	Amount            int    `json:"amount"`
	Email             string `json:"email"`
	AuthorizationCode string `json:"authorization_code"`

	// Optional
	Reference         string          `json:"reference,omitempty"`
	Currency          types.Currency  `json:"currency,omitempty"`
	Metadata          types.Metadata  `json:"metadata,omitempty"`
	Channels          []types.Channel `json:"channels,omitempty"`
	Subaccount        string          `json:"subaccount,omitempty"`
	TransactionCharge int             `json:"transaction_charge,omitempty"`
	Bearer            types.Bearer    `json:"bearer,omitempty"`
	Queue             bool            `json:"queue,omitempty"`
}

type TransactionChargeAuthorizationResponse struct {
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
	Plan            types.Plan          `json:"plan"`
}

func (c *Client) Initialize(req *TransactionChargeAuthorizationRequest) (*types.Response[TransactionChargeAuthorizationResponse], error) {
	return net.Post[TransactionChargeAuthorizationRequest, TransactionChargeAuthorizationResponse](
		c.client,
		c.secret,
		fmt.Sprintf("%s%s", transactionBasePath, transactionChargeAutorizationPath),
		req,
	)
}
