// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// SSHKeys - OK
type SSHKeys struct {
	Data []SSHKey `json:"data"`
}

func (o *SSHKeys) GetData() []SSHKey {
	if o == nil {
		return []SSHKey{}
	}
	return o.Data
}