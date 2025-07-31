package charges

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// SubmitAddress submits address to continue a charge
func (c *Client) SubmitAddress(ctx context.Context, req *SubmitAddressRequest) (*SubmitAddressResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("submit address request cannot be nil")
	}

	if err := validateSubmitAddressRequest(req); err != nil {
		return nil, err
	}

	url := c.baseURL + chargesBasePath + "/submit_address"
	return net.Post[SubmitAddressRequest, ChargeData](ctx, c.client, c.secret, url, req)
}

// validateSubmitAddressRequest validates the submit address request
func validateSubmitAddressRequest(req *SubmitAddressRequest) error {
	if req.Address == "" {
		return fmt.Errorf("address is required")
	}

	if req.Reference == "" {
		return fmt.Errorf("reference is required")
	}

	if req.City == "" {
		return fmt.Errorf("city is required")
	}

	if req.State == "" {
		return fmt.Errorf("state is required")
	}

	if req.ZipCode == "" {
		return fmt.Errorf("zip code is required")
	}

	// Basic validation for address length
	if len(req.Address) < 5 {
		return fmt.Errorf("address must be at least 5 characters long")
	}

	if len(req.Address) > 200 {
		return fmt.Errorf("address must not exceed 200 characters")
	}

	// Basic validation for city
	if len(req.City) < 2 {
		return fmt.Errorf("city must be at least 2 characters long")
	}

	if len(req.City) > 50 {
		return fmt.Errorf("city must not exceed 50 characters")
	}

	// Basic validation for state
	if len(req.State) < 2 {
		return fmt.Errorf("state must be at least 2 characters long")
	}

	if len(req.State) > 50 {
		return fmt.Errorf("state must not exceed 50 characters")
	}

	// Basic validation for zip code
	if len(req.ZipCode) < 3 {
		return fmt.Errorf("zip code must be at least 3 characters long")
	}

	if len(req.ZipCode) > 20 {
		return fmt.Errorf("zip code must not exceed 20 characters")
	}

	return nil
}
