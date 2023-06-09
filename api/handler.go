package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIHandler struct {
	key           string
	defaultRegion string
}

func NewAPIHandler(key, defaultRegion string) *APIHandler {
	return &APIHandler{
		key:           key,
		defaultRegion: defaultRegion,
	}
}

func (api *APIHandler) GetDefaultRegion() string {
	return api.defaultRegion
}

func (api *APIHandler) Get(ctx context.Context, url string) (*http.Response, error) {
	return api.makeAPICall(ctx, http.MethodGet, url, nil)
}

func (api *APIHandler) Post(ctx context.Context, url string, data interface{}) (*http.Response, error) {
	return api.makeAPICall(ctx, http.MethodPost, url, data)
}

func (api *APIHandler) Delete(ctx context.Context, url string) (*http.Response, error) {
	return api.makeAPICall(ctx, http.MethodDelete, url, nil)
}

func (api *APIHandler) makeAPICall(ctx context.Context, method, url string, data interface{}) (*http.Response, error) {
	var reader io.Reader
	if data != nil {
		raw, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(raw)
	}
	httpReq, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("https://cloud.lambdalabs.com/api/v1/%s", url), reader)
	if err != nil {
		return nil, err
	}

	httpReq.SetBasicAuth(api.key, "")
	res, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		// Get the error
		defer res.Body.Close()
		// should be json
		apiError := &APIError{}
		if err := json.NewDecoder(res.Body).Decode(&apiError); err != nil {
			return nil, err
		}
		return nil, apiError
	}
	return res, nil
}
