// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsAsinRequestBuilder struct{ BaseRequestBuilder }

// Asin action undocumented
func (b *WorkbookFunctionsRequestBuilder) Asin(reqObj *WorkbookFunctionsAsinRequestParameter) *WorkbookFunctionsAsinRequestBuilder {
	bb := &WorkbookFunctionsAsinRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/asin"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsAsinRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsAsinRequestBuilder) Request() *WorkbookFunctionsAsinRequest {
	return &WorkbookFunctionsAsinRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsAsinRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsAsinRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
