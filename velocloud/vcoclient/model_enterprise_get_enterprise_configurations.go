package vcoclient

type EnterpriseGetEnterpriseConfigurations struct {
	EnterpriseId      int      `json:"enterpriseId,omitempty"`
	ConfigurationType string   `json:"configurationType,omitempty"`
	Name              string   `json:"name,omitempty"`
	With              []string `json:"with,omitempty"`
}
