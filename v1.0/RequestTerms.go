// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// TermsAndConditionsRequestBuilder is request builder for TermsAndConditions
type TermsAndConditionsRequestBuilder struct{ BaseRequestBuilder }

// Request returns TermsAndConditionsRequest
func (b *TermsAndConditionsRequestBuilder) Request() *TermsAndConditionsRequest {
	return &TermsAndConditionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TermsAndConditionsRequest is request for TermsAndConditions
type TermsAndConditionsRequest struct{ BaseRequest }

// Get performs GET request for TermsAndConditions
func (r *TermsAndConditionsRequest) Get(ctx context.Context) (resObj *TermsAndConditions, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TermsAndConditions
func (r *TermsAndConditionsRequest) Update(ctx context.Context, reqObj *TermsAndConditions) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TermsAndConditions
func (r *TermsAndConditionsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for TermsAndConditions
func (r *TermsAndConditionsRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj TermsAndConditions
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for TermsAndConditions
func (r *TermsAndConditionsRequest) BatchUpdate(batch *BatchRequest, reqObj *TermsAndConditions) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for TermsAndConditions
func (r *TermsAndConditionsRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

// TermsAndConditionsAcceptanceStatusRequestBuilder is request builder for TermsAndConditionsAcceptanceStatus
type TermsAndConditionsAcceptanceStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns TermsAndConditionsAcceptanceStatusRequest
func (b *TermsAndConditionsAcceptanceStatusRequestBuilder) Request() *TermsAndConditionsAcceptanceStatusRequest {
	return &TermsAndConditionsAcceptanceStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TermsAndConditionsAcceptanceStatusRequest is request for TermsAndConditionsAcceptanceStatus
type TermsAndConditionsAcceptanceStatusRequest struct{ BaseRequest }

// Get performs GET request for TermsAndConditionsAcceptanceStatus
func (r *TermsAndConditionsAcceptanceStatusRequest) Get(ctx context.Context) (resObj *TermsAndConditionsAcceptanceStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TermsAndConditionsAcceptanceStatus
func (r *TermsAndConditionsAcceptanceStatusRequest) Update(ctx context.Context, reqObj *TermsAndConditionsAcceptanceStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TermsAndConditionsAcceptanceStatus
func (r *TermsAndConditionsAcceptanceStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for TermsAndConditionsAcceptanceStatus
func (r *TermsAndConditionsAcceptanceStatusRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj TermsAndConditionsAcceptanceStatus
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for TermsAndConditionsAcceptanceStatus
func (r *TermsAndConditionsAcceptanceStatusRequest) BatchUpdate(batch *BatchRequest, reqObj *TermsAndConditionsAcceptanceStatus) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for TermsAndConditionsAcceptanceStatus
func (r *TermsAndConditionsAcceptanceStatusRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

// TermsAndConditionsAssignmentRequestBuilder is request builder for TermsAndConditionsAssignment
type TermsAndConditionsAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns TermsAndConditionsAssignmentRequest
func (b *TermsAndConditionsAssignmentRequestBuilder) Request() *TermsAndConditionsAssignmentRequest {
	return &TermsAndConditionsAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TermsAndConditionsAssignmentRequest is request for TermsAndConditionsAssignment
type TermsAndConditionsAssignmentRequest struct{ BaseRequest }

// Get performs GET request for TermsAndConditionsAssignment
func (r *TermsAndConditionsAssignmentRequest) Get(ctx context.Context) (resObj *TermsAndConditionsAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TermsAndConditionsAssignment
func (r *TermsAndConditionsAssignmentRequest) Update(ctx context.Context, reqObj *TermsAndConditionsAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TermsAndConditionsAssignment
func (r *TermsAndConditionsAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for TermsAndConditionsAssignment
func (r *TermsAndConditionsAssignmentRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj TermsAndConditionsAssignment
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for TermsAndConditionsAssignment
func (r *TermsAndConditionsAssignmentRequest) BatchUpdate(batch *BatchRequest, reqObj *TermsAndConditionsAssignment) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for TermsAndConditionsAssignment
func (r *TermsAndConditionsAssignmentRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
