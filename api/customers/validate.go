package customers

import (
	"context"
	"errors"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CustomerValidateRequest struct {
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Type          string  `json:"type"` // Only "bank_account" is supported
	Value         string  `json:"value"`
	Country       string  `json:"country"`
	BVN           string  `json:"bvn"`
	BankCode      string  `json:"bank_code"`      // Required if type is bank_account
	AccountNumber string  `json:"account_number"` // Required if type is bank_account
	MiddleName    *string `json:"middle_name,omitempty"`
}


type CustomerValidateResponse struct {
	Message string `json:"message"`
}

func (c *Client) Validate(ctx context.Context, code string, req *CustomerValidateRequest) (*types.Response[CustomerValidateResponse], error) {
	if code == "" {
		return nil, errors.New("customer code is required")
	}

	if req == nil {
		return nil, errors.New("request cannot be nil")
	}


	path := fmt.Sprintf("%s/%s/identification", customerBasePath, code)

	return net.Post[CustomerValidateRequest, CustomerValidateResponse](
		ctx,
		c.client,
		c.secret,
		path,
		req,
		c.baseURL,
	)
}
