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

// Group is navigation property
func (b *EducationClassRequestBuilder) Group() *GroupRequestBuilder {
	bb := &GroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/group"
	return bb
}

// Members returns request builder for EducationUser collection
func (b *EducationClassRequestBuilder) Members() *EducationClassMembersCollectionRequestBuilder {
	bb := &EducationClassMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/members"
	return bb
}

// EducationClassMembersCollectionRequestBuilder is request builder for EducationUser collection
type EducationClassMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationUser collection
func (b *EducationClassMembersCollectionRequestBuilder) Request() *EducationClassMembersCollectionRequest {
	return &EducationClassMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationUser item
func (b *EducationClassMembersCollectionRequestBuilder) ID(id string) *EducationUserRequestBuilder {
	bb := &EducationUserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationClassMembersCollectionRequest is request for EducationUser collection
type EducationClassMembersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationUser collection
func (r *EducationClassMembersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EducationUser, error) {
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
	var values []EducationUser
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
			value  []EducationUser
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

// GetN performs GET request for EducationUser collection, max N pages
func (r *EducationClassMembersCollectionRequest) GetN(ctx context.Context, n int) ([]EducationUser, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EducationUser collection
func (r *EducationClassMembersCollectionRequest) Get(ctx context.Context) ([]EducationUser, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EducationUser collection
func (r *EducationClassMembersCollectionRequest) Add(ctx context.Context, reqObj *EducationUser) (resObj *EducationUser, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for EducationUser collection
func (r *EducationClassMembersCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []EducationUser
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for EducationUser collection
func (r *EducationClassMembersCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *EducationUser) error {
	var resObj []EducationUser
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// Schools returns request builder for EducationSchool collection
func (b *EducationClassRequestBuilder) Schools() *EducationClassSchoolsCollectionRequestBuilder {
	bb := &EducationClassSchoolsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/schools"
	return bb
}

// EducationClassSchoolsCollectionRequestBuilder is request builder for EducationSchool collection
type EducationClassSchoolsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationSchool collection
func (b *EducationClassSchoolsCollectionRequestBuilder) Request() *EducationClassSchoolsCollectionRequest {
	return &EducationClassSchoolsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationSchool item
func (b *EducationClassSchoolsCollectionRequestBuilder) ID(id string) *EducationSchoolRequestBuilder {
	bb := &EducationSchoolRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationClassSchoolsCollectionRequest is request for EducationSchool collection
type EducationClassSchoolsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationSchool collection
func (r *EducationClassSchoolsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EducationSchool, error) {
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
	var values []EducationSchool
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
			value  []EducationSchool
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

// GetN performs GET request for EducationSchool collection, max N pages
func (r *EducationClassSchoolsCollectionRequest) GetN(ctx context.Context, n int) ([]EducationSchool, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EducationSchool collection
func (r *EducationClassSchoolsCollectionRequest) Get(ctx context.Context) ([]EducationSchool, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EducationSchool collection
func (r *EducationClassSchoolsCollectionRequest) Add(ctx context.Context, reqObj *EducationSchool) (resObj *EducationSchool, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for EducationSchool collection
func (r *EducationClassSchoolsCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []EducationSchool
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for EducationSchool collection
func (r *EducationClassSchoolsCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *EducationSchool) error {
	var resObj []EducationSchool
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// Teachers returns request builder for EducationUser collection
func (b *EducationClassRequestBuilder) Teachers() *EducationClassTeachersCollectionRequestBuilder {
	bb := &EducationClassTeachersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/teachers"
	return bb
}

// EducationClassTeachersCollectionRequestBuilder is request builder for EducationUser collection
type EducationClassTeachersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationUser collection
func (b *EducationClassTeachersCollectionRequestBuilder) Request() *EducationClassTeachersCollectionRequest {
	return &EducationClassTeachersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationUser item
func (b *EducationClassTeachersCollectionRequestBuilder) ID(id string) *EducationUserRequestBuilder {
	bb := &EducationUserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationClassTeachersCollectionRequest is request for EducationUser collection
type EducationClassTeachersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationUser collection
func (r *EducationClassTeachersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EducationUser, error) {
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
	var values []EducationUser
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
			value  []EducationUser
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

// GetN performs GET request for EducationUser collection, max N pages
func (r *EducationClassTeachersCollectionRequest) GetN(ctx context.Context, n int) ([]EducationUser, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EducationUser collection
func (r *EducationClassTeachersCollectionRequest) Get(ctx context.Context) ([]EducationUser, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EducationUser collection
func (r *EducationClassTeachersCollectionRequest) Add(ctx context.Context, reqObj *EducationUser) (resObj *EducationUser, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for EducationUser collection
func (r *EducationClassTeachersCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []EducationUser
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for EducationUser collection
func (r *EducationClassTeachersCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *EducationUser) error {
	var resObj []EducationUser
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// Classes returns request builder for EducationClass collection
func (b *EducationRootRequestBuilder) Classes() *EducationRootClassesCollectionRequestBuilder {
	bb := &EducationRootClassesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/classes"
	return bb
}

// EducationRootClassesCollectionRequestBuilder is request builder for EducationClass collection
type EducationRootClassesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationClass collection
func (b *EducationRootClassesCollectionRequestBuilder) Request() *EducationRootClassesCollectionRequest {
	return &EducationRootClassesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationClass item
func (b *EducationRootClassesCollectionRequestBuilder) ID(id string) *EducationClassRequestBuilder {
	bb := &EducationClassRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationRootClassesCollectionRequest is request for EducationClass collection
type EducationRootClassesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationClass collection
func (r *EducationRootClassesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EducationClass, error) {
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
	var values []EducationClass
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
			value  []EducationClass
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

// GetN performs GET request for EducationClass collection, max N pages
func (r *EducationRootClassesCollectionRequest) GetN(ctx context.Context, n int) ([]EducationClass, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EducationClass collection
func (r *EducationRootClassesCollectionRequest) Get(ctx context.Context) ([]EducationClass, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EducationClass collection
func (r *EducationRootClassesCollectionRequest) Add(ctx context.Context, reqObj *EducationClass) (resObj *EducationClass, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for EducationClass collection
func (r *EducationRootClassesCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []EducationClass
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for EducationClass collection
func (r *EducationRootClassesCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *EducationClass) error {
	var resObj []EducationClass
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// Me is navigation property
func (b *EducationRootRequestBuilder) Me() *EducationUserRequestBuilder {
	bb := &EducationUserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/me"
	return bb
}

// Schools returns request builder for EducationSchool collection
func (b *EducationRootRequestBuilder) Schools() *EducationRootSchoolsCollectionRequestBuilder {
	bb := &EducationRootSchoolsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/schools"
	return bb
}

// EducationRootSchoolsCollectionRequestBuilder is request builder for EducationSchool collection
type EducationRootSchoolsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationSchool collection
func (b *EducationRootSchoolsCollectionRequestBuilder) Request() *EducationRootSchoolsCollectionRequest {
	return &EducationRootSchoolsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationSchool item
func (b *EducationRootSchoolsCollectionRequestBuilder) ID(id string) *EducationSchoolRequestBuilder {
	bb := &EducationSchoolRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationRootSchoolsCollectionRequest is request for EducationSchool collection
type EducationRootSchoolsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationSchool collection
func (r *EducationRootSchoolsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EducationSchool, error) {
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
	var values []EducationSchool
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
			value  []EducationSchool
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

// GetN performs GET request for EducationSchool collection, max N pages
func (r *EducationRootSchoolsCollectionRequest) GetN(ctx context.Context, n int) ([]EducationSchool, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EducationSchool collection
func (r *EducationRootSchoolsCollectionRequest) Get(ctx context.Context) ([]EducationSchool, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EducationSchool collection
func (r *EducationRootSchoolsCollectionRequest) Add(ctx context.Context, reqObj *EducationSchool) (resObj *EducationSchool, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for EducationSchool collection
func (r *EducationRootSchoolsCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []EducationSchool
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for EducationSchool collection
func (r *EducationRootSchoolsCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *EducationSchool) error {
	var resObj []EducationSchool
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// Users returns request builder for EducationUser collection
func (b *EducationRootRequestBuilder) Users() *EducationRootUsersCollectionRequestBuilder {
	bb := &EducationRootUsersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/users"
	return bb
}

// EducationRootUsersCollectionRequestBuilder is request builder for EducationUser collection
type EducationRootUsersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationUser collection
func (b *EducationRootUsersCollectionRequestBuilder) Request() *EducationRootUsersCollectionRequest {
	return &EducationRootUsersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationUser item
func (b *EducationRootUsersCollectionRequestBuilder) ID(id string) *EducationUserRequestBuilder {
	bb := &EducationUserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationRootUsersCollectionRequest is request for EducationUser collection
type EducationRootUsersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationUser collection
func (r *EducationRootUsersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EducationUser, error) {
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
	var values []EducationUser
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
			value  []EducationUser
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

// GetN performs GET request for EducationUser collection, max N pages
func (r *EducationRootUsersCollectionRequest) GetN(ctx context.Context, n int) ([]EducationUser, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EducationUser collection
func (r *EducationRootUsersCollectionRequest) Get(ctx context.Context) ([]EducationUser, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EducationUser collection
func (r *EducationRootUsersCollectionRequest) Add(ctx context.Context, reqObj *EducationUser) (resObj *EducationUser, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for EducationUser collection
func (r *EducationRootUsersCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []EducationUser
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for EducationUser collection
func (r *EducationRootUsersCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *EducationUser) error {
	var resObj []EducationUser
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// Classes returns request builder for EducationClass collection
func (b *EducationSchoolRequestBuilder) Classes() *EducationSchoolClassesCollectionRequestBuilder {
	bb := &EducationSchoolClassesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/classes"
	return bb
}

// EducationSchoolClassesCollectionRequestBuilder is request builder for EducationClass collection
type EducationSchoolClassesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationClass collection
func (b *EducationSchoolClassesCollectionRequestBuilder) Request() *EducationSchoolClassesCollectionRequest {
	return &EducationSchoolClassesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationClass item
func (b *EducationSchoolClassesCollectionRequestBuilder) ID(id string) *EducationClassRequestBuilder {
	bb := &EducationClassRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationSchoolClassesCollectionRequest is request for EducationClass collection
type EducationSchoolClassesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationClass collection
func (r *EducationSchoolClassesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EducationClass, error) {
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
	var values []EducationClass
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
			value  []EducationClass
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

// GetN performs GET request for EducationClass collection, max N pages
func (r *EducationSchoolClassesCollectionRequest) GetN(ctx context.Context, n int) ([]EducationClass, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EducationClass collection
func (r *EducationSchoolClassesCollectionRequest) Get(ctx context.Context) ([]EducationClass, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EducationClass collection
func (r *EducationSchoolClassesCollectionRequest) Add(ctx context.Context, reqObj *EducationClass) (resObj *EducationClass, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for EducationClass collection
func (r *EducationSchoolClassesCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []EducationClass
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for EducationClass collection
func (r *EducationSchoolClassesCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *EducationClass) error {
	var resObj []EducationClass
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// Users returns request builder for EducationUser collection
func (b *EducationSchoolRequestBuilder) Users() *EducationSchoolUsersCollectionRequestBuilder {
	bb := &EducationSchoolUsersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/users"
	return bb
}

// EducationSchoolUsersCollectionRequestBuilder is request builder for EducationUser collection
type EducationSchoolUsersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationUser collection
func (b *EducationSchoolUsersCollectionRequestBuilder) Request() *EducationSchoolUsersCollectionRequest {
	return &EducationSchoolUsersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationUser item
func (b *EducationSchoolUsersCollectionRequestBuilder) ID(id string) *EducationUserRequestBuilder {
	bb := &EducationUserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationSchoolUsersCollectionRequest is request for EducationUser collection
type EducationSchoolUsersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationUser collection
func (r *EducationSchoolUsersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EducationUser, error) {
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
	var values []EducationUser
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
			value  []EducationUser
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

// GetN performs GET request for EducationUser collection, max N pages
func (r *EducationSchoolUsersCollectionRequest) GetN(ctx context.Context, n int) ([]EducationUser, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EducationUser collection
func (r *EducationSchoolUsersCollectionRequest) Get(ctx context.Context) ([]EducationUser, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EducationUser collection
func (r *EducationSchoolUsersCollectionRequest) Add(ctx context.Context, reqObj *EducationUser) (resObj *EducationUser, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for EducationUser collection
func (r *EducationSchoolUsersCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []EducationUser
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for EducationUser collection
func (r *EducationSchoolUsersCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *EducationUser) error {
	var resObj []EducationUser
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// Classes returns request builder for EducationClass collection
func (b *EducationUserRequestBuilder) Classes() *EducationUserClassesCollectionRequestBuilder {
	bb := &EducationUserClassesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/classes"
	return bb
}

// EducationUserClassesCollectionRequestBuilder is request builder for EducationClass collection
type EducationUserClassesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationClass collection
func (b *EducationUserClassesCollectionRequestBuilder) Request() *EducationUserClassesCollectionRequest {
	return &EducationUserClassesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationClass item
func (b *EducationUserClassesCollectionRequestBuilder) ID(id string) *EducationClassRequestBuilder {
	bb := &EducationClassRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationUserClassesCollectionRequest is request for EducationClass collection
type EducationUserClassesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationClass collection
func (r *EducationUserClassesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EducationClass, error) {
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
	var values []EducationClass
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
			value  []EducationClass
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

// GetN performs GET request for EducationClass collection, max N pages
func (r *EducationUserClassesCollectionRequest) GetN(ctx context.Context, n int) ([]EducationClass, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EducationClass collection
func (r *EducationUserClassesCollectionRequest) Get(ctx context.Context) ([]EducationClass, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EducationClass collection
func (r *EducationUserClassesCollectionRequest) Add(ctx context.Context, reqObj *EducationClass) (resObj *EducationClass, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for EducationClass collection
func (r *EducationUserClassesCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []EducationClass
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for EducationClass collection
func (r *EducationUserClassesCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *EducationClass) error {
	var resObj []EducationClass
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// Schools returns request builder for EducationSchool collection
func (b *EducationUserRequestBuilder) Schools() *EducationUserSchoolsCollectionRequestBuilder {
	bb := &EducationUserSchoolsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/schools"
	return bb
}

// EducationUserSchoolsCollectionRequestBuilder is request builder for EducationSchool collection
type EducationUserSchoolsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationSchool collection
func (b *EducationUserSchoolsCollectionRequestBuilder) Request() *EducationUserSchoolsCollectionRequest {
	return &EducationUserSchoolsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationSchool item
func (b *EducationUserSchoolsCollectionRequestBuilder) ID(id string) *EducationSchoolRequestBuilder {
	bb := &EducationSchoolRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationUserSchoolsCollectionRequest is request for EducationSchool collection
type EducationUserSchoolsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationSchool collection
func (r *EducationUserSchoolsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EducationSchool, error) {
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
	var values []EducationSchool
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
			value  []EducationSchool
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

// GetN performs GET request for EducationSchool collection, max N pages
func (r *EducationUserSchoolsCollectionRequest) GetN(ctx context.Context, n int) ([]EducationSchool, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EducationSchool collection
func (r *EducationUserSchoolsCollectionRequest) Get(ctx context.Context) ([]EducationSchool, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EducationSchool collection
func (r *EducationUserSchoolsCollectionRequest) Add(ctx context.Context, reqObj *EducationSchool) (resObj *EducationSchool, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for EducationSchool collection
func (r *EducationUserSchoolsCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []EducationSchool
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for EducationSchool collection
func (r *EducationUserSchoolsCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *EducationSchool) error {
	var resObj []EducationSchool
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// User is navigation property
func (b *EducationUserRequestBuilder) User() *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/user"
	return bb
}
