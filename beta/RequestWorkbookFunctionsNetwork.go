// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

//
type WorkbookFunctionsNetworkDaysRequestBuilder struct{ BaseRequestBuilder }

// NetworkDays action undocumented
func (b *WorkbookFunctionsRequestBuilder) NetworkDays(reqObj *WorkbookFunctionsNetworkDaysRequestParameter) *WorkbookFunctionsNetworkDaysRequestBuilder {
	bb := &WorkbookFunctionsNetworkDaysRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/networkDays"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsNetworkDaysRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsNetworkDaysRequestBuilder) Request() *WorkbookFunctionsNetworkDaysRequest {
	return &WorkbookFunctionsNetworkDaysRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsNetworkDaysRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsNetworkDaysRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}

//
type WorkbookFunctionsNetworkDays_IntlRequestBuilder struct{ BaseRequestBuilder }

// NetworkDays_Intl action undocumented
func (b *WorkbookFunctionsRequestBuilder) NetworkDays_Intl(reqObj *WorkbookFunctionsNetworkDays_IntlRequestParameter) *WorkbookFunctionsNetworkDays_IntlRequestBuilder {
	bb := &WorkbookFunctionsNetworkDays_IntlRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/networkDays_Intl"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsNetworkDays_IntlRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsNetworkDays_IntlRequestBuilder) Request() *WorkbookFunctionsNetworkDays_IntlRequest {
	return &WorkbookFunctionsNetworkDays_IntlRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsNetworkDays_IntlRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *WorkbookFunctionsNetworkDays_IntlRequest) BatchPost(batch *BatchRequest) error {
	var resObj *WorkbookFunctionResult
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
