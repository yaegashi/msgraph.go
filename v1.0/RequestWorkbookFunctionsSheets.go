// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsSheetsRequestBuilder struct{ BaseRequestBuilder }

// Sheets action undocumented
func (b *WorkbookFunctionsRequestBuilder) Sheets(reqObj *WorkbookFunctionsSheetsRequestParameter) *WorkbookFunctionsSheetsRequestBuilder {
	bb := &WorkbookFunctionsSheetsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/sheets"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsSheetsRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsSheetsRequestBuilder) Request() *WorkbookFunctionsSheetsRequest {
	return &WorkbookFunctionsSheetsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsSheetsRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsSheetsRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
