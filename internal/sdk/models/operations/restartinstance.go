// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/squat/terraform-provider-lambda/internal/sdk/models/shared"
	"net/http"
)

type RestartInstanceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Restart *shared.ResponseRestart
	// Unauthorized.
	ErrorResponseBody *shared.ErrorResponseBody
}

func (o *RestartInstanceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RestartInstanceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RestartInstanceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RestartInstanceResponse) GetRestart() *shared.ResponseRestart {
	if o == nil {
		return nil
	}
	return o.Restart
}

func (o *RestartInstanceResponse) GetErrorResponseBody() *shared.ErrorResponseBody {
	if o == nil {
		return nil
	}
	return o.ErrorResponseBody
}