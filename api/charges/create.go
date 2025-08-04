package charges

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CreateChargeRequest struct {
	Email             string               `json:"email"`
	Amount            string               `json:"amount"`
	SplitCode         *string              `json:"split_code,omitempty"`
	Subaccount        *string              `json:"subaccount,omitempty"`
	TransactionCharge *int                 `json:"transaction_charge,omitempty"`
	Bearer            *string              `json:"bearer,omitempty"`
	Bank              *BankDetails         `json:"bank,omitempty"`
	BankTransfer      *BankTransferDetails `json:"bank_transfer,omitempty"`
	USSD              *USSDDetails         `json:"ussd,omitempty"`
	MobileMoney       *MobileMoneyDetails  `json:"mobile_money,omitempty"`
	QR                *QRDetails           `json:"qr,omitempty"`
	AuthorizationCode *string              `json:"authorization_code,omitempty"`
	PIN               *string              `json:"pin,omitempty"`
	Metadata          map[string]any       `json:"metadata,omitempty"`
	Reference         *string              `json:"reference,omitempty"`
	DeviceID          *string              `json:"device_id,omitempty"`
	Birthday          *string              `json:"birthday,omitempty"`
}

type CreateChargeRequestBuilder struct {
	req *CreateChargeRequest
}

func NewCreateChargeRequest(email, amount string) *CreateChargeRequestBuilder {
	return &CreateChargeRequestBuilder{
		req: &CreateChargeRequest{
			Email:  email,
			Amount: amount,
		},
	}
}

func (b *CreateChargeRequestBuilder) SplitCode(splitCode string) *CreateChargeRequestBuilder {
	b.req.SplitCode = &splitCode

	return b
}

func (b *CreateChargeRequestBuilder) Subaccount(subaccount string) *CreateChargeRequestBuilder {
	b.req.Subaccount = &subaccount

	return b
}

func (b *CreateChargeRequestBuilder) TransactionCharge(charge int) *CreateChargeRequestBuilder {
	b.req.TransactionCharge = &charge

	return b
}

func (b *CreateChargeRequestBuilder) Bearer(bearer string) *CreateChargeRequestBuilder {
	b.req.Bearer = &bearer

	return b
}

func (b *CreateChargeRequestBuilder) Bank(bank *BankDetails) *CreateChargeRequestBuilder {
	b.req.Bank = bank

	return b
}

func (b *CreateChargeRequestBuilder) BankTransfer(bankTransfer *BankTransferDetails) *CreateChargeRequestBuilder {
	b.req.BankTransfer = bankTransfer

	return b
}

func (b *CreateChargeRequestBuilder) USSD(ussd *USSDDetails) *CreateChargeRequestBuilder {
	b.req.USSD = ussd

	return b
}

func (b *CreateChargeRequestBuilder) MobileMoney(mobileMoney *MobileMoneyDetails) *CreateChargeRequestBuilder {
	b.req.MobileMoney = mobileMoney

	return b
}

func (b *CreateChargeRequestBuilder) QR(qr *QRDetails) *CreateChargeRequestBuilder {
	b.req.QR = qr

	return b
}

func (b *CreateChargeRequestBuilder) AuthorizationCode(authCode string) *CreateChargeRequestBuilder {
	b.req.AuthorizationCode = &authCode

	return b
}

func (b *CreateChargeRequestBuilder) PIN(pin string) *CreateChargeRequestBuilder {
	b.req.PIN = &pin

	return b
}

func (b *CreateChargeRequestBuilder) Metadata(metadata map[string]any) *CreateChargeRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *CreateChargeRequestBuilder) Reference(reference string) *CreateChargeRequestBuilder {
	b.req.Reference = &reference

	return b
}

func (b *CreateChargeRequestBuilder) DeviceID(deviceID string) *CreateChargeRequestBuilder {
	b.req.DeviceID = &deviceID

	return b
}

func (b *CreateChargeRequestBuilder) Birthday(birthday string) *CreateChargeRequestBuilder {
	b.req.Birthday = &birthday

	return b
}

func (b *CreateChargeRequestBuilder) Build() *CreateChargeRequest {
	return b.req
}

type BankDetails struct {
	Code          string `json:"code"`
	AccountNumber string `json:"account_number"`
}

type BankTransferDetails struct {
	AccountExpiresAt *time.Time `json:"account_expires_at,omitempty"`
}

type USSDDetails struct {
	Type string `json:"type"`
}

type MobileMoneyDetails struct {
	Phone    string `json:"phone"`
	Provider string `json:"provider"`
}

type QRDetails struct {
	Provider string `json:"provider"`
}

type CreateChargeResponse = types.Response[types.ChargeData]

func (c *Client) Create(ctx context.Context, builder *CreateChargeRequestBuilder) (*CreateChargeResponse, error) {
	return net.Post[CreateChargeRequest, types.ChargeData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
