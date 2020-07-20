// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// ActiveDirectoryWindowsAutopilotDeploymentProfileRequestBuilder is request builder for ActiveDirectoryWindowsAutopilotDeploymentProfile
type ActiveDirectoryWindowsAutopilotDeploymentProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns ActiveDirectoryWindowsAutopilotDeploymentProfileRequest
func (b *ActiveDirectoryWindowsAutopilotDeploymentProfileRequestBuilder) Request() *ActiveDirectoryWindowsAutopilotDeploymentProfileRequest {
	return &ActiveDirectoryWindowsAutopilotDeploymentProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ActiveDirectoryWindowsAutopilotDeploymentProfileRequest is request for ActiveDirectoryWindowsAutopilotDeploymentProfile
type ActiveDirectoryWindowsAutopilotDeploymentProfileRequest struct{ BaseRequest }

// Get performs GET request for ActiveDirectoryWindowsAutopilotDeploymentProfile
func (r *ActiveDirectoryWindowsAutopilotDeploymentProfileRequest) Get(ctx context.Context) (resObj *ActiveDirectoryWindowsAutopilotDeploymentProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ActiveDirectoryWindowsAutopilotDeploymentProfile
func (r *ActiveDirectoryWindowsAutopilotDeploymentProfileRequest) Update(ctx context.Context, reqObj *ActiveDirectoryWindowsAutopilotDeploymentProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ActiveDirectoryWindowsAutopilotDeploymentProfile
func (r *ActiveDirectoryWindowsAutopilotDeploymentProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for ActiveDirectoryWindowsAutopilotDeploymentProfile
func (r *ActiveDirectoryWindowsAutopilotDeploymentProfileRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj ActiveDirectoryWindowsAutopilotDeploymentProfile
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for ActiveDirectoryWindowsAutopilotDeploymentProfile
func (r *ActiveDirectoryWindowsAutopilotDeploymentProfileRequest) BatchUpdate(batch *BatchRequest, reqObj *ActiveDirectoryWindowsAutopilotDeploymentProfile) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for ActiveDirectoryWindowsAutopilotDeploymentProfile
func (r *ActiveDirectoryWindowsAutopilotDeploymentProfileRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
