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
