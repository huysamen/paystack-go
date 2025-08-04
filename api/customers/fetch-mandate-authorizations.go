package customers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type MandateAuthorization struct {
	ID                int               `json:"id"`
	Status            string            `json:"status"`
	MandateID         int               `json:"mandate_id"`
	AuthorizationID   int               `json:"authorization_id"`
	AuthorizationCode string            `json:"authorization_code"`
	IntegrationID     int               `json:"integration_id"`
	AccountNumber     string            `json:"account_number"`
	BankCode          string            `json:"bank_code"`
	BankName          *string           `json:"bank_name"`
	Customer          CustomerReference `json:"customer"`
}

type FetchMandateAuthorizationsResponse = types.Response[[]MandateAuthorization]

func (c *Client) FetchMandateAuthorizations(ctx context.Context, customerID string) (*FetchMandateAuthorizationsResponse, error) {
	path := fmt.Sprintf("%s/%s/directdebit-mandate-authorizations", basePath, customerID)

	return net.Get[[]MandateAuthorization](ctx, c.Client, c.Secret, path, c.BaseURL)
}
