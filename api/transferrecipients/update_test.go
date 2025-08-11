package transferrecipients

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransferRecipients_Update_JSONDeserialization(t *testing.T) {
	// Use create response as the update response envelope is the same (single recipient)
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transferrecipients", "create_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	// Reuse list response shape to ensure struct compatibility for update response data
	var rsp UpdateResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}

func TestTransferRecipients_Update_Builder(t *testing.T) {
	r := NewUpdateRequestBuilder("name").Email("e@x.com").Build()
	assert.Equal(t, "name", r.Name)
}
