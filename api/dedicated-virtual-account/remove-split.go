package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// RemoveSplitFromDedicatedAccountRequest represents the request to remove split from dedicated account
type RemoveSplitFromDedicatedAccountRequest struct {
	AccountNumber string `json:"account_number"`
}

// RemoveSplitFromDedicatedAccountResponse represents the response from removing split from dedicated account
type RemoveSplitFromDedicatedAccountResponse struct {
	Status  bool                     `json:"status"`
	Message string                   `json:"message"`
	Data    *DedicatedVirtualAccount `json:"data"`
}

// RemoveSplitFromDedicatedAccountBuilder builds requests for removing split from dedicated accounts
type RemoveSplitFromDedicatedAccountBuilder struct {
	request *RemoveSplitFromDedicatedAccountRequest
}

// NewRemoveSplitFromDedicatedAccountBuilder creates a new builder for removing split from dedicated accounts
func NewRemoveSplitFromDedicatedAccountBuilder() *RemoveSplitFromDedicatedAccountBuilder {
	return &RemoveSplitFromDedicatedAccountBuilder{
		request: &RemoveSplitFromDedicatedAccountRequest{},
	}
}

// AccountNumber sets the account number for removing split from the dedicated account
func (b *RemoveSplitFromDedicatedAccountBuilder) AccountNumber(accountNumber string) *RemoveSplitFromDedicatedAccountBuilder {
	b.request.AccountNumber = accountNumber
	return b
}

// Build returns the built request
func (b *RemoveSplitFromDedicatedAccountBuilder) Build() *RemoveSplitFromDedicatedAccountRequest {
	return b.request
}

// RemoveSplit removes split payment setup from a dedicated virtual account
func (c *Client) RemoveSplit(ctx context.Context, builder *RemoveSplitFromDedicatedAccountBuilder) (*types.Response[DedicatedVirtualAccount], error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	endpoint := dedicatedVirtualAccountBasePath + "/split"
	resp, err := net.DeleteWithBody[RemoveSplitFromDedicatedAccountRequest, DedicatedVirtualAccount](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
