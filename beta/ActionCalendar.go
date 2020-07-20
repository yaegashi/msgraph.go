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

// CalendarGetScheduleRequestParameter undocumented
type CalendarGetScheduleRequestParameter struct {
	// Schedules undocumented
	Schedules []string `json:"Schedules,omitempty"`
	// EndTime undocumented
	EndTime *DateTimeTimeZone `json:"EndTime,omitempty"`
	// StartTime undocumented
	StartTime *DateTimeTimeZone `json:"StartTime,omitempty"`
	// AvailabilityViewInterval undocumented
	AvailabilityViewInterval *int `json:"AvailabilityViewInterval,omitempty"`
}

// CalendarSharingMessageAcceptRequestParameter undocumented
type CalendarSharingMessageAcceptRequestParameter struct {
}

// CalendarPermissions returns request builder for CalendarPermission collection
func (b *CalendarRequestBuilder) CalendarPermissions() *CalendarCalendarPermissionsCollectionRequestBuilder {
	bb := &CalendarCalendarPermissionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/calendarPermissions"
	return bb
}

// CalendarCalendarPermissionsCollectionRequestBuilder is request builder for CalendarPermission collection
type CalendarCalendarPermissionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CalendarPermission collection
func (b *CalendarCalendarPermissionsCollectionRequestBuilder) Request() *CalendarCalendarPermissionsCollectionRequest {
	return &CalendarCalendarPermissionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CalendarPermission item
func (b *CalendarCalendarPermissionsCollectionRequestBuilder) ID(id string) *CalendarPermissionRequestBuilder {
	bb := &CalendarPermissionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CalendarCalendarPermissionsCollectionRequest is request for CalendarPermission collection
type CalendarCalendarPermissionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CalendarPermission collection
func (r *CalendarCalendarPermissionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]CalendarPermission, error) {
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
	var values []CalendarPermission
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
			value  []CalendarPermission
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

// GetN performs GET request for CalendarPermission collection, max N pages
func (r *CalendarCalendarPermissionsCollectionRequest) GetN(ctx context.Context, n int) ([]CalendarPermission, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for CalendarPermission collection
func (r *CalendarCalendarPermissionsCollectionRequest) Get(ctx context.Context) ([]CalendarPermission, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for CalendarPermission collection
func (r *CalendarCalendarPermissionsCollectionRequest) Add(ctx context.Context, reqObj *CalendarPermission) (resObj *CalendarPermission, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for CalendarPermission collection
func (r *CalendarCalendarPermissionsCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []CalendarPermission
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for CalendarPermission collection
func (r *CalendarCalendarPermissionsCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *CalendarPermission) error {
	var resObj []CalendarPermission
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// CalendarView returns request builder for Event collection
func (b *CalendarRequestBuilder) CalendarView() *CalendarCalendarViewCollectionRequestBuilder {
	bb := &CalendarCalendarViewCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/calendarView"
	return bb
}

// CalendarCalendarViewCollectionRequestBuilder is request builder for Event collection
type CalendarCalendarViewCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Event collection
func (b *CalendarCalendarViewCollectionRequestBuilder) Request() *CalendarCalendarViewCollectionRequest {
	return &CalendarCalendarViewCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Event item
func (b *CalendarCalendarViewCollectionRequestBuilder) ID(id string) *EventRequestBuilder {
	bb := &EventRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CalendarCalendarViewCollectionRequest is request for Event collection
type CalendarCalendarViewCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Event collection
func (r *CalendarCalendarViewCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Event, error) {
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
	var values []Event
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
			value  []Event
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

// GetN performs GET request for Event collection, max N pages
func (r *CalendarCalendarViewCollectionRequest) GetN(ctx context.Context, n int) ([]Event, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Event collection
func (r *CalendarCalendarViewCollectionRequest) Get(ctx context.Context) ([]Event, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Event collection
func (r *CalendarCalendarViewCollectionRequest) Add(ctx context.Context, reqObj *Event) (resObj *Event, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for Event collection
func (r *CalendarCalendarViewCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []Event
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for Event collection
func (r *CalendarCalendarViewCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *Event) error {
	var resObj []Event
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// Events returns request builder for Event collection
func (b *CalendarRequestBuilder) Events() *CalendarEventsCollectionRequestBuilder {
	bb := &CalendarEventsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/events"
	return bb
}

// CalendarEventsCollectionRequestBuilder is request builder for Event collection
type CalendarEventsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Event collection
func (b *CalendarEventsCollectionRequestBuilder) Request() *CalendarEventsCollectionRequest {
	return &CalendarEventsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Event item
func (b *CalendarEventsCollectionRequestBuilder) ID(id string) *EventRequestBuilder {
	bb := &EventRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CalendarEventsCollectionRequest is request for Event collection
type CalendarEventsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Event collection
func (r *CalendarEventsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Event, error) {
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
	var values []Event
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
			value  []Event
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

// GetN performs GET request for Event collection, max N pages
func (r *CalendarEventsCollectionRequest) GetN(ctx context.Context, n int) ([]Event, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Event collection
func (r *CalendarEventsCollectionRequest) Get(ctx context.Context) ([]Event, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Event collection
func (r *CalendarEventsCollectionRequest) Add(ctx context.Context, reqObj *Event) (resObj *Event, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for Event collection
func (r *CalendarEventsCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []Event
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for Event collection
func (r *CalendarEventsCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *Event) error {
	var resObj []Event
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// MultiValueExtendedProperties returns request builder for MultiValueLegacyExtendedProperty collection
func (b *CalendarRequestBuilder) MultiValueExtendedProperties() *CalendarMultiValueExtendedPropertiesCollectionRequestBuilder {
	bb := &CalendarMultiValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/multiValueExtendedProperties"
	return bb
}

// CalendarMultiValueExtendedPropertiesCollectionRequestBuilder is request builder for MultiValueLegacyExtendedProperty collection
type CalendarMultiValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MultiValueLegacyExtendedProperty collection
func (b *CalendarMultiValueExtendedPropertiesCollectionRequestBuilder) Request() *CalendarMultiValueExtendedPropertiesCollectionRequest {
	return &CalendarMultiValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MultiValueLegacyExtendedProperty item
func (b *CalendarMultiValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *MultiValueLegacyExtendedPropertyRequestBuilder {
	bb := &MultiValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CalendarMultiValueExtendedPropertiesCollectionRequest is request for MultiValueLegacyExtendedProperty collection
type CalendarMultiValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MultiValueLegacyExtendedProperty collection
func (r *CalendarMultiValueExtendedPropertiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MultiValueLegacyExtendedProperty, error) {
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
	var values []MultiValueLegacyExtendedProperty
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
			value  []MultiValueLegacyExtendedProperty
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

// GetN performs GET request for MultiValueLegacyExtendedProperty collection, max N pages
func (r *CalendarMultiValueExtendedPropertiesCollectionRequest) GetN(ctx context.Context, n int) ([]MultiValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MultiValueLegacyExtendedProperty collection
func (r *CalendarMultiValueExtendedPropertiesCollectionRequest) Get(ctx context.Context) ([]MultiValueLegacyExtendedProperty, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MultiValueLegacyExtendedProperty collection
func (r *CalendarMultiValueExtendedPropertiesCollectionRequest) Add(ctx context.Context, reqObj *MultiValueLegacyExtendedProperty) (resObj *MultiValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for MultiValueLegacyExtendedProperty collection
func (r *CalendarMultiValueExtendedPropertiesCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []MultiValueLegacyExtendedProperty
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for MultiValueLegacyExtendedProperty collection
func (r *CalendarMultiValueExtendedPropertiesCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *MultiValueLegacyExtendedProperty) error {
	var resObj []MultiValueLegacyExtendedProperty
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// SingleValueExtendedProperties returns request builder for SingleValueLegacyExtendedProperty collection
func (b *CalendarRequestBuilder) SingleValueExtendedProperties() *CalendarSingleValueExtendedPropertiesCollectionRequestBuilder {
	bb := &CalendarSingleValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/singleValueExtendedProperties"
	return bb
}

// CalendarSingleValueExtendedPropertiesCollectionRequestBuilder is request builder for SingleValueLegacyExtendedProperty collection
type CalendarSingleValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SingleValueLegacyExtendedProperty collection
func (b *CalendarSingleValueExtendedPropertiesCollectionRequestBuilder) Request() *CalendarSingleValueExtendedPropertiesCollectionRequest {
	return &CalendarSingleValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SingleValueLegacyExtendedProperty item
func (b *CalendarSingleValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *SingleValueLegacyExtendedPropertyRequestBuilder {
	bb := &SingleValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CalendarSingleValueExtendedPropertiesCollectionRequest is request for SingleValueLegacyExtendedProperty collection
type CalendarSingleValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SingleValueLegacyExtendedProperty collection
func (r *CalendarSingleValueExtendedPropertiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SingleValueLegacyExtendedProperty, error) {
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
	var values []SingleValueLegacyExtendedProperty
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
			value  []SingleValueLegacyExtendedProperty
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

// GetN performs GET request for SingleValueLegacyExtendedProperty collection, max N pages
func (r *CalendarSingleValueExtendedPropertiesCollectionRequest) GetN(ctx context.Context, n int) ([]SingleValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SingleValueLegacyExtendedProperty collection
func (r *CalendarSingleValueExtendedPropertiesCollectionRequest) Get(ctx context.Context) ([]SingleValueLegacyExtendedProperty, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SingleValueLegacyExtendedProperty collection
func (r *CalendarSingleValueExtendedPropertiesCollectionRequest) Add(ctx context.Context, reqObj *SingleValueLegacyExtendedProperty) (resObj *SingleValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for SingleValueLegacyExtendedProperty collection
func (r *CalendarSingleValueExtendedPropertiesCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []SingleValueLegacyExtendedProperty
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for SingleValueLegacyExtendedProperty collection
func (r *CalendarSingleValueExtendedPropertiesCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *SingleValueLegacyExtendedProperty) error {
	var resObj []SingleValueLegacyExtendedProperty
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}

// Calendars returns request builder for Calendar collection
func (b *CalendarGroupRequestBuilder) Calendars() *CalendarGroupCalendarsCollectionRequestBuilder {
	bb := &CalendarGroupCalendarsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/calendars"
	return bb
}

// CalendarGroupCalendarsCollectionRequestBuilder is request builder for Calendar collection
type CalendarGroupCalendarsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Calendar collection
func (b *CalendarGroupCalendarsCollectionRequestBuilder) Request() *CalendarGroupCalendarsCollectionRequest {
	return &CalendarGroupCalendarsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Calendar item
func (b *CalendarGroupCalendarsCollectionRequestBuilder) ID(id string) *CalendarRequestBuilder {
	bb := &CalendarRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CalendarGroupCalendarsCollectionRequest is request for Calendar collection
type CalendarGroupCalendarsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Calendar collection
func (r *CalendarGroupCalendarsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Calendar, error) {
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
	var values []Calendar
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
			value  []Calendar
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

// GetN performs GET request for Calendar collection, max N pages
func (r *CalendarGroupCalendarsCollectionRequest) GetN(ctx context.Context, n int) ([]Calendar, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Calendar collection
func (r *CalendarGroupCalendarsCollectionRequest) Get(ctx context.Context) ([]Calendar, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Calendar collection
func (r *CalendarGroupCalendarsCollectionRequest) Add(ctx context.Context, reqObj *Calendar) (resObj *Calendar, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BatchGet adds Get operation to Batch for Calendar collection
func (r *CalendarGroupCalendarsCollectionRequest) BatchGet(batch *BatchRequest) error {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	var resObj []Calendar
	return batch.Add("GET", strings.TrimPrefix(r.baseURL+query, defaultBaseURL), nil, resObj)
}

// BatchAdd adds Add operation to Batch for Calendar collection
func (r *CalendarGroupCalendarsCollectionRequest) BatchAdd(batch *BatchRequest, reqObj *Calendar) error {
	var resObj []Calendar
	return batch.Add("POST", strings.TrimPrefix(r.baseURL, defaultBaseURL), reqObj, resObj)
}
