package virtualterminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// AssignDestination assigns destinations to a virtual terminal
func (c *Client) AssignDestination(ctx context.Context, code string, req *AssignDestinationRequest) (*[]VirtualTerminalDestination, error) {
	if err := validateCode(code); err != nil {
		return nil, err
	}
	if err := validateAssignDestinationRequest(req); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s/destination/assign", virtualTerminalBasePath, code)
	resp, err := net.Post[AssignDestinationRequest, []VirtualTerminalDestination](
		ctx, c.client, c.secret, endpoint, req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
