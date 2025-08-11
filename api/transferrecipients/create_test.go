package transferrecipients

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/huysamen/paystack-go/enums"
	"github.com/stretchr/testify/require"
)

func TestTransferRecipients_Create_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transferrecipients", "create_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp CreateResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
}

func TestTransferRecipients_Create_Builder(t *testing.T) {
	_ = NewCreateRequestBuilder(enums.TransferRecipientTypeNuban, "name", "0123", "058").Build()
}
