package webhookendpoints

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEventsCatalog_JSONDeserialization(t *testing.T) {
	var rsp EventsCatalogResponse
	decodeFixture(t, "events_catalog_200.json", &rsp)

	require.True(t, rsp.Status.Bool())
	require.Len(t, rsp.Data["transactions"], 2)
	assert.Equal(t, "charge.failed", rsp.Data["transactions"][0].EventName.String())
	assert.False(t, rsp.Data["transactions"][1].Description.Valid)
	require.Len(t, rsp.Data["refunds"], 1)
	assert.Equal(t, "A processed refund", rsp.Data["refunds"][0].Description.String())
}

func TestEventsCatalog_RequestShape(t *testing.T) {
	c, rec := recordingClient(t, `{"status":true,"message":"ok","data":{}}`)

	_, err := c.EventsCatalog(context.Background())
	require.NoError(t, err)

	assert.Equal(t, "GET", rec.method)
	assert.Equal(t, "/integration/webhooks/events-catalog", rec.rawPath)
}
