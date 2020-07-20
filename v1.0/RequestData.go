// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// DataPolicyOperationRequestBuilder is request builder for DataPolicyOperation
type DataPolicyOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns DataPolicyOperationRequest
func (b *DataPolicyOperationRequestBuilder) Request() *DataPolicyOperationRequest {
	return &DataPolicyOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DataPolicyOperationRequest is request for DataPolicyOperation
type DataPolicyOperationRequest struct{ BaseRequest }

// Get performs GET request for DataPolicyOperation
func (r *DataPolicyOperationRequest) Get(ctx context.Context) (resObj *DataPolicyOperation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DataPolicyOperation
func (r *DataPolicyOperationRequest) Update(ctx context.Context, reqObj *DataPolicyOperation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DataPolicyOperation
func (r *DataPolicyOperationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for DataPolicyOperation
func (r *DataPolicyOperationRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj DataPolicyOperation
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for DataPolicyOperation
func (r *DataPolicyOperationRequest) BatchUpdate(batch *BatchRequest, reqObj *DataPolicyOperation) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for DataPolicyOperation
func (r *DataPolicyOperationRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
