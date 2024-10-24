// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Update struct {
	// Unique identifier (ID) of an instance
	ID string `json:"id"`
	// User-provided name for the instance
	Name *string `json:"name"`
}

func (o *Update) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Update) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
