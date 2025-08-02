package types

// Customer represents a Paystack customer
type Customer struct {
	ID                       uint64   `json:"id"`
	FirstName                string   `json:"first_name"`
	LastName                 string   `json:"last_name"`
	Email                    string   `json:"email"`
	CustomerCode             string   `json:"customer_code"`
	Phone                    string   `json:"phone"`
	Metadata                 Metadata `json:"metadata"`
	RiskAction               string   `json:"risk_action"`
	InternationalFormatPhone string   `json:"international_format_phone"`
}
