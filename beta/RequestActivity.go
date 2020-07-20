// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"strings"
)

// ActivityHistoryItemRequestBuilder is request builder for ActivityHistoryItem
type ActivityHistoryItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns ActivityHistoryItemRequest
func (b *ActivityHistoryItemRequestBuilder) Request() *ActivityHistoryItemRequest {
	return &ActivityHistoryItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ActivityHistoryItemRequest is request for ActivityHistoryItem
type ActivityHistoryItemRequest struct{ BaseRequest }

// Get performs GET request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Get(ctx context.Context) (resObj *ActivityHistoryItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Update(ctx context.Context, reqObj *ActivityHistoryItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj ActivityHistoryItem
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) BatchUpdate(batch *BatchRequest, reqObj *ActivityHistoryItem) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}

// ActivityStatisticsRequestBuilder is request builder for ActivityStatistics
type ActivityStatisticsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ActivityStatisticsRequest
func (b *ActivityStatisticsRequestBuilder) Request() *ActivityStatisticsRequest {
	return &ActivityStatisticsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ActivityStatisticsRequest is request for ActivityStatistics
type ActivityStatisticsRequest struct{ BaseRequest }

// Get performs GET request for ActivityStatistics
func (r *ActivityStatisticsRequest) Get(ctx context.Context) (resObj *ActivityStatistics, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ActivityStatistics
func (r *ActivityStatisticsRequest) Update(ctx context.Context, reqObj *ActivityStatistics) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ActivityStatistics
func (r *ActivityStatisticsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BatchGet adds Get operation to Batch for ActivityStatistics
func (r *ActivityStatisticsRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj ActivityStatistics
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchUpdate adds Update operation to Batch for ActivityStatistics
func (r *ActivityStatisticsRequest) BatchUpdate(batch *BatchRequest, reqObj *ActivityStatistics) error {
	return batch.Add("PATCH", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, nil)
}

// BatchDelete adds Delete operation to Batch for ActivityStatistics
func (r *ActivityStatisticsRequest) BatchDelete(batch *BatchRequest) error {
	return batch.Add("DELETE", strings.TrimPrefix(r.baseURL, defaultBaseURL), nil, nil)
}
