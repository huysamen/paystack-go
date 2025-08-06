package terminal

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type sendEventRequest struct {
	Type   enums.TerminalEventType   `json:"type"`
	Action enums.TerminalEventAction `json:"action"`
	Data   types.TerminalEventData   `json:"data"`
}

type SendEventRequestBuilder struct {
	req *sendEventRequest
}

func NewSendEventRequestBuilder(eventType enums.TerminalEventType, action enums.TerminalEventAction, data types.TerminalEventData) *SendEventRequestBuilder {
	return &SendEventRequestBuilder{
		req: &sendEventRequest{
			Type:   eventType,
			Action: action,
			Data:   data,
		},
	}
}

func (b *SendEventRequestBuilder) Build() *sendEventRequest {
	return b.req
}

type SendEventResponseData = types.TerminalEventResult
type SendEventResponse = types.Response[SendEventResponseData]

func (c *Client) SendEvent(ctx context.Context, terminalID string, builder SendEventRequestBuilder) (*SendEventResponse, error) {
	endpoint := fmt.Sprintf("%s/%s/event", basePath, terminalID)

	return net.Post[sendEventRequest, SendEventResponseData](ctx, c.Client, c.Secret, endpoint, builder.Build(), c.BaseURL)
}
