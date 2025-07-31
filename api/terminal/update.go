package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Update updates a terminal's details
func (c *Client) Update(ctx context.Context, terminalID string, req *TerminalUpdateRequest) (*TerminalUpdateResponse, error) {
	if err := validateTerminalID(terminalID); err != nil {
		return nil, err
	}
	if err := validateUpdateRequest(req); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s", terminalBasePath, terminalID)
	resp, err := net.Put[TerminalUpdateRequest, TerminalUpdateResponse](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
