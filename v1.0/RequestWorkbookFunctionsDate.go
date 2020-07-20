// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsDateRequestBuilder struct{ BaseRequestBuilder }

// Date action undocumented
func (b *WorkbookFunctionsRequestBuilder) Date(reqObj *WorkbookFunctionsDateRequestParameter) *WorkbookFunctionsDateRequestBuilder {
	bb := &WorkbookFunctionsDateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/date"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsDateRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsDateRequestBuilder) Request() *WorkbookFunctionsDateRequest {
	return &WorkbookFunctionsDateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsDateRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsDateRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
