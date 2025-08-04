package applepay

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type unregisterDomainRequest struct {
	DomainName string `json:"domainName"`
}

type UnregisterDomainRequestBuilder struct {
	req *unregisterDomainRequest
}

func NewUnregisterDomainRequestBuilder(domainName string) UnregisterDomainRequestBuilder {
	return UnregisterDomainRequestBuilder{
		req: &unregisterDomainRequest{
			DomainName: domainName,
		},
	}
}

func (b *UnregisterDomainRequestBuilder) Build() *unregisterDomainRequest {
	return b.req
}

type UnregisterDomainResponseData = any
type UnregisterDomainResponse = types.Response[UnregisterDomainResponseData]

func (c *Client) UnregisterDomain(ctx context.Context, builder UnregisterDomainRequestBuilder) (*UnregisterDomainResponse, error) {
	return net.DeleteWithBody[unregisterDomainRequest, UnregisterDomainResponseData](ctx, c.Client, c.Secret, unregisterPath, builder.Build(), c.BaseURL)
}
