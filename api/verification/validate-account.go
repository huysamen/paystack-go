package verification

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ValidateAccount validates an account using additional verification data
func (c *Client) ValidateAccount(ctx context.Context, builder *AccountValidateRequestBuilder) (*types.Response[AccountValidation], error) {
	req := builder.Build()
	return net.Post[AccountValidateRequest, AccountValidation](ctx, c.Client, c.Secret, accountValidateBasePath, req, c.BaseURL)
}
