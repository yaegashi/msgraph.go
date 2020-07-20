// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsMidRequestBuilder struct{ BaseRequestBuilder }

// Mid action undocumented
func (b *WorkbookFunctionsRequestBuilder) Mid(reqObj *WorkbookFunctionsMidRequestParameter) *WorkbookFunctionsMidRequestBuilder {
	bb := &WorkbookFunctionsMidRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/mid"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsMidRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsMidRequestBuilder) Request() *WorkbookFunctionsMidRequest {
	return &WorkbookFunctionsMidRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsMidRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsMidRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
