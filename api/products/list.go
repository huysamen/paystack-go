package products

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

type listRequest struct {
	PerPage *int    `json:"perPage,omitempty"`
	Page    *int    `json:"page,omitempty"`
	From    *string `json:"from,omitempty"`
	To      *string `json:"to,omitempty"`
}

type ListRequestBuilder struct {
	req *listRequest
}

func NewListRequestBuilder() *ListRequestBuilder {
	return &ListRequestBuilder{
		req: &listRequest{},
	}
}

func (b *ListRequestBuilder) PerPage(perPage int) *ListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *ListRequestBuilder) Page(page int) *ListRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *ListRequestBuilder) From(from string) *ListRequestBuilder {
	b.req.From = &from

	return b
}

func (b *ListRequestBuilder) To(to string) *ListRequestBuilder {
	b.req.To = &to

	return b
}

func (b *ListRequestBuilder) DateRange(from, to string) *ListRequestBuilder {
	b.req.From = &from
	b.req.To = &to

	return b
}

func (b *ListRequestBuilder) Build() *listRequest {
	return b.req
}

func (r *listRequest) toQuery() string {
	params := url.Values{}
	if r.PerPage != nil {
		params.Set("perPage", strconv.Itoa(*r.PerPage))
	}
	if r.Page != nil {
		params.Set("page", strconv.Itoa(*r.Page))
	}
	if r.From != nil {
		params.Set("from", *r.From)
	}
	if r.To != nil {
		params.Set("to", *r.To)
	}

	return params.Encode()
}

type ListProduct struct {
	ID                 data.Int        `json:"id,omitempty"`
	Integration        data.Int        `json:"integration,omitempty"`
	Name               data.String     `json:"name"`
	Description        data.String     `json:"description"`
	ProductCode        data.String     `json:"product_code,omitempty"`
	Price              data.Int        `json:"price"`
	Currency           enums.Currency  `json:"currency"`
	Quantity           data.NullInt    `json:"quantity,omitempty"`
	QuantitySold       data.NullInt    `json:"quantity_sold,omitempty"`
	Type               data.String     `json:"type,omitempty"`  // good, service
	Files              []any           `json:"files,omitempty"` // Array in list response
	IsShippable        data.Bool       `json:"is_shippable,omitempty"`
	ShippingFields     types.Metadata  `json:"shipping_fields,omitempty"` // Object in list response
	Unlimited          data.Bool       `json:"unlimited,omitempty"`
	Domain             data.String     `json:"domain,omitempty"`
	Active             data.Bool       `json:"active,omitempty"`
	InStock            data.Bool       `json:"in_stock,omitempty"`
	Metadata           types.Metadata  `json:"metadata,omitempty"` // Object in list response
	Slug               data.String     `json:"slug,omitempty"`
	SuccessMessage     data.NullString `json:"success_message,omitempty"`
	RedirectURL        data.NullString `json:"redirect_url,omitempty"`
	SplitCode          data.NullString `json:"split_code,omitempty"`
	NotificationEmails []data.String   `json:"notification_emails,omitempty"`
	MinimumOrderable   data.NullInt    `json:"minimum_orderable,omitempty"`
	MaximumOrderable   data.NullInt    `json:"maximum_orderable,omitempty"`
	LowStockAlert      data.Bool       `json:"low_stock_alert,omitempty"`
	DigitalAssets      []any           `json:"digital_assets,omitempty"`  // Array in list response
	VariantOptions     []any           `json:"variant_options,omitempty"` // Array in list response
	CreatedAt          data.NullTime   `json:"createdAt,omitempty"`
	UpdatedAt          data.NullTime   `json:"updatedAt,omitempty"`
}

type ListResponseData = []ListProduct
type ListResponse = types.Response[ListResponseData]

func (c *Client) List(ctx context.Context, builder ListRequestBuilder) (*ListResponse, error) {
	req := builder.Build()
	path := basePath

	if req != nil {
		if query := req.toQuery(); query != "" {
			path += "?" + query
		}
	}

	return net.Get[ListResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
