package vcoclient

type ConfigurationDeleteConfigurationResult struct {
	// The id of the deleted object.
	Id int32 `json:"id,omitempty"`
	// An error message explaining why the method failed
	Error_ string `json:"error,omitempty"`
	// The number of rows modified
	Rows int32 `json:"rows"`
}
