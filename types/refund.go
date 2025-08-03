package types

// RefundStatus represents the status of a refund
type RefundStatus string

const (
	RefundStatusPending   RefundStatus = "pending"
	RefundStatusProcessed RefundStatus = "processed"
	RefundStatusFailed    RefundStatus = "failed"
)

// String returns the string representation of RefundStatus
func (s RefundStatus) String() string {
	return string(s)
}

// RefundChannel represents the payment channel for a refund
type RefundChannel string

const (
	RefundChannelCard         RefundChannel = "card"
	RefundChannelBank         RefundChannel = "bank"
	RefundChannelUSSD         RefundChannel = "ussd"
	RefundChannelQR           RefundChannel = "qr"
	RefundChannelMobileMoney  RefundChannel = "mobile_money"
	RefundChannelBankTransfer RefundChannel = "bank_transfer"
	RefundChannelApplePay     RefundChannel = "apple_pay"
	RefundChannelMigs         RefundChannel = "migs"
)

// String returns the string representation of RefundChannel
func (c RefundChannel) String() string {
	return string(c)
}

// Refund represents a refund object
type Refund struct {
	ID             int           `json:"id"`
	Integration    int           `json:"integration"`
	Domain         string        `json:"domain"`
	Transaction    int           `json:"transaction"`
	Dispute        *int          `json:"dispute"`
	Settlement     *int          `json:"settlement"`
	Amount         int           `json:"amount"`
	DeductedAmount int           `json:"deducted_amount"`
	Currency       string        `json:"currency"`
	Channel        RefundChannel `json:"channel"`
	FullyDeducted  bool          `json:"fully_deducted"`
	Status         RefundStatus  `json:"status"`
	RefundedBy     string        `json:"refunded_by"`
	RefundedAt     *DateTime     `json:"refunded_at"`
	ExpectedAt     *DateTime     `json:"expected_at"`
	CreatedAt      *DateTime     `json:"created_at"`
	UpdatedAt      *DateTime     `json:"updated_at"`
	CustomerNote   *string       `json:"customer_note"`
	MerchantNote   *string       `json:"merchant_note"`
}
