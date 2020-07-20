// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// IdentityProviderRequestBuilder is request builder for IdentityProvider
type IdentityProviderRequestBuilder struct{ BaseRequestBuilder }

// Request returns IdentityProviderRequest
func (b *IdentityProviderRequestBuilder) Request() *IdentityProviderRequest {
	return &IdentityProviderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IdentityProviderRequest is request for IdentityProvider
type IdentityProviderRequest struct{ BaseRequest }

// Get performs GET request for IdentityProvider
func (r *IdentityProviderRequest) Get(ctx context.Context) (resObj *IdentityProvider, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IdentityProvider
func (r *IdentityProviderRequest) Update(ctx context.Context, reqObj *IdentityProvider) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IdentityProvider
func (r *IdentityProviderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for IdentityProvider
func (r *IdentityProviderRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj IdentityProvider
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for IdentityProvider
func (r *IdentityProviderRequest) BatchUpdate(batch *BatchRequest, reqObj *IdentityProvider) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for IdentityProvider
func (r *IdentityProviderRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
