package paymentpages

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CheckSlugAvailabilityResponse represents the response from checking slug availability
type CheckSlugAvailabilityResponse = types.Response[any]

// CheckSlugAvailability checks the availability of a slug for a payment page
func (c *Client) CheckSlugAvailability(ctx context.Context, slug string) (*CheckSlugAvailabilityResponse, error) {
	return net.Get[any](ctx, c.Client, c.Secret, basePath+"/check_slug_availability/"+slug, c.BaseURL)
}
