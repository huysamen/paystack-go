package customers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Account represents customer account details for direct debit
type Account struct {
	Number   string `json:"number"`
	BankCode string `json:"bank_code"`
}

// Address represents customer address information
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

// Request and Response types
type DirectDebitInitializeRequest struct {
	Account Account `json:"account"`
	Address Address `json:"address"`
}

type DirectDebitInitializeResponse struct {
	RedirectURL string `json:"redirect_url"`
	AccessCode  string `json:"access_code"`
	Reference   string `json:"reference"`
}

// Builder for DirectDebitInitializeRequest
type DirectDebitInitializeRequestBuilder struct {
	account Account
	address Address
}

// NewInitializeDirectDebitRequest creates a new builder for direct debit initialization
func NewInitializeDirectDebitRequest(accountNumber, bankCode, street, city, state string) *DirectDebitInitializeRequestBuilder {
	return &DirectDebitInitializeRequestBuilder{
		account: Account{
			Number:   accountNumber,
			BankCode: bankCode,
		},
		address: Address{
			Street: street,
			City:   city,
			State:  state,
		},
	}
}

// Build creates the DirectDebitInitializeRequest
func (b *DirectDebitInitializeRequestBuilder) Build() *DirectDebitInitializeRequest {
	return &DirectDebitInitializeRequest{
		Account: b.account,
		Address: b.address,
	}
}

// InitializeDirectDebit initializes direct debit for a customer
func (c *Client) InitializeDirectDebit(ctx context.Context, customerID string, builder *DirectDebitInitializeRequestBuilder) (*types.Response[DirectDebitInitializeResponse], error) {
	path := fmt.Sprintf("%s/%s/initialize-direct-debit", basePath, customerID)
	return net.Post[DirectDebitInitializeRequest, DirectDebitInitializeResponse](ctx, c.Client, c.Secret, path, builder.Build(), c.BaseURL)
}
