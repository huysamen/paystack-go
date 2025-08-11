package applepay

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type registerDomainRequest struct {
	DomainName string `json:"domainName"`
}

type RegisterDomainRequestBuilder struct {
	req *registerDomainRequest
}

func NewRegisterDomainRequestBuilder(domainName string) RegisterDomainRequestBuilder {
	return RegisterDomainRequestBuilder{
		req: &registerDomainRequest{
			DomainName: domainName,
		},
	}
}

func (b *RegisterDomainRequestBuilder) Build() *registerDomainRequest {
	return b.req
}

type RegisterDomainResponseData = any
type RegisterDomainResponse = types.Response[RegisterDomainResponseData]

func (c *Client) RegisterDomain(ctx context.Context, builder RegisterDomainRequestBuilder) (*RegisterDomainResponse, error) {
	return net.Post[registerDomainRequest, RegisterDomainResponseData](ctx, c.Client, c.Secret, registerPath, builder.Build(), c.BaseURL)
}
