// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// EntityRequestBuilder is request builder for Entity
type EntityRequestBuilder struct{ BaseRequestBuilder }

// Request returns EntityRequest
func (b *EntityRequestBuilder) Request() *EntityRequest {
	return &EntityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EntityRequest is request for Entity
type EntityRequest struct{ BaseRequest }

// Get performs GET request for Entity
func (r *EntityRequest) Get(ctx context.Context) (resObj *Entity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Entity
func (r *EntityRequest) Update(ctx context.Context, reqObj *Entity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Entity
func (r *EntityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for Entity
func (r *EntityRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj Entity
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for Entity
func (r *EntityRequest) BatchUpdate(batch *BatchRequest, reqObj *Entity) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for Entity
func (r *EntityRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
