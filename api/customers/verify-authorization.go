package customers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CustomerReference struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}

type VerifyAuthorizationResponseData struct {
	AuthorizationCode string            `json:"authorization_code"`
	Channel           types.Channel     `json:"channel"`
	Bank              string            `json:"bank"`
	Active            bool              `json:"active"`
	Customer          CustomerReference `json:"customer"`
}

type VerifyAuthorizationResponse = types.Response[VerifyAuthorizationResponseData]

func (c *Client) VerifyAuthorization(ctx context.Context, reference string) (*VerifyAuthorizationResponse, error) {
	path := fmt.Sprintf("%s/authorization/verify/%s", basePath, reference)

	return net.Get[VerifyAuthorizationResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
