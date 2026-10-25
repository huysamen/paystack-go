package paystack

import (
	"math"
	"testing"

	"github.com/huysamen/paystack-go/enums"
	"github.com/stretchr/testify/assert"
)

func TestNewMoney(t *testing.T) {
	money := NewMoney(MoneyConfig{
		Amount:        50.00,
		Currency:      enums.CurrencyGHS,
		FeePercentage: 1.95,
	})

	assert.Equal(t, enums.CurrencyGHS, money.Currency)
	assert.Equal(t, 1.95, money.FeePercentage)
}

func TestAmount_Subunits(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		expected int
	}{
		{"50.00", 50.00, 5000},
		{"0.00", 0.00, 0},
		{"0.01", 0.01, 1},
		{"0.10", 0.10, 10},
		{"100.99", 100.99, 10099},
		{"-50.00", -50.00, -5000},
		{"1.005", 1.005, 100},
		{"1.004", 1.004, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			money := NewMoney(MoneyConfig{
				Amount:        tt.amount,
				Currency:      enums.CurrencyGHS,
				FeePercentage: 1.95,
			})

			assert.Equal(t, tt.expected, money.Amount.Subunits())
		})
	}
}

func TestAmount_Fee(t *testing.T) {
	tests := []struct {
		name          string
		amount        float64
		feePercentage float64
		expected      float64
	}{
		{"50.00 with 1.95%", 50.00, 1.95, 1.00},
		{"100.00 with 1.95%", 100.00, 1.95, 1.99},
		{"0.00 with 1.95%", 0.00, 1.95, 0.00},
		{"50.00 with 0%", 50.00, 0.00, 0.00},
		{"1.005 with 1.95%", 1.005, 1.95, 0.02},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			money := NewMoney(MoneyConfig{
				Amount:        tt.amount,
				Currency:      enums.CurrencyGHS,
				FeePercentage: tt.feePercentage,
			})

			assert.InDelta(t, tt.expected, money.Amount.Fee(), 0.001)
		})
	}
}

func TestAmount_Total(t *testing.T) {
	tests := []struct {
		name          string
		amount        float64
		feePercentage float64
		expected      float64
	}{
		{"50.00 with 1.95%", 50.00, 1.95, 51.00},
		{"100.00 with 1.95%", 100.00, 1.95, 101.99},
		{"0.00 with 1.95%", 0.00, 1.95, 0.00},
		{"50.00 with 0%", 50.00, 0.00, 50.00},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			money := NewMoney(MoneyConfig{
				Amount:        tt.amount,
				Currency:      enums.CurrencyGHS,
				FeePercentage: tt.feePercentage,
			})

			assert.InDelta(t, tt.expected, money.Amount.Total(), 0.001)
		})
	}
}

func TestAmount_TotalSubunits(t *testing.T) {
	tests := []struct {
		name          string
		amount        float64
		feePercentage float64
		expected      int
	}{
		{"50.00 with 1.95%", 50.00, 1.95, 5100},
		{"100.00 with 1.95%", 100.00, 1.95, 10199},
		{"0.00 with 1.95%", 0.00, 1.95, 0},
		{"50.00 with 0%", 50.00, 0.00, 5000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			money := NewMoney(MoneyConfig{
				Amount:        tt.amount,
				Currency:      enums.CurrencyGHS,
				FeePercentage: tt.feePercentage,
			})

			assert.Equal(t, tt.expected, money.Amount.TotalSubunits())
		})
	}
}

func TestMoney_Consistency(t *testing.T) {
	money := NewMoney(MoneyConfig{
		Amount:        50.00,
		Currency:      enums.CurrencyGHS,
		FeePercentage: 1.95,
	})

	subunits := money.Amount.Subunits()
	totalSubunits := money.Amount.TotalSubunits()

	assert.Equal(t, 5000, subunits)
	assert.Equal(t, 5100, totalSubunits)
	assert.Equal(t, 100, totalSubunits-subunits)
}

func TestMoney_FeeMatchesTotal(t *testing.T) {
	amounts := []float64{0.01, 0.29, 1.15, 2.675, 10.075, 19.99, 50.00, 1234.565, 99999.99}
	feePercentages := []float64{0, 1.5, 1.95, 3.9}

	for _, value := range amounts {
		for _, feePercentage := range feePercentages {
			money := NewMoney(MoneyConfig{
				Amount:        value,
				Currency:      enums.CurrencyGHS,
				FeePercentage: feePercentage,
			})

			subunits := money.Amount.Subunits()
			totalSubunits := money.Amount.TotalSubunits()

			// Fee and Total must agree with the subunit values actually charged.
			assert.Equal(t, totalSubunits-subunits, int(math.Round(money.Amount.Fee()*100)))
			assert.Equal(t, totalSubunits, int(math.Round(money.Amount.Total()*100)))

			// After Paystack deducts its percentage, the merchant nets at least
			// the amount, and no more than one subunit over.
			net := float64(totalSubunits) * (1 - feePercentage/100)
			assert.GreaterOrEqual(t, net, float64(subunits)-1e-6)
			assert.Less(t, net, float64(subunits)+1)
		}
	}
}
