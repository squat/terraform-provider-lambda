// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type TerminateData struct {
	// List of instances that were terminated. Note: this list might not contain all instances requested to be terminated.
	TerminatedInstances []Instance `json:"terminated_instances"`
}

func (o *TerminateData) GetTerminatedInstances() []Instance {
	if o == nil {
		return []Instance{}
	}
	return o.TerminatedInstances
}

// ResponseTerminate - OK
type ResponseTerminate struct {
	Data TerminateData `json:"data"`
}

func (o *ResponseTerminate) GetData() TerminateData {
	if o == nil {
		return TerminateData{}
	}
	return o.Data
}
