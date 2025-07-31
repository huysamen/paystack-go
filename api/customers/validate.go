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

func (r *CustomerValidateRequest) Validate() error {
	if r.FirstName == "" {
		return errors.New("first_name is required")
	}
	if r.LastName == "" {
		return errors.New("last_name is required")
	}
	if r.Type == "" {
		return errors.New("type is required")
	}
	if r.Type != "bank_account" {
		return errors.New("only 'bank_account' type is supported")
	}
	if r.Value == "" {
		return errors.New("value is required")
	}
	if r.Country == "" {
		return errors.New("country is required")
	}
	if r.BVN == "" {
		return errors.New("bvn is required")
	}
	if r.Type == "bank_account" {
		if r.BankCode == "" {
			return errors.New("bank_code is required for bank_account type")
		}
		if r.AccountNumber == "" {
			return errors.New("account_number is required for bank_account type")
		}
	}
	return nil
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

	if err := req.Validate(); err != nil {
		return nil, err
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
