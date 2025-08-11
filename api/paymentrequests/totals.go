package paymentrequests

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
	"github.com/huysamen/paystack-go/types/data"
)

// CurrencyAmount represents a currency amount pair in totals response
type CurrencyAmount struct {
	Currency data.String `json:"currency"`
	Amount   data.Int    `json:"amount"`
}

// TotalsResponseData represents the response data for payment request totals
type TotalsResponseData struct {
	Pending    []CurrencyAmount `json:"pending"`
	Successful []CurrencyAmount `json:"successful"`
	Total      []CurrencyAmount `json:"total"`
}

type TotalsResponse = types.Response[TotalsResponseData]

func (c *Client) GetTotals(ctx context.Context) (*TotalsResponse, error) {
	return net.Get[TotalsResponseData](ctx, c.Client, c.Secret, basePath+"/totals", c.BaseURL)
}
