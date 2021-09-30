package vcoclient

type ConfigurationGetConfiguration struct {
	Id            int      `json:"id,omitempty"`
	EnteerpriseId int      `json:"enterpriseId,omitempty"`
	With          []string `json:"with,omitempty"`
}
