// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// AadUserConversationMemberRequestBuilder is request builder for AadUserConversationMember
type AadUserConversationMemberRequestBuilder struct{ BaseRequestBuilder }

// Request returns AadUserConversationMemberRequest
func (b *AadUserConversationMemberRequestBuilder) Request() *AadUserConversationMemberRequest {
	return &AadUserConversationMemberRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AadUserConversationMemberRequest is request for AadUserConversationMember
type AadUserConversationMemberRequest struct{ BaseRequest }

// Get performs GET request for AadUserConversationMember
func (r *AadUserConversationMemberRequest) Get(ctx context.Context) (resObj *AadUserConversationMember, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AadUserConversationMember
func (r *AadUserConversationMemberRequest) Update(ctx context.Context, reqObj *AadUserConversationMember) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AadUserConversationMember
func (r *AadUserConversationMemberRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for AadUserConversationMember
func (r *AadUserConversationMemberRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj AadUserConversationMember
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for AadUserConversationMember
func (r *AadUserConversationMemberRequest) BatchUpdate(batch *BatchRequest, reqObj *AadUserConversationMember) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for AadUserConversationMember
func (r *AadUserConversationMemberRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
