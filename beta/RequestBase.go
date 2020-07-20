// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// BaseItemRequestBuilder is request builder for BaseItem
type BaseItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns BaseItemRequest
func (b *BaseItemRequestBuilder) Request() *BaseItemRequest {
	return &BaseItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BaseItemRequest is request for BaseItem
type BaseItemRequest struct{ BaseRequest }

// Get performs GET request for BaseItem
func (r *BaseItemRequest) Get(ctx context.Context) (resObj *BaseItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BaseItem
func (r *BaseItemRequest) Update(ctx context.Context, reqObj *BaseItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BaseItem
func (r *BaseItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for BaseItem
func (r *BaseItemRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj BaseItem
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for BaseItem
func (r *BaseItemRequest) BatchUpdate(batch *BatchRequest, reqObj *BaseItem) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for BaseItem
func (r *BaseItemRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
