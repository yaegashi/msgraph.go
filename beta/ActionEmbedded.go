// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// EmbeddedSIMActivationCodePoolAssignRequestParameter undocumented
type EmbeddedSIMActivationCodePoolAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []EmbeddedSIMActivationCodePoolAssignment `json:"assignments,omitempty"`
}

// Assignments returns request builder for EmbeddedSIMActivationCodePoolAssignment collection
func (b *EmbeddedSIMActivationCodePoolRequestBuilder) Assignments() *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder {
	bb := &EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder is request builder for EmbeddedSIMActivationCodePoolAssignment collection
type EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EmbeddedSIMActivationCodePoolAssignment collection
func (b *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder) Request() *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest {
	return &EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EmbeddedSIMActivationCodePoolAssignment item
func (b *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder) ID(id string) *EmbeddedSIMActivationCodePoolAssignmentRequestBuilder {
	bb := &EmbeddedSIMActivationCodePoolAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest is request for EmbeddedSIMActivationCodePoolAssignment collection
type EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EmbeddedSIMActivationCodePoolAssignment collection
func (r *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EmbeddedSIMActivationCodePoolAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EmbeddedSIMActivationCodePoolAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []EmbeddedSIMActivationCodePoolAssignment
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for EmbeddedSIMActivationCodePoolAssignment collection, max N pages
func (r *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]EmbeddedSIMActivationCodePoolAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EmbeddedSIMActivationCodePoolAssignment collection
func (r *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest) Get(ctx context.Context) ([]EmbeddedSIMActivationCodePoolAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EmbeddedSIMActivationCodePoolAssignment collection
func (r *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *EmbeddedSIMActivationCodePoolAssignment) (resObj *EmbeddedSIMActivationCodePoolAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for EmbeddedSIMActivationCodePoolAssignment collection
func (r *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []EmbeddedSIMActivationCodePoolAssignment
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for EmbeddedSIMActivationCodePoolAssignment collection
func (r *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *EmbeddedSIMActivationCodePoolAssignment) error {
	var resObj []EmbeddedSIMActivationCodePoolAssignment
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// DeviceStates returns request builder for EmbeddedSIMDeviceState collection
func (b *EmbeddedSIMActivationCodePoolRequestBuilder) DeviceStates() *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder {
	bb := &EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceStates"
	return bb
}

// EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder is request builder for EmbeddedSIMDeviceState collection
type EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EmbeddedSIMDeviceState collection
func (b *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder) Request() *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest {
	return &EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EmbeddedSIMDeviceState item
func (b *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder) ID(id string) *EmbeddedSIMDeviceStateRequestBuilder {
	bb := &EmbeddedSIMDeviceStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest is request for EmbeddedSIMDeviceState collection
type EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EmbeddedSIMDeviceState collection
func (r *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EmbeddedSIMDeviceState, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EmbeddedSIMDeviceState
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []EmbeddedSIMDeviceState
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for EmbeddedSIMDeviceState collection, max N pages
func (r *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest) GetN(ctx context.Context, n int) ([]EmbeddedSIMDeviceState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EmbeddedSIMDeviceState collection
func (r *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest) Get(ctx context.Context) ([]EmbeddedSIMDeviceState, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EmbeddedSIMDeviceState collection
func (r *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest) Add(ctx context.Context, reqObj *EmbeddedSIMDeviceState) (resObj *EmbeddedSIMDeviceState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for EmbeddedSIMDeviceState collection
func (r *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []EmbeddedSIMDeviceState
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for EmbeddedSIMDeviceState collection
func (r *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *EmbeddedSIMDeviceState) error {
	var resObj []EmbeddedSIMDeviceState
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}
