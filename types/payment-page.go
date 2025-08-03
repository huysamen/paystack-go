package types

// PaymentPage represents a payment page
type PaymentPage struct {
	ID                int           `json:"id"`
	Name              string        `json:"name"`
	Description       string        `json:"description,omitempty"`
	Amount            *int          `json:"amount,omitempty"`
	Currency          string        `json:"currency"`
	Slug              string        `json:"slug"`
	Type              string        `json:"type"`
	FixedAmount       bool          `json:"fixed_amount"`
	RedirectURL       string        `json:"redirect_url,omitempty"`
	SuccessMessage    string        `json:"success_message,omitempty"`
	NotificationEmail string        `json:"notification_email,omitempty"`
	CollectPhone      bool          `json:"collect_phone"`
	Active            bool          `json:"active"`
	Published         bool          `json:"published"`
	Migrate           bool          `json:"migrate"`
	CustomFields      []CustomField `json:"custom_fields,omitempty"`
	SplitCode         string        `json:"split_code,omitempty"`
	Plan              *int          `json:"plan,omitempty"`
	Products          []Product     `json:"products,omitempty"`
	Integration       int           `json:"integration"`
	Domain            string        `json:"domain"`
	Metadata          *Metadata     `json:"metadata,omitempty"`
	CreatedAt         string        `json:"createdAt,omitempty"`
	UpdatedAt         string        `json:"updatedAt,omitempty"`
}
