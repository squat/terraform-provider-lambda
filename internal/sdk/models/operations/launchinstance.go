// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/squat/terraform-provider-lambda/internal/sdk/models/shared"
	"net/http"
)

type LaunchInstanceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Launch *shared.ResponseLaunch
	// Unauthorized.
	ErrorResponseBody *shared.ErrorResponseBody
}

func (o *LaunchInstanceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LaunchInstanceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LaunchInstanceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *LaunchInstanceResponse) GetLaunch() *shared.ResponseLaunch {
	if o == nil {
		return nil
	}
	return o.Launch
}

func (o *LaunchInstanceResponse) GetErrorResponseBody() *shared.ErrorResponseBody {
	if o == nil {
		return nil
	}
	return o.ErrorResponseBody
}
