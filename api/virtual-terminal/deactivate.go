package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Deactivate deactivates a virtual terminal
func (c *Client) Deactivate(ctx context.Context, code string) (*types.Response[interface{}], error) {
	if err := validateCode(code); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s/deactivate", virtualTerminalBasePath, code)
	resp, err := net.Put[interface{}, interface{}](
		ctx, c.client, c.secret, endpoint, nil, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
