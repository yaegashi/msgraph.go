// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsFloor_MathRequestBuilder struct{ BaseRequestBuilder }

// Floor_Math action undocumented
func (b *WorkbookFunctionsRequestBuilder) Floor_Math(reqObj *WorkbookFunctionsFloor_MathRequestParameter) *WorkbookFunctionsFloor_MathRequestBuilder {
	bb := &WorkbookFunctionsFloor_MathRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/floor_Math"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsFloor_MathRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsFloor_MathRequestBuilder) Request() *WorkbookFunctionsFloor_MathRequest {
	return &WorkbookFunctionsFloor_MathRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsFloor_MathRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsFloor_MathRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}

//
type WorkbookFunctionsFloor_PreciseRequestBuilder struct{ BaseRequestBuilder }

// Floor_Precise action undocumented
func (b *WorkbookFunctionsRequestBuilder) Floor_Precise(reqObj *WorkbookFunctionsFloor_PreciseRequestParameter) *WorkbookFunctionsFloor_PreciseRequestBuilder {
	bb := &WorkbookFunctionsFloor_PreciseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/floor_Precise"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsFloor_PreciseRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsFloor_PreciseRequestBuilder) Request() *WorkbookFunctionsFloor_PreciseRequest {
	return &WorkbookFunctionsFloor_PreciseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsFloor_PreciseRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsFloor_PreciseRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
