package disputes

import (
"context"
"errors"

"github.com/huysamen/paystack-go/net"
"github.com/huysamen/paystack-go/types"
)

// UpdateDisputeRequest represents the request to update a dispute
type UpdateDisputeRequest struct {
	RefundAmount     *int    `json:"refund_amount,omitempty"`
	UploadedFileName *string `json:"uploaded_filename,omitempty"`
}

// UpdateDisputeResponse represents the response from updating a dispute
type UpdateDisputeResponse = types.Response[Dispute]

// UpdateDisputeBuilder builds requests for updating disputes
type UpdateDisputeBuilder struct {
	request *UpdateDisputeRequest
}

// NewUpdateDisputeBuilder creates a new builder for updating disputes
func NewUpdateDisputeBuilder() *UpdateDisputeBuilder {
	return &UpdateDisputeBuilder{
		request: &UpdateDisputeRequest{},
	}
}

// RefundAmount sets the refund amount
func (b *UpdateDisputeBuilder) RefundAmount(amount int) *UpdateDisputeBuilder {
	b.request.RefundAmount = &amount
	return b
}

// UploadedFileName sets the uploaded file name
func (b *UpdateDisputeBuilder) UploadedFileName(fileName string) *UpdateDisputeBuilder {
	b.request.UploadedFileName = &fileName
	return b
}

// Build returns the built request
func (b *UpdateDisputeBuilder) Build() *UpdateDisputeRequest {
	return b.request
}

// Update updates the details of a dispute on your integration
func (c *Client) Update(ctx context.Context, disputeID string, builder *UpdateDisputeBuilder) (*types.Response[Dispute], error) {
	if disputeID == "" {
		return nil, errors.New("dispute ID is required")
	}

	if builder == nil {
		return nil, errors.New(ErrBuilderRequired)
	}

	endpoint := c.baseURL + disputesBasePath + "/" + disputeID
	req := builder.Build()

	resp, err := net.Put[UpdateDisputeRequest, Dispute](ctx, c.client, c.secret, endpoint, req, c.baseURL)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
