// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// SynchronizationRequestBuilder is request builder for Synchronization
type SynchronizationRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationRequest
func (b *SynchronizationRequestBuilder) Request() *SynchronizationRequest {
	return &SynchronizationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationRequest is request for Synchronization
type SynchronizationRequest struct{ BaseRequest }

// Get performs GET request for Synchronization
func (r *SynchronizationRequest) Get(ctx context.Context) (resObj *Synchronization, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Synchronization
func (r *SynchronizationRequest) Update(ctx context.Context, reqObj *Synchronization) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Synchronization
func (r *SynchronizationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for Synchronization
func (r *SynchronizationRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj Synchronization
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for Synchronization
func (r *SynchronizationRequest) BatchUpdate(batch *BatchRequest, reqObj *Synchronization) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for Synchronization
func (r *SynchronizationRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

// SynchronizationJobRequestBuilder is request builder for SynchronizationJob
type SynchronizationJobRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationJobRequest
func (b *SynchronizationJobRequestBuilder) Request() *SynchronizationJobRequest {
	return &SynchronizationJobRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationJobRequest is request for SynchronizationJob
type SynchronizationJobRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationJob
func (r *SynchronizationJobRequest) Get(ctx context.Context) (resObj *SynchronizationJob, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationJob
func (r *SynchronizationJobRequest) Update(ctx context.Context, reqObj *SynchronizationJob) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationJob
func (r *SynchronizationJobRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for SynchronizationJob
func (r *SynchronizationJobRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj SynchronizationJob
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for SynchronizationJob
func (r *SynchronizationJobRequest) BatchUpdate(batch *BatchRequest, reqObj *SynchronizationJob) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for SynchronizationJob
func (r *SynchronizationJobRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

// SynchronizationSchemaRequestBuilder is request builder for SynchronizationSchema
type SynchronizationSchemaRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationSchemaRequest
func (b *SynchronizationSchemaRequestBuilder) Request() *SynchronizationSchemaRequest {
	return &SynchronizationSchemaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationSchemaRequest is request for SynchronizationSchema
type SynchronizationSchemaRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationSchema
func (r *SynchronizationSchemaRequest) Get(ctx context.Context) (resObj *SynchronizationSchema, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationSchema
func (r *SynchronizationSchemaRequest) Update(ctx context.Context, reqObj *SynchronizationSchema) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationSchema
func (r *SynchronizationSchemaRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for SynchronizationSchema
func (r *SynchronizationSchemaRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj SynchronizationSchema
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for SynchronizationSchema
func (r *SynchronizationSchemaRequest) BatchUpdate(batch *BatchRequest, reqObj *SynchronizationSchema) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for SynchronizationSchema
func (r *SynchronizationSchemaRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

// SynchronizationTemplateRequestBuilder is request builder for SynchronizationTemplate
type SynchronizationTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationTemplateRequest
func (b *SynchronizationTemplateRequestBuilder) Request() *SynchronizationTemplateRequest {
	return &SynchronizationTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationTemplateRequest is request for SynchronizationTemplate
type SynchronizationTemplateRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationTemplate
func (r *SynchronizationTemplateRequest) Get(ctx context.Context) (resObj *SynchronizationTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationTemplate
func (r *SynchronizationTemplateRequest) Update(ctx context.Context, reqObj *SynchronizationTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationTemplate
func (r *SynchronizationTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for SynchronizationTemplate
func (r *SynchronizationTemplateRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj SynchronizationTemplate
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for SynchronizationTemplate
func (r *SynchronizationTemplateRequest) BatchUpdate(batch *BatchRequest, reqObj *SynchronizationTemplate) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for SynchronizationTemplate
func (r *SynchronizationTemplateRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

//
type SynchronizationJobCollectionValidateCredentialsRequestBuilder struct{ BaseRequestBuilder }

// ValidateCredentials action undocumented
func (b *SynchronizationJobsCollectionRequestBuilder) ValidateCredentials(reqObj *SynchronizationJobCollectionValidateCredentialsRequestParameter) *SynchronizationJobCollectionValidateCredentialsRequestBuilder {
	bb := &SynchronizationJobCollectionValidateCredentialsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateCredentials"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationJobCollectionValidateCredentialsRequest struct{ BaseRequest }

//
func (b *SynchronizationJobCollectionValidateCredentialsRequestBuilder) Request() *SynchronizationJobCollectionValidateCredentialsRequest {
	return &SynchronizationJobCollectionValidateCredentialsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationJobCollectionValidateCredentialsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *SynchronizationJobCollectionValidateCredentialsRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type SynchronizationJobPauseRequestBuilder struct{ BaseRequestBuilder }

// Pause action undocumented
func (b *SynchronizationJobRequestBuilder) Pause(reqObj *SynchronizationJobPauseRequestParameter) *SynchronizationJobPauseRequestBuilder {
	bb := &SynchronizationJobPauseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/pause"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationJobPauseRequest struct{ BaseRequest }

//
func (b *SynchronizationJobPauseRequestBuilder) Request() *SynchronizationJobPauseRequest {
	return &SynchronizationJobPauseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationJobPauseRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *SynchronizationJobPauseRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type SynchronizationJobStartRequestBuilder struct{ BaseRequestBuilder }

// Start action undocumented
func (b *SynchronizationJobRequestBuilder) Start(reqObj *SynchronizationJobStartRequestParameter) *SynchronizationJobStartRequestBuilder {
	bb := &SynchronizationJobStartRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/start"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationJobStartRequest struct{ BaseRequest }

//
func (b *SynchronizationJobStartRequestBuilder) Request() *SynchronizationJobStartRequest {
	return &SynchronizationJobStartRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationJobStartRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *SynchronizationJobStartRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type SynchronizationJobStopRequestBuilder struct{ BaseRequestBuilder }

// Stop action undocumented
func (b *SynchronizationJobRequestBuilder) Stop(reqObj *SynchronizationJobStopRequestParameter) *SynchronizationJobStopRequestBuilder {
	bb := &SynchronizationJobStopRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/stop"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationJobStopRequest struct{ BaseRequest }

//
func (b *SynchronizationJobStopRequestBuilder) Request() *SynchronizationJobStopRequest {
	return &SynchronizationJobStopRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationJobStopRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *SynchronizationJobStopRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type SynchronizationJobApplyRequestBuilder struct{ BaseRequestBuilder }

// Apply action undocumented
func (b *SynchronizationJobRequestBuilder) Apply(reqObj *SynchronizationJobApplyRequestParameter) *SynchronizationJobApplyRequestBuilder {
	bb := &SynchronizationJobApplyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/apply"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationJobApplyRequest struct{ BaseRequest }

//
func (b *SynchronizationJobApplyRequestBuilder) Request() *SynchronizationJobApplyRequest {
	return &SynchronizationJobApplyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationJobApplyRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *SynchronizationJobApplyRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type SynchronizationJobRestartRequestBuilder struct{ BaseRequestBuilder }

// Restart action undocumented
func (b *SynchronizationJobRequestBuilder) Restart(reqObj *SynchronizationJobRestartRequestParameter) *SynchronizationJobRestartRequestBuilder {
	bb := &SynchronizationJobRestartRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/restart"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationJobRestartRequest struct{ BaseRequest }

//
func (b *SynchronizationJobRestartRequestBuilder) Request() *SynchronizationJobRestartRequest {
	return &SynchronizationJobRestartRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationJobRestartRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *SynchronizationJobRestartRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type SynchronizationJobValidateCredentialsRequestBuilder struct{ BaseRequestBuilder }

// ValidateCredentials action undocumented
func (b *SynchronizationJobRequestBuilder) ValidateCredentials(reqObj *SynchronizationJobValidateCredentialsRequestParameter) *SynchronizationJobValidateCredentialsRequestBuilder {
	bb := &SynchronizationJobValidateCredentialsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateCredentials"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationJobValidateCredentialsRequest struct{ BaseRequest }

//
func (b *SynchronizationJobValidateCredentialsRequestBuilder) Request() *SynchronizationJobValidateCredentialsRequest {
	return &SynchronizationJobValidateCredentialsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationJobValidateCredentialsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *SynchronizationJobValidateCredentialsRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}

//
type SynchronizationSchemaParseExpressionRequestBuilder struct{ BaseRequestBuilder }

// ParseExpression action undocumented
func (b *SynchronizationSchemaRequestBuilder) ParseExpression(reqObj *SynchronizationSchemaParseExpressionRequestParameter) *SynchronizationSchemaParseExpressionRequestBuilder {
	bb := &SynchronizationSchemaParseExpressionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/parseExpression"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationSchemaParseExpressionRequest struct{ BaseRequest }

//
func (b *SynchronizationSchemaParseExpressionRequestBuilder) Request() *SynchronizationSchemaParseExpressionRequest {
	return &SynchronizationSchemaParseExpressionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationSchemaParseExpressionRequest) Post(ctx context.Context) (resObj *ParseExpressionResponse, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
func (r *SynchronizationSchemaParseExpressionRequest) BatchPost(batch *BatchRequest) error {
	var resObj *ParseExpressionResponse
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, resObj)
}
