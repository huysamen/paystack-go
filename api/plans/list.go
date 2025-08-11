package plans

import (
	"context"
	"net/url"
	"strconv"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

type listRequest struct {
	PerPage  *int            `json:"perPage,omitempty"`
	Page     *int            `json:"page,omitempty"`
	Status   *string         `json:"status,omitempty"`
	Interval *enums.Interval `json:"interval,omitempty"`
	Amount   *int            `json:"amount,omitempty"`
}

type ListRequestBuilder struct {
	req *listRequest
}

func NewListRequestBuilder() *ListRequestBuilder {
	return &ListRequestBuilder{
		req: &listRequest{},
	}
}

func (b *ListRequestBuilder) PerPage(perPage int) *ListRequestBuilder {
	b.req.PerPage = &perPage

	return b
}

func (b *ListRequestBuilder) Page(page int) *ListRequestBuilder {
	b.req.Page = &page

	return b
}

func (b *ListRequestBuilder) Status(status string) *ListRequestBuilder {
	b.req.Status = &status

	return b
}

func (b *ListRequestBuilder) Interval(interval enums.Interval) *ListRequestBuilder {
	b.req.Interval = &interval

	return b
}

func (b *ListRequestBuilder) Amount(amount int) *ListRequestBuilder {
	b.req.Amount = &amount

	return b
}

func (b *ListRequestBuilder) Build() *listRequest {
	return b.req
}

func (r *listRequest) toQuery() string {
	params := url.Values{}

	if r.PerPage != nil {
		params.Add("perPage", strconv.Itoa(*r.PerPage))
	}
	if r.Page != nil {
		params.Add("page", strconv.Itoa(*r.Page))
	}
	if r.Status != nil {
		params.Add("status", *r.Status)
	}
	if r.Interval != nil {
		params.Add("interval", r.Interval.String())
	}
	if r.Amount != nil {
		params.Add("amount", strconv.Itoa(*r.Amount))
	}

	return params.Encode()
}

type PlanSubscription struct {
	Customer         data.Int            `json:"customer"`
	Plan             data.Int            `json:"plan"`
	Integration      data.Int            `json:"integration"`
	Domain           data.String         `json:"domain"`
	Start            data.NullInt        `json:"start,omitempty"`
	Status           data.String         `json:"status"`
	Quantity         data.Int            `json:"quantity"`
	Amount           data.Int            `json:"amount"`
	SubscriptionCode data.String         `json:"subscription_code"`
	EmailToken       data.String         `json:"email_token"`
	Authorization    types.Authorization `json:"authorization"`
	EasyCronID       data.NullString     `json:"easy_cron_id"`
	CronExpression   data.String         `json:"cron_expression"`
	NextPaymentDate  data.Time           `json:"next_payment_date"`
	OpenInvoice      data.NullString     `json:"open_invoice"`
	ID               data.Uint           `json:"id"`
	CreatedAt        data.Time           `json:"createdAt"`
	UpdatedAt        data.Time           `json:"updatedAt"`
}

type ListPlan struct {
	ID                data.Uint          `json:"id"`
	Domain            data.String        `json:"domain"`
	Name              data.String        `json:"name"`
	PlanCode          data.String        `json:"plan_code"`
	Description       data.NullString    `json:"description"`
	Amount            data.Int           `json:"amount"`
	Interval          enums.Interval     `json:"interval"`
	SendInvoices      data.Bool          `json:"send_invoices"`
	SendSms           data.Bool          `json:"send_sms"`
	HostedPage        data.Bool          `json:"hosted_page"`
	HostedPageURL     data.NullString    `json:"hosted_page_url,omitempty"`
	HostedPageSummary data.NullString    `json:"hosted_page_summary,omitempty"`
	Currency          enums.Currency     `json:"currency"`
	Integration       data.Int           `json:"integration"`
	CreatedAt         data.Time          `json:"createdAt"`
	UpdatedAt         data.Time          `json:"updatedAt"`
	Subscriptions     []PlanSubscription `json:"subscriptions,omitempty"`
}

type ListResponseData = []ListPlan
type ListResponse = types.Response[ListResponseData]

func (c *Client) List(ctx context.Context, builder ListRequestBuilder) (*ListResponse, error) {
	path := basePath

	req := builder.Build()
	if query := req.toQuery(); query != "" {
		path += "?" + query
	}

	return net.Get[ListResponseData](ctx, c.Client, c.Secret, path, c.BaseURL)
}
