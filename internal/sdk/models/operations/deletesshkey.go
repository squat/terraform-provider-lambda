// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/squat/terraform-provider-lambda/internal/sdk/models/shared"
	"net/http"
)

type DeleteSSHKeyRequest struct {
	// The unique identifier (ID) of the SSH key
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteSSHKeyRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteSSHKeyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Request parameters were invalid.
	ErrorResponseBody *shared.ErrorResponseBody
}

func (o *DeleteSSHKeyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteSSHKeyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteSSHKeyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteSSHKeyResponse) GetErrorResponseBody() *shared.ErrorResponseBody {
	if o == nil {
		return nil
	}
	return o.ErrorResponseBody
}
