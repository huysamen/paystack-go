package transferscontrol

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransfersControl_FinalizeDisableOTP_JSONDeserialization(t *testing.T) {
	p := filepath.Join("..", "..", "resources", "examples", "responses", "transferscontrol", "finalize_otp_200.json")
	b, err := os.ReadFile(p)
	require.NoError(t, err)

	var rsp FinalizeDisableOTPResponse
	err = json.Unmarshal(b, &rsp)
	require.NoError(t, err)
	assert.True(t, rsp.Status.Bool())
}

func TestTransfersControl_FinalizeDisableOTP_Builder(t *testing.T) {
	r := NewFinalizeDisableOTPRequestBuilder().OTP("123456").Build()
	assert.Equal(t, "123456", r.OTP)
}
