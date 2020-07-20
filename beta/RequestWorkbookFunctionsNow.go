// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsNowRequestBuilder struct{ BaseRequestBuilder }

// Now action undocumented
func (b *WorkbookFunctionsRequestBuilder) Now(reqObj *WorkbookFunctionsNowRequestParameter) *WorkbookFunctionsNowRequestBuilder {
	bb := &WorkbookFunctionsNowRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/now"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsNowRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsNowRequestBuilder) Request() *WorkbookFunctionsNowRequest {
	return &WorkbookFunctionsNowRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsNowRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsNowRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
