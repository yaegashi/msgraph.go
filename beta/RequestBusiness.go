// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// BusinessFlowRequestBuilder is request builder for BusinessFlow
type BusinessFlowRequestBuilder struct{ BaseRequestBuilder }

// Request returns BusinessFlowRequest
func (b *BusinessFlowRequestBuilder) Request() *BusinessFlowRequest {
	return &BusinessFlowRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BusinessFlowRequest is request for BusinessFlow
type BusinessFlowRequest struct{ BaseRequest }

// Get performs GET request for BusinessFlow
func (r *BusinessFlowRequest) Get(ctx context.Context) (resObj *BusinessFlow, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BusinessFlow
func (r *BusinessFlowRequest) Update(ctx context.Context, reqObj *BusinessFlow) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BusinessFlow
func (r *BusinessFlowRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for BusinessFlow
func (r *BusinessFlowRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj BusinessFlow
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for BusinessFlow
func (r *BusinessFlowRequest) BatchUpdate(batch *BatchRequest, reqObj *BusinessFlow) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for BusinessFlow
func (r *BusinessFlowRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

// BusinessFlowTemplateRequestBuilder is request builder for BusinessFlowTemplate
type BusinessFlowTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns BusinessFlowTemplateRequest
func (b *BusinessFlowTemplateRequestBuilder) Request() *BusinessFlowTemplateRequest {
	return &BusinessFlowTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BusinessFlowTemplateRequest is request for BusinessFlowTemplate
type BusinessFlowTemplateRequest struct{ BaseRequest }

// Get performs GET request for BusinessFlowTemplate
func (r *BusinessFlowTemplateRequest) Get(ctx context.Context) (resObj *BusinessFlowTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BusinessFlowTemplate
func (r *BusinessFlowTemplateRequest) Update(ctx context.Context, reqObj *BusinessFlowTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BusinessFlowTemplate
func (r *BusinessFlowTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for BusinessFlowTemplate
func (r *BusinessFlowTemplateRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj BusinessFlowTemplate
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for BusinessFlowTemplate
func (r *BusinessFlowTemplateRequest) BatchUpdate(batch *BatchRequest, reqObj *BusinessFlowTemplate) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for BusinessFlowTemplate
func (r *BusinessFlowTemplateRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

//
type BusinessFlowRecordDecisionsRequestBuilder struct{ BaseRequestBuilder }

// RecordDecisions action undocumented
func (b *BusinessFlowRequestBuilder) RecordDecisions(reqObj *BusinessFlowRecordDecisionsRequestParameter) *BusinessFlowRecordDecisionsRequestBuilder {
	bb := &BusinessFlowRecordDecisionsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/recordDecisions"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type BusinessFlowRecordDecisionsRequest struct{ BaseRequest }

//
func (b *BusinessFlowRecordDecisionsRequestBuilder) Request() *BusinessFlowRecordDecisionsRequest {
	return &BusinessFlowRecordDecisionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *BusinessFlowRecordDecisionsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
func (r *BusinessFlowRecordDecisionsRequest) BatchPost(batch *BatchRequest) error {
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), r.requestObject, nil)
}
