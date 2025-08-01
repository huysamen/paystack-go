package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CheckSlugAvailability checks the availability of a slug for a payment page
func (c *Client) CheckSlugAvailability(ctx context.Context, slug string) (*types.Response[any], error) {

	return net.Get[any](
		ctx, c.client, c.secret, paymentPagesBasePath+"/check_slug_availability/"+slug, c.baseURL,
	)
}
