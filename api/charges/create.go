package charges

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/utils"
)

// Create initiates a payment by integrating multiple payment channels
func (c *Client) Create(ctx context.Context, req *CreateChargeRequest) (*CreateChargeResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("create charge request cannot be nil")
	}

	if err := validateCreateChargeRequest(req); err != nil {
		return nil, err
	}

	url := c.baseURL + chargesBasePath
	return net.Post[CreateChargeRequest, ChargeData](ctx, c.client, c.secret, url, req)
}

// validateCreateChargeRequest validates the create charge request
func validateCreateChargeRequest(req *CreateChargeRequest) error {
	if req.Email == "" {
		return fmt.Errorf("email is required")
	}

	if req.Amount == "" {
		return fmt.Errorf("amount is required")
	}

	if err := utils.ValidateEmail(req.Email); err != nil {
		return fmt.Errorf("invalid email: %w", err)
	}

	// Validate that at least one payment method is provided
	hasPaymentMethod := req.Bank != nil ||
		req.BankTransfer != nil ||
		req.USSD != nil ||
		req.MobileMoney != nil ||
		req.QR != nil ||
		req.AuthorizationCode != nil

	if !hasPaymentMethod {
		return fmt.Errorf("at least one payment method must be provided (bank, bank_transfer, ussd, mobile_money, qr, or authorization_code)")
	}

	// Validate bank details if provided
	if req.Bank != nil {
		if req.Bank.Code == "" {
			return fmt.Errorf("bank code is required when bank details are provided")
		}
		if req.Bank.AccountNumber == "" {
			return fmt.Errorf("bank account number is required when bank details are provided")
		}
	}

	// Validate USSD details if provided
	if req.USSD != nil {
		if req.USSD.Type == "" {
			return fmt.Errorf("USSD type is required when USSD details are provided")
		}
	}

	// Validate mobile money details if provided
	if req.MobileMoney != nil {
		if req.MobileMoney.Phone == "" {
			return fmt.Errorf("phone number is required when mobile money details are provided")
		}
		if req.MobileMoney.Provider == "" {
			return fmt.Errorf("provider is required when mobile money details are provided")
		}
	}

	// Validate QR details if provided
	if req.QR != nil {
		if req.QR.Provider == "" {
			return fmt.Errorf("provider is required when QR details are provided")
		}
		// Validate allowed QR providers
		validProviders := map[string]bool{
			"scan-to-pay": true,
			"visa":        true,
		}
		if !validProviders[req.QR.Provider] {
			return fmt.Errorf("invalid QR provider: %s. Valid providers are: scan-to-pay, visa", req.QR.Provider)
		}
	}

	// Validate bearer if provided
	if req.Bearer != nil {
		validBearers := map[string]bool{
			"account":    true,
			"subaccount": true,
		}
		if !validBearers[*req.Bearer] {
			return fmt.Errorf("invalid bearer: %s. Valid bearers are: account, subaccount", *req.Bearer)
		}
	}

	return nil
}
