package vcoclient

type ConfigurationCloneEnterpriseTemplate struct {
	EnterpriseId      int    `json:"enterpriseId,omitempty"`
	ConfigurationType string `json:"configurationType,omitempty"`
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
}
