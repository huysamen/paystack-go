package transferscontrol

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransfersControl_ResendOTP_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transferscontrol", "resend_otp_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp ResendOTPResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
	assert.True(t, rsp.Status.Bool())
}

func TestTransfersControl_ResendOTP_Builder(t *testing.T) {
	r := NewResendOTPRequestBuilder().TransferCode("TRF_x").Reason("resend").Build()
	assert.Equal(t, "TRF_x", r.TransferCode)
	assert.Equal(t, "resend", r.Reason)
}
