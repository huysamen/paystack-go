package charges

import (
	"context"
	"time"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CreateChargeRequest represents the request to create a charge
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

// CreateChargeRequestBuilder provides a fluent interface for building CreateChargeRequest
type CreateChargeRequestBuilder struct {
	req *CreateChargeRequest
}

// NewCreateChargeRequest creates a new builder for CreateChargeRequest
func NewCreateChargeRequest(email, amount string) *CreateChargeRequestBuilder {
	return &CreateChargeRequestBuilder{
		req: &CreateChargeRequest{
			Email:  email,
			Amount: amount,
		},
	}
}

// SplitCode sets the split code
func (b *CreateChargeRequestBuilder) SplitCode(splitCode string) *CreateChargeRequestBuilder {
	b.req.SplitCode = &splitCode
	return b
}

// Subaccount sets the subaccount ID
func (b *CreateChargeRequestBuilder) Subaccount(subaccount string) *CreateChargeRequestBuilder {
	b.req.Subaccount = &subaccount
	return b
}

// TransactionCharge sets the transaction charge amount
func (b *CreateChargeRequestBuilder) TransactionCharge(charge int) *CreateChargeRequestBuilder {
	b.req.TransactionCharge = &charge
	return b
}

// Bearer sets the bearer of transaction charges
func (b *CreateChargeRequestBuilder) Bearer(bearer string) *CreateChargeRequestBuilder {
	b.req.Bearer = &bearer
	return b
}

// Bank sets the bank details for bank charging
func (b *CreateChargeRequestBuilder) Bank(bank *BankDetails) *CreateChargeRequestBuilder {
	b.req.Bank = bank
	return b
}

// BankTransfer sets the bank transfer details
func (b *CreateChargeRequestBuilder) BankTransfer(bankTransfer *BankTransferDetails) *CreateChargeRequestBuilder {
	b.req.BankTransfer = bankTransfer
	return b
}

// USSD sets the USSD details
func (b *CreateChargeRequestBuilder) USSD(ussd *USSDDetails) *CreateChargeRequestBuilder {
	b.req.USSD = ussd
	return b
}

// MobileMoney sets the mobile money details
func (b *CreateChargeRequestBuilder) MobileMoney(mobileMoney *MobileMoneyDetails) *CreateChargeRequestBuilder {
	b.req.MobileMoney = mobileMoney
	return b
}

// QR sets the QR details
func (b *CreateChargeRequestBuilder) QR(qr *QRDetails) *CreateChargeRequestBuilder {
	b.req.QR = qr
	return b
}

// AuthorizationCode sets the authorization code for repeat charges
func (b *CreateChargeRequestBuilder) AuthorizationCode(authCode string) *CreateChargeRequestBuilder {
	b.req.AuthorizationCode = &authCode
	return b
}

// PIN sets the PIN for card charges
func (b *CreateChargeRequestBuilder) PIN(pin string) *CreateChargeRequestBuilder {
	b.req.PIN = &pin
	return b
}

// Metadata sets the transaction metadata
func (b *CreateChargeRequestBuilder) Metadata(metadata map[string]any) *CreateChargeRequestBuilder {
	b.req.Metadata = metadata
	return b
}

// Reference sets the transaction reference
func (b *CreateChargeRequestBuilder) Reference(reference string) *CreateChargeRequestBuilder {
	b.req.Reference = &reference
	return b
}

// DeviceID sets the device ID
func (b *CreateChargeRequestBuilder) DeviceID(deviceID string) *CreateChargeRequestBuilder {
	b.req.DeviceID = &deviceID
	return b
}

// Birthday sets the birthday for verification
func (b *CreateChargeRequestBuilder) Birthday(birthday string) *CreateChargeRequestBuilder {
	b.req.Birthday = &birthday
	return b
}

// Build returns the constructed CreateChargeRequest
func (b *CreateChargeRequestBuilder) Build() *CreateChargeRequest {
	return b.req
}

// BankDetails represents bank account details for charging
type BankDetails struct {
	Code          string `json:"code"`
	AccountNumber string `json:"account_number"`
}

// BankTransferDetails represents bank transfer payment details
type BankTransferDetails struct {
	AccountExpiresAt *time.Time `json:"account_expires_at,omitempty"`
}

// USSDDetails represents USSD payment details
type USSDDetails struct {
	Type string `json:"type"`
}

// MobileMoneyDetails represents mobile money payment details
type MobileMoneyDetails struct {
	Phone    string `json:"phone"`
	Provider string `json:"provider"`
}

// QRDetails represents QR payment details
type QRDetails struct {
	Provider string `json:"provider"`
}

type CreateChargeResponse = types.Response[ChargeData]

// Create initiates a payment by integrating multiple payment channels
func (c *Client) Create(ctx context.Context, builder *CreateChargeRequestBuilder) (*CreateChargeResponse, error) {
	return net.Post[CreateChargeRequest, ChargeData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
