package charge

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type createRequest struct {
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

type CreateRequestBuilder struct {
	req *createRequest
}

func NewCreateRequestBuilder(email, amount string) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		req: &createRequest{
			Email:  email,
			Amount: amount,
		},
	}
}

func (b *CreateRequestBuilder) SplitCode(splitCode string) *CreateRequestBuilder {
	b.req.SplitCode = &splitCode

	return b
}

func (b *CreateRequestBuilder) Subaccount(subaccount string) *CreateRequestBuilder {
	b.req.Subaccount = &subaccount

	return b
}

func (b *CreateRequestBuilder) TransactionCharge(charge int) *CreateRequestBuilder {
	b.req.TransactionCharge = &charge

	return b
}

func (b *CreateRequestBuilder) Bearer(bearer string) *CreateRequestBuilder {
	b.req.Bearer = &bearer

	return b
}

func (b *CreateRequestBuilder) Bank(bank *BankDetails) *CreateRequestBuilder {
	b.req.Bank = bank

	return b
}

func (b *CreateRequestBuilder) BankTransfer(bankTransfer *BankTransferDetails) *CreateRequestBuilder {
	b.req.BankTransfer = bankTransfer

	return b
}

func (b *CreateRequestBuilder) USSD(ussd *USSDDetails) *CreateRequestBuilder {
	b.req.USSD = ussd

	return b
}

func (b *CreateRequestBuilder) MobileMoney(mobileMoney *MobileMoneyDetails) *CreateRequestBuilder {
	b.req.MobileMoney = mobileMoney

	return b
}

func (b *CreateRequestBuilder) QR(qr *QRDetails) *CreateRequestBuilder {
	b.req.QR = qr

	return b
}

func (b *CreateRequestBuilder) AuthorizationCode(authCode string) *CreateRequestBuilder {
	b.req.AuthorizationCode = &authCode

	return b
}

func (b *CreateRequestBuilder) PIN(pin string) *CreateRequestBuilder {
	b.req.PIN = &pin

	return b
}

func (b *CreateRequestBuilder) Metadata(metadata map[string]any) *CreateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

func (b *CreateRequestBuilder) Reference(reference string) *CreateRequestBuilder {
	b.req.Reference = &reference

	return b
}

func (b *CreateRequestBuilder) DeviceID(deviceID string) *CreateRequestBuilder {
	b.req.DeviceID = &deviceID

	return b
}

func (b *CreateRequestBuilder) Birthday(birthday string) *CreateRequestBuilder {
	b.req.Birthday = &birthday

	return b
}

func (b *CreateRequestBuilder) Build() *createRequest {
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

type CreateChargeResponseData = types.ChargeData
type CreateChargeResponse = types.Response[CreateChargeResponseData]

func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateChargeResponse, error) {
	return net.Post[createRequest, CreateChargeResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
