// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// OnenoteRequestBuilder is request builder for Onenote
type OnenoteRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnenoteRequest
func (b *OnenoteRequestBuilder) Request() *OnenoteRequest {
	return &OnenoteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnenoteRequest is request for Onenote
type OnenoteRequest struct{ BaseRequest }

// Get performs GET request for Onenote
func (r *OnenoteRequest) Get(ctx context.Context) (resObj *Onenote, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Onenote
func (r *OnenoteRequest) Update(ctx context.Context, reqObj *Onenote) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Onenote
func (r *OnenoteRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for Onenote
func (r *OnenoteRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj Onenote
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for Onenote
func (r *OnenoteRequest) BatchUpdate(batch *BatchRequest, reqObj *Onenote) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for Onenote
func (r *OnenoteRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

// OnenoteOperationRequestBuilder is request builder for OnenoteOperation
type OnenoteOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnenoteOperationRequest
func (b *OnenoteOperationRequestBuilder) Request() *OnenoteOperationRequest {
	return &OnenoteOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnenoteOperationRequest is request for OnenoteOperation
type OnenoteOperationRequest struct{ BaseRequest }

// Get performs GET request for OnenoteOperation
func (r *OnenoteOperationRequest) Get(ctx context.Context) (resObj *OnenoteOperation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnenoteOperation
func (r *OnenoteOperationRequest) Update(ctx context.Context, reqObj *OnenoteOperation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnenoteOperation
func (r *OnenoteOperationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for OnenoteOperation
func (r *OnenoteOperationRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj OnenoteOperation
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for OnenoteOperation
func (r *OnenoteOperationRequest) BatchUpdate(batch *BatchRequest, reqObj *OnenoteOperation) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for OnenoteOperation
func (r *OnenoteOperationRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

// OnenotePageRequestBuilder is request builder for OnenotePage
type OnenotePageRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnenotePageRequest
func (b *OnenotePageRequestBuilder) Request() *OnenotePageRequest {
	return &OnenotePageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnenotePageRequest is request for OnenotePage
type OnenotePageRequest struct{ BaseRequest }

// Get performs GET request for OnenotePage
func (r *OnenotePageRequest) Get(ctx context.Context) (resObj *OnenotePage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnenotePage
func (r *OnenotePageRequest) Update(ctx context.Context, reqObj *OnenotePage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnenotePage
func (r *OnenotePageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for OnenotePage
func (r *OnenotePageRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj OnenotePage
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for OnenotePage
func (r *OnenotePageRequest) BatchUpdate(batch *BatchRequest, reqObj *OnenotePage) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for OnenotePage
func (r *OnenotePageRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

// OnenoteResourceRequestBuilder is request builder for OnenoteResource
type OnenoteResourceRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnenoteResourceRequest
func (b *OnenoteResourceRequestBuilder) Request() *OnenoteResourceRequest {
	return &OnenoteResourceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnenoteResourceRequest is request for OnenoteResource
type OnenoteResourceRequest struct{ BaseRequest }

// Get performs GET request for OnenoteResource
func (r *OnenoteResourceRequest) Get(ctx context.Context) (resObj *OnenoteResource, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnenoteResource
func (r *OnenoteResourceRequest) Update(ctx context.Context, reqObj *OnenoteResource) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnenoteResource
func (r *OnenoteResourceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for OnenoteResource
func (r *OnenoteResourceRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj OnenoteResource
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for OnenoteResource
func (r *OnenoteResourceRequest) BatchUpdate(batch *BatchRequest, reqObj *OnenoteResource) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for OnenoteResource
func (r *OnenoteResourceRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

// OnenoteSectionRequestBuilder is request builder for OnenoteSection
type OnenoteSectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnenoteSectionRequest
func (b *OnenoteSectionRequestBuilder) Request() *OnenoteSectionRequest {
	return &OnenoteSectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnenoteSectionRequest is request for OnenoteSection
type OnenoteSectionRequest struct{ BaseRequest }

// Get performs GET request for OnenoteSection
func (r *OnenoteSectionRequest) Get(ctx context.Context) (resObj *OnenoteSection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnenoteSection
func (r *OnenoteSectionRequest) Update(ctx context.Context, reqObj *OnenoteSection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnenoteSection
func (r *OnenoteSectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for OnenoteSection
func (r *OnenoteSectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj OnenoteSection
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for OnenoteSection
func (r *OnenoteSectionRequest) BatchUpdate(batch *BatchRequest, reqObj *OnenoteSection) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for OnenoteSection
func (r *OnenoteSectionRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

//
type OnenotePageOnenotePatchContentRequestBuilder struct{ BaseRequestBuilder }

// OnenotePatchContent action undocumented
func (b *OnenotePageRequestBuilder) OnenotePatchContent(reqObj *OnenotePageOnenotePatchContentRequestParameter) *OnenotePageOnenotePatchContentRequestBuilder {
	bb := &OnenotePageOnenotePatchContentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/onenotePatchContent"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OnenotePageOnenotePatchContentRequest struct{ BaseRequest }

//
func (b *OnenotePageOnenotePatchContentRequestBuilder) Request() *OnenotePageOnenotePatchContentRequest {
	return &OnenotePageOnenotePatchContentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OnenotePageOnenotePatchContentRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *OnenotePageOnenotePatchContentRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type OnenotePageCopyToSectionRequestBuilder struct{ BaseRequestBuilder }

// CopyToSection action undocumented
func (b *OnenotePageRequestBuilder) CopyToSection(reqObj *OnenotePageCopyToSectionRequestParameter) *OnenotePageCopyToSectionRequestBuilder {
	bb := &OnenotePageCopyToSectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/copyToSection"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OnenotePageCopyToSectionRequest struct{ BaseRequest }

//
func (b *OnenotePageCopyToSectionRequestBuilder) Request() *OnenotePageCopyToSectionRequest {
	return &OnenotePageCopyToSectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OnenotePageCopyToSectionRequest) Post(ctx context.Context) (resObj *OnenoteOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *OnenotePageCopyToSectionRequest) BatchPost(batch *BatchRequest) error {
	var resObj *OnenoteOperation
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}

//
type OnenoteSectionCopyToNotebookRequestBuilder struct{ BaseRequestBuilder }

// CopyToNotebook action undocumented
func (b *OnenoteSectionRequestBuilder) CopyToNotebook(reqObj *OnenoteSectionCopyToNotebookRequestParameter) *OnenoteSectionCopyToNotebookRequestBuilder {
	bb := &OnenoteSectionCopyToNotebookRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/copyToNotebook"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OnenoteSectionCopyToNotebookRequest struct{ BaseRequest }

//
func (b *OnenoteSectionCopyToNotebookRequestBuilder) Request() *OnenoteSectionCopyToNotebookRequest {
	return &OnenoteSectionCopyToNotebookRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OnenoteSectionCopyToNotebookRequest) Post(ctx context.Context) (resObj *OnenoteOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *OnenoteSectionCopyToNotebookRequest) BatchPost(batch *BatchRequest) error {
	var resObj *OnenoteOperation
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}

//
type OnenoteSectionCopyToSectionGroupRequestBuilder struct{ BaseRequestBuilder }

// CopyToSectionGroup action undocumented
func (b *OnenoteSectionRequestBuilder) CopyToSectionGroup(reqObj *OnenoteSectionCopyToSectionGroupRequestParameter) *OnenoteSectionCopyToSectionGroupRequestBuilder {
	bb := &OnenoteSectionCopyToSectionGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/copyToSectionGroup"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OnenoteSectionCopyToSectionGroupRequest struct{ BaseRequest }

//
func (b *OnenoteSectionCopyToSectionGroupRequestBuilder) Request() *OnenoteSectionCopyToSectionGroupRequest {
	return &OnenoteSectionCopyToSectionGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OnenoteSectionCopyToSectionGroupRequest) Post(ctx context.Context) (resObj *OnenoteOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *OnenoteSectionCopyToSectionGroupRequest) BatchPost(batch *BatchRequest) error {
	var resObj *OnenoteOperation
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
