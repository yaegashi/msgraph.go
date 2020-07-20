// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// LookupResultRowRequestBuilder is request builder for LookupResultRow
type LookupResultRowRequestBuilder struct{ BaseRequestBuilder }

// Request returns LookupResultRowRequest
func (b *LookupResultRowRequestBuilder) Request() *LookupResultRowRequest {
	return &LookupResultRowRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// LookupResultRowRequest is request for LookupResultRow
type LookupResultRowRequest struct{ BaseRequest }

// Get performs GET request for LookupResultRow
func (r *LookupResultRowRequest) Get(ctx context.Context) (resObj *LookupResultRow, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for LookupResultRow
func (r *LookupResultRowRequest) Update(ctx context.Context, reqObj *LookupResultRow) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for LookupResultRow
func (r *LookupResultRowRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for LookupResultRow
func (r *LookupResultRowRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj LookupResultRow
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for LookupResultRow
func (r *LookupResultRowRequest) BatchUpdate(batch *BatchRequest, reqObj *LookupResultRow) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for LookupResultRow
func (r *LookupResultRowRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
