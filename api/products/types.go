package products

import (
	"github.com/huysamen/paystack-go/types"
)

// Response type aliases
type CreateProductResponse = types.Response[types.Product]
type ListProductsResponse = types.Response[[]types.Product]
type FetchProductResponse = types.Response[types.Product]
type UpdateProductResponse = types.Response[types.Product]
