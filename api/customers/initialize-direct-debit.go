package customers

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type DirectDebitInitializeRequest struct {
	Account Account `json:"account"`
	Address Address `json:"address"`
}

func (r *DirectDebitInitializeRequest) Validate() error {
	if r.Account.Number == "" {
		return errors.New("account number is required")
	}
	if r.Account.BankCode == "" {
		return errors.New("bank code is required")
	}
	if r.Address.Street == "" {
		return errors.New("address street is required")
	}
	if r.Address.City == "" {
		return errors.New("address city is required")
	}
	if r.Address.State == "" {
		return errors.New("address state is required")
	}
	return nil
}

type DirectDebitInitializeResponse struct {
	RedirectURL string `json:"redirect_url"`
	AccessCode  string `json:"access_code"`
	Reference   string `json:"reference"`
}

func (c *Client) InitializeDirectDebit(ctx context.Context, customerID string, req *DirectDebitInitializeRequest) (*types.Response[DirectDebitInitializeResponse], error) {
	if customerID == "" {
		return nil, errors.New("customer ID is required")
	}

	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s/initialize-direct-debit", customerBasePath, customerID)

	return net.Post[DirectDebitInitializeRequest, DirectDebitInitializeResponse](
		ctx,
		c.client,
		c.secret,
		path,
		req,
		c.baseURL,
	)
}
