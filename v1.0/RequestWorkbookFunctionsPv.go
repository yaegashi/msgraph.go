// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsPvRequestBuilder struct{ BaseRequestBuilder }

// Pv action undocumented
func (b *WorkbookFunctionsRequestBuilder) Pv(reqObj *WorkbookFunctionsPvRequestParameter) *WorkbookFunctionsPvRequestBuilder {
	bb := &WorkbookFunctionsPvRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/pv"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsPvRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsPvRequestBuilder) Request() *WorkbookFunctionsPvRequest {
	return &WorkbookFunctionsPvRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsPvRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsPvRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
