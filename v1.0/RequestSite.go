// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// SiteRequestBuilder is request builder for Site
type SiteRequestBuilder struct{ BaseRequestBuilder }

// Request returns SiteRequest
func (b *SiteRequestBuilder) Request() *SiteRequest {
	return &SiteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SiteRequest is request for Site
type SiteRequest struct{ BaseRequest }

// Get performs GET request for Site
func (r *SiteRequest) Get(ctx context.Context) (resObj *Site, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Site
func (r *SiteRequest) Update(ctx context.Context, reqObj *Site) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Site
func (r *SiteRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for Site
func (r *SiteRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj Site
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for Site
func (r *SiteRequest) BatchUpdate(batch *BatchRequest, reqObj *Site) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for Site
func (r *SiteRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
