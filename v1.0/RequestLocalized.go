// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// LocalizedNotificationMessageRequestBuilder is request builder for LocalizedNotificationMessage
type LocalizedNotificationMessageRequestBuilder struct{ BaseRequestBuilder }

// Request returns LocalizedNotificationMessageRequest
func (b *LocalizedNotificationMessageRequestBuilder) Request() *LocalizedNotificationMessageRequest {
	return &LocalizedNotificationMessageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// LocalizedNotificationMessageRequest is request for LocalizedNotificationMessage
type LocalizedNotificationMessageRequest struct{ BaseRequest }

// Get performs GET request for LocalizedNotificationMessage
func (r *LocalizedNotificationMessageRequest) Get(ctx context.Context) (resObj *LocalizedNotificationMessage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for LocalizedNotificationMessage
func (r *LocalizedNotificationMessageRequest) Update(ctx context.Context, reqObj *LocalizedNotificationMessage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for LocalizedNotificationMessage
func (r *LocalizedNotificationMessageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for LocalizedNotificationMessage
func (r *LocalizedNotificationMessageRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj LocalizedNotificationMessage
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for LocalizedNotificationMessage
func (r *LocalizedNotificationMessageRequest) BatchUpdate(batch *BatchRequest, reqObj *LocalizedNotificationMessage) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for LocalizedNotificationMessage
func (r *LocalizedNotificationMessageRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
