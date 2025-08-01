package applepay

import (
	"github.com/huysamen/paystack-go/types"
)

// Domain represents an Apple Pay registered domain
type Domain struct {
	DomainName string `json:"domainName"`
}

// RegisterDomainRequest represents the request to register an Apple Pay domain
type RegisterDomainRequest struct {
	DomainName string `json:"domainName"`
}

// UnregisterDomainRequest represents the request to unregister an Apple Pay domain
type UnregisterDomainRequest struct {
	DomainName string `json:"domainName"`
}

// ListDomainsRequest represents the request to list Apple Pay domains
type ListDomainsRequest struct {
	UseCursor *bool   `json:"use_cursor,omitempty"`
	Next      *string `json:"next,omitempty"`
	Previous  *string `json:"previous,omitempty"`
}

// ListDomainsRequestBuilder provides a fluent interface for building ListDomainsRequest
type ListDomainsRequestBuilder struct {
	req *ListDomainsRequest
}

// NewListDomainsRequest creates a new builder for ListDomainsRequest
func NewListDomainsRequest() *ListDomainsRequestBuilder {
	return &ListDomainsRequestBuilder{
		req: &ListDomainsRequest{},
	}
}

// UseCursor sets whether to use cursor-based pagination
func (b *ListDomainsRequestBuilder) UseCursor(useCursor bool) *ListDomainsRequestBuilder {
	b.req.UseCursor = &useCursor
	return b
}

// Next sets the cursor for next page
func (b *ListDomainsRequestBuilder) Next(next string) *ListDomainsRequestBuilder {
	b.req.Next = &next
	return b
}

// Previous sets the cursor for previous page
func (b *ListDomainsRequestBuilder) Previous(previous string) *ListDomainsRequestBuilder {
	b.req.Previous = &previous
	return b
}

// Build returns the constructed ListDomainsRequest
func (b *ListDomainsRequestBuilder) Build() *ListDomainsRequest {
	return b.req
}

// RegisterDomainResponse represents the response from registering an Apple Pay domain
type RegisterDomainResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ListDomainsResponse represents the response from listing Apple Pay domains
type ListDomainsResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		DomainNames []string `json:"domainNames"`
	} `json:"data"`
	Meta *types.Meta `json:"meta,omitempty"`
}

// UnregisterDomainResponse represents the response from unregistering an Apple Pay domain
type UnregisterDomainResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
