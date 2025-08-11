package transactions

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactions_Export_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transactions", "export_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp ExportResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)

	assert.True(t, rsp.Status.Bool())
	assert.NotNil(t, rsp.Data)
}

func TestTransactions_Export_Builder(t *testing.T) {
	from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)
	q := NewExportRequestBuilder().PerPage(10).Page(1).From(from).To(to).Customer(1).Status("success").Amount(1000).Settled(true).Settlement(1).PaymentPage(2).Build().toQuery()
	assert.Contains(t, q, "perPage=10")
	assert.Contains(t, q, "page=1")
	assert.Contains(t, q, "customer=1")
	assert.Contains(t, q, "status=success")
	assert.Contains(t, q, "amount=1000")
	assert.Contains(t, q, "settled=true")
}
