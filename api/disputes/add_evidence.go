package disputes

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type addEvidenceRequest struct {
	CustomerEmail   string     `json:"customer_email"`
	CustomerName    string     `json:"customer_name"`
	CustomerPhone   string     `json:"customer_phone"`
	ServiceDetails  string     `json:"service_details"`
	DeliveryAddress *string    `json:"delivery_address,omitempty"`
	DeliveryDate    *time.Time `json:"delivery_date,omitempty"`
}

type AddEvidenceRequestBuilder struct {
	req *addEvidenceRequest
}

func NewAddEvidenceRequestBuilder(customerEmail, customerName, customerPhone, serviceDetails string) *AddEvidenceRequestBuilder {
	return &AddEvidenceRequestBuilder{
		req: &addEvidenceRequest{
			CustomerEmail:  customerEmail,
			CustomerName:   customerName,
			CustomerPhone:  customerPhone,
			ServiceDetails: serviceDetails,
		},
	}
}

func (b *AddEvidenceRequestBuilder) DeliveryAddress(address string) *AddEvidenceRequestBuilder {
	b.req.DeliveryAddress = &address

	return b
}

func (b *AddEvidenceRequestBuilder) DeliveryDate(date time.Time) *AddEvidenceRequestBuilder {
	b.req.DeliveryDate = &date

	return b
}

func (b *AddEvidenceRequestBuilder) Build() *addEvidenceRequest {
	return b.req
}

type AddEvidenceRequestData = types.Evidence
type AddEvidenceResponse = types.Response[AddEvidenceRequestData]

func (c *Client) AddEvidence(ctx context.Context, disputeID string, builder AddEvidenceRequestBuilder) (*AddEvidenceResponse, error) {
	return net.Post[addEvidenceRequest, AddEvidenceRequestData](ctx, c.Client, c.Secret, basePath+"/"+disputeID+"/evidence", builder.Build(), c.BaseURL)
}
