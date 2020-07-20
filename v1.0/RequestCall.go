// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// CallRequestBuilder is request builder for Call
type CallRequestBuilder struct{ BaseRequestBuilder }

// Request returns CallRequest
func (b *CallRequestBuilder) Request() *CallRequest {
	return &CallRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CallRequest is request for Call
type CallRequest struct{ BaseRequest }

// Get performs GET request for Call
func (r *CallRequest) Get(ctx context.Context) (resObj *Call, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Call
func (r *CallRequest) Update(ctx context.Context, reqObj *Call) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Call
func (r *CallRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for Call
func (r *CallRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj Call
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for Call
func (r *CallRequest) BatchUpdate(batch *BatchRequest, reqObj *Call) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for Call
func (r *CallRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

//
type CallAnswerRequestBuilder struct{ BaseRequestBuilder }

// Answer action undocumented
func (b *CallRequestBuilder) Answer(reqObj *CallAnswerRequestParameter) *CallAnswerRequestBuilder {
	bb := &CallAnswerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/answer"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallAnswerRequest struct{ BaseRequest }

//
func (b *CallAnswerRequestBuilder) Request() *CallAnswerRequest {
	return &CallAnswerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *CallAnswerRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *CallAnswerRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type CallChangeScreenSharingRoleRequestBuilder struct{ BaseRequestBuilder }

// ChangeScreenSharingRole action undocumented
func (b *CallRequestBuilder) ChangeScreenSharingRole(reqObj *CallChangeScreenSharingRoleRequestParameter) *CallChangeScreenSharingRoleRequestBuilder {
	bb := &CallChangeScreenSharingRoleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/changeScreenSharingRole"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallChangeScreenSharingRoleRequest struct{ BaseRequest }

//
func (b *CallChangeScreenSharingRoleRequestBuilder) Request() *CallChangeScreenSharingRoleRequest {
	return &CallChangeScreenSharingRoleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *CallChangeScreenSharingRoleRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *CallChangeScreenSharingRoleRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type CallKeepAliveRequestBuilder struct{ BaseRequestBuilder }

// KeepAlive action undocumented
func (b *CallRequestBuilder) KeepAlive(reqObj *CallKeepAliveRequestParameter) *CallKeepAliveRequestBuilder {
	bb := &CallKeepAliveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/keepAlive"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallKeepAliveRequest struct{ BaseRequest }

//
func (b *CallKeepAliveRequestBuilder) Request() *CallKeepAliveRequest {
	return &CallKeepAliveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *CallKeepAliveRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *CallKeepAliveRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type CallMuteRequestBuilder struct{ BaseRequestBuilder }

// Mute action undocumented
func (b *CallRequestBuilder) Mute(reqObj *CallMuteRequestParameter) *CallMuteRequestBuilder {
	bb := &CallMuteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/mute"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallMuteRequest struct{ BaseRequest }

//
func (b *CallMuteRequestBuilder) Request() *CallMuteRequest {
	return &CallMuteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *CallMuteRequest) Post(ctx context.Context) (resObj *MuteParticipantOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *CallMuteRequest) BatchPost(batch *BatchRequest) error {
	var resObj *MuteParticipantOperation
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}

//
type CallPlayPromptRequestBuilder struct{ BaseRequestBuilder }

// PlayPrompt action undocumented
func (b *CallRequestBuilder) PlayPrompt(reqObj *CallPlayPromptRequestParameter) *CallPlayPromptRequestBuilder {
	bb := &CallPlayPromptRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/playPrompt"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallPlayPromptRequest struct{ BaseRequest }

//
func (b *CallPlayPromptRequestBuilder) Request() *CallPlayPromptRequest {
	return &CallPlayPromptRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *CallPlayPromptRequest) Post(ctx context.Context) (resObj *PlayPromptOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *CallPlayPromptRequest) BatchPost(batch *BatchRequest) error {
	var resObj *PlayPromptOperation
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}

//
type CallRecordResponseRequestBuilder struct{ BaseRequestBuilder }

// RecordResponse action undocumented
func (b *CallRequestBuilder) RecordResponse(reqObj *CallRecordResponseRequestParameter) *CallRecordResponseRequestBuilder {
	bb := &CallRecordResponseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/recordResponse"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallRecordResponseRequest struct{ BaseRequest }

//
func (b *CallRecordResponseRequestBuilder) Request() *CallRecordResponseRequest {
	return &CallRecordResponseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *CallRecordResponseRequest) Post(ctx context.Context) (resObj *RecordOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *CallRecordResponseRequest) BatchPost(batch *BatchRequest) error {
	var resObj *RecordOperation
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}

//
type CallRedirectRequestBuilder struct{ BaseRequestBuilder }

// Redirect action undocumented
func (b *CallRequestBuilder) Redirect(reqObj *CallRedirectRequestParameter) *CallRedirectRequestBuilder {
	bb := &CallRedirectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/redirect"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallRedirectRequest struct{ BaseRequest }

//
func (b *CallRedirectRequestBuilder) Request() *CallRedirectRequest {
	return &CallRedirectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *CallRedirectRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *CallRedirectRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type CallRejectRequestBuilder struct{ BaseRequestBuilder }

// Reject action undocumented
func (b *CallRequestBuilder) Reject(reqObj *CallRejectRequestParameter) *CallRejectRequestBuilder {
	bb := &CallRejectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/reject"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallRejectRequest struct{ BaseRequest }

//
func (b *CallRejectRequestBuilder) Request() *CallRejectRequest {
	return &CallRejectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *CallRejectRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *CallRejectRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type CallSubscribeToToneRequestBuilder struct{ BaseRequestBuilder }

// SubscribeToTone action undocumented
func (b *CallRequestBuilder) SubscribeToTone(reqObj *CallSubscribeToToneRequestParameter) *CallSubscribeToToneRequestBuilder {
	bb := &CallSubscribeToToneRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/subscribeToTone"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallSubscribeToToneRequest struct{ BaseRequest }

//
func (b *CallSubscribeToToneRequestBuilder) Request() *CallSubscribeToToneRequest {
	return &CallSubscribeToToneRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *CallSubscribeToToneRequest) Post(ctx context.Context) (resObj *SubscribeToToneOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *CallSubscribeToToneRequest) BatchPost(batch *BatchRequest) error {
	var resObj *SubscribeToToneOperation
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}

//
type CallTransferRequestBuilder struct{ BaseRequestBuilder }

// Transfer action undocumented
func (b *CallRequestBuilder) Transfer(reqObj *CallTransferRequestParameter) *CallTransferRequestBuilder {
	bb := &CallTransferRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/transfer"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallTransferRequest struct{ BaseRequest }

//
func (b *CallTransferRequestBuilder) Request() *CallTransferRequest {
	return &CallTransferRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *CallTransferRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *CallTransferRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type CallUnmuteRequestBuilder struct{ BaseRequestBuilder }

// Unmute action undocumented
func (b *CallRequestBuilder) Unmute(reqObj *CallUnmuteRequestParameter) *CallUnmuteRequestBuilder {
	bb := &CallUnmuteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unmute"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CallUnmuteRequest struct{ BaseRequest }

//
func (b *CallUnmuteRequestBuilder) Request() *CallUnmuteRequest {
	return &CallUnmuteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *CallUnmuteRequest) Post(ctx context.Context) (resObj *UnmuteParticipantOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *CallUnmuteRequest) BatchPost(batch *BatchRequest) error {
	var resObj *UnmuteParticipantOperation
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
