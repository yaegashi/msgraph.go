// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsXirrRequestBuilder struct{ BaseRequestBuilder }

// Xirr action undocumented
func (b *WorkbookFunctionsRequestBuilder) Xirr(reqObj *WorkbookFunctionsXirrRequestParameter) *WorkbookFunctionsXirrRequestBuilder {
	bb := &WorkbookFunctionsXirrRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/xirr"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsXirrRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsXirrRequestBuilder) Request() *WorkbookFunctionsXirrRequest {
	return &WorkbookFunctionsXirrRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsXirrRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsXirrRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
