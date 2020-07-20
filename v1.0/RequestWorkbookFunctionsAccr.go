// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsAccrIntRequestBuilder struct{ BaseRequestBuilder }

// AccrInt action undocumented
func (b *WorkbookFunctionsRequestBuilder) AccrInt(reqObj *WorkbookFunctionsAccrIntRequestParameter) *WorkbookFunctionsAccrIntRequestBuilder {
	bb := &WorkbookFunctionsAccrIntRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/accrInt"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsAccrIntRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsAccrIntRequestBuilder) Request() *WorkbookFunctionsAccrIntRequest {
	return &WorkbookFunctionsAccrIntRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsAccrIntRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsAccrIntRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}

//
type WorkbookFunctionsAccrIntMRequestBuilder struct{ BaseRequestBuilder }

// AccrIntM action undocumented
func (b *WorkbookFunctionsRequestBuilder) AccrIntM(reqObj *WorkbookFunctionsAccrIntMRequestParameter) *WorkbookFunctionsAccrIntMRequestBuilder {
	bb := &WorkbookFunctionsAccrIntMRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/accrIntM"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsAccrIntMRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsAccrIntMRequestBuilder) Request() *WorkbookFunctionsAccrIntMRequest {
	return &WorkbookFunctionsAccrIntMRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsAccrIntMRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsAccrIntMRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
