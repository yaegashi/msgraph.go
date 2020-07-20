// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// RemoteActionAuditRequestBuilder is request builder for RemoteActionAudit
type RemoteActionAuditRequestBuilder struct{ BaseRequestBuilder }

// Request returns RemoteActionAuditRequest
func (b *RemoteActionAuditRequestBuilder) Request() *RemoteActionAuditRequest {
	return &RemoteActionAuditRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RemoteActionAuditRequest is request for RemoteActionAudit
type RemoteActionAuditRequest struct{ BaseRequest }

// Get performs GET request for RemoteActionAudit
func (r *RemoteActionAuditRequest) Get(ctx context.Context) (resObj *RemoteActionAudit, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RemoteActionAudit
func (r *RemoteActionAuditRequest) Update(ctx context.Context, reqObj *RemoteActionAudit) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RemoteActionAudit
func (r *RemoteActionAuditRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for RemoteActionAudit
func (r *RemoteActionAuditRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj RemoteActionAudit
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for RemoteActionAudit
func (r *RemoteActionAuditRequest) BatchUpdate(batch *BatchRequest, reqObj *RemoteActionAudit) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for RemoteActionAudit
func (r *RemoteActionAuditRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

// RemoteAssistancePartnerRequestBuilder is request builder for RemoteAssistancePartner
type RemoteAssistancePartnerRequestBuilder struct{ BaseRequestBuilder }

// Request returns RemoteAssistancePartnerRequest
func (b *RemoteAssistancePartnerRequestBuilder) Request() *RemoteAssistancePartnerRequest {
	return &RemoteAssistancePartnerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RemoteAssistancePartnerRequest is request for RemoteAssistancePartner
type RemoteAssistancePartnerRequest struct{ BaseRequest }

// Get performs GET request for RemoteAssistancePartner
func (r *RemoteAssistancePartnerRequest) Get(ctx context.Context) (resObj *RemoteAssistancePartner, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RemoteAssistancePartner
func (r *RemoteAssistancePartnerRequest) Update(ctx context.Context, reqObj *RemoteAssistancePartner) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RemoteAssistancePartner
func (r *RemoteAssistancePartnerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for RemoteAssistancePartner
func (r *RemoteAssistancePartnerRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj RemoteAssistancePartner
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for RemoteAssistancePartner
func (r *RemoteAssistancePartnerRequest) BatchUpdate(batch *BatchRequest, reqObj *RemoteAssistancePartner) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for RemoteAssistancePartner
func (r *RemoteAssistancePartnerRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

//
type RemoteAssistancePartnerBeginOnboardingRequestBuilder struct{ BaseRequestBuilder }

// BeginOnboarding action undocumented
func (b *RemoteAssistancePartnerRequestBuilder) BeginOnboarding(reqObj *RemoteAssistancePartnerBeginOnboardingRequestParameter) *RemoteAssistancePartnerBeginOnboardingRequestBuilder {
	bb := &RemoteAssistancePartnerBeginOnboardingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/beginOnboarding"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type RemoteAssistancePartnerBeginOnboardingRequest struct{ BaseRequest }

//
func (b *RemoteAssistancePartnerBeginOnboardingRequestBuilder) Request() *RemoteAssistancePartnerBeginOnboardingRequest {
	return &RemoteAssistancePartnerBeginOnboardingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *RemoteAssistancePartnerBeginOnboardingRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *RemoteAssistancePartnerBeginOnboardingRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type RemoteAssistancePartnerDisconnectRequestBuilder struct{ BaseRequestBuilder }

// Disconnect action undocumented
func (b *RemoteAssistancePartnerRequestBuilder) Disconnect(reqObj *RemoteAssistancePartnerDisconnectRequestParameter) *RemoteAssistancePartnerDisconnectRequestBuilder {
	bb := &RemoteAssistancePartnerDisconnectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/disconnect"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type RemoteAssistancePartnerDisconnectRequest struct{ BaseRequest }

//
func (b *RemoteAssistancePartnerDisconnectRequestBuilder) Request() *RemoteAssistancePartnerDisconnectRequest {
	return &RemoteAssistancePartnerDisconnectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *RemoteAssistancePartnerDisconnectRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *RemoteAssistancePartnerDisconnectRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}
