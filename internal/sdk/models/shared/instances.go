// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Instances struct {
	Data []Instance `json:"data"`
}

func (o *Instances) GetData() []Instance {
	if o == nil {
		return []Instance{}
	}
	return o.Data
}
