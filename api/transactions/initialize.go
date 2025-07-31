package transactions

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/utils"
)

type TransactionInitializeRequest struct {
	// Required fields
	Amount int    `json:"amount" validate:"required,min=1"`
	Email  string `json:"email" validate:"required,email"`

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

// Validate checks if the request has all required fields
func (r *TransactionInitializeRequest) Validate() error {
	var errs []error

	// Validate amount
	if err := utils.ValidateAmount(r.Amount); err != nil {
		errs = append(errs, err)
	}

	// Validate email
	if err := utils.ValidateEmail(r.Email); err != nil {
		errs = append(errs, err)
	}

	// Return combined errors
	return utils.CombineValidationErrors(errs...)
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

	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
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
