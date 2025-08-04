package paymentpages

import (
	"context"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// AddProductsToPageRequest represents the request to add products to a payment page
type AddProductsToPageRequest struct {
	Product []int `json:"product"`
}

// AddProductsToPageRequestBuilder provides a fluent interface for building AddProductsToPageRequest
type AddProductsToPageRequestBuilder struct {
	req *AddProductsToPageRequest
}

// NewAddProductsToPageRequest creates a new builder for AddProductsToPageRequest
func NewAddProductsToPageRequest() *AddProductsToPageRequestBuilder {
	return &AddProductsToPageRequestBuilder{
		req: &AddProductsToPageRequest{
			Product: []int{},
		},
	}
}

// Products sets the list of product IDs to add
func (b *AddProductsToPageRequestBuilder) Products(productIDs []int) *AddProductsToPageRequestBuilder {
	b.req.Product = productIDs
	return b
}

// AddProduct adds a single product ID to the list
func (b *AddProductsToPageRequestBuilder) AddProduct(productID int) *AddProductsToPageRequestBuilder {
	b.req.Product = append(b.req.Product, productID)
	return b
}

// Build returns the constructed AddProductsToPageRequest
func (b *AddProductsToPageRequestBuilder) Build() *AddProductsToPageRequest {
	return b.req
}

// AddProductsToPageResponse represents the response from adding products to a payment page
type AddProductsToPageResponse = types.Response[types.PaymentPage]

// AddProducts adds products to a payment page using the builder pattern
func (c *Client) AddProducts(ctx context.Context, pageID int, builder *AddProductsToPageRequestBuilder) (*AddProductsToPageResponse, error) {
	return net.Post[AddProductsToPageRequest, types.PaymentPage](ctx, c.Client, c.Secret, basePath+"/"+strconv.Itoa(pageID)+"/product", builder.Build(), c.BaseURL)
}
