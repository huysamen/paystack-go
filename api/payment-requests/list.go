package paymentrequests

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ListPaymentRequestsRequest represents the request to list payment requests
type ListPaymentRequestsRequest struct {
	PerPage        int    `json:"perPage,omitempty"`
	Page           int    `json:"page,omitempty"`
	Customer       string `json:"customer,omitempty"`
	Status         string `json:"status,omitempty"`
	Currency       string `json:"currency,omitempty"`
	IncludeArchive string `json:"include_archive,omitempty"`
	From           string `json:"from,omitempty"`
	To             string `json:"to,omitempty"`
}

// ListPaymentRequestsRequestBuilder provides a fluent interface for building ListPaymentRequestsRequest
type ListPaymentRequestsRequestBuilder struct {
	req *ListPaymentRequestsRequest
}

// NewListPaymentRequestsRequest creates a new builder for ListPaymentRequestsRequest
func NewListPaymentRequestsRequest() *ListPaymentRequestsRequestBuilder {
	return &ListPaymentRequestsRequestBuilder{
		req: &ListPaymentRequestsRequest{},
	}
}

// PerPage sets the number of payment requests per page
func (b *ListPaymentRequestsRequestBuilder) PerPage(perPage int) *ListPaymentRequestsRequestBuilder {
	b.req.PerPage = perPage
	return b
}

// Page sets the page number
func (b *ListPaymentRequestsRequestBuilder) Page(page int) *ListPaymentRequestsRequestBuilder {
	b.req.Page = page
	return b
}

// Customer sets the customer filter
func (b *ListPaymentRequestsRequestBuilder) Customer(customer string) *ListPaymentRequestsRequestBuilder {
	b.req.Customer = customer
	return b
}

// Status sets the status filter
func (b *ListPaymentRequestsRequestBuilder) Status(status string) *ListPaymentRequestsRequestBuilder {
	b.req.Status = status
	return b
}

// Currency sets the currency filter
func (b *ListPaymentRequestsRequestBuilder) Currency(currency string) *ListPaymentRequestsRequestBuilder {
	b.req.Currency = currency
	return b
}

// IncludeArchive sets whether to include archived requests
func (b *ListPaymentRequestsRequestBuilder) IncludeArchive(includeArchive string) *ListPaymentRequestsRequestBuilder {
	b.req.IncludeArchive = includeArchive
	return b
}

// From sets the start date filter
func (b *ListPaymentRequestsRequestBuilder) From(from string) *ListPaymentRequestsRequestBuilder {
	b.req.From = from
	return b
}

// To sets the end date filter
func (b *ListPaymentRequestsRequestBuilder) To(to string) *ListPaymentRequestsRequestBuilder {
	b.req.To = to
	return b
}

// DateRange sets both from and to dates for convenience
func (b *ListPaymentRequestsRequestBuilder) DateRange(from, to string) *ListPaymentRequestsRequestBuilder {
	b.req.From = from
	b.req.To = to
	return b
}

// Build returns the constructed ListPaymentRequestsRequest
func (b *ListPaymentRequestsRequestBuilder) Build() *ListPaymentRequestsRequest {
	return b.req
}

// ListPaymentRequestsResponse represents the response from listing payment requests
type ListPaymentRequestsResponse struct {
	Status  bool             `json:"status"`
	Message string           `json:"message"`
	Data    []PaymentRequest `json:"data"`
	Meta    *types.Meta      `json:"meta,omitempty"`
}

// List retrieves payment requests available on your integration
func (c *Client) List(ctx context.Context, builder *ListPaymentRequestsRequestBuilder) (*types.Response[[]PaymentRequest], error) {
	var req *ListPaymentRequestsRequest
	if builder != nil {
		req = builder.Build()
	}

	params := url.Values{}

	if req != nil {
		if req.PerPage > 0 {
			params.Set("perPage", strconv.Itoa(req.PerPage))
		}
		if req.Page > 0 {
			params.Set("page", strconv.Itoa(req.Page))
		}
		if req.Customer != "" {
			params.Set("customer", req.Customer)
		}
		if req.Status != "" {
			params.Set("status", req.Status)
		}
		if req.Currency != "" {
			params.Set("currency", req.Currency)
		}
		if req.IncludeArchive != "" {
			params.Set("include_archive", req.IncludeArchive)
		}
		if req.From != "" {
			params.Set("from", req.From)
		}
		if req.To != "" {
			params.Set("to", req.To)
		}
	}

	endpoint := paymentRequestsBasePath
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	return net.Get[[]PaymentRequest](
		ctx, c.client, c.secret, endpoint, c.baseURL,
	)
}
