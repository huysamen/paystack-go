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
