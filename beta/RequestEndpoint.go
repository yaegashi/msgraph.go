// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// EndpointRequestBuilder is request builder for Endpoint
type EndpointRequestBuilder struct{ BaseRequestBuilder }

// Request returns EndpointRequest
func (b *EndpointRequestBuilder) Request() *EndpointRequest {
	return &EndpointRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EndpointRequest is request for Endpoint
type EndpointRequest struct{ BaseRequest }

// Get performs GET request for Endpoint
func (r *EndpointRequest) Get(ctx context.Context) (resObj *Endpoint, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Endpoint
func (r *EndpointRequest) Update(ctx context.Context, reqObj *Endpoint) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Endpoint
func (r *EndpointRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for Endpoint
func (r *EndpointRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj Endpoint
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for Endpoint
func (r *EndpointRequest) BatchUpdate(batch *BatchRequest, reqObj *Endpoint) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for Endpoint
func (r *EndpointRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
