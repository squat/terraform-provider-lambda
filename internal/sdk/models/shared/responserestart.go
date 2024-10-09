// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type RestartData struct {
	// List of instances that were restarted. Note: this list might not contain all instances requested to be restarted.
	RestartedInstances []Instance `json:"restarted_instances"`
}

func (o *RestartData) GetRestartedInstances() []Instance {
	if o == nil {
		return []Instance{}
	}
	return o.RestartedInstances
}

// ResponseRestart - OK
type ResponseRestart struct {
	Data RestartData `json:"data"`
}

func (o *ResponseRestart) GetData() RestartData {
	if o == nil {
		return RestartData{}
	}
	return o.Data
}
