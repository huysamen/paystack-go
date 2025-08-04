package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CheckSlugAvailabilityResponse = types.Response[any]

func (c *Client) CheckSlugAvailability(ctx context.Context, slug string) (*CheckSlugAvailabilityResponse, error) {
	return net.Get[any](ctx, c.Client, c.Secret, basePath+"/check_slug_availability/"+slug, c.BaseURL)
}
