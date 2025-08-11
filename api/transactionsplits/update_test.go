package transactionsplits

import (
	"testing"

	"github.com/huysamen/paystack-go/enums"
	"github.com/stretchr/testify/assert"
)

func TestTransactionSplits_Update_Builder(t *testing.T) {
	ub := NewUpdateRequestBuilder().Name("N").Active(true).BearerType(enums.TransactionSplitBearerTypeAccount).BearerSubaccount("SUB_X").Build()
	assert.NotNil(t, ub.Name)
}
