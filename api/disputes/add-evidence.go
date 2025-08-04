package disputes

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type AddEvidenceRequest struct {
	CustomerEmail   string     `json:"customer_email"`
	CustomerName    string     `json:"customer_name"`
	CustomerPhone   string     `json:"customer_phone"`
	ServiceDetails  string     `json:"service_details"`
	DeliveryAddress *string    `json:"delivery_address,omitempty"`
	DeliveryDate    *time.Time `json:"delivery_date,omitempty"`
}

type AddEvidenceBuilder struct {
	req *AddEvidenceRequest
}

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

func (b *AddEvidenceBuilder) DeliveryAddress(address string) *AddEvidenceBuilder {
	b.req.DeliveryAddress = &address

	return b
}

func (b *AddEvidenceBuilder) DeliveryDate(date time.Time) *AddEvidenceBuilder {
	b.req.DeliveryDate = &date

	return b
}

func (b *AddEvidenceBuilder) Build() *AddEvidenceRequest {
	return b.req
}

type AddEvidenceResponse = types.Response[types.Evidence]

func (c *Client) AddEvidence(ctx context.Context, disputeID string, builder *AddEvidenceBuilder) (*AddEvidenceResponse, error) {
	return net.Post[AddEvidenceRequest, types.Evidence](ctx, c.Client, c.Secret, basePath+"/"+disputeID+"/evidence", builder.Build(), c.BaseURL)
}
