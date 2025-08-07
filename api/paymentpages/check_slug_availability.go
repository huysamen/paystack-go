package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CheckSlugAvailabilityResponseData = any
type CheckSlugAvailabilityResponse = types.Response[CheckSlugAvailabilityResponseData]

func (c *Client) CheckSlugAvailability(ctx context.Context, slug string) (*CheckSlugAvailabilityResponse, error) {
	return net.Get[CheckSlugAvailabilityResponseData](ctx, c.Client, c.Secret, basePath+"/check_slug_availability/"+slug, c.BaseURL)
}
