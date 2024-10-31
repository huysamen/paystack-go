package net

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/huysamen/paystack-go/types"
)

const apiURL = "https://api.paystack.co"

// todo: can probably reduce json unmarshalling to single method

func Get[O any](client *http.Client, secret, path string) (*types.Response[O], error) {
	body, err := doReq(client, http.MethodGet, secret, path, nil)
	if err != nil {
		return nil, err
	}

	rsp := new(types.Response[O])

	if len(body) > 0 {
		err = json.Unmarshal(body, rsp)
		if err != nil {
			return nil, err
		}
	}

	return rsp, nil
}

func Put[I any, O any](client *http.Client, secret, path string, payload *I) (*types.Response[O], error) {
	return putOrPost[I, O](client, http.MethodPut, secret, path, payload)
}

func Post[I any, O any](client *http.Client, secret, path string, payload *I) (*types.Response[O], error) {
	return putOrPost[I, O](client, http.MethodPost, secret, path, payload)
}

func putOrPost[I any, O any](client *http.Client, method, secret, path string, payload *I) (*types.Response[O], error) {
	body, err := doReq(client, method, secret, path, payload)
	if err != nil {
		return nil, err
	}

	rsp := new(types.Response[O])
	data := new(O)

	if len(body) > 0 {
		err = json.Unmarshal(body, data)
		if err != nil {
			return nil, err
		}
	}

	rsp.Data = *data

	return rsp, nil
}

func doReq(client *http.Client, method, secret, path string, data any) ([]byte, error) {
	var req *http.Request
	var err error

	if data != nil {
		d, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method, apiURL+path, bytes.NewBuffer(d))
	} else {
		req, err = http.NewRequest(method, apiURL+path, nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+secret)

	if data != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rsp.Body.Close() }()

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	switch rsp.StatusCode {
	case http.StatusOK:
		return body, nil
	default:
		// todo
	}

	return nil, nil
}
