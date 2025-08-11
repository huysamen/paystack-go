package paymentpages

import (
	"context"
	"strconv"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type addProductsRequest struct {
	Product []int `json:"product"`
}

type AddProductsRequestBuilder struct {
	req *addProductsRequest
}

func NewAddProductsRequestBuilder() *AddProductsRequestBuilder {
	return &AddProductsRequestBuilder{
		req: &addProductsRequest{
			Product: []int{},
		},
	}
}

func (b *AddProductsRequestBuilder) Products(productIDs []int) *AddProductsRequestBuilder {
	b.req.Product = productIDs

	return b
}

func (b *AddProductsRequestBuilder) AddProduct(productID int) *AddProductsRequestBuilder {
	b.req.Product = append(b.req.Product, productID)

	return b
}

func (b *AddProductsRequestBuilder) Build() *addProductsRequest {
	return b.req
}

type AddProductsResponseData = types.PaymentPage
type AddProductsResponse = types.Response[AddProductsResponseData]

func (c *Client) AddProducts(ctx context.Context, pageID int, builder AddProductsRequestBuilder) (*AddProductsResponse, error) {
	return net.Post[addProductsRequest, AddProductsResponseData](ctx, c.Client, c.Secret, basePath+"/"+strconv.Itoa(pageID)+"/product", builder.Build(), c.BaseURL)
}
