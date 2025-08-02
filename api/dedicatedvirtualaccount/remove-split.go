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
	return net.DeleteWithBody[RemoveSplitFromDedicatedAccountRequest, DedicatedVirtualAccount](ctx, c.Client, c.Secret, basePath+"/split", builder.Build(), c.BaseURL)
}
