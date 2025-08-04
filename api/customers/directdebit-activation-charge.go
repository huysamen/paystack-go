package customers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type DirectDebitActivationChargeRequest struct {
	AuthorizationID int `json:"authorization_id"`
}

type DirectDebitActivationChargeRequestBuilder struct {
	req *DirectDebitActivationChargeRequest
}

func NewDirectDebitActivationChargeRequest(authorizationID int) *DirectDebitActivationChargeRequestBuilder {
	return &DirectDebitActivationChargeRequestBuilder{
		req: &DirectDebitActivationChargeRequest{
			AuthorizationID: authorizationID,
		},
	}
}

func (b *DirectDebitActivationChargeRequestBuilder) Build() *DirectDebitActivationChargeRequest {
	return b.req
}

type DirectDebitActivationChargeResponseData = any
type DirectDebitActivationChargeResponse = types.Response[DirectDebitActivationChargeResponseData]

func (c *Client) DirectDebitActivationCharge(ctx context.Context, customerID string, builder DirectDebitActivationChargeRequestBuilder) (*DirectDebitActivationChargeResponse, error) {
	path := fmt.Sprintf("%s/%s/directdebit-activation-charge", basePath, customerID)

	return net.Put[DirectDebitActivationChargeRequest, DirectDebitActivationChargeResponseData](ctx, c.Client, c.Secret, path, builder.Build(), c.BaseURL)
}
