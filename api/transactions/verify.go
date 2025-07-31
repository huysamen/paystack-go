package transactions

import (
	"context"
	"fmt"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionVerifyResponse struct {
	ID                 int64               `json:"id"`
	Domain             string              `json:"domain"`
	Status             string              `json:"status"`
	Reference          string              `json:"reference"`
	ReceiptNumber      string              `json:"receipt_number"`
	Amount             int                 `json:"amount"`
	Message            string              `json:"message"`
	GatewayResponse    string              `json:"gateway_response"`
	PaidAt             time.Time           `json:"paid_at"`
	CreatedAt          time.Time           `json:"created_at"`
	Channel            types.Channel       `json:"channel"`
	Currency           types.Currency      `json:"currency"`
	IPAddress          string              `json:"ip_address"`
	Metadata           types.Metadata      `json:"metadata"`
	Log                types.Log           `json:"log"`
	Fees               int                 `json:"fees"`
	FeesSplit          any                 `json:"fees_split"`
	Authorization      types.Authorization `json:"authorization"`
	Customer           types.Customer      `json:"customer"`
	Plan               string              `json:"plan"`
	Split              types.Split         `json:"split"`
	OrderID            any                 `json:"order_id"`
	RequestedAmount    int                 `json:"requested_amount"`
	PosTransactionData any                 `json:"pos_transaction_data"`
	Source             any                 `json:"source"`
	FeesBreakdown      any                 `json:"fees_breakdown"`
	Connect            any                 `json:"connect"`
	TransactionDate    time.Time           `json:"transaction_date"`
	PlanObject         types.Plan          `json:"plan_object"`
	Subaccount         types.Subaccount    `json:"subaccount"`
}

func (c *Client) Verify(ctx context.Context, reference string) (*types.Response[TransactionVerifyResponse], error) {
	return net.Get[TransactionVerifyResponse](
		ctx,
		c.client,
		c.secret,
		fmt.Sprintf("%s%s/%s", transactionBasePath, transactionVerifyPath, reference),
		c.baseURL,
	)
}
