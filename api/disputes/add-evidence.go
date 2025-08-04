package disputes

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// AddEvidenceRequest represents the request to add evidence to a dispute
type AddEvidenceRequest struct {
	CustomerEmail   string     `json:"customer_email"`
	CustomerName    string     `json:"customer_name"`
	CustomerPhone   string     `json:"customer_phone"`
	ServiceDetails  string     `json:"service_details"`
	DeliveryAddress *string    `json:"delivery_address,omitempty"`
	DeliveryDate    *time.Time `json:"delivery_date,omitempty"`
}

// AddEvidenceBuilder builds requests for adding evidence to disputes
type AddEvidenceBuilder struct {
	req *AddEvidenceRequest
}

// NewAddEvidenceBuilder creates a new builder for adding evidence
func NewAddEvidenceBuilder(customerEmail, customerName, customerPhone, serviceDetails string) *AddEvidenceBuilder {
	return &AddEvidenceBuilder{
		req: &AddEvidenceRequest{
			CustomerEmail:  customerEmail,
			CustomerName:   customerName,
			CustomerPhone:  customerPhone,
			ServiceDetails: serviceDetails,
		},
	}
}

// DeliveryAddress sets the delivery address
func (b *AddEvidenceBuilder) DeliveryAddress(address string) *AddEvidenceBuilder {
	b.req.DeliveryAddress = &address

	return b
}

// DeliveryDate sets the delivery date
func (b *AddEvidenceBuilder) DeliveryDate(date time.Time) *AddEvidenceBuilder {
	b.req.DeliveryDate = &date

	return b
}

// Build returns the built request
func (b *AddEvidenceBuilder) Build() *AddEvidenceRequest {
	return b.req
}

// AddEvidenceResponse represents the response from adding evidence to a dispute
type AddEvidenceResponse = types.Response[types.Evidence]

// AddEvidence provides evidence for a dispute
func (c *Client) AddEvidence(ctx context.Context, disputeID string, builder *AddEvidenceBuilder) (*AddEvidenceResponse, error) {
	return net.Post[AddEvidenceRequest, types.Evidence](ctx, c.Client, c.Secret, basePath+"/"+disputeID+"/evidence", builder.Build(), c.BaseURL)
}
