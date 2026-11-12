package webhookendpoints

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestList_JSONDeserialization(t *testing.T) {
	var rsp ListResponse
	decodeFixture(t, "list_200.json", &rsp)

	require.True(t, rsp.Status.Bool())
	require.Len(t, rsp.Data, 1)
	ep := rsp.Data[0]
	assert.Equal(t, "WHK_1a2b3c4d", ep.WebhookURLCode.String())
	assert.Equal(t, "Order fulfillment", ep.Name.String())
	assert.Equal(t, "https://example.com/webhooks/paystack", ep.WebhookURL.String())
	assert.True(t, ep.Enabled.Bool())
	assert.Equal(t, int64(2), ep.EventCount.Int64())

	require.NotNil(t, rsp.Meta)
	assert.Equal(t, "V0hLXzk5", rsp.Meta.Next.String())
	assert.False(t, rsp.Meta.Previous.Valid)
	assert.Equal(t, int64(20), rsp.Meta.PerPage.Int64())
}

func TestList_RequestShape(t *testing.T) {
	c, rec := recordingClient(t, `{"status":true,"message":"ok","data":[]}`)

	_, err := c.List(context.Background(), *NewListRequestBuilder().Limit(20).Next("abc"))
	require.NoError(t, err)

	assert.Equal(t, "GET", rec.method)
	assert.Equal(t, "/integration/webhooks/endpoints", rec.rawPath)
	assert.Equal(t, "limit=20&next=abc", rec.query)
}

func TestList_NoParams(t *testing.T) {
	c, rec := recordingClient(t, `{"status":true,"message":"ok","data":[]}`)

	_, err := c.List(context.Background(), *NewListRequestBuilder())
	require.NoError(t, err)

	assert.Empty(t, rec.query)
}

func TestList_Previous(t *testing.T) {
	c, rec := recordingClient(t, `{"status":true,"message":"ok","data":[]}`)

	_, err := c.List(context.Background(), *NewListRequestBuilder().Previous("xyz"))
	require.NoError(t, err)

	assert.Equal(t, "previous=xyz", rec.query)
}
