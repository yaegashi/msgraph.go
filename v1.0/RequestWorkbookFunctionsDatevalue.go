// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsDatevalueRequestBuilder struct{ BaseRequestBuilder }

// Datevalue action undocumented
func (b *WorkbookFunctionsRequestBuilder) Datevalue(reqObj *WorkbookFunctionsDatevalueRequestParameter) *WorkbookFunctionsDatevalueRequestBuilder {
	bb := &WorkbookFunctionsDatevalueRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/datevalue"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsDatevalueRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsDatevalueRequestBuilder) Request() *WorkbookFunctionsDatevalueRequest {
	return &WorkbookFunctionsDatevalueRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsDatevalueRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsDatevalueRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
