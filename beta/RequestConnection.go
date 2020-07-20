// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// ConnectionOperationRequestBuilder is request builder for ConnectionOperation
type ConnectionOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConnectionOperationRequest
func (b *ConnectionOperationRequestBuilder) Request() *ConnectionOperationRequest {
	return &ConnectionOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConnectionOperationRequest is request for ConnectionOperation
type ConnectionOperationRequest struct{ BaseRequest }

// Get performs GET request for ConnectionOperation
func (r *ConnectionOperationRequest) Get(ctx context.Context) (resObj *ConnectionOperation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConnectionOperation
func (r *ConnectionOperationRequest) Update(ctx context.Context, reqObj *ConnectionOperation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConnectionOperation
func (r *ConnectionOperationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for ConnectionOperation
func (r *ConnectionOperationRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj ConnectionOperation
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for ConnectionOperation
func (r *ConnectionOperationRequest) BatchUpdate(batch *BatchRequest, reqObj *ConnectionOperation) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for ConnectionOperation
func (r *ConnectionOperationRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
