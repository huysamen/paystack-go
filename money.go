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

func (a amount) Fee() float64 {
	rawFee := a.value * a.feePercentage / 100
	return math.Round(rawFee*100) / 100
}

func (a amount) Total() float64 {
	return math.Round((a.value+a.Fee())*100) / 100
}

func (a amount) TotalSubunits() int {
	return int(math.Round(a.Total() * 100))
}
