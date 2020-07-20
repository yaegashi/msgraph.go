// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsNaRequestBuilder struct{ BaseRequestBuilder }

// Na action undocumented
func (b *WorkbookFunctionsRequestBuilder) Na(reqObj *WorkbookFunctionsNaRequestParameter) *WorkbookFunctionsNaRequestBuilder {
	bb := &WorkbookFunctionsNaRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/na"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsNaRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsNaRequestBuilder) Request() *WorkbookFunctionsNaRequest {
	return &WorkbookFunctionsNaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsNaRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsNaRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
