package applepay

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type UnregisterDomainRequest struct {
	DomainName string `json:"domainName"`
}

type UnregisterDomainRequestBuilder struct {
	req *UnregisterDomainRequest
}

func NewUnregisterDomainRequest(domainName string) UnregisterDomainRequestBuilder {
	return UnregisterDomainRequestBuilder{
		req: &UnregisterDomainRequest{
			DomainName: domainName,
		},
	}
}

func (b *UnregisterDomainRequestBuilder) Build() *UnregisterDomainRequest {
	return b.req
}

type UnregisterDomainResponseData = any
type UnregisterDomainResponse = types.Response[UnregisterDomainResponseData]

func (c *Client) UnregisterDomain(ctx context.Context, builder UnregisterDomainRequestBuilder) (*UnregisterDomainResponse, error) {
	return net.DeleteWithBody[UnregisterDomainRequest, UnregisterDomainResponseData](ctx, c.Client, c.Secret, unregisterPath, builder.Build(), c.BaseURL)
}
