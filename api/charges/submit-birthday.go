package charges

import (
	"context"
	"fmt"
	"time"

	"github.com/huysamen/paystack-go/net"
)

// SubmitBirthday submits birthday when requested
func (c *Client) SubmitBirthday(ctx context.Context, req *SubmitBirthdayRequest) (*SubmitBirthdayResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("submit birthday request cannot be nil")
	}

	if err := validateSubmitBirthdayRequest(req); err != nil {
		return nil, err
	}

	url := c.baseURL + chargesBasePath + "/submit_birthday"
	return net.Post[SubmitBirthdayRequest, ChargeData](ctx, c.client, c.secret, url, req)
}

// validateSubmitBirthdayRequest validates the submit birthday request
func validateSubmitBirthdayRequest(req *SubmitBirthdayRequest) error {
	if req.Birthday == "" {
		return fmt.Errorf("birthday is required")
	}

	if req.Reference == "" {
		return fmt.Errorf("reference is required")
	}

	// Validate birthday format (YYYY-MM-DD)
	_, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		return fmt.Errorf("invalid birthday format. Expected YYYY-MM-DD format (e.g., 1990-12-25)")
	}

	// Validate that the birthday is not in the future
	birthday, _ := time.Parse("2006-01-02", req.Birthday)
	if birthday.After(time.Now()) {
		return fmt.Errorf("birthday cannot be in the future")
	}

	// Validate reasonable age range (not more than 150 years ago)
	minDate := time.Now().AddDate(-150, 0, 0)
	if birthday.Before(minDate) {
		return fmt.Errorf("birthday is too far in the past")
	}

	return nil
}
