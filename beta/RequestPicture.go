// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PictureRequestBuilder is request builder for Picture
type PictureRequestBuilder struct{ BaseRequestBuilder }

// Request returns PictureRequest
func (b *PictureRequestBuilder) Request() *PictureRequest {
	return &PictureRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PictureRequest is request for Picture
type PictureRequest struct{ BaseRequest }

// Get performs GET request for Picture
func (r *PictureRequest) Get(ctx context.Context) (resObj *Picture, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Picture
func (r *PictureRequest) Update(ctx context.Context, reqObj *Picture) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Picture
func (r *PictureRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
