package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// AddSplitCode adds a split code to a virtual terminal
func (c *Client) AddSplitCode(ctx context.Context, code string, req *AddSplitCodeRequest) (*any, error) {
	if err := validateCode(code); err != nil {
		return nil, err
	}
	if err := validateAddSplitCodeRequest(req); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s/split_code", virtualTerminalBasePath, code)
	resp, err := net.Put[AddSplitCodeRequest, any](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
