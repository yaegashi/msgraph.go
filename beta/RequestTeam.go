// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// TeamRequestBuilder is request builder for Team
type TeamRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamRequest
func (b *TeamRequestBuilder) Request() *TeamRequest {
	return &TeamRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamRequest is request for Team
type TeamRequest struct{ BaseRequest }

// Get performs GET request for Team
func (r *TeamRequest) Get(ctx context.Context) (resObj *Team, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Team
func (r *TeamRequest) Update(ctx context.Context, reqObj *Team) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Team
func (r *TeamRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for Team
func (r *TeamRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj Team
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for Team
func (r *TeamRequest) BatchUpdate(batch *BatchRequest, reqObj *Team) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for Team
func (r *TeamRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

//
type TeamCloneRequestBuilder struct{ BaseRequestBuilder }

// Clone action undocumented
func (b *TeamRequestBuilder) Clone(reqObj *TeamCloneRequestParameter) *TeamCloneRequestBuilder {
	bb := &TeamCloneRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/clone"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TeamCloneRequest struct{ BaseRequest }

//
func (b *TeamCloneRequestBuilder) Request() *TeamCloneRequest {
	return &TeamCloneRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TeamCloneRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *TeamCloneRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type TeamArchiveRequestBuilder struct{ BaseRequestBuilder }

// Archive action undocumented
func (b *TeamRequestBuilder) Archive(reqObj *TeamArchiveRequestParameter) *TeamArchiveRequestBuilder {
	bb := &TeamArchiveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/archive"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TeamArchiveRequest struct{ BaseRequest }

//
func (b *TeamArchiveRequestBuilder) Request() *TeamArchiveRequest {
	return &TeamArchiveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TeamArchiveRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *TeamArchiveRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type TeamUnarchiveRequestBuilder struct{ BaseRequestBuilder }

// Unarchive action undocumented
func (b *TeamRequestBuilder) Unarchive(reqObj *TeamUnarchiveRequestParameter) *TeamUnarchiveRequestBuilder {
	bb := &TeamUnarchiveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unarchive"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TeamUnarchiveRequest struct{ BaseRequest }

//
func (b *TeamUnarchiveRequestBuilder) Request() *TeamUnarchiveRequest {
	return &TeamUnarchiveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TeamUnarchiveRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *TeamUnarchiveRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}
