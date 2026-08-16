package paystack

import (
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
		expected int64
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
		{"50.00 with 1.95%", 50.00, 1.95, 0.98},
		{"100.00 with 1.95%", 100.00, 1.95, 1.95},
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
		{"50.00 with 1.95%", 50.00, 1.95, 50.98},
		{"100.00 with 1.95%", 100.00, 1.95, 101.95},
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
		expected      int64
	}{
		{"50.00 with 1.95%", 50.00, 1.95, 5098},
		{"100.00 with 1.95%", 100.00, 1.95, 10195},
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

	assert.Equal(t, int64(5000), subunits)
	assert.Equal(t, int64(5098), totalSubunits)
	assert.Equal(t, int64(98), totalSubunits-subunits)
}
