package customers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Request and Response types
type DirectDebitActivationChargeRequest struct {
	AuthorizationID int `json:"authorization_id"`
}

type DirectDebitActivationChargeResponse struct {
	Message string `json:"message"`
}

// Builder for DirectDebitActivationChargeRequest
type DirectDebitActivationChargeRequestBuilder struct {
	authorizationID int
}

// NewDirectDebitActivationChargeRequest creates a new builder for direct debit activation charge
func NewDirectDebitActivationChargeRequest(authorizationID int) *DirectDebitActivationChargeRequestBuilder {
	return &DirectDebitActivationChargeRequestBuilder{
		authorizationID: authorizationID,
	}
}

// Build creates the DirectDebitActivationChargeRequest
func (b *DirectDebitActivationChargeRequestBuilder) Build() *DirectDebitActivationChargeRequest {
	return &DirectDebitActivationChargeRequest{
		AuthorizationID: b.authorizationID,
	}
}

// DirectDebitActivationCharge creates an activation charge for a direct debit authorization
func (c *Client) DirectDebitActivationCharge(ctx context.Context, customerID string, builder *DirectDebitActivationChargeRequestBuilder) (*types.Response[DirectDebitActivationChargeResponse], error) {
	if customerID == "" {
		return nil, fmt.Errorf("customer ID is required")
	}

	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	path := fmt.Sprintf("%s/%s/directdebit-activation-charge", customerBasePath, customerID)

	return net.Put[DirectDebitActivationChargeRequest, DirectDebitActivationChargeResponse](
		ctx,
		c.client,
		c.secret,
		path,
		req,
		c.baseURL,
	)
}
