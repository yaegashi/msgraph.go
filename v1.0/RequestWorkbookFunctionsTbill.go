// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsTbillEqRequestBuilder struct{ BaseRequestBuilder }

// TbillEq action undocumented
func (b *WorkbookFunctionsRequestBuilder) TbillEq(reqObj *WorkbookFunctionsTbillEqRequestParameter) *WorkbookFunctionsTbillEqRequestBuilder {
	bb := &WorkbookFunctionsTbillEqRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/tbillEq"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsTbillEqRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsTbillEqRequestBuilder) Request() *WorkbookFunctionsTbillEqRequest {
	return &WorkbookFunctionsTbillEqRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsTbillEqRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsTbillEqRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}

//
type WorkbookFunctionsTbillPriceRequestBuilder struct{ BaseRequestBuilder }

// TbillPrice action undocumented
func (b *WorkbookFunctionsRequestBuilder) TbillPrice(reqObj *WorkbookFunctionsTbillPriceRequestParameter) *WorkbookFunctionsTbillPriceRequestBuilder {
	bb := &WorkbookFunctionsTbillPriceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/tbillPrice"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsTbillPriceRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsTbillPriceRequestBuilder) Request() *WorkbookFunctionsTbillPriceRequest {
	return &WorkbookFunctionsTbillPriceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsTbillPriceRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsTbillPriceRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}

//
type WorkbookFunctionsTbillYieldRequestBuilder struct{ BaseRequestBuilder }

// TbillYield action undocumented
func (b *WorkbookFunctionsRequestBuilder) TbillYield(reqObj *WorkbookFunctionsTbillYieldRequestParameter) *WorkbookFunctionsTbillYieldRequestBuilder {
	bb := &WorkbookFunctionsTbillYieldRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/tbillYield"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsTbillYieldRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsTbillYieldRequestBuilder) Request() *WorkbookFunctionsTbillYieldRequest {
	return &WorkbookFunctionsTbillYieldRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsTbillYieldRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsTbillYieldRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
