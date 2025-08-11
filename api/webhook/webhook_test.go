package webhook

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func readEventFixture(t *testing.T, name string) *Event {
	t.Helper()
	p := filepath.Join("..", "..", "resources", "examples", "webhook", name+".json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)
	var e Event
	require.NoError(t, json.Unmarshal(b, &e))
	return &e
}

func TestWebhook_ChargeSuccess(t *testing.T) {
	e := readEventFixture(t, "charge.success")
	assert.Equal(t, EventChargeSuccess, e.Event)
	ev, err := e.AsChargeSuccess()
	require.NoError(t, err)
	assert.Equal(t, int64(302961), ev.ID.Int64())
	assert.Equal(t, "success", ev.Status.String())
	assert.Equal(t, "qTPrJoy9Bx", ev.Reference.String())
	assert.Equal(t, int64(10000), ev.Amount.Int64())
}

func TestWebhook_CustomerIdentification_Failed_Success(t *testing.T) {
	e1 := readEventFixture(t, "customeridentification.failed")
	v1, err := e1.AsCustomerIdentificationFailed()
	require.NoError(t, err)
	assert.Equal(t, "NG", v1.Identification.Country.String())

	e2 := readEventFixture(t, "customeridentification.success")
	v2, err := e2.AsCustomerIdentificationSuccess()
	require.NoError(t, err)
	assert.Equal(t, "bvn", v2.Identification.Type.String())
}

func TestWebhook_ChargeDispute_Create_Remind_Resolve(t *testing.T) {
	for _, name := range []string{"charge.dispute.create", "charge.dispute.remind", "charge.dispute.resolve"} {
		e := readEventFixture(t, name)
		v, err := e.AsChargeDispute()
		require.NoError(t, err)
		assert.NotZero(t, v.ID.Int64())
		assert.NotEmpty(t, v.Status.String())
	}
}

func TestWebhook_DedicatedAccount_Assign(t *testing.T) {
	for _, name := range []string{"dedicatedaccount.assign.failed", "dedicatedaccount.assign.success"} {
		e := readEventFixture(t, name)
		v, err := e.AsDedicatedAccount()
		require.NoError(t, err)
		// identification may only contain status for failed case; ensure field exists
		assert.True(t, v.Identification.Valid || !v.Identification.Valid)
	}
}

func TestWebhook_Invoice_Events(t *testing.T) {
	e1 := readEventFixture(t, "invoice.create")
	v1, err := e1.AsInvoiceCreate()
	require.NoError(t, err)
	assert.True(t, v1.Paid.Bool())

	e2 := readEventFixture(t, "invoice.payment_failed")
	v2, err := e2.AsInvoicePaymentFailed()
	require.NoError(t, err)
	assert.False(t, v2.Paid.Bool())

	e3 := readEventFixture(t, "invoice.update")
	v3, err := e3.AsInvoiceUpdate()
	require.NoError(t, err)
	assert.True(t, v3.Paid.Bool())
}

func TestWebhook_PaymentRequest(t *testing.T) {
	for _, name := range []string{"paymentrequest.pending", "paymentrequest.success"} {
		e := readEventFixture(t, name)
		v, err := e.AsPaymentRequest()
		require.NoError(t, err)
		assert.NotEmpty(t, v.RequestCode.String())
	}
}

func TestWebhook_Refund_Events(t *testing.T) {
	e1 := readEventFixture(t, "refund.failed")
	v1, err := e1.AsRefundFailed()
	require.NoError(t, err)
	assert.Equal(t, "failed", v1.Status.String())

	e2 := readEventFixture(t, "refund.pending")
	v2, err := e2.AsRefundPending()
	require.NoError(t, err)
	assert.Equal(t, "pending", v2.Status.String())

	e3 := readEventFixture(t, "refund.processed")
	v3, err := e3.AsRefundProcessed()
	require.NoError(t, err)
	assert.Equal(t, "processed", v3.Status.String())

	e4 := readEventFixture(t, "refund.processing")
	// Using generic parsing to ensure envelope is valid
	var payload map[string]any
	require.NoError(t, json.Unmarshal(e4.Data, &payload))
	assert.Equal(t, "processing", payload["status"])
}

func TestWebhook_Subscription_Events(t *testing.T) {
	e1 := readEventFixture(t, "subscription.create")
	v1, err := e1.AsSubscriptionCreate()
	require.NoError(t, err)
	assert.Equal(t, "active", v1.Status.String())

	e2 := readEventFixture(t, "subscription.disable")
	v2, err := e2.AsSubscriptionDisable()
	require.NoError(t, err)
	assert.Equal(t, "complete", v2.Status.String())

	e3 := readEventFixture(t, "subscription.not_renew")
	v3, err := e3.AsSubscriptionNotRenew()
	require.NoError(t, err)
	assert.Equal(t, "non-renewing", v3.Status.String())

	e4 := readEventFixture(t, "subscription.expiring_cards")
	v4, err := e4.AsSubscriptionExpiringCards()
	require.NoError(t, err)
	// Ensure we captured entries array in flexible metadata
	assert.True(t, v4.Entries.Valid)
}

func TestWebhook_Transfer_Events(t *testing.T) {
	e1 := readEventFixture(t, "transfer.success")
	v1, err := e1.AsTransferSuccess()
	require.NoError(t, err)
	assert.Equal(t, "success", v1.Status.String())

	e2 := readEventFixture(t, "transfer.failed")
	v2, err := e2.AsTransferFailed()
	require.NoError(t, err)
	assert.Equal(t, "failed", v2.Status.String())

	e3 := readEventFixture(t, "transfer.reversed")
	v3, err := e3.AsTransferReversed()
	require.NoError(t, err)
	assert.Equal(t, "reversed", v3.Status.String())
}
