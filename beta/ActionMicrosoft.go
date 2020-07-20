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

// ContainedApps returns request builder for MobileContainedApp collection
func (b *MicrosoftStoreForBusinessAppRequestBuilder) ContainedApps() *MicrosoftStoreForBusinessAppContainedAppsCollectionRequestBuilder {
	bb := &MicrosoftStoreForBusinessAppContainedAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/containedApps"
	return bb
}

// MicrosoftStoreForBusinessAppContainedAppsCollectionRequestBuilder is request builder for MobileContainedApp collection
type MicrosoftStoreForBusinessAppContainedAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileContainedApp collection
func (b *MicrosoftStoreForBusinessAppContainedAppsCollectionRequestBuilder) Request() *MicrosoftStoreForBusinessAppContainedAppsCollectionRequest {
	return &MicrosoftStoreForBusinessAppContainedAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileContainedApp item
func (b *MicrosoftStoreForBusinessAppContainedAppsCollectionRequestBuilder) ID(id string) *MobileContainedAppRequestBuilder {
	bb := &MobileContainedAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MicrosoftStoreForBusinessAppContainedAppsCollectionRequest is request for MobileContainedApp collection
type MicrosoftStoreForBusinessAppContainedAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileContainedApp collection
func (r *MicrosoftStoreForBusinessAppContainedAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MobileContainedApp, error) {
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
	var values []MobileContainedApp
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
			value  []MobileContainedApp
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

// GetN performs GET request for MobileContainedApp collection, max N pages
func (r *MicrosoftStoreForBusinessAppContainedAppsCollectionRequest) GetN(ctx context.Context, n int) ([]MobileContainedApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MobileContainedApp collection
func (r *MicrosoftStoreForBusinessAppContainedAppsCollectionRequest) Get(ctx context.Context) ([]MobileContainedApp, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MobileContainedApp collection
func (r *MicrosoftStoreForBusinessAppContainedAppsCollectionRequest) Add(ctx context.Context, reqObj *MobileContainedApp) (resObj *MobileContainedApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for MobileContainedApp collection
func (r *MicrosoftStoreForBusinessAppContainedAppsCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []MobileContainedApp
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for MobileContainedApp collection
func (r *MicrosoftStoreForBusinessAppContainedAppsCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *MobileContainedApp) error {
	var resObj []MobileContainedApp
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}
