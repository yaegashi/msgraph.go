// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// ReportRootRequestBuilder is request builder for ReportRoot
type ReportRootRequestBuilder struct{ BaseRequestBuilder }

// Request returns ReportRootRequest
func (b *ReportRootRequestBuilder) Request() *ReportRootRequest {
	return &ReportRootRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ReportRootRequest is request for ReportRoot
type ReportRootRequest struct{ BaseRequest }

// Get performs GET request for ReportRoot
func (r *ReportRootRequest) Get(ctx context.Context) (resObj *ReportRoot, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ReportRoot
func (r *ReportRootRequest) Update(ctx context.Context, reqObj *ReportRoot) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ReportRoot
func (r *ReportRootRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for ReportRoot
func (r *ReportRootRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj ReportRoot
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for ReportRoot
func (r *ReportRootRequest) BatchUpdate(batch *BatchRequest, reqObj *ReportRoot) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for ReportRoot
func (r *ReportRootRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
