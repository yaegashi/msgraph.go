// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsAveDevRequestBuilder struct{ BaseRequestBuilder }

// AveDev action undocumented
func (b *WorkbookFunctionsRequestBuilder) AveDev(reqObj *WorkbookFunctionsAveDevRequestParameter) *WorkbookFunctionsAveDevRequestBuilder {
	bb := &WorkbookFunctionsAveDevRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/aveDev"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsAveDevRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsAveDevRequestBuilder) Request() *WorkbookFunctionsAveDevRequest {
	return &WorkbookFunctionsAveDevRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsAveDevRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsAveDevRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
