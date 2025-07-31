package virtualterminal

import (
	"fmt"
	"net/http"
)

const virtualTerminalBasePath = "/virtual_terminal"

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new virtual terminal client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}

// Validation functions

func validateCreateRequest(req *CreateVirtualTerminalRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}

func validateUpdateRequest(req *UpdateVirtualTerminalRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}

func validateAssignDestinationRequest(req *AssignDestinationRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if len(req.Destinations) == 0 {
		return fmt.Errorf("destinations are required")
	}
	for i, dest := range req.Destinations {
		if dest.Target == "" {
			return fmt.Errorf("destination target is required at index %d", i)
		}
		if dest.Name == "" {
			return fmt.Errorf("destination name is required at index %d", i)
		}
	}
	return nil
}

func validateUnassignDestinationRequest(req *UnassignDestinationRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if len(req.Targets) == 0 {
		return fmt.Errorf("targets are required")
	}
	return nil
}

func validateAddSplitCodeRequest(req *AddSplitCodeRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.SplitCode == "" {
		return fmt.Errorf("split code is required")
	}
	return nil
}

func validateRemoveSplitCodeRequest(req *RemoveSplitCodeRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.SplitCode == "" {
		return fmt.Errorf("split code is required")
	}
	return nil
}

func validateCode(code string) error {
	if code == "" {
		return fmt.Errorf("virtual terminal code is required")
	}
	return nil
}
