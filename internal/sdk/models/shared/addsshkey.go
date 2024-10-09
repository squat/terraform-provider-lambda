// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// AddSSHKey - The name for the SSH key. Optionally, an existing public key can be supplied for the `public_key` property. If the `public_key` property is omitted, a new key pair is generated. The private key is returned in the response.
type AddSSHKey struct {
	// Name of the SSH key
	Name string `json:"name"`
	// Public key for the SSH key
	PublicKey *string `json:"public_key,omitempty"`
}

func (o *AddSSHKey) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AddSSHKey) GetPublicKey() *string {
	if o == nil {
		return nil
	}
	return o.PublicKey
}
