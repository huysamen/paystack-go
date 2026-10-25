package paystack

import (
	"math"

	"github.com/huysamen/paystack-go/enums"
)

type Currency = enums.Currency

const (
	CurrencyZAR Currency = enums.CurrencyZAR
	CurrencyNGN Currency = enums.CurrencyNGN
	CurrencyUSD Currency = enums.CurrencyUSD
	CurrencyGHS Currency = enums.CurrencyGHS
	CurrencyKES Currency = enums.CurrencyKES
)

type MoneyConfig struct {
	Amount        float64
	Currency      enums.Currency
	FeePercentage float64
}

type Money struct {
	Amount        amount
	Currency      enums.Currency
	FeePercentage float64
}

type amount struct {
	value         float64
	feePercentage float64
}

func NewMoney(config MoneyConfig) *Money {
	return &Money{
		Amount:        amount{value: config.Amount, feePercentage: config.FeePercentage},
		Currency:      config.Currency,
		FeePercentage: config.FeePercentage,
	}
}

func (a amount) Subunits() int {
	return int(math.Round(a.value * 100))
}

// Fee is the amount added on top so that, after Paystack deducts its
// percentage from the charged total, the merchant still receives the
// full amount.
func (a amount) Fee() float64 {
	return float64(a.TotalSubunits()-a.Subunits()) / 100
}

// Total is the amount to charge the customer, fee included.
func (a amount) Total() float64 {
	return float64(a.TotalSubunits()) / 100
}

// TotalSubunits grosses the amount up as amount / (1 - fee%), since Paystack
// takes its fee from the charged total. It rounds up so the merchant never
// nets less than the amount. The epsilon keeps exact results from being
// pushed up a subunit by float error.
func (a amount) TotalSubunits() int {
	gross := float64(a.Subunits()) / (1 - a.feePercentage/100)
	return int(math.Ceil(gross - 1e-9))
}
