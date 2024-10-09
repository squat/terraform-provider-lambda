// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// ErrorResponseBody - Unauthorized.
type ErrorResponseBody struct {
	Error Error `json:"error"`
	// Details about errors on a per-parameter basis
	FieldErrors map[string]Error `json:"field_errors,omitempty"`
}

func (o *ErrorResponseBody) GetError() Error {
	if o == nil {
		return Error{}
	}
	return o.Error
}

func (o *ErrorResponseBody) GetFieldErrors() map[string]Error {
	if o == nil {
		return nil
	}
	return o.FieldErrors
}