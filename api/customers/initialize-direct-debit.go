package customers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type Account struct {
	Number   string `json:"number"`
	BankCode string `json:"bank_code"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

type InitializeDirectDebitRequest struct {
	Account Account `json:"account"`
	Address Address `json:"address"`
}

type InitializeDirectDebitRequestBuilder struct {
	req *InitializeDirectDebitRequest
}

func NewInitializeDirectDebitRequest(accountNumber, bankCode, street, city, state string) *InitializeDirectDebitRequestBuilder {
	return &InitializeDirectDebitRequestBuilder{
		req: &InitializeDirectDebitRequest{
			Account: Account{
				Number:   accountNumber,
				BankCode: bankCode,
			},
			Address: Address{
				Street: street,
				City:   city,
				State:  state,
			},
		},
	}
}

func (b *InitializeDirectDebitRequestBuilder) Build() *InitializeDirectDebitRequest {
	return b.req
}

type InitializeDirectDebitResponseData struct {
	RedirectURL string `json:"redirect_url"`
	AccessCode  string `json:"access_code"`
	Reference   string `json:"reference"`
}

type InitializeDirectDebitResponse = types.Response[InitializeDirectDebitResponseData]

func (c *Client) InitializeDirectDebit(ctx context.Context, customerID string, builder InitializeDirectDebitRequestBuilder) (*InitializeDirectDebitResponse, error) {
	path := fmt.Sprintf("%s/%s/initialize-direct-debit", basePath, customerID)
	return net.Post[InitializeDirectDebitRequest, InitializeDirectDebitResponseData](ctx, c.Client, c.Secret, path, builder.Build(), c.BaseURL)
}
