package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type DecommissionDeviceRequest struct {
	SerialNumber string `json:"serial_number"`
}

type DecommissionDeviceRequestBuilder struct {
	req *DecommissionDeviceRequest
}

func NewDecommissionDeviceRequest(serialNumber string) *DecommissionDeviceRequestBuilder {
	return &DecommissionDeviceRequestBuilder{
		req: &DecommissionDeviceRequest{
			SerialNumber: serialNumber,
		},
	}
}

func (b *DecommissionDeviceRequestBuilder) Build() *DecommissionDeviceRequest {
	return b.req
}

type DecommissionDeviceResponseData = any
type DecommissionDeviceResponse = types.Response[DecommissionDeviceResponseData]

func (c *Client) DecommissionDevice(ctx context.Context, builder DecommissionDeviceRequestBuilder) (*DecommissionDeviceResponse, error) {
	endpoint := fmt.Sprintf("%s/decommission_device", basePath)

	return net.Post[DecommissionDeviceRequest, DecommissionDeviceResponseData](ctx, c.Client, c.Secret, endpoint, builder.Build(), c.BaseURL)
}
