// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/squat/terraform-provider-lambda/internal/sdk/models/shared"
	"net/http"
)

type InstanceTypesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	InstanceTypes *shared.InstanceTypes
	// Unauthorized.
	ErrorResponseBody *shared.ErrorResponseBody
}

func (o *InstanceTypesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *InstanceTypesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *InstanceTypesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *InstanceTypesResponse) GetInstanceTypes() *shared.InstanceTypes {
	if o == nil {
		return nil
	}
	return o.InstanceTypes
}

func (o *InstanceTypesResponse) GetErrorResponseBody() *shared.ErrorResponseBody {
	if o == nil {
		return nil
	}
	return o.ErrorResponseBody
}
