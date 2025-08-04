package applepay

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type RegisterDomainRequest struct {
	DomainName string `json:"domainName"`
}

type RegisterDomainRequestBuilder struct {
	req *RegisterDomainRequest
}

func NewRegisterDomainRequest(domainName string) RegisterDomainRequestBuilder {
	return RegisterDomainRequestBuilder{
		req: &RegisterDomainRequest{
			DomainName: domainName,
		},
	}
}

func (b *RegisterDomainRequestBuilder) Build() *RegisterDomainRequest {
	return b.req
}

type RegisterDomainResponseData = any
type RegisterDomainResponse = types.Response[RegisterDomainResponseData]

func (c *Client) RegisterDomain(ctx context.Context, builder RegisterDomainRequestBuilder) (*RegisterDomainResponse, error) {
	return net.Post[RegisterDomainRequest, RegisterDomainResponseData](ctx, c.Client, c.Secret, registerPath, builder.Build(), c.BaseURL)
}
