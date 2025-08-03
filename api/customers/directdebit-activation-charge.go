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

type DirectDebitActivationChargeResponse = types.Response[any]

// DirectDebitActivationCharge creates an activation charge for a direct debit authorization
func (c *Client) DirectDebitActivationCharge(ctx context.Context, customerID string, builder *DirectDebitActivationChargeRequestBuilder) (*DirectDebitActivationChargeResponse, error) {
	path := fmt.Sprintf("%s/%s/directdebit-activation-charge", basePath, customerID)

	return net.Put[DirectDebitActivationChargeRequest, any](ctx, c.Client, c.Secret, path, builder.Build(), c.BaseURL)
}
