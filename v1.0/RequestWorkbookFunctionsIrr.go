// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsIrrRequestBuilder struct{ BaseRequestBuilder }

// Irr action undocumented
func (b *WorkbookFunctionsRequestBuilder) Irr(reqObj *WorkbookFunctionsIrrRequestParameter) *WorkbookFunctionsIrrRequestBuilder {
	bb := &WorkbookFunctionsIrrRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/irr"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsIrrRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsIrrRequestBuilder) Request() *WorkbookFunctionsIrrRequest {
	return &WorkbookFunctionsIrrRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsIrrRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsIrrRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
