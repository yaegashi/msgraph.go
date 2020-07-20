// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsTimeRequestBuilder struct{ BaseRequestBuilder }

// Time action undocumented
func (b *WorkbookFunctionsRequestBuilder) Time(reqObj *WorkbookFunctionsTimeRequestParameter) *WorkbookFunctionsTimeRequestBuilder {
	bb := &WorkbookFunctionsTimeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/time"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsTimeRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsTimeRequestBuilder) Request() *WorkbookFunctionsTimeRequest {
	return &WorkbookFunctionsTimeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsTimeRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsTimeRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
