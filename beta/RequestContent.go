// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// ContentTypeRequestBuilder is request builder for ContentType
type ContentTypeRequestBuilder struct{ BaseRequestBuilder }

// Request returns ContentTypeRequest
func (b *ContentTypeRequestBuilder) Request() *ContentTypeRequest {
	return &ContentTypeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ContentTypeRequest is request for ContentType
type ContentTypeRequest struct{ BaseRequest }

// Get performs GET request for ContentType
func (r *ContentTypeRequest) Get(ctx context.Context) (resObj *ContentType, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ContentType
func (r *ContentTypeRequest) Update(ctx context.Context, reqObj *ContentType) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ContentType
func (r *ContentTypeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for ContentType
func (r *ContentTypeRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj ContentType
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for ContentType
func (r *ContentTypeRequest) BatchUpdate(batch *BatchRequest, reqObj *ContentType) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for ContentType
func (r *ContentTypeRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
