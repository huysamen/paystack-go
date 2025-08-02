package disputes

import (
"context"
"errors"

"github.com/huysamen/paystack-go/net"
"github.com/huysamen/paystack-go/types"
)

// ResolveDisputeRequest represents the request to resolve a dispute
type ResolveDisputeRequest struct {
	Resolution       DisputeResolution `json:"resolution"`
	Message          string            `json:"message"`
	RefundAmount     int               `json:"refund_amount"`
	UploadedFileName string            `json:"uploaded_filename"`
	Evidence         *int              `json:"evidence,omitempty"`
}

// ResolveDisputeResponse represents the response from resolving a dispute
type ResolveDisputeResponse = types.Response[Dispute]

// ResolveDisputeBuilder builds requests for resolving disputes
type ResolveDisputeBuilder struct {
	request *ResolveDisputeRequest
}

// NewResolveDisputeBuilder creates a new builder for resolving disputes
func NewResolveDisputeBuilder(resolution DisputeResolution, message string, refundAmount int, uploadedFileName string) *ResolveDisputeBuilder {
	return &ResolveDisputeBuilder{
		request: &ResolveDisputeRequest{
			Resolution:       resolution,
			Message:          message,
			RefundAmount:     refundAmount,
			UploadedFileName: uploadedFileName,
		},
	}
}

// Evidence sets the evidence ID
func (b *ResolveDisputeBuilder) Evidence(evidence int) *ResolveDisputeBuilder {
	b.request.Evidence = &evidence
	return b
}

// Build returns the built request
func (b *ResolveDisputeBuilder) Build() *ResolveDisputeRequest {
	return b.request
}

// Resolve resolves a dispute on your integration
func (c *Client) Resolve(ctx context.Context, disputeID string, builder *ResolveDisputeBuilder) (*types.Response[Dispute], error) {
	if disputeID == "" {
		return nil, errors.New("dispute ID is required")
	}

	if builder == nil {
		return nil, ErrBuilderRequired
	}

	endpoint := c.baseURL + disputesBasePath + "/" + disputeID + "/resolve"
	req := builder.Build()

	resp, err := net.Put[ResolveDisputeRequest, Dispute](ctx, c.client, c.secret, endpoint, req, c.baseURL)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
