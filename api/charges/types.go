package charges

import (
	"time"

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

// CreateChargeResponse represents the response from creating a charge
type CreateChargeResponse = types.Response[ChargeData]

// SubmitPINRequest represents the request to submit PIN for a charge
type SubmitPINRequest struct {
	PIN       string `json:"pin"`
	Reference string `json:"reference"`
}

// SubmitPINResponse represents the response from submitting PIN
type SubmitPINResponse = types.Response[ChargeData]

// SubmitOTPRequest represents the request to submit OTP for a charge
type SubmitOTPRequest struct {
	OTP       string `json:"otp"`
	Reference string `json:"reference"`
}

// SubmitOTPResponse represents the response from submitting OTP
type SubmitOTPResponse = types.Response[ChargeData]

// SubmitPhoneRequest represents the request to submit phone number for a charge
type SubmitPhoneRequest struct {
	Phone     string `json:"phone"`
	Reference string `json:"reference"`
}

// SubmitPhoneResponse represents the response from submitting phone
type SubmitPhoneResponse = types.Response[ChargeData]

// SubmitBirthdayRequest represents the request to submit birthday for a charge
type SubmitBirthdayRequest struct {
	Birthday  string `json:"birthday"`
	Reference string `json:"reference"`
}

// SubmitBirthdayResponse represents the response from submitting birthday
type SubmitBirthdayResponse = types.Response[ChargeData]

// SubmitAddressRequest represents the request to submit address for a charge
type SubmitAddressRequest struct {
	Address   string `json:"address"`
	Reference string `json:"reference"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
}

// SubmitAddressResponse represents the response from submitting address
type SubmitAddressResponse = types.Response[ChargeData]

// CheckPendingChargeResponse represents the response from checking a pending charge
type CheckPendingChargeResponse = types.Response[ChargeData]

// ChargeData represents the charge data in API responses
type ChargeData struct {
	ID              int             `json:"id"`
	Domain          string          `json:"domain"`
	Status          string          `json:"status"`
	Reference       string          `json:"reference"`
	Amount          int             `json:"amount"`
	Message         string          `json:"message"`
	GatewayResponse string          `json:"gateway_response"`
	PaidAt          *types.DateTime `json:"paid_at"`
	CreatedAt       *types.DateTime `json:"created_at"`
	Channel         string          `json:"channel"`
	Currency        string          `json:"currency"`
	IPAddress       string          `json:"ip_address"`
	Metadata        map[string]any  `json:"metadata"`
	Log             any             `json:"log"`
	Fees            int             `json:"fees"`
	RequestedAmount int             `json:"requested_amount"`
	TransactionDate *types.DateTime `json:"transaction_date"`
	Plan            any             `json:"plan"`
	Authorization   *Authorization  `json:"authorization"`
	Customer        *Customer       `json:"customer"`
}

// Authorization represents authorization details in charge response
type Authorization struct {
	AuthorizationCode string `json:"authorization_code"`
	Bin               string `json:"bin"`
	Last4             string `json:"last4"`
	ExpMonth          string `json:"exp_month"`
	ExpYear           string `json:"exp_year"`
	Channel           string `json:"channel"`
	CardType          string `json:"card_type"`
	Bank              string `json:"bank"`
	CountryCode       string `json:"country_code"`
	Brand             string `json:"brand"`
	Reusable          bool   `json:"reusable"`
	Signature         string `json:"signature"`
	AccountName       string `json:"account_name"`
}

// Customer represents customer details in charge response
type Customer struct {
	ID                       int            `json:"id"`
	FirstName                string         `json:"first_name"`
	LastName                 string         `json:"last_name"`
	Email                    string         `json:"email"`
	CustomerCode             string         `json:"customer_code"`
	Phone                    string         `json:"phone"`
	Metadata                 map[string]any `json:"metadata"`
	RiskAction               string         `json:"risk_action"`
	InternationalFormatPhone string         `json:"international_format_phone"`
}
