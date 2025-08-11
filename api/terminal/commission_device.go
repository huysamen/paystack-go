package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CommissionDeviceRequest struct {
	SerialNumber string `json:"serial_number"`
}

type CommissionDeviceRequestBuilder struct {
	req *CommissionDeviceRequest
}

func NewCommissionDeviceRequestBuilder(serialNumber string) *CommissionDeviceRequestBuilder {
	return &CommissionDeviceRequestBuilder{
		req: &CommissionDeviceRequest{
			SerialNumber: serialNumber,
		},
	}
}

func (b *CommissionDeviceRequestBuilder) Build() *CommissionDeviceRequest {
	return b.req
}

type CommissionDeviceResponseData = types.Terminal
type CommissionDeviceResponse = types.Response[CommissionDeviceResponseData]

func (c *Client) CommissionDevice(ctx context.Context, builder CommissionDeviceRequestBuilder) (*CommissionDeviceResponse, error) {
	endpoint := fmt.Sprintf("%s/commission_device", basePath)

	return net.Post[CommissionDeviceRequest, types.Terminal](ctx, c.Client, c.Secret, endpoint, builder.Build(), c.BaseURL)
}
