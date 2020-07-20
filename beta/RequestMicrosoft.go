// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// MicrosoftStoreForBusinessAppRequestBuilder is request builder for MicrosoftStoreForBusinessApp
type MicrosoftStoreForBusinessAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns MicrosoftStoreForBusinessAppRequest
func (b *MicrosoftStoreForBusinessAppRequestBuilder) Request() *MicrosoftStoreForBusinessAppRequest {
	return &MicrosoftStoreForBusinessAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MicrosoftStoreForBusinessAppRequest is request for MicrosoftStoreForBusinessApp
type MicrosoftStoreForBusinessAppRequest struct{ BaseRequest }

// Get performs GET request for MicrosoftStoreForBusinessApp
func (r *MicrosoftStoreForBusinessAppRequest) Get(ctx context.Context) (resObj *MicrosoftStoreForBusinessApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MicrosoftStoreForBusinessApp
func (r *MicrosoftStoreForBusinessAppRequest) Update(ctx context.Context, reqObj *MicrosoftStoreForBusinessApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MicrosoftStoreForBusinessApp
func (r *MicrosoftStoreForBusinessAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for MicrosoftStoreForBusinessApp
func (r *MicrosoftStoreForBusinessAppRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj MicrosoftStoreForBusinessApp
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for MicrosoftStoreForBusinessApp
func (r *MicrosoftStoreForBusinessAppRequest) BatchUpdate(batch *BatchRequest, reqObj *MicrosoftStoreForBusinessApp) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for MicrosoftStoreForBusinessApp
func (r *MicrosoftStoreForBusinessAppRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
