// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsSlnRequestBuilder struct{ BaseRequestBuilder }

// Sln action undocumented
func (b *WorkbookFunctionsRequestBuilder) Sln(reqObj *WorkbookFunctionsSlnRequestParameter) *WorkbookFunctionsSlnRequestBuilder {
	bb := &WorkbookFunctionsSlnRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/sln"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsSlnRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsSlnRequestBuilder) Request() *WorkbookFunctionsSlnRequest {
	return &WorkbookFunctionsSlnRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsSlnRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsSlnRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
