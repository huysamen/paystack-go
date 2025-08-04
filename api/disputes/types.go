package disputes

import (
	"github.com/huysamen/paystack-go/types"
)

type TransactionDisputeData struct {
	History  []types.DisputeHistory `json:"history"`
	Messages []types.DisputeMessage `json:"messages"`
	Dispute  *types.Dispute         `json:"dispute,omitempty"`
}

type UploadURLData struct {
	SignedURL string `json:"signedUrl"`
	FileName  string `json:"fileName"`
	ExpiresIn int    `json:"expiresIn"`
}

type ExportData struct {
	Path      string          `json:"path"`
	ExpiresAt *types.DateTime `json:"expires_at,omitempty"`
}
