package disputes

import (
	"github.com/huysamen/paystack-go/types"
)

// TransactionDisputeData represents transaction dispute data
type TransactionDisputeData struct {
	History  []types.DisputeHistory `json:"history"`
	Messages []types.DisputeMessage `json:"messages"`
	Dispute  *types.Dispute         `json:"dispute,omitempty"`
}

// UploadURLData represents upload URL data
type UploadURLData struct {
	SignedURL string `json:"signedUrl"`
	FileName  string `json:"fileName"`
	ExpiresIn int    `json:"expiresIn"`
}

// ExportData represents export data
type ExportData struct {
	Path      string          `json:"path"`
	ExpiresAt *types.DateTime `json:"expires_at,omitempty"`
}
