package paymentpages

import (
	"context"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type AddProductsToPageRequest struct {
	Product []int `json:"product"`
}

type AddProductsToPageRequestBuilder struct {
	req *AddProductsToPageRequest
}

func NewAddProductsToPageRequest() *AddProductsToPageRequestBuilder {
	return &AddProductsToPageRequestBuilder{
		req: &AddProductsToPageRequest{
			Product: []int{},
		},
	}
}

func (b *AddProductsToPageRequestBuilder) Products(productIDs []int) *AddProductsToPageRequestBuilder {
	b.req.Product = productIDs

	return b
}

func (b *AddProductsToPageRequestBuilder) AddProduct(productID int) *AddProductsToPageRequestBuilder {
	b.req.Product = append(b.req.Product, productID)

	return b
}

func (b *AddProductsToPageRequestBuilder) Build() *AddProductsToPageRequest {
	return b.req
}

type AddProductsToPageResponse = types.Response[types.PaymentPage]

func (c *Client) AddProducts(ctx context.Context, pageID int, builder *AddProductsToPageRequestBuilder) (*AddProductsToPageResponse, error) {
	return net.Post[AddProductsToPageRequest, types.PaymentPage](ctx, c.Client, c.Secret, basePath+"/"+strconv.Itoa(pageID)+"/product", builder.Build(), c.BaseURL)
}
