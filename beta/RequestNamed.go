// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// NamedLocationRequestBuilder is request builder for NamedLocation
type NamedLocationRequestBuilder struct{ BaseRequestBuilder }

// Request returns NamedLocationRequest
func (b *NamedLocationRequestBuilder) Request() *NamedLocationRequest {
	return &NamedLocationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// NamedLocationRequest is request for NamedLocation
type NamedLocationRequest struct{ BaseRequest }

// Get performs GET request for NamedLocation
func (r *NamedLocationRequest) Get(ctx context.Context) (resObj *NamedLocation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for NamedLocation
func (r *NamedLocationRequest) Update(ctx context.Context, reqObj *NamedLocation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for NamedLocation
func (r *NamedLocationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for NamedLocation
func (r *NamedLocationRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj NamedLocation
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for NamedLocation
func (r *NamedLocationRequest) BatchUpdate(batch *BatchRequest, reqObj *NamedLocation) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for NamedLocation
func (r *NamedLocationRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
