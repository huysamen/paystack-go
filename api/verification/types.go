package verification

// AccountResolution represents the result of resolving an account number
type AccountResolution struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
}

// AccountValidation represents the result of validating an account
type AccountValidation struct {
	Verified            bool   `json:"verified"`
	VerificationMessage string `json:"verificationMessage"`
}

// CardBINResolution represents the result of resolving a card BIN
type CardBINResolution struct {
	BIN          string `json:"bin"`
	Brand        string `json:"brand"`
	SubBrand     string `json:"sub_brand"`
	CountryCode  string `json:"country_code"`
	CountryName  string `json:"country_name"`
	CardType     string `json:"card_type"`
	Bank         string `json:"bank"`
	LinkedBankID int    `json:"linked_bank_id"`
}

// Account Resolution

// AccountResolveRequest represents the request to resolve an account number
type AccountResolveRequest struct {
	AccountNumber string `json:"account_number"` // Required: account number
	BankCode      string `json:"bank_code"`      // Required: bank code
}

// AccountResolveResponse represents the response from resolving an account
type AccountResolveResponse struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Data    AccountResolution `json:"data"`
}

// Account Validation

// AccountValidateRequest represents the request to validate an account
type AccountValidateRequest struct {
	AccountName    string  `json:"account_name"`              // Required: customer's name
	AccountNumber  string  `json:"account_number"`            // Required: account number
	AccountType    string  `json:"account_type"`              // Required: personal or business
	BankCode       string  `json:"bank_code"`                 // Required: bank code
	CountryCode    string  `json:"country_code"`              // Required: two digit ISO code
	DocumentType   string  `json:"document_type"`             // Required: identity document type
	DocumentNumber *string `json:"document_number,omitempty"` // Optional: identity document number
}

// AccountValidateRequestBuilder provides a fluent interface for building AccountValidateRequest
type AccountValidateRequestBuilder struct {
	req *AccountValidateRequest
}

// NewAccountValidateRequest creates a new builder for AccountValidateRequest
func NewAccountValidateRequest(accountName, accountNumber, accountType, bankCode, countryCode, documentType string) *AccountValidateRequestBuilder {
	return &AccountValidateRequestBuilder{
		req: &AccountValidateRequest{
			AccountName:   accountName,
			AccountNumber: accountNumber,
			AccountType:   accountType,
			BankCode:      bankCode,
			CountryCode:   countryCode,
			DocumentType:  documentType,
		},
	}
}

// DocumentNumber sets the optional document number
func (b *AccountValidateRequestBuilder) DocumentNumber(documentNumber string) *AccountValidateRequestBuilder {
	b.req.DocumentNumber = &documentNumber
	return b
}

// Build returns the constructed AccountValidateRequest
func (b *AccountValidateRequestBuilder) Build() *AccountValidateRequest {
	return b.req
}

// AccountValidateResponse represents the response from validating an account
type AccountValidateResponse struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Data    AccountValidation `json:"data"`
}

// Card BIN Resolution

// CardBINResolveResponse represents the response from resolving a card BIN
type CardBINResolveResponse struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Data    CardBINResolution `json:"data"`
}
