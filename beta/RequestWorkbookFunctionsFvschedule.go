// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsFvscheduleRequestBuilder struct{ BaseRequestBuilder }

// Fvschedule action undocumented
func (b *WorkbookFunctionsRequestBuilder) Fvschedule(reqObj *WorkbookFunctionsFvscheduleRequestParameter) *WorkbookFunctionsFvscheduleRequestBuilder {
	bb := &WorkbookFunctionsFvscheduleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/fvschedule"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsFvscheduleRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsFvscheduleRequestBuilder) Request() *WorkbookFunctionsFvscheduleRequest {
	return &WorkbookFunctionsFvscheduleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsFvscheduleRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsFvscheduleRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
