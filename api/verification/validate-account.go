package verification

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// ValidateAccount validates an account using additional verification data
func (c *Client) ValidateAccount(ctx context.Context, req *AccountValidateRequest) (*AccountValidateResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.AccountName == "" {
		return nil, fmt.Errorf("account_name is required")
	}
	if req.AccountNumber == "" {
		return nil, fmt.Errorf("account_number is required")
	}
	if req.AccountType == "" {
		return nil, fmt.Errorf("account_type is required")
	}
	if req.BankCode == "" {
		return nil, fmt.Errorf("bank_code is required")
	}
	if req.CountryCode == "" {
		return nil, fmt.Errorf("country_code is required")
	}
	if req.DocumentType == "" {
		return nil, fmt.Errorf("document_type is required")
	}

	resp, err := net.Post[AccountValidateRequest, AccountValidateResponse](ctx, c.client, c.secret, accountValidateBasePath, req, c.baseURL)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// ValidateAccountWithBuilder validates an account using the builder pattern
func (c *Client) ValidateAccountWithBuilder(ctx context.Context, builder *AccountValidateRequestBuilder) (*AccountValidateResponse, error) {
	if builder == nil {
		return nil, fmt.Errorf("builder cannot be nil")
	}
	return c.ValidateAccount(ctx, builder.Build())
}
