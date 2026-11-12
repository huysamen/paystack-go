package webhookendpoints

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPing_JSONDeserialization(t *testing.T) {
	var rsp PingResponse
	decodeFixture(t, "ping_200.json", &rsp)

	require.True(t, rsp.Status.Bool())
	assert.True(t, rsp.Data.OK.Bool())
	require.True(t, rsp.Data.HTTPStatus.Valid)
	assert.Equal(t, int64(200), rsp.Data.HTTPStatus.Int)
	assert.Equal(t, int64(143), rsp.Data.LatencyMS.Int64())
}

func TestPing_UnreachableDeserializes(t *testing.T) {
	var rsp PingResponse
	decodeFixture(t, "ping_unreachable_200.json", &rsp)

	assert.False(t, rsp.Data.OK.Bool())
	assert.False(t, rsp.Data.HTTPStatus.Valid)
}

func TestPing_429Deserializes(t *testing.T) {
	var rsp PingResponse
	decodeFixture(t, "ping_429.json", &rsp)

	assert.False(t, rsp.Status.Bool())
	assert.NotEmpty(t, rsp.Message)
}

func TestPing_RequestShape(t *testing.T) {
	c, rec := recordingClient(t, `{"status":true,"message":"ok","data":{}}`)

	_, err := c.Ping(context.Background(), "WHK_1")
	require.NoError(t, err)

	assert.Equal(t, "POST", rec.method)
	assert.Equal(t, "/integration/webhooks/endpoints/WHK_1/ping", rec.rawPath)
	assert.Empty(t, rec.body)
}
